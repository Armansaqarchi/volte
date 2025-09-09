package service

import (
	"volte/backend/chain"
	"volte/backend/crypto/zkproofs"
)

type keyValDatabase interface {
	Get(key string) ([]byte, error)
	Set(key string, value []byte) error
}

type VotingService struct {
	keyValDB        keyValDatabase
	contractManager *chain.EthereumChainHandler

	volteGroth16      *zkproofs.Groth16
	CipherTextGroth16 *zkproofs.Groth16
	TallyGroth16      *zkproofs.Groth16
}

func NewVotingService() *VotingService {
	// Initialize a KV DB
	// Initialize ethereum contract client
	// fetch Groth16 specs from redis
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
	// pre-filter invalid votes to reduce gas fee as much as possible
	// check nullifier proof (via contract RPC call)
	// submit vote value and update incremental merkle tree

	// Note: use locking to avoid race condition
}
