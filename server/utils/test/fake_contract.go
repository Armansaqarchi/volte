package test

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"volte/backend/chain"
)

type FakeVolteSession struct {
	nullifierMerkleRoots map[string] /* eventID */ []byte /* NullifierRootHash */
	voteMerkleRoots      map[string] /* eventID */ []byte /* VoteRootHash */
	eventHashes          map[string] /* eventID */ []byte /* EventDetailsHash */
	usedNullifiers       map[string] /* eventID */ []byte /* UsedNullifiers */
}

func NewFakeVolteContract() chain.VolteSessionHandler {
	return &FakeVolteSession{
		nullifierMerkleRoots: make(map[string][]byte, 10),
		voteMerkleRoots:      make(map[string][]byte, 10),
		eventHashes:          make(map[string][]byte, 10),
		usedNullifiers:       make(map[string][]byte, 10),
	}
}

func (v *FakeVolteSession) SetNullifierMerkleRoot(eventID string, value []byte) (*types.Transaction, error) {
	v.nullifierMerkleRoots[eventID] = value
	return nil, nil
}
func (v *FakeVolteSession) SetVoteMerkleRoot(eventID string, value []byte) (*types.Transaction, error) {
	v.voteMerkleRoots[eventID] = value
	return nil, nil
}
func (v *FakeVolteSession) SetEventHash(eventID string, value []byte) (*types.Transaction, error) {
	v.eventHashes[eventID] = value
	return nil, nil
}
func (v *FakeVolteSession) SetUsedNullifiers(eventID string, value []byte) (*types.Transaction, error) {
	v.usedNullifiers[eventID] = value
	return nil, nil
}
func (v *FakeVolteSession) GetNullifierMerkleRoot(eventID string) ([]byte, error) {
	return v.nullifierMerkleRoots[eventID], nil
}
func (v *FakeVolteSession) GetVoteMerkleRoot(eventID string) ([]byte, error) {
	return v.voteMerkleRoots[eventID], nil
}
func (v *FakeVolteSession) GetEventHash(eventID string) ([]byte, error) {
	return v.eventHashes[eventID], nil
}
func (v *FakeVolteSession) GetUsedNullifiers(eventID string) ([]byte, error) {
	return v.usedNullifiers[eventID], nil
}

type FakeContractManager struct {
	contractHandler chain.VolteSessionHandler
}

func NewFakeContractManager() chain.ContractHandler {
	return &FakeContractManager{
		contractHandler: NewFakeVolteContract(),
	}
}

func (cm *FakeContractManager) GetClient() *ethclient.Client {
	return nil
}
func (cm *FakeContractManager) GetFromAddress() common.Address {
	return common.Address{}
}
func (cm *FakeContractManager) GetVolteContract() chain.VolteSessionHandler {
	return cm.contractHandler
}
