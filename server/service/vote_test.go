package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"volte/backend/chain/contracts"
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

func TestVotingService_VoteAndGetTallyScore(t *testing.T) {
	service := newFakeService(t)
	// Just a random event id with no encryption
	eventID := "33201053"
	m := uint64(0)

	root, _ := big.NewInt(0).SetString("2673035909906541938567902697210471713711742177182271385831434770364824838270", 10)
	if _, err := service.contractHandler.GetVolteContract().SetVoteMerkleRoot("33201053", root); err != nil {
		t.Fatal(err)
	}
	test.GetFakeChain().Backend.Commit()

	proofs := contracts.VolteContractVoteSubmission{
		EventID: eventID,
		Proofs: contracts.VolteContractProofs{
			Ballot: contracts.VolteContractBallotProof{
				Proof: contracts.VolteContractProof{
					Arx:  newBigIntFromString("9651545063954101077555658626007997755534362126980517841526368944250080351657"),
					Ary:  newBigIntFromString("1887932183960159063428064252056510771503690997109463346422270939021286922823"),
					Brx0: newBigIntFromString("12606867278310455749420446216668233647451630469380438991863017726617992962209"),
					Brx1: newBigIntFromString("217281117135355954356942095698500228319988868317343834516215804831346573439"),
					Bry0: newBigIntFromString("8284011433997405132236062370976449320871583690283528604952734595566414176806"),
					Bry1: newBigIntFromString("4135535540885055950134505064159603698790342691250936423511574753219909625104"),
					Cx:   newBigIntFromString("8474655738312475715889936512422086822038691411011435435126876831869196969369"),
					Cy:   newBigIntFromString("9758642134775720342067673185092262409084830033843993092728193796663044779017"),
				},
				Input: [16]*big.Int{
					newBigIntFromString("10460265752165522035"),
					newBigIntFromString("1007005022369442759"),
					newBigIntFromString("16219709270715832329"),
					newBigIntFromString("1098792919491332617"),
					newBigIntFromString("9590098487146431018"),
					newBigIntFromString("14892204839263526492"),
					newBigIntFromString("17394270043169223033"),
					newBigIntFromString("486073702419636823"),
					newBigIntFromString("13766268267297692530"),
					newBigIntFromString("13688256442378508665"),
					newBigIntFromString("14330586739089358222"),
					newBigIntFromString("3271505827464271239"),
					newBigIntFromString("2755087811367306977"),
					newBigIntFromString("6830942266724961423"),
					newBigIntFromString("13141105381902129043"),
					newBigIntFromString("1204908303416686802"),
				},
				CommitmentX:    newBigIntFromString("5463388527642851962734914983266376471450204296801422563285612236956608360819"),
				CommitmentY:    newBigIntFromString("15078915136568211978979162468536897084224948435903513907298060305871900350386"),
				CommitmentPokX: newBigIntFromString("14549953999564351437331625258649823046442469087546481959082733202369927619883"),
				CommitmentPokY: newBigIntFromString("11741489163135189610632801170682059292092503093620264253794541223982864377902"),
			},

			Membership: contracts.VolteContractMembershipProof{
				Proof: contracts.VolteContractProof{
					Arx:  newBigIntFromString("20190044418467755639135349476383604937264227105734655466668601385843656222949"),
					Ary:  newBigIntFromString("8487808233679998213464417942127489117597494124431878113464805994989003533607"),
					Brx0: newBigIntFromString("20729479858541298397238118687113059054905085449344709780583095705950427577914"),
					Brx1: newBigIntFromString("17423343450931144909301321138327298411894707919439300030740014215931141738175"),
					Bry0: newBigIntFromString("2182734738548382303975464309465437227973969583950334754059707373542303155521"),
					Bry1: newBigIntFromString("18481734051868214109455791505658397154657794191080390952816899383550516826765"),
					Cx:   newBigIntFromString("13225832168902722584586394134805946427609425972858932700551639269647323297235"),
					Cy:   newBigIntFromString("17571199486813906357848855680444080602662853718177079762304265445416721239469"),
				},
				Input: [2]*big.Int{
					newBigIntFromString("2673035909906541938567902697210471713711742177182271385831434770364824838270"),
					newBigIntFromString("19812511382630843747406943278967657446632006987648881852033406093539356710614"),
				},
			},

			Nullifier: contracts.VolteContractNullifierProof{
				Proof: contracts.VolteContractProof{
					Arx:  newBigIntFromString("17901375588089764233557400033861995640680680516553037339742983926110337224166"),
					Ary:  newBigIntFromString("14596865370548702332391164350858050927721831961627078913208265631744833600592"),
					Brx0: newBigIntFromString("9600165183033776171746205778281748437688496531967518525750078902875126590043"),
					Brx1: newBigIntFromString("3473164375368330787770601528752351002915037951824166415842025901497219285098"),
					Bry0: newBigIntFromString("4275250395403744442760577962714373175965840906852524181482176447740129610764"),
					Bry1: newBigIntFromString("10230542104736223586882831924854018937990921842392273254560634479677635280185"),
					Cx:   newBigIntFromString("107916262805368103018519733251533747970267096739992313406760952111835917298"),
					Cy:   newBigIntFromString("19719733751900284111313197815732314127948790976978394499742762562398764818970"),
				},
				Input: [3]*big.Int{
					newBigIntFromString("19812511382630843747406943278967657446632006987648881852033406093539356710614"),
					newBigIntFromString("255283938693201877191845615451726749197"),
					newBigIntFromString("10859401122032956014596409203061453637372336210313702122816716006917952629282"),
				},
			},
		},
	}
	request := createFakePostRequest(t, proofs, "/vote")
	ctx, recorder := newTestGinContext()
	ctx.Request = request
	service.Vote(ctx)
	assert.Equal(t, recorder.Code, http.StatusOK)
	// Commit chain so the transactions will be processed and the contract states will be updated.

	test.GetFakeChain().Backend.Commit()
	request = createFakePostRequest(t, nil, "/tally")
	ctx, recorder = newTestGinContext()
	ctx.Request = request
	ctx.AddParam("id", eventID)

	service.GetTallyScore(ctx)

	var tallyResponse map[string]any
	if err := json.Unmarshal(recorder.Body.Bytes(), &tallyResponse); err != nil {
		t.Fatal(err)
	}
	fmt.Println("tally response : ", tallyResponse)
	assert.Equal(t, tallyResponse["score"], m)
}
