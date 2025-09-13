package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"volte/backend/chain"
	"volte/backend/crypto/constraintsys"
	"volte/backend/crypto/zkproofs"
	"volte/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/txaty/go-merkletree"
)

type keyValDatabase interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, value interface{}, duration time.Duration) error
}

type VotingService struct {
	keyValDB        keyValDatabase
	contractHandler chain.ContractHandler

	volteGroth16      *zkproofs.Groth16
	cipherTextGroth16 *zkproofs.Groth16
	tallyGroth16      *zkproofs.Groth16
}

func NewVotingService(keyValueDB keyValDatabase, contractManager *chain.EthereumContractHandler) *VotingService {
	// Initialize a KV DB
	// Initialize ethereum contract client
	// fetch Groth16 specs from redis
	return &VotingService{
		keyValDB:          keyValueDB,
		contractHandler:   contractManager,
		volteGroth16:      zkproofs.SetupNewGroth16(constraintsys.NewVolteBLS12377R1CS()),
		cipherTextGroth16: zkproofs.SetupNewGroth16(constraintsys.NewVolteBLS12377R1CS()),
		tallyGroth16:      zkproofs.SetupNewGroth16(constraintsys.NewVolteBLS12377R1CS()),
	}

}

func (v *VotingService) AddMemberToEvent(ctx *gin.Context) {
	// add security step
	eventId, err := strconv.ParseInt(ctx.Query("event_id"), 10, 64)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to parse event_id, err : %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Failed to parse event_id",
		})
	}
	eventBytes, err := v.keyValDB.Get(
		ctx, fmt.Sprintf(fmt.Sprintf("volte:models:events:%d", eventId)),
	)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get event %d from redis, err : %s", eventId, err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to fetch event %d", eventId),
		})
		return
	}
	var event models.Event
	if err := json.Unmarshal(eventBytes, &event); err != nil {
		slog.Error(fmt.Sprintf("Failed to unmarshal event, err : %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if event.Revoked {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Event has been revoked due to security problems.",
		})
		return
	}
	eventHash, err := v.contractHandler.GetVolteContract().GetEventHash(big.NewInt(eventId))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get event hash, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to get event %s hash", event.ID),
		})
		return
	}
	if !bytes.Equal(eventHash, event.CalculateEventHash()) {
		slog.Warn(fmt.Sprintf("inconsistent event hash between chain and server for event : %s", event.ID))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "inconsistent event hash between chain and server for event",
		})
		return
	}
	identitySecret := ctx.Query("identity_secret")
	eventBytes, err = json.Marshal(event)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to encode event, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to encode event"),
		})
	}
	event.VoteMembers = append(event.VoteMembers, identitySecret)
	_, err = v.contractHandler.GetVolteContract().SetEventHash(
		big.NewInt(eventId), event.CalculateEventHash(),
	)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to set event hash, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to set event hash on chain"),
		})
		return
	}
	if err := v.keyValDB.Set(ctx, "volte:models:events:%s", eventBytes, 10*time.Second); err != nil {
		slog.Error(fmt.Sprintf("Failed to store event hash, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to store event"),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Member successfully added.",
	})
}

func (v *VotingService) StartEvent(ctx *gin.Context) {
	eventID, err := strconv.ParseInt(ctx.Query("event_id"), 10, 64)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to parse event_id, err : %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Failed to parse event_id"})
	}
	eventBytes, err := v.keyValDB.Get(ctx, fmt.Sprintf(fmt.Sprintf("volte:models:events_%d", eventID)))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get event %d from redis, err : %s", eventID, err.Error()))
		ctx.JSON(
			http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to fetch event %d", eventID)},
		)
		return
	}
	var event models.Event
	if err := json.Unmarshal(eventBytes, &event); err != nil {
		slog.Error(fmt.Sprintf("Failed to unmarshal event, err : %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Failed to unmarshal event")})
		return
	}
	if err := event.StartEvent(); err != nil {
		slog.Error(fmt.Sprintf("Failed to start event, err : %s", err.Error()))
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Failed to start event, err : %s", event.ID)},
		)
	}
	var commitments []merkletree.DataBlock
	for _, member := range event.VoteMembers {
		commitments = append(commitments, models.Commitment(member))
	}
	commitmentsTree, err := merkletree.New(nil, commitments)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to create commitments tree, err : %s", err.Error()))
		ctx.JSON(
			http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to create commitments tree")},
		)
	}
	commitmentsTreeBytes, err := json.Marshal(commitmentsTree)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to marshal commitments, err : %s", err.Error()))
		ctx.JSON(
			http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to marshal commitments")},
		)
	}
	if err := v.keyValDB.Set(ctx,
		fmt.Sprintf("volte:models:events:trees:%d", eventID), commitmentsTreeBytes, 10*time.Second,
	); err != nil {
		slog.Error(fmt.Sprintf("Failed to store commitments for event %s, err : %s", event.ID, err.Error()))
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event has started",
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
