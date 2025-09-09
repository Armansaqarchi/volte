package chain

import (
	"context"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log/slog"
	"volte/backend/chain/contracts"
)

var (
	walletPrivateKey = flag.String(
		"wallet_private_key",
		"",
		"Wallet private key for signing write transactions and read operations.",
	)
	chainRpcNodeUrl = flag.String(
		"chain_rpc_node_url", "", "Url of RPC node on blockchain used to submit transactions.",
	)
	contractAddress = flag.String("contract_address", "", "Contract address.")
)

type EthereumChainHandler struct {
	client      *ethclient.Client // Ethereum client for RPC communication.
	fromAddress common.Address    // Server wallet address.
	// List of contracts.
	volteContract *contracts.ContractSession
}

func NewEthereumChainHandler() *EthereumChainHandler {
	client, err := ethclient.Dial(*chainRpcNodeUrl)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to Dial chain rpc node. err : %s", err))
		panic(err)
	}
	// Load private key.
	privateKey, err := crypto.HexToECDSA(*walletPrivateKey)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to load private key. err : %s", err))
		panic(err)
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	volteContract, err := contracts.NewContract(common.HexToAddress(*contractAddress), client)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to create contract. err : %s", err))
		panic(err)
	}

	return &EthereumChainHandler{
		client:      client,
		fromAddress: fromAddress,
		volteContract: &contracts.ContractSession{
			Contract: volteContract,
			CallOpts: bind.CallOpts{
				From:    fromAddress,
				Pending: false,
				Context: context.Background(),
			},
		},
	}
}

func (e *EthereumChainHandler) GetVolteContract() *contracts.ContractSession {
	return e.volteContract
}
