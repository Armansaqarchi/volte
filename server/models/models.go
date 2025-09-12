package models

import (
	"fmt"
	"time"
)

// Commitment is just a string with Serialize() function, making it compatible with merkle tree
type Commitment string

func (e Commitment) Serialize() ([]byte, error) {
	return []byte(e), nil
}

type Event struct {
	ID          string         `json:"id"`
	Name        string         `json:"name"`
	Admin       string         `json:"admin"`
	Duration    time.Duration  `json:"duration"`
	StartTime   time.Time      `json:"start_time"`
	VoteOptions []string       `json:"votes"`        // List of possible vote values.
	VoteMembers []string       `json:"vote_members"` // List of members eligible for voting.
	Tally       map[string]int `json:"tally"`        // The overall score for each VoteOption after event ending
	Revoked     bool           `json:"revoked"`      // Whether event has been revoked because of security issues.
}

// CalculateEventHash is used to store hash of the event in blockchain for security purposes.
func (e *Event) CalculateEventHash() []byte {
	return []byte("")
}

func (e *Event) StartEvent() error {
	if e.StartTime.IsZero() {
		e.StartTime = time.Now()
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
