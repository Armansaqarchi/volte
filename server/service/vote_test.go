package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/txaty/go-merkletree"
	"log/slog"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"volte/backend/models"
	"volte/backend/utils/test"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func stringPtr(str string) *string {
	s := new(string)
	*s = str
	return s
}

func newFakeService(t *testing.T) *VotingService {
	eventCollection = stringPtr("testEventCollection")
	commitmentCollection = stringPtr("testCommitmentCollection")
	database = stringPtr("testDatabase")

	test.CreateNewFakeMongoServer(t)
	return &VotingService{
		mongoClient:     test.NewFakeMongoClient(),
		contractHandler: test.NewFakeContractManager(),
		// Add groth16 mocks as well
	}
}

func newTestGinContext() (*gin.Context, *httptest.ResponseRecorder) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{URL: &url.URL{}, Header: http.Header{}, RemoteAddr: "7.7.7.7"}
	return ctx, recorder
}

func createFakePostRequest(t *testing.T, body map[string]interface{}, target string) *http.Request {
	jsonData, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("Failed to convert body to json: %v", err)
	}
	postRequest := httptest.NewRequest("POST", target, bytes.NewBuffer(jsonData))
	postRequest.Header.Add("Content-Type", "application/json")
	return postRequest
}

func createFakeGetRequest(baseURL string, routeParams map[string]string,
	queryParams map[string]string) (*http.Request, error) {
	// Construct the URL with the base URL and route parameters.
	routeURL := baseURL
	for key, value := range routeParams {
		routeURL = strings.Replace(routeURL, fmt.Sprintf("{%s}", key), value, -1)
	}

	// Construct the query parameters.
	queryValues := url.Values{}
	for key, value := range queryParams {
		queryValues.Add(key, value)
	}

	// Construct the full URL with query parameters.
	fullURL := fmt.Sprintf("%s?%s", routeURL, queryValues.Encode())
	req := httptest.NewRequest("GET", fullURL, nil)

	return req, nil
}

func TestAddMemberToEvent(t *testing.T) {
	service := newFakeService(t)

	eventId := "2"
	commitment := "some random commitment"

	event := &models.Event{
		ID:          eventId,
		Name:        "test_event",
		Admin:       "AdminFake",
		Duration:    time.Hour * 24 * 3,
		StartTime:   time.Now().Add(-time.Hour * 24 * 1),
		VoteOptions: make([]string, 0),
		VoteMembers: make([]string, 0),
		Tally:       make(map[string]int),
		Revoked:     false,
	}
	eventsCollection := service.mongoClient.GetClient().Database(*database).Collection(*eventCollection)
	if _, err := eventsCollection.InsertOne(t.Context(), event); err != nil {
		t.Fatal(err)
	}
	_, err := service.contractHandler.GetVolteContract().SetEventHash(event.ID, event.CalculateEventHash())
	if err != nil {
		t.Fatal(err)
	}

	ctx, recorder := newTestGinContext()
	postRequest := createFakePostRequest(t, map[string]interface{}{}, "/target")
	ctx.AddParam("event_id", event.ID)
	ctx.AddParam("commitment", commitment)
	ctx.Request = postRequest
	service.AddMemberToEvent(ctx)
	assert.Equal(t, recorder.Code, http.StatusOK)
	if err := eventsCollection.FindOne(ctx, bson.M{"_id": eventId}).Decode(&event); err != nil {

		slog.Error(fmt.Sprintf("Couldnt find event %s", strconv.FormatInt(eventId, 10)))
		t.Fatal(err.Error())
	}
	assert.Equal(t, event.VoteMembers, []string{commitment})
	eventHash, err := service.contractHandler.GetVolteContract().GetEventHash(big.NewInt(event.ID))
	if err != nil {
		t.Fatal(err.Error())
	}
	assert.Equal(t, event.CalculateEventHash(), eventHash)
}

func TestStartEvent(t *testing.T) {
	service := newFakeService(t)

	eventId := "2"

	event := &models.Event{
		ID:          eventId,
		Name:        "test_event",
		Admin:       "AdminFake",
		Duration:    time.Hour * 24 * 3,
		VoteOptions: make([]string, 0),
		VoteMembers: []string{"FakeMember1", "FakeMember2", "FakeMember3", "FakeMember4"},
		Tally:       make(map[string]int),
		Revoked:     false,
	}
	commitmentsCollection := service.mongoClient.GetClient().Database(*database).Collection(*commitmentCollection)
	eventsCollection := service.mongoClient.GetClient().Database(*database).Collection(*eventCollection)
	if _, err := eventsCollection.InsertOne(t.Context(), event); err != nil {
		t.Fatal(err)
	}
	_, err := service.contractHandler.GetVolteContract().SetEventHash(event.ID, event.CalculateEventHash())
	if err != nil {
		t.Fatal(err)
	}

	ctx, recorder := newTestGinContext()
	postRequest := createFakePostRequest(t, map[string]interface{}{}, "/target")
	ctx.AddParam("event_id", event.ID)
	ctx.Request = postRequest
	service.StartEvent(ctx)
	assert.Equal(t, recorder.Code, http.StatusOK)

	eventHash, err := service.contractHandler.GetVolteContract().GetEventHash(event.ID)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, event.CalculateEventHash(), eventHash)
	var got models.EventTree
	if err := commitmentsCollection.FindOne(ctx, bson.M{"_id": eventId}).Decode(&got); err != nil {
		slog.Error(fmt.Sprintf("Couldnt find event %s", eventId))
	}
	var expectedCommitments []merkletree.DataBlock
	for _, member := range event.VoteMembers {
		expectedCommitments = append(expectedCommitments, models.Commitment(member))
	}
	expectedCommitmentsTree, _ := merkletree.New(
		&merkletree.Config{Mode: merkletree.ModeProofGenAndTreeBuild}, expectedCommitments,
	)
	expected := models.EventTree{ID: eventId, Tree: expectedCommitmentsTree}
	assert.Equal(t, expected.Tree.Proofs, got.Tree.Proofs)
	assert.Equal(t, expected.Tree.Root, got.Tree.Root)
	assert.Equal(t, expected.Tree.Leaves, got.Tree.Leaves)
	assert.Equal(t, expected.Tree.Nodes, got.Tree.Nodes)

}
