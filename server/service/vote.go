package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"strconv"
	"volte/backend/chain"
	"volte/backend/crypto/constraintsys"
	"volte/backend/crypto/zkproofs"
	"volte/backend/models"

	"github.com/gin-gonic/gin"
)

type keyValDatabase interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}

type VotingService struct {
	keyValDB        keyValDatabase
	contractManager *chain.EthereumChainHandler

	volteGroth16      *zkproofs.Groth16
	cipherTextGroth16 *zkproofs.Groth16
	tallyGroth16      *zkproofs.Groth16
}

func NewVotingService(keyValueDB keyValDatabase, contractManager *chain.EthereumChainHandler) *VotingService {
	// Initialize a KV DB
	// Initialize ethereum contract client
	// fetch Groth16 specs from redis
	return &VotingService{
		keyValDB:          keyValueDB,
		contractManager:   contractManager,
		volteGroth16:      zkproofs.SetupNewGroth16(constraintsys.NewVolteBLS12377R1CS()),
		cipherTextGroth16: zkproofs.SetupNewGroth16(constraintsys.NewVolteBLS12377R1CS()),
		tallyGroth16:      zkproofs.SetupNewGroth16(constraintsys.NewVolteBLS12377R1CS()),
	}

}

func (v *VotingService) AddMemberToEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Query("event_id"), 10, 64)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to parse event_id, err : %s", err.Error()))
	}
	eventBytes, err := v.keyValDB.Get(fmt.Sprintf(fmt.Sprintf("volte_models_events_%d", eventId)))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get event %d from redis, err : %s", eventId, err.Error()))
	}
	var event models.Event
	if err := json.Unmarshal(eventBytes, &event); err != nil {
		slog.Error(fmt.Sprintf("Failed to unmarshal event, err : %s", err.Error()))
	}
	if event.Revoked {
		ctx.JSON(http.StatusForbidden, map[string]string{
			"message": "Event has been revoked due to security problems.",
		})
		return
	}
	eventHash, err := v.contractManager.GetVolteContract().GetEventHash(big.NewInt(eventId))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get event hash, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Failed to get event %s hash", event.ID),
		})
		return
	}
	if !bytes.Equal(eventHash, event.CalculateEventHash()) {
		slog.Warn(fmt.Sprintf("inconsistent event hash between chain and server for event : %s", event.ID))
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "inconsistent event hash between chain and server for event",
		})
	}
	identitySecret := ctx.Query("identity_secret")
	eventBytes, err = json.Marshal(event)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to encode event, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Failed to encode event"),
		})
	}
	event.VoteMembers = append(event.VoteMembers, identitySecret)
	_, err = v.contractManager.GetVolteContract().SetEventHash(
		big.NewInt(eventId), event.CalculateEventHash(),
	)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to set event hash, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Failed to set event hash on chain"),
		})
	}
	if err := v.keyValDB.Set("volte_models_events_%s", eventBytes); err != nil {
		slog.Error(fmt.Sprintf("Failed to store event hash, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": fmt.Sprintf("Failed to store event"),
		})
	}
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "Member successfully added.",
	})
}

func (v *VotingService) CreateEvent() {
	// Check authority
	// event specification
	// create event
	// create a corresponding incremental merkle trie
	// store the root hash value inside db
	// store the specifications hash inside the db so event specifications cannot change
}

func (v *VotingService) RemoveEvent() {
	// Check is owner
	// get event id
	// remove event from db
	// remove event's spec hash from chain
}

func (v *VotingService) Vote() {
	// check authority
	// pre-filter invalid votes to reduce gas fee as much as possible
	// check nullifier proof (via contract RPC call)
	// submit vote value and update incremental merkle tree

	// Note: use locking to avoid race condition
}
