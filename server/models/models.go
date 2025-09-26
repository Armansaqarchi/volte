package models

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/txaty/go-merkletree"
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

// Commitment is just a string with Serialize() function, making it compatible with merkle tree
type Commitment string

func (e Commitment) Serialize() ([]byte, error) {
	return []byte(e), nil
}

type Event struct {
	ID          string         `json:"-" bson:"_id"` // auto-generated id for mongo
	Name        string         `json:"name" bson:"name"`
	Admin       string         `json:"admin" bson:"admin"` // Owner of this event.
	Duration    time.Duration  `json:"duration" bson:"duration"`
	StartTime   time.Time      `json:"start_time" bson:"start_time"`
	VoteOptions []string       `json:"votes" bson:"vote_options"`        // List of possible vote values.
	VoteMembers []string       `json:"vote_members" bson:"vote_members"` // List of members eligible for voting.
	Tally       map[string]int `json:"tally" bson:"tally"`               // The overall score for each VoteOption after event ending.
	Revoked     bool           `json:"revoked" bson:"revoked"`           // Whether event has been revoked because of security issues.
}

type EventTree struct {
	ID   string                 `json:"id" bson:"_id"`
	Tree *merkletree.MerkleTree `json:"tree" bson:"tree"`
}

// UnmarshalBSON implements the bson.Unmarshaler interface for EventTree.
func (e *EventTree) UnmarshalBSON(data []byte) error {
	// Define a temporary alias type to break the recursion.
	type eventTreeAlias EventTree
	var alias eventTreeAlias

	// This will not cause recursion because eventTreeAlias does not have a custom unmarshaler.
	if err := bson.Unmarshal(data, &alias); err != nil {
		return err
	}

	// Copy the data from the alias back to the original struct.
	*e = EventTree(alias)

	if e.Tree != nil && e.Tree.Config != nil {
		if e.Tree.Config.RunInParallel {
			e.Tree.HashFunc = merkletree.DefaultHashFuncParallel
		} else {
			e.Tree.HashFunc = merkletree.DefaultHashFunc
		}
	}

	return nil
}

// CalculateEventHash is used to store hash of the event in blockchain for security purposes.
func (e *Event) CalculateEventHash() []byte {
	jsonData, _ := json.Marshal(map[string]any{
		"ID":          e.ID,
		"Name":        e.Name,
		"Admin":       e.Admin,
		"Duration":    e.Duration,
		"StartTime":   e.StartTime.UTC().Truncate(time.Millisecond),
		"Votes":       e.VoteOptions,
		"VoteMembers": e.VoteMembers,
		"Tally":       e.Tally,
		"Revoked":     e.Revoked,
	})
	h := sha256.New()
	h.Write(jsonData)
	return h.Sum(nil)
}

func (e *Event) MarshalBinary() ([]byte, error) {
	jsonData, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (e *Event) StartEvent() error {
	if e.StartTime.IsZero() {
		e.StartTime = time.Now()
		return nil
	}
	return fmt.Errorf(fmt.Sprintf("Event %d has already started!", e.ID))
}

func (e *Event) EventEnded() bool {
	return time.Now().After(e.StartTime.Add(e.Duration))
}

// Nullifier is just a string with Serialize() function, making it compatible with incremental merkle tree
type Nullifier string

func (n Nullifier) Serialize() ([]byte, error) {
	return []byte(n), nil
}
