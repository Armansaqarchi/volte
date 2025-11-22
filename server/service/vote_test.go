package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"math"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"volte/backend/chain/contracts"
	"volte/backend/crypto/utils"
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
		Admin:       "fake_id",
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
	eventID := "3523196653250260958887739657950671762678466692388251624290163732010351636053"
	m := 99
	x := big.NewInt(30)

	proofs := contracts.VolteContractVoteSubmission{
		Sender:  service.contractHandler.GetFromAddress(),
		EventID: eventID,
		Proofs: contracts.VolteContractProofs{
			Ballot: contracts.VolteContractBallotProof{
				Proof: contracts.VolteContractProof{
					Arx:  newBigIntFromString("18901350265255527581095513162312611932070578267688684455001092092092699306713"),
					Ary:  newBigIntFromString("3973684204030300040690544173926537110512864262307661299031782109964113989725"),
					Brx1: newBigIntFromString("456791595468160720123795738298254930436516128847963769132229398988356376030"),
					Brx0: newBigIntFromString("13913927961784605313610362288468304063108896762953313755212627088358501898468"),
					Bry1: newBigIntFromString("20961664021107855764888910908507914831244791958992598287726959701312203587233"),
					Bry0: newBigIntFromString("430818465704284692024495878845876013841578472946195639572152467970619270220"),
					Cx:   newBigIntFromString("16101854476106560621809940543357164739975974176878439052171173818605894443285"),
					Cy:   newBigIntFromString("8923679615216762351321608510532073342724028507604092068314420184267808105245"),
				},
				Input: [16]*big.Int{
					newBigIntFromString("16068448641403006952"),
					newBigIntFromString("11401819955764306896"),
					newBigIntFromString("14034514348447911692"),
					newBigIntFromString("1161599979058448361"),
					newBigIntFromString("11848544127381919838"),
					newBigIntFromString("10110847344335578921"),
					newBigIntFromString("13360974023958266782"),
					newBigIntFromString("571327339866239689"),
					newBigIntFromString("14145516526586811472"),
					newBigIntFromString("6909929774512124840"),
					newBigIntFromString("8889541373244783464"),
					newBigIntFromString("1339997705979671789"),
					newBigIntFromString("3157877771170468305"),
					newBigIntFromString("16144328623067908114"),
					newBigIntFromString("10722905077406537962"),
					newBigIntFromString("2404946929832722734"),
				},
				CommitmentX:    newBigIntFromString("11781742062718661832194844182016150645681761708591013256714969837857880377122"),
				CommitmentY:    newBigIntFromString("1549744844124333981450801949290951318683081156734707012053285393097950706085"),
				CommitmentPokX: newBigIntFromString("12591565341333321479450689931096834240842206009783585007411990446512892346420"),
				CommitmentPokY: newBigIntFromString("19804122319499367039252227027992652037022224898454221268218417509484251223509"),
			},
			Membership: contracts.VolteContractMembershipProof{
				Proof: contracts.VolteContractProof{
					Arx:  newBigIntFromString("218307417371761318217510504769136052452728445114696001036883765377063469724"),
					Ary:  newBigIntFromString("20090163883772003059825487989424674536934171275888101565903843320568431785156"),
					Brx1: newBigIntFromString("20672981127299851928153146471267540643191282834601289424189154306151666205115"),
					Brx0: newBigIntFromString("4172086382342582216562454009513565487757169922717923889675064727095550155749"),
					Bry1: newBigIntFromString("17939853704919616285850575439447615788332440851643822120260123931481880196413"),
					Bry0: newBigIntFromString("21425700694512897993682652343295206699287853314112930253564539516723211313028"),
					Cx:   newBigIntFromString("14912915543876692391392469854197380978673110738187575389273652791900426988366"),
					Cy:   newBigIntFromString("11014960880009887574094976234921461537265317658907218783575181959733851237046"),
				},
				Input: [2]*big.Int{
					newBigIntFromString("1445849190805089689712507754685671168673091553372230482406710117687961492369"),
					newBigIntFromString("4137760094704180852789719500758563423980885922685717827383305955441808899436"),
				},
			},
			Nullifier: contracts.VolteContractNullifierProof{
				Proof: contracts.VolteContractProof{
					Arx:  newBigIntFromString("19299394707588237889842204045380945831468246783948553771064559703390041220853"),
					Ary:  newBigIntFromString("4071137541844626537200040579510633652956775219111126463167504656912249659108"),
					Brx1: newBigIntFromString("19151520548966540575565064419782787574054224241906696011061027507795307865775"),
					Brx0: newBigIntFromString("9974598127107197347902395079144830182776403908580955216438534257515427271736"),
					Bry1: newBigIntFromString("2500492010900186065866422064845897335014621085330341328994588041624033165681"),
					Bry0: newBigIntFromString("17660984402210788913374345393465049457076285707709524303644798707434662042479"),
					Cx:   newBigIntFromString("10805355095395483666986219112228748527423776665232009971299604346256688772744"),
					Cy:   newBigIntFromString("1870504571597614792738126687608435171464683251354744603329741393782069333371"),
				},
				Input: [2]*big.Int{
					newBigIntFromString("3523196653250260958887739657950671762678466692388251624290163732010351636053"),
					newBigIntFromString("10858051838952645709440492871522823286939885463470476254383445187743108776413"),
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
	request = createFakePostRequest(t, nil, fmt.Sprintf("/tally?eventID=%s", eventID))
	ctx, recorder = newTestGinContext()
	ctx.Request = request
	service.GetTallyScore(ctx)
	var tallyResponse map[string][]*big.Int
	if err := json.Unmarshal(recorder.Body.Bytes(), &tallyResponse); err != nil {
		t.Fatal(err)
	}
	tallyScore := tallyResponse["score"]
	C1, _ := utils.MakeG1Affine(tallyScore[0].String(), tallyScore[1].String())
	C2, _ := utils.MakeG1Affine(tallyScore[2].String(), tallyScore[3].String())
	// The private key used to decrypt the elgamal encryption.
	// This is bound and is derived before the encryption phase.
	C1x := utils.G1MulAffine(&C1, x)
	// m.(G) = C2 - x.(C1)
	M := utils.G1AddAffine(&C2, C1x.Neg(&C1x))
	G1 := utils.GenerateBaseECC()
	// for calculating the raw m, we use an optimized brute force approach called BSGS.
	// The Baby step giant step algorithm, calculates the raw m in O(squared(n)).
	// In our case, it takes about O(radical(2 ** n)) which in case is about 2*16 operations.
	got, _ := utils.BSGS(M, G1, uint64(math.Pow(2, 32)))
	assert.Equal(t, got, m)
}
