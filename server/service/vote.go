package service

type KeyValDatabase interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}
type VotingService struct {
	keyValDB KeyValDatabase
}

func NewVotingService() *VotingService {
	return &VotingService{}
}

func (v *VotingService) AddMemberToEvent() {
	// Get event id
	// Check is owner
	// add member to the event
	// update the spec hash in blockchain
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
	// check nullifier proof (via contract RPC call)
	// submit vote value
}
