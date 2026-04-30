package service

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/txaty/go-merkletree"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"log/slog"
	"math"
	"math/big"
	"net/http"
	"time"
	"volte/backend/chain"
	"volte/backend/crypto/utils"
	"volte/backend/databases"
	"volte/backend/models"
	"volte/backend/models/service/dto"
)

var (
	database                       = flag.String("event_database", "events", "Database to use")
	eventCollection                = flag.String("event_collection", "events", "Collection to use")
	commitmentCollection           = flag.String("commitment_collection", "commitments", "Collection to use")
	commitmentMerklePathCollection = flag.String(
		"commitment_merkle_path_collection", "commitment_merkle_path", "Collection to use",
	)
)

type VotingService struct {
	mongoClient     *databases.MongoClient
	redisClient     *databases.RedisClientProvider
	contractHandler chain.ContractHandler
}

func NewVotingService(
	mongoClient *databases.MongoClient,
	contractManager chain.ContractHandler,
	redis *databases.RedisClientProvider) *VotingService {

	return &VotingService{
		mongoClient:     mongoClient,
		contractHandler: contractManager,
		redisClient:     redis,
	}
}

func (v *VotingService) isEventValid(ctx *gin.Context, event *models.Event) bool {
	eventHash, err := v.contractHandler.GetVolteContract().GetEventHash(event.ID)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get event hash, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to get event %s hash", event.ID),
		})
		return false
	}
	if !bytes.Equal(eventHash, event.CalculateEventHash()) {
		slog.Warn(fmt.Sprintf(
			"inconsistent event hash between chain and server for event : %s. Expected : %s, got: %s",
			event.ID, event.CalculateEventHash(), eventHash),
		)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "inconsistent event hash between chain and server for event!",
		})
		return false
	}
	return true
}

func (v *VotingService) CreateEvent(ctx *gin.Context) {
	session := sessions.Default(ctx)
	adminID := session.Get("user")
	fmt.Println("admin id ", adminID)
	if adminID == nil {
		slog.Warn("Attempt to create event without active session")
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: user not logged in"})
		return
	}

	eventsCollection := v.mongoClient.GetClient().Database(*database).Collection(*eventCollection)

	var req struct {
		Name        string        `json:"name" binding:"required"`
		Duration    time.Duration `json:"duration" binding:"required"`
		VoteOptions []string      `json:"vote_options" binding:"required,min=1"`
		Question    string        `json:"question" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		slog.Error(fmt.Sprintf("Invalid request to create event: %s", err.Error()))
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid input: " + err.Error()})
		return
	}

	// events with more than two options are not implemented yet.
	if len(req.VoteOptions) > 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "votings can only have 2 options for now",
		})
	}

	event := models.Event{
		ID:          uuid.New().String(),
		Name:        req.Name,
		Admin:       adminID.(string),
		Question:    req.Question,
		Duration:    req.Duration,
		StartTime:   nil,
		VoteOptions: req.VoteOptions,
		VoteMembers: []string{},
		Tally:       make(map[string]int),
		Revoked:     false,
	}

	if _, err := eventsCollection.InsertOne(ctx, event); err != nil {
		slog.Error(fmt.Sprintf("Failed to insert event: %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create event"})
		return
	}
	_, err := v.contractHandler.GetVolteContract().SetEventHash(event.ID, event.CalculateEventHash())
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to set event hash on chain, err : %s", err.Error()))

		_, _ = eventsCollection.DeleteOne(ctx, bson.M{"_id": event.ID})
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Event creation failed: could not register hash on blockchain",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event successfully created",
		"event":   event,
	})
}

func (v *VotingService) DeleteEvent(ctx *gin.Context) {
	session := sessions.Default(ctx)
	adminID := session.Get("user")

	if adminID == nil {
		slog.Warn("Attempt to delete event without active session")
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: user not logged in"})
		return
	}
	eventsCollection := v.mongoClient.GetClient().Database(*database).Collection(*eventCollection)
	eventID := ctx.Param("id")

	var event models.Event
	if err := eventsCollection.FindOne(ctx, bson.M{"_id": eventID}).Decode(&event); err != nil {
		slog.Error(fmt.Sprintf("Failed to get event by event_id, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("No such event %s found", event.ID),
		})
	}

	if adminID != event.Admin {
		ctx.JSON(
			http.StatusUnauthorized, gin.H{"message": "Unauthorized: user must be admin to perform this operation."},
		)
		return
	}

	if _, err := eventsCollection.DeleteOne(ctx, &bson.M{"_id": eventID}); err != nil {
		slog.Error(fmt.Sprintf("Failed to delete event: %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete event"})
		return
	}
	_, err := v.contractHandler.GetVolteContract().SetEventHash(eventID, []byte(""))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to set event hash on chain, err : %s", err.Error()))

		_, _ = eventsCollection.DeleteOne(ctx, bson.M{"_id": eventID})
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Event creation failed: could not register hash on blockchain",
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Event successfully deleted.",
		"id":      eventID,
	})
}

func (v *VotingService) AddMemberToEvent(ctx *gin.Context) {
	// Add security step
	eventsCollection := v.mongoClient.GetClient().Database(*database).Collection(*eventCollection)
	usersCollection := v.mongoClient.GetClient().Database(*database).Collection(*usersCollection)
	eventId := ctx.Param("id")

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
		return
	}

	if count, err := usersCollection.CountDocuments(ctx, bson.M{"_id": commitment}); err != nil {
		slog.Error(fmt.Sprintf("Failed to count users on commitment, err : %s", err.Error()))
	} else if count == 0 {
		slog.Error("User not found.")
		ctx.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
		return
	}

	if commitment == event.Admin {
		slog.Error("commitment belongs to the event admin")
		ctx.JSON(http.StatusAlreadyReported, gin.H{"message": "User is already a member of the event"})
		return
	}

	for _, member := range event.VoteMembers {
		if member == commitment {
			slog.Error("user is already a member of the event!")
			ctx.JSON(http.StatusAlreadyReported, gin.H{"message": "User is already a member of the event"})
			return
		}
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
	commitmentMerklePathCollection := v.mongoClient.GetClient().Database(
		*database).Collection(*commitmentMerklePathCollection)
	eventID := ctx.Param("id")
	if eventID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event_id"})
		return
	}
	var event models.Event
	if err := eventsCollection.FindOne(ctx, bson.M{"_id": eventID}).Decode(&event); err != nil {
		slog.Error(fmt.Sprintf("Failed to get event %s, err : %s", eventID, err.Error()))
		ctx.JSON(
			http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to fetch event %s", eventID)},
		)
		return
	}
	session := sessions.Default(ctx)
	if session.Get("user") != event.Admin {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: user not logged in"})
		return
	}

	if len(event.VoteMembers) < 1 {
		ctx.JSON(http.StatusForbidden, gin.H{
			"message": "Total eligible voters must be greater than 1.",
		})
		return
	}
	if err := event.StartEvent(); err != nil {
		fmt.Println(err)
		slog.Error(fmt.Sprintf("Failed to start event, err : %s", err.Error()))
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Failed to start event, err : %s", event.ID)},
		)
		return
	}
	var commitments []merkletree.DataBlock
	fmt.Println("admin : ", event.Admin)
	commitments = append(commitments, models.Commitment(event.Admin))
	for _, member := range event.VoteMembers {
		commitments = append(commitments, models.Commitment(member))
	}
	// define sha256(0) representing empty nodes that fills total members
	// up to 2^8 to match the setup arguments.

	emptyNodeVal, _ := utils.MimcHash([]byte("0"))
	for i := 0; i < int(math.Pow(2, 8))-len(commitments); i++ {
		commitments = append(commitments, models.Commitment(emptyNodeVal))
	}

	commitmentsTree, err := merkletree.New(
		&merkletree.Config{Mode: merkletree.ModeProofGenAndTreeBuild, HashFunc: func(args ...[]byte) ([]byte, error) {
			// For gnarks internal MimC hash, use `utils.MimcHash`
			var argsBigInt []*big.Int
			for _, arg := range args {
				argString, ok := big.NewInt(0).SetString(string(arg), 10)
				if !ok {
					return nil, fmt.Errorf("failed to create merkle tree")
				}
				argsBigInt = append(argsBigInt, argString)
			}
			hash, err := utils.MiMC7MultiHashCircomFr(argsBigInt)
			return []byte(hash.String()), err
		}, DisableLeafHashing: true,
		},
		commitments,
	)

	if err != nil {
		if errors.Is(err, merkletree.ErrInvalidNumOfDataBlocks) {
			slog.Error(fmt.Sprintf("Failed to create commitments tree, err : %s", err.Error()))
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Members must be at least 2"})
			return
		}
		slog.Error(fmt.Sprintf("Failed to create commitments tree, err : %s", err.Error()))
		ctx.JSON(
			http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Failed to create commitments tree")},
		)
		return
	}
	fmt.Println("commitments : ", commitments)
	fmt.Println("root : ", string(commitmentsTree.Root))

	eventCommitmentTree := &models.EventTree{ID: eventID, Tree: commitmentsTree}
	if _, err := eventsCollection.UpdateOne(ctx, bson.M{"_id": event.ID}, bson.M{"$set": event}); err != nil {
		slog.Error(fmt.Sprintf("Failed to store event %s, err : %s", event.ID, err.Error()))
	}
	if _, err := commitmentsCollection.InsertOne(ctx, eventCommitmentTree); err != nil {
		slog.Error(fmt.Sprintf("Failed to store commitments for event %s, err : %s", event.ID, err.Error()))
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"message": fmt.Sprintf("Failed to store commitments for event %s", event.ID)},
		)
		return
	}

	if _, err := commitmentMerklePathCollection.InsertOne(ctx, &models.EventTreeProofsDto{
		ID:      eventID,
		LeafMap: eventCommitmentTree.Tree.LeafMap,
		Proofs:  eventCommitmentTree.Tree.Proofs,
		Root:    eventCommitmentTree.Tree.Root,
	}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to store commitments for event %s", event.ID),
		})
		return
	}
	_, err = v.contractHandler.GetVolteContract().SetEventHash(event.ID, event.CalculateEventHash())
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to set event hash, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to set event hash on chain"),
		})
		return
	}

	root, ok := big.NewInt(0).SetString(string(eventCommitmentTree.Tree.Root), 10)
	if !ok {
		slog.Error("Failed to Set root string, err")
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to set event hash on root"),
		})
		return
	}
	fmt.Println("Root : ", root)
	if _, err := v.contractHandler.GetVolteContract().SetVoteMerkleRoot(eventID, root); err != nil {
		slog.Error(fmt.Sprintf("Failed to set vote merkle root, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to set vote merkle root"),
		})
	}
	slog.Info(fmt.Sprintf("Successfully updated merkleRoot for event %s", event.ID))
	ctx.JSON(http.StatusOK, gin.H{"message": "Event has started", "start": event.StartTime})
}

func (v *VotingService) UserEvents(ctx *gin.Context) {
	session := sessions.Default(ctx)
	commitment, ok := session.Get("user").(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to parse user session name"})
		return
	}
	eventsCollection := v.mongoClient.GetClient().Database(*database).Collection(*eventCollection)
	if cur, err := eventsCollection.Find(ctx, bson.M{
		"$or": []bson.M{
			{"admin": commitment},
			{"vote_members": commitment},
		}}); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to fetch commitments for user %s", commitment),
		})
		return
	} else {
		var events []models.Event
		if err := cur.All(ctx, &events); err != nil {
			slog.Error(fmt.Sprintf("Failed to get events, err : %s", err.Error()))
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("failed to decode events, err : %s", err.Error()),
			})
			return
		}
		fmt.Println(events)
		ctx.JSON(http.StatusOK, gin.H{
			"data": events,
		})
	}
}

func (v *VotingService) MembershipDetails(ctx *gin.Context) {
	session := sessions.Default(ctx)
	member := models.Commitment(session.Get("user").(string))
	eventId := ctx.Param("id")
	if eventId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event_id"})
		return
	}
	var eventTree models.EventTreeProofsDto
	commitmentsPathCollection := v.mongoClient.GetClient().Database(*database).Collection(*commitmentMerklePathCollection)
	if err := commitmentsPathCollection.FindOne(ctx, bson.M{"_id": eventId}).Decode(&eventTree); err != nil {
		fmt.Println(err)
		slog.Error(fmt.Sprintf("Couldnt find event %s", eventId))
	}

	commitmentIdx := eventTree.LeafMap[string(member)]
	ctx.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"root":  eventTree.Root,
			"proof": eventTree.Proofs[commitmentIdx],
		},
	})
}

func (v *VotingService) UserEvent(ctx *gin.Context) {
	session := sessions.Default(ctx)
	userId := session.Get("user").(string)
	eventId := ctx.Param("event_id")
	eventsCollection := v.mongoClient.GetClient().Database(*database).Collection(*eventCollection)
	if res := eventsCollection.FindOne(ctx, bson.M{
		"_id": eventId,
		"$or": []bson.M{{"admin": userId}, {"vote_members": userId}},
	},
	); res.Err() != nil {
		slog.Error(fmt.Sprintf("Failed to fetch event, err : %s", res.Err()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Failed to fetch commitments for user %s", userId),
		})
		return
	} else {
		var event models.Event
		if err := res.Decode(&event); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Failed to decode event, err : %s", err.Error()),
			})
			return
		} else {
			ctx.JSON(http.StatusOK, event)
		}
	}
}

func (v *VotingService) EndEvent(ctx *gin.Context) {
	session := sessions.Default(ctx)
	adminID := session.Get("user")

	if adminID == nil {
		slog.Warn("Attempt to delete event without active session")
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized: user not logged in"})
		return
	}
	eventsCollection := v.mongoClient.GetClient().Database(*database).Collection(*eventCollection)
	eventID := ctx.Param("id")

	var event models.Event
	if err := eventsCollection.FindOneAndUpdate(ctx, bson.M{"_id": eventID, "admin": adminID},
		bson.M{"$set": bson.M{"force_end": true}}).Decode(&event); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			ctx.JSON(http.StatusForbidden, gin.H{
				"error": "you dont have access to this event or the event does not exist",
			})
			return
		}
		slog.Error(fmt.Sprintf("Failed to get event by event_id, err : %s", err))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("No such event %s found", event.ID),
		})
		return
	}

	_, err := v.contractHandler.GetVolteContract().SetEventHash(eventID, event.CalculateEventHash())
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to set event hash on chain, err : %s", err.Error()))

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "Event creation failed: could not register hash on blockchain",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Event successfully ended.",
		"id":      eventID,
	})
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

func (v *VotingService) parseVolteContractVoteSubmission() {

}

func (v *VotingService) Vote(ctx *gin.Context) {
	var proofsDto dto.VolteContractVoteSubmissionDTO
	if err := ctx.ShouldBindJSON(&proofsDto); err != nil {
		slog.Error(fmt.Sprintf("Failed to bind json, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status":  "failure",
			"message": fmt.Sprintf("Internal server error. err : %s", err),
		})
		return
	}
	proofs := dto.ConvertVolteContractVoteSubmissionDTO(proofsDto)
	fmt.Println(fmt.Sprintf("event:%s:voter:%s", proofs.EventID, proofs.Proofs.Nullifier.Input[2].String()))
	// Third element of nullifier input represents the actual nullifier hash
	status := v.redisClient.RedisClient.SetArgs(ctx, fmt.Sprintf(
		"event:%s:voter:%s",
		proofs.EventID, proofs.Proofs.Nullifier.Input[2].String()),
		1, redis.SetArgs{Mode: string(redis.NX)})
	// Nullifier already exists
	if errors.Is(status.Err(), redis.Nil) {
		ctx.JSON(http.StatusAlreadyReported, gin.H{
			"status":  "failure",
			"message": "nullifier hash already exists",
		})
		return
	} else if status.Err() != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "failure"})
		return
	}
	slog.Info(fmt.Sprintf("proof : %v", proofs))
	fmt.Println(proofs.Proofs)
	if txn, err := v.contractHandler.GetVolteContract().Vote(*proofs); err != nil {
		slog.Error(fmt.Sprintf("Failed to verify vote, err : %s", err.Error()))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": "failure"})
	} else {
		slog.Info(fmt.Sprintln("GAS: ", txn.Gas()))
		slog.Info(fmt.Sprintln("GAS PRICE: ", txn.GasPrice()))
		slog.Info(fmt.Sprintln("txn data length: ", len(txn.Data())))
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "Accepted"})
}

func (v *VotingService) GetTallyScore(ctx *gin.Context) {
	eventID := ctx.Param("id")
	tallyScore, err := v.contractHandler.GetTallyScore(eventID)
	if err != nil {
		slog.Error("Failed to get tally score, err : %w", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"status": "failure",
		})
		return
	}
	slog.Info("Got status : ", tallyScore)
	ctx.JSON(http.StatusOK, gin.H{
		"score": tallyScore.Score,
		"total": tallyScore.Total,
	})
}
