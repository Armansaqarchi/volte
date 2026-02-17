package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"
	"volte/backend/models"
	"volte/backend/utils/test"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/txaty/go-merkletree"
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
		contractHandler: test.NewFakeContractHandler(),
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

func createFakePostRequest(t *testing.T, body interface{}, target string) *http.Request {
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

// newTestGinContextWithSession creates a Gin context with an embedded session for testing.
func newTestGinContextWithSession() (*gin.Context, *httptest.ResponseRecorder, sessions.Session) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)

	req := httptest.NewRequest("GET", "/", nil)
	ctx.Request = req

	store := cookie.NewStore([]byte(*sessionSecret))
	sessionName := "user_session"

	sessionMiddleware := sessions.Sessions(sessionName, store)
	sessionMiddleware(ctx)

	session := sessions.Default(ctx)

	return ctx, recorder, session
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
		VoteOptions: make([]string, 0),
		VoteMembers: make([]string, 0),
		Tally:       make(map[string]int),
		Revoked:     false,
	}
	startTime := time.Now().Add(-time.Hour * 24 * 1)
	event.StartTime = &startTime
	eventsCollection := service.mongoClient.GetClient().Database(*database).Collection(*eventCollection)
	if _, err := eventsCollection.InsertOne(t.Context(), event); err != nil {
		t.Fatal(err)
	}
	_, err := service.contractHandler.GetVolteContract().SetEventHash(event.ID, event.CalculateEventHash())
	if err != nil {
		t.Fatal(err)
	}
	test.GetFakeChain().Backend.Commit()

	ctx, recorder := newTestGinContext()
	postRequest := createFakePostRequest(t, map[string]interface{}{}, "/target")
	ctx.AddParam("id", event.ID)
	ctx.AddParam("commitment", commitment)
	ctx.Request = postRequest
	service.AddMemberToEvent(ctx)
	assert.Equal(t, recorder.Code, http.StatusOK)
	if err := eventsCollection.FindOne(ctx, bson.M{"_id": eventId}).Decode(&event); err != nil {

		slog.Error(fmt.Sprintf("Couldnt find event %s", eventId))
		t.Fatal(err.Error())
	}
	assert.Equal(t, event.VoteMembers, []string{commitment})
	test.Chain.Backend.Commit()
	eventHash, err := service.contractHandler.GetVolteContract().GetEventHash(eventId)
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
		Admin:       "321425363245",
		Duration:    time.Hour * 24 * 3,
		VoteOptions: make([]string, 0),
		VoteMembers: []string{"FakeMember1", "FakeMember2", "FakeMember3", "FakeMember4"},
		Tally:       make(map[string]int),
		Revoked:     false,
	}
	commitmentsCollection := service.mongoClient.GetClient().Database(*database).Collection(*commitmentCollection)
	eventsCollection := service.mongoClient.GetClient().Database(*database).Collection(*eventCollection)
	commitmentsMerklePathCollection := service.mongoClient.GetClient().Database(
		*database).Collection(*commitmentMerklePathCollection)
	if _, err := eventsCollection.InsertOne(t.Context(), event); err != nil {
		t.Fatal(err)
	}
	_, err := service.contractHandler.GetVolteContract().SetEventHash(event.ID, event.CalculateEventHash())
	if err != nil {
		t.Fatal(err)
	}
	test.GetFakeChain().Backend.Commit()
	ctx, recorder, session := newTestGinContextWithSession()
	session.Set("user", "fake_id")
	postRequest := createFakePostRequest(t, map[string]interface{}{}, "/target")
	ctx.AddParam("id", event.ID)
	ctx.Request = postRequest

	service.StartEvent(ctx)
	assert.Equal(t, recorder.Code, http.StatusOK)

	eventHash, err := service.contractHandler.GetVolteContract().GetEventHash(event.ID)
	if err != nil {
		panic(err)
	}
	assert.Equal(t, event.CalculateEventHash(), eventHash)
	var got models.EventTree
	var gotCommitmentPath models.EventTreeProofsDto
	if err := commitmentsCollection.FindOne(ctx, bson.M{"_id": eventId}).Decode(&got); err != nil {
		slog.Error(fmt.Sprintf("Couldnt find commitment %s", eventId))
	}
	if err := commitmentsMerklePathCollection.FindOne(ctx, bson.M{"_id": eventId}).Decode(&gotCommitmentPath); err != nil {
		slog.Error(fmt.Sprintf("Couldnt find commitment merkle path %s.", eventId))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to fetch the event details.",
		})
		return
	}
	var expectedCommitments []merkletree.DataBlock
	expectedCommitments = append(expectedCommitments, models.Commitment(event.Admin))
	for _, member := range event.VoteMembers {
		expectedCommitments = append(expectedCommitments, models.Commitment(member))
	}
	expectedCommitmentsTree, _ := merkletree.New(
		&merkletree.Config{Mode: merkletree.ModeProofGenAndTreeBuild}, expectedCommitments,
	)
	expected := models.EventTree{ID: eventId, Tree: expectedCommitmentsTree}
	assert.Equal(t, expected.Tree.Proofs, gotCommitmentPath.Proofs)
	assert.Equal(t, expected.Tree.Proofs, got.Tree.Proofs)
	assert.Equal(t, expected.Tree.Root, got.Tree.Root)
	assert.Equal(t, expected.Tree.Leaves, got.Tree.Leaves)
	assert.Equal(t, expected.Tree.Nodes, got.Tree.Nodes)
}

func TestUserEvents(t *testing.T) {
	service := newFakeService(t)
	authService := NewFakeAuthService(t)

	ctx, recorder, _ := newTestGinContextWithSession()

	req := createFakePostRequest(t, map[string]any{
		"username": "example_username",
		"password": "exmpale_password",
	}, "/auth/signup")
	ctx.Request = req
	authService.Register(ctx)
	type Result struct {
		data models.User
	}
	var user Result
	if err := json.Unmarshal(recorder.Body.Bytes(), &user); err != nil {
		t.Fatal(err)
	}

	ctx, recorder, _ = newTestGinContextWithSession()
	req, err := createFakeGetRequest(
		"/users/:key/events", map[string]string{"key": user.data.Commitment}, nil,
	)
	ctx.Request = req
	if err != nil {
		t.Fatal(err)
	}
	service.UserEvents(ctx)
}

func TestCreateEvent(t *testing.T) {
	service := newFakeService(t)
	authService := NewFakeAuthService(t)

	ctx, recorder, _ := newTestGinContextWithSession()

	req := createFakePostRequest(t, map[string]any{
		"username":   "example_username",
		"password":   "example_password",
		"commitment": "fake_commitment",
	}, "/auth/signup")
	ctx.Request = req
	authService.Register(ctx)
	type Result struct {
		Data models.User `json:"data"`
	}
	var user Result
	if err := json.Unmarshal(recorder.Body.Bytes(), &user); err != nil {
		t.Fatal(err)
	}

	ctx, recorder, session := newTestGinContextWithSession()
	req = createFakePostRequest(t, map[string]any{
		"name":         "Project Approval Vote",
		"duration":     360000000,
		"question":     "ey baba",
		"vote_options": []string{"Approve", "Reject"},
	}, "/auth/signup")
	ctx.Request = req
	session.Set("user", user.Data.Commitment)
	service.CreateEvent(ctx)
}

//func TestVotingService_VoteAndGetTallyScore(t *testing.T) {
//	service := newFakeService(t)
//	// Just a random event id with no encryption
//	eventID := "285361107209702954985467434449777005623"
//	m := uint64(1)
//
//	root, _ := big.NewInt(0).SetString("13594411463883921251454988740927603831334698588320922729025732314421451405961", 10)
//	if _, err := service.contractHandler.GetVolteContract().SetVoteMerkleRoot(eventID, root); err != nil {
//		t.Fatal(err)
//	}
//	test.GetFakeChain().Backend.Commit()
//
//	proofs := contracts.VolteContractVoteSubmission{
//		EventID: eventID,
//		Proofs: contracts.VolteContractProofs{
//			Ballot:     *proof.RunBallotProof(),
//			Membership: *proof.RunMerklePathProof(),
//			Nullifier:  *proof.RunNullifierProof(),
//		},
//	}
//	request := createFakePostRequest(t, proofs, "/vote")
//	ctx, recorder := newTestGinContext()
//	ctx.Request = request
//	service.Vote(ctx)
//	assert.Equal(t, recorder.Code, http.StatusOK)
//	// Commit chain so the transactions will be processed and the contract states will be updated.
//
//	test.GetFakeChain().Backend.Commit()
//	request = createFakePostRequest(t, nil, "/tally")
//	ctx, recorder = newTestGinContext()
//	ctx.Request = request
//	ctx.AddParam("id", eventID)
//
//	service.GetTallyScore(ctx)
//
//	var tallyResponse map[string]any
//	if err := json.Unmarshal(recorder.Body.Bytes(), &tallyResponse); err != nil {
//		t.Fatal(err)
//	}
//	fmt.Println("tally response : ", tallyResponse)
//	assert.Equal(t, tallyResponse["score"], m)
//}
