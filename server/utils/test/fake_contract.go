package test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-iden3-crypto/v2/babyjub"
	"log/slog"
	"math"
	"math/big"
	"volte/backend/chain"
	"volte/backend/chain/contracts"
	"volte/backend/crypto/utils"
)

type FakeContractHandler struct {
	contractHandler chain.VolteSessionHandler
	fromAddress     common.Address
	tallyPrivateKey *big.Int
}

func NewFakeContractHandler() chain.ContractHandler {
	fakeChain := GetFakeChain()
	key := &big.Int{}
	key, ok := key.SetString(*chain.ElgamalPrivateKey, 10)
	if !ok {
		panic("Failed to parse elgamal private key.")
	}
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
		fromAddress:     fakeChain.FromAddr,
		tallyPrivateKey: key,
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
func (cm *FakeContractHandler) GetTallyScore(eventID string) (*chain.TallyScore, error) {
	scores, err := cm.GetVolteContract().GetTallyScore(eventID)
	if err != nil {
		slog.Error("Failed to get tally score")
		return nil, fmt.Errorf("failed to get tally score, %w", err)
	}

	votes, err := cm.contractHandler.GetTotalEventVotes(eventID)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get total votes, err : %s", err.Error()))
		return nil, fmt.Errorf("failed to get total votes, %w", err)
	}

	// Contract returns BabyJubJub points (C1, C2) as 4 field elements.
	// scores[0]=C1.x, scores[1]=C1.y, scores[2]=C2.x, scores[3]=C2.y
	slog.Info(
		scores[0].String(),
		scores[1].String(),
		scores[2].String(),
		scores[3].String(),
	)
	C1 := babyjub.NewPoint()
	C2 := babyjub.NewPoint()
	C1.X = chain.NewIntFromString(scores[0].String())
	C1.Y = chain.NewIntFromString(scores[1].String())
	C2.X = chain.NewIntFromString(scores[2].String())
	C2.Y = chain.NewIntFromString(scores[3].String())
	score, err := utils.DecryptM_BSGS(C1, C2, chain.NewIntFromString(*chain.ElgamalPrivateKey), uint64(math.Pow(2, 32)))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to decrypt tally score, err : %s", err.Error()))
		return nil, fmt.Errorf("failed to decrypt tally score, %w", err)
	}

	return &chain.TallyScore{
		Score: score,
		Total: votes.Uint64(),
	}, nil
}
