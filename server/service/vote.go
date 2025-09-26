package service

import (
	"bytes"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"volte/backend/chain"
	"volte/backend/crypto/constraintsys"
	"volte/backend/crypto/zkproofs"
	"volte/backend/databases"
	"volte/backend/models"

	"github.com/gin-gonic/gin"
	"github.com/txaty/go-merkletree"
	"go.mongodb.org/mongo-driver/v2/bson"
)

var (
	database             = flag.String("event_database", "events", "Database to use")
	eventCollection      = flag.String("event_collection", "events", "Collection to use")
	commitmentCollection = flag.String("commitment_collection", "commitments", "Collection to use")
)

type VotingService struct {
	mongoClient     *databases.MongoClient
	contractHandler chain.ContractHandler

	volteGroth16      *zkproofs.Groth16
	cipherTextGroth16 *zkproofs.Groth16
	tallyGroth16      *zkproofs.Groth16
}

type VoteZKRequest struct {
}

type VoteRequest struct {
	// Consists of two dot separated encrypted values C1.C2. See https://eprint.iacr.org/2018/930.pdf
	CipherText       string `json:"cipher_text"`
	Nullifier        string `json:"nullifier"`
	EventID          string `json:"event_id"`
	MembershipRootID string `json:"membership_root_id"`
}

func NewVotingService(mongoClient *databases.MongoClient, contractManager *chain.EthereumContractHandler) *VotingService {
	return &VotingService{
		mongoClient:       mongoClient,
		contractHandler:   contractManager,
		volteGroth16:      zkproofs.SetupNewGroth16(constraintsys.NewVolteBLS12377R1CS()),
		cipherTextGroth16: zkproofs.SetupNewGroth16(constraintsys.NewVolteBLS12377R1CS()),
		tallyGroth16:      zkproofs.SetupNewGroth16(constraintsys.NewVolteBLS12377R1CS()),
	}

}

func (v *VotingService) isEventValid(ctx *gin.Context, event *models.Event) bool {
	eventHash, err := v.contractHandler.GetVolteContract().GetEventHash(event.ID)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get event hash, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to get event %d hash", event.ID),
		})
		return false
	}
	fmt.Println(event)
	if !bytes.Equal(eventHash, event.CalculateEventHash()) {
		slog.Warn(fmt.Sprintf(
			"inconsistent event hash between chain and server for event : %d. Expected : %s, got: %s",
			event.ID, event.CalculateEventHash(), eventHash),
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "inconsistent event hash between chain and server for event!",
		})
		return false
	}
	return true
}

func (v *VotingService) AddMemberToEvent(ctx *gin.Context) {
	// Add security step
	eventsCollection := v.mongoClient.GetClient().Database(*database).Collection(*eventCollection)
	eventId := ctx.Param("event_id")

	var event models.Event
	if err := eventsCollection.FindOne(ctx, bson.M{"_id": eventId}).Decode(&event); err != nil {
		slog.Error(fmt.Sprintf("Failed to get event by event_id, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("No such event %d found", eventId),
		})
	}

	if !v.isEventValid(ctx, &event) {
		return
	}
	if event.Revoked {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Event has been revoked due to security problems.",
		})
		return
	}
	commitment := ctx.Param("commitment")
	if commitment == "" {
		slog.Error("Commitment must not be empty!.")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Identity secret must be provided!"})
	}
	event.VoteMembers = append(event.VoteMembers, commitment)

	if _, err := eventsCollection.UpdateOne(
		ctx, bson.M{"_id": eventId}, bson.D{{"$set", bson.D{{"vote_members", event.VoteMembers}}}},
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to perform operation on event, err : %s", err.Error()),
		})
	}
	_, err := v.contractHandler.GetVolteContract().SetEventHash(eventId, event.CalculateEventHash())
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to set event hash, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to set event hash on chain"),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Member successfully added."})
}

func (v *VotingService) StartEvent(ctx *gin.Context) {
	eventsCollection := v.mongoClient.GetClient().Database(*database).Collection(*eventCollection)
	commitmentsCollection := v.mongoClient.GetClient().Database(*database).Collection(*commitmentCollection)
	eventID := ctx.Param("event_id")

	var event models.Event
	if err := eventsCollection.FindOne(ctx, bson.M{"_id": eventID}).Decode(&event); err != nil {
		slog.Error(fmt.Sprintf("Failed to get event %d, err : %s", eventID, err.Error()))
		ctx.JSON(
			http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to fetch event %d", eventID)},
		)
		return
	}
	if len(event.VoteMembers) <= 1 {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Total eligible voters must be greater than 1.",
		})
	}
	if err := event.StartEvent(); err != nil {
		slog.Error(fmt.Sprintf("Failed to start event, err : %s", err.Error()))
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Failed to start event, err : %d", event.ID)},
		)
		return
	}
	var commitments []merkletree.DataBlock
	for _, member := range event.VoteMembers {
		commitments = append(commitments, models.Commitment(member))
	}
	if _, err := eventsCollection.UpdateOne(ctx, bson.M{"_id": event.ID}, bson.M{"$set": event}); err != nil {
		slog.Error(fmt.Sprintf("Failed to store event %d, err : %s", event.ID, err.Error()))
	}
	commitmentsTree, err := merkletree.New(&merkletree.Config{Mode: merkletree.ModeProofGenAndTreeBuild}, commitments)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to create commitments tree, err : %s", err.Error()))
		ctx.JSON(
			http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to create commitments tree")},
		)
		return
	}
	eventCommitmentTree := &models.EventTree{ID: eventID, Tree: commitmentsTree}

	if _, err := commitmentsCollection.InsertOne(ctx, eventCommitmentTree); err != nil {
		slog.Error(fmt.Sprintf("Failed to store commitments for event %s, err : %s", event.ID, err.Error()))
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Failed to store commitments for event %s", event.ID)},
		)
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Event has started"})
}

func (v *VotingService) RemoveMemberFromEvent(ctx *gin.Context) {
	// Add security step
	eventsCollection := v.mongoClient.GetClient().Database(*database).Collection(*eventCollection)
	eventId := ctx.Param("event_id")

	var event models.Event
	if err := eventsCollection.FindOne(ctx, bson.M{"_id": eventId}).Decode(&event); err != nil {
		slog.Error(fmt.Sprintf("Failed to get event by event_id, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("No such event %s found", eventId),
		})
	}

	if !v.isEventValid(ctx, &event) {
		return
	}
	if event.Revoked {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Event has been revoked due to security problems.",
		})
		return
	}
	commitment := ctx.Param("commitment")
	if commitment == "" {
		slog.Error("Commitment must not be empty!.")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Identity secret must be provided!"})
	}
	foundCommitment := false
	for idx, voteMember := range event.VoteMembers {
		if commitment == voteMember {
			event.VoteMembers = append(event.VoteMembers[:idx], event.VoteMembers[idx+1:]...)
			foundCommitment = true
		}
	}
	if !foundCommitment {
		slog.Info(fmt.Sprintf("Couldn't find commitment %s for event %s", commitment, eventId))
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": fmt.Sprintf("Couldn't find commitment %s for event %s", commitment, eventId),
		})
		return
	}

	if _, err := eventsCollection.UpdateOne(
		ctx, bson.M{"_id": eventId}, bson.D{{"$set", bson.D{{"vote_members", event.VoteMembers}}}},
	); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to perform operation on event, err : %s", err.Error()),
		})
	}
	_, err := v.contractHandler.GetVolteContract().SetEventHash(eventId, event.CalculateEventHash())
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to set event hash, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to set event hash on chain"),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Member successfully added."})
}

func (v *VotingService) Vote(ctx *gin.Context) {
	return
}

//func (v *VotingService) RemoveEvent() {
//	// Check is owner
//	// get event id
//	// remove event from db
//	// remove event's spec hash from chain
//}
//
//func (v *VotingService) Vote() {
//	// check authority
//	// pre-filter invalid votes to reduce gas fee as much as possible
//	// check nullifier proof (via contract RPC call)
//	// submit vote value and update incremental merkle tree
//
//	// Note: use locking to avoid race condition
//}
