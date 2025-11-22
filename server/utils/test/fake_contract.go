package test

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"volte/backend/chain"
	"volte/backend/chain/contracts"
)

type FakeContractHandler struct {
	contractHandler chain.VolteSessionHandler
	fromAddress     common.Address
}

func NewFakeContractHandler() chain.ContractHandler {
	fakeChain := GetFakeChain()
	return &FakeContractHandler{
		contractHandler: &contracts.VolteSession{
			Contract: fakeChain.VolteContract,
			CallOpts: bind.CallOpts{
				From: fakeChain.FromAddr,
				// Don't wait for the transactions to mine, this is intended only for testing comfortability
				Pending: true,
				Context: context.Background(),
			},
			TransactOpts: *fakeChain.Auth,
		},
		fromAddress: fakeChain.FromAddr,
	}
}

func (cm *FakeContractHandler) GetClient() *ethclient.Client {
	return nil
}
func (cm *FakeContractHandler) GetFromAddress() common.Address {
	return cm.fromAddress
}
func (cm *FakeContractHandler) GetVolteContract() chain.VolteSessionHandler {
	return cm.contractHandler
}
