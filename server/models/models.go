package models

type Event struct {
	EventID   string `json:"event_id"`
	EventName string `json:"event_name"`
}

// CalculateEventHash is used to store hash of the event in blockchain for security purposes.
func (e Event) CalculateEventHash() string {
	return ""
}

type Nullifier string

func (n Nullifier) Serialize() ([]byte, error) {
	return []byte(n), nil
}

type EventVote string

func (e EventVote) Serialize() ([]byte, error) {
	return []byte(e), nil
}
