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
	"strings"
	"testing"
	"time"
	"volte/backend/chain/contracts"

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

func newBigIntFromString(str string) *big.Int {
	bigInt := big.NewInt(0)
	bigInt.SetString(str, 10)
	return bigInt
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
	test.GetFakeChain().Backend.Commit()

	ctx, recorder := newTestGinContext()
	postRequest := createFakePostRequest(t, map[string]interface{}{}, "/target")
	ctx.AddParam("event_id", event.ID)
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
	test.GetFakeChain().Backend.Commit()
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

func TestVotingService_Vote(t *testing.T) {
	service := newFakeService(t)
	eventID := "2"

	proofs := contracts.VolteContractVoteSubmission{
		Sender:  service.contractHandler.GetFromAddress(),
		EventID: eventID,
		Proofs: contracts.VolteContractProofs{
			Ballot: contracts.VolteContractBallotProof{
				Proof: contracts.VolteContractProof{
					Arx:  newBigIntFromString("18170799973635857261294822456123490767858300849315309338461975281292057361897"),
					Ary:  newBigIntFromString("14244284926918058364760444392977669585174379885827289936114452693528992997163"),
					Brx1: newBigIntFromString("8205706887668846080512646517510214276711411686303786932562449557988537815781"),
					Brx0: newBigIntFromString("12990447600059071286245895009349311798673194421436131053870023739348184059458"),
					Bry1: newBigIntFromString("9671061850765303257270787633253906737828422360121872275518493683153384565306"),
					Bry0: newBigIntFromString("21221193920824920418096999215436075741166768735025429467622161681613828356068"),
					Cx:   newBigIntFromString("1517193873151208544848678992355534033633141954616508000290254744282263078371"),
					Cy:   newBigIntFromString("528379084656834411015156790430078756051847835499804695870892096254824837298"),
				},
				Input: [16]*big.Int{
					newBigIntFromString("16068448641403006952"), newBigIntFromString("11401819955764306896"),
					newBigIntFromString("14034514348447911692"), newBigIntFromString("1161599979058448361"),
					newBigIntFromString("11848544127381919838"), newBigIntFromString("10110847344335578921"),
					newBigIntFromString("13360974023958266782"), newBigIntFromString("571327339866239689"),
					newBigIntFromString("14145516526586811472"), newBigIntFromString("6909929774512124840"),
					newBigIntFromString("8889541373244783464"), newBigIntFromString("1339997705979671789"),
					newBigIntFromString("3157877771170468305"), newBigIntFromString("16144328623067908114"),
					newBigIntFromString("10722905077406537962"), newBigIntFromString("2404946929832722734"),
				},
				CommitmentX:    newBigIntFromString("9157316909797196729284159350080547527003504213246209792610856577653016457815"),
				CommitmentY:    newBigIntFromString("10977044080192009214233157132198410281161392619246670423037870656121352481613"),
				CommitmentPokX: newBigIntFromString("94186758941001159818991885158478630669303642981805140890713727155036236730"),
				CommitmentPokY: newBigIntFromString("7757756877188287045888631358520676915214476184722814968395082837614480348173"),
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
}
