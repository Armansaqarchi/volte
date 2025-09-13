package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alicebob/miniredis"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"volte/backend/chain"
	"volte/backend/databases"

	"github.com/gin-gonic/gin"
)

type FakeVolteSession struct {
	nullifierMerkleRoots map[string] /* eventID */ []byte /* NullifierRootHash */
	voteMerkleRoots      map[string] /* eventID */ []byte /* VoteRootHash */
	eventHashes          map[string] /* eventID */ []byte /* EventDetailsHash */
	usedNullifiers       map[string] /* eventID */ []byte /* UsedNullifiers */
}

func NewFakeVolteContract() chain.VolteSessionHandler {
	return &FakeVolteSession{
		nullifierMerkleRoots: make(map[string][]byte),
		voteMerkleRoots:      make(map[string][]byte),
		eventHashes:          make(map[string][]byte),
		usedNullifiers:       make(map[string][]byte),
	}
}

func (v *FakeVolteSession) SetNullifierMerkleRoot(eventID *big.Int, value []byte) (*types.Transaction, error) {
	v.nullifierMerkleRoots[eventID.String()] = value
	return nil, nil
}
func (v *FakeVolteSession) SetVoteMerkleRoot(eventID *big.Int, value []byte) (*types.Transaction, error) {
	v.voteMerkleRoots[eventID.String()] = value
	return nil, nil
}
func (v *FakeVolteSession) SetEventHash(eventID *big.Int, value []byte) (*types.Transaction, error) {
	v.eventHashes[eventID.String()] = value
	return nil, nil
}
func (v *FakeVolteSession) GetNullifierMerkleRoot(eventID *big.Int) ([]byte, error) {
	return v.nullifierMerkleRoots[eventID.String()], nil
}
func (v *FakeVolteSession) GetVoteMerkleRoot(eventID *big.Int) ([]byte, error) {
	return v.voteMerkleRoots[eventID.String()], nil
}
func (v *FakeVolteSession) GetEventHash(eventID *big.Int) ([]byte, error) {
	return v.eventHashes[eventID.String()], nil
}

type FakeContractManager struct {
	contractHandler chain.ContractHandler
}

func NewFakeContractManager() chain.ContractHandler {
	return &FakeContractManager{
		contractHandler: NewFakeContractManager(),
	}
}

func (cm *FakeContractManager) GetClient() *ethclient.Client {
	return nil
}
func (cm *FakeContractManager) GetFromAddress() common.Address {
	return common.Address{}
}
func (cm *FakeContractManager) GetVolteContract() chain.VolteSessionHandler {
	return NewFakeVolteContract()
}

func newFakeService() *VotingService {
	redisServer := miniredis.NewMiniRedis()
	redisClient := databases.RedisClientProvider{
		RedisClient: redis.NewClient(&redis.Options{Addr: redisServer.Addr()}),
	}
	return &VotingService{
		keyValDB:        redisClient,
		contractHandler: NewFakeContractManager(),
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
	service := newFakeService()

	ctx, recorder := newTestGinContext()
	postRequest := createFakePostRequest(t,
		map[string]interface{}{
			"event_id": 2,
		}, "/target",
	)
	ctx.Request = postRequest
	service.AddMemberToEvent(ctx)
	fmt.Println(recorder.Code)
}

func TestStartEvent(t *testing.T) {

}
