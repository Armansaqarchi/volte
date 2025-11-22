package test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"time"
	"volte/backend/chain/contracts"
	"volte/backend/chain/contracts/ballot"
	"volte/backend/chain/contracts/membership"
	"volte/backend/chain/contracts/nullifier"

	"log/slog"
	"math/big"
)

type FakeChain struct {
	ContractAddr  common.Address
	FromAddr      common.Address
	Auth          *bind.TransactOpts
	Backend       *simulated.Backend
	VolteContract *contracts.Volte
	Transaction   *types.Transaction
}

var Chain *FakeChain

func GetFakeChain() *FakeChain {
	if Chain != nil {
		return Chain
	}
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	addr := crypto.PubkeyToAddress(privateKey.PublicKey)
	slog.Info("Deployer: ", addr.Hex())

	genesisAlloc := types.GenesisAlloc{
		addr: {Balance: big.NewInt(0).Mul(big.NewInt(1000), big.NewInt(1e18))},
	}
	backend := simulated.NewBackend(genesisAlloc)
	chainID, err := backend.Client().ChainID(context.Background())
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to gain chainID, err : %s", err.Error()))
		panic(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	ballotContractAddr, _, _, err := ballot.DeployVolte(auth, backend.Client())
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to gain ballot, err : %s", err.Error()))
	}
	nullifierContractAddr, _, _, err := nullifier.DeployVolte(auth, backend.Client())
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to gain ballot, err : %s", err.Error()))
	}
	membershipContractAddr, _, _, err := membership.DeployVolte(auth, backend.Client())
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to gain ballot, err : %s", err.Error()))
	}
	contractAddr, tx, volteContract, err := contracts.DeployVolte(
		auth, backend.Client(), ballotContractAddr, membershipContractAddr, nullifierContractAddr,
	)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to deploy contract, err : %s", err.Error()))
		panic(err)
	}

	// Auto commitment of transactions.
	// This way, we won't need to commit on every Transaction done in test mode.
	go func() {
		for {
			time.Sleep(1 * time.Second)
			backend.Commit()
		}
	}()
	slog.Info(fmt.Sprintf("Successfully deployed contract at address %s", contractAddr))
	backend.Commit()

	Chain = &FakeChain{
		ContractAddr:  contractAddr,
		FromAddr:      addr,
		Auth:          auth,
		Backend:       backend,
		VolteContract: volteContract,
		Transaction:   tx,
	}
	return Chain
}
