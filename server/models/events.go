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

func (c Commitment) Serialize() ([]byte, error) {
	return []byte(c), nil
}

func (c Commitment) ToLeafValue() ([]byte, error) {
	// We assume that leaves are derived by applying merkletree default hash function on commitments.
	memberByte, err := c.Serialize()
	if err != nil {
		return nil, err
	}
	return merkletree.DefaultHashFunc(memberByte)
}

type Event struct {
	ID          string         `json:"id" bson:"_id"` // auto-generated id for mongo
	Name        string         `json:"name" bson:"name"`
	Question    string         `json:"question" bson:"question"`
	Admin       string         `json:"admin" bson:"admin"` // Owner of this event.
	Duration    time.Duration  `json:"duration" bson:"duration"`
	StartTime   *time.Time     `json:"startTime" bson:"start_time"`
	ForceEnd    bool           `json:"forceEnd" bson:"force_end"`
	VoteOptions []string       `json:"voteOptions" bson:"vote_options"` // List of possible vote values.
	VoteMembers []string       `json:"voteMembers" bson:"vote_members"` // List of members eligible for voting.
	Tally       map[string]int `json:"tally" bson:"tally"`              // The overall score for each VoteOption after event ending.
	Revoked     bool           `json:"revoked" bson:"revoked"`          // Whether event has been revoked because of security issues.
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
	var startEventHash string
	if e.StartTime != nil {
		startEventHash = e.StartTime.UTC().Truncate(time.Millisecond).String()
	} else {
		startEventHash = ""
	}
	jsonData, _ := json.Marshal(map[string]any{
		"Commitment":   e.ID,
		"Name":         e.Name,
		"Question":     e.Question,
		"Admin":        e.Admin,
		"Duration":     e.Duration,
		"StartTime":    startEventHash,
		"ForceDeleted": e.ForceEnd,
		"Votes":        e.VoteOptions,
		"VoteMembers":  e.VoteMembers,
		"Tally":        e.Tally,
		"Revoked":      e.Revoked,
	})
	h := sha256.New()
	h.Write(jsonData)
	return h.Sum(nil)
}

// EventTreeProofsDto holds proofs for event members separately.
// On each request, the proofs are fetched efficiently.
type EventTreeProofsDto struct {
	ID      string              `json:"id" bson:"_id"`
	LeafMap map[string]int      `json:"-" bson:"-"`
	Proofs  []*merkletree.Proof `json:"proofs" bson:"proofs"`
	Root    []byte              `json:"root" bson:"root"`
}

func (e *EventTreeProofsDto) MarshalBSON() ([]byte, error) {
	leafMapJSON, err := json.Marshal(e.LeafMap)
	if err != nil {
		return nil, fmt.Errorf("marshal LeafMap: %w", err)
	}
	aux := struct {
		ID          string              `bson:"_id,omitempty"`
		LeafMapJSON []byte              `bson:"leaf_map"`
		Proofs      []*merkletree.Proof `bson:"proofs"`
		Root        []byte              `bson:"root"`
	}{
		ID:          e.ID,
		LeafMapJSON: leafMapJSON,
		Proofs:      e.Proofs,
		Root:        e.Root,
	}
	fmt.Println("root : ", e.Root)

	return bson.Marshal(aux)
}

func (e *EventTreeProofsDto) UnmarshalBSON(data []byte) error {

	var aux struct {
		ID          string              `bson:"_id,omitempty"`
		LeafMapJSON []byte              `bson:"leaf_map"`
		Proofs      []*merkletree.Proof `bson:"proofs"`
		Root        []byte              `bson:"root"`
	}

	if err := bson.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("unmarshal EventTreeProofsDto (bson): %w", err)
	}

	e.ID = aux.ID
	e.Proofs = aux.Proofs
	e.Root = aux.Root

	fmt.Println("root unmarshalling : ", aux.Root)

	// Handle missing / empty leaf_map gracefully
	if len(aux.LeafMapJSON) == 0 {
		e.LeafMap = map[string]int{}
		return nil
	}

	// Decode JSON back into LeafMap
	leafMap := make(map[string]int)
	if err := json.Unmarshal(aux.LeafMapJSON, &leafMap); err != nil {
		return fmt.Errorf("unmarshal LeafMap (json): %w", err)
	}

	e.LeafMap = leafMap
	return nil
}

func (e *Event) MarshalBinary() ([]byte, error) {
	jsonData, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func (e *Event) StartEvent() error {
	if e.StartTime == nil || e.StartTime.IsZero() {
		nowTime := time.Now()
		e.StartTime = &nowTime
		return nil
	}
	return fmt.Errorf(fmt.Sprintf("Event %s has already started!", e.ID))
}

func (e *Event) EventEnded() bool {
	return time.Now().After(e.StartTime.Add(e.Duration))
}

// Nullifier is just a string with Serialize() function, making it compatible with incremental merkle tree
type Nullifier string

func (n Nullifier) Serialize() ([]byte, error) {
	return []byte(n), nil
}
