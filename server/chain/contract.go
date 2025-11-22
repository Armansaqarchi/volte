package chain

import (
	"context"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log/slog"
	"math/big"
	"strconv"
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

type Proof struct {
	Arx  []byte `json:"arx"`
	Ary  []byte `json:"ary"`
	Brx0 []byte `json:"brx0"`
	Brx1 []byte `json:"brx1"`
	Bry0 []byte `json:"bry0"`
	Bry1 []byte `json:"bry1"`
	Cx   []byte `json:"cx"`
	Cy   []byte `json:"cy"`
}

type BallotProof struct {
	Proof          Proof
	Input          [16][]byte
	CommitmentX    []byte
	CommitmentY    []byte
	CommitmentPokX []byte
	CommitmentPokY []byte
}

type MembershipProof struct {
	Proof Proof
	Input [2][]byte
}

type NullifierProof struct {
	Proof Proof
	Input [2][]byte
}

var (
	chainID = flag.String("eth_based_network_chain_id", "11155111", "Chain Commitment")
)

type VolteSessionHandler interface {
	Vote(proof contracts.VolteContractVoteSubmission) (*types.Transaction, error)
	GetTallyScore(eventID *big.Int) ([4]*big.Int, error)
	SetVoteMerkleRoot(eventID string, value []byte) (*types.Transaction, error)
	SetEventHash(eventID string, value []byte) (*types.Transaction, error)
	GetVoteMerkleRoot(eventID string) ([]byte, error)
	GetEventHash(eventID string) ([]byte, error)
}

type ContractHandler interface {
	GetClient() *ethclient.Client
	GetFromAddress() common.Address
	GetVolteContract() VolteSessionHandler
}

type EthereumContractHandler struct {
	client      *ethclient.Client // Ethereum client for RPC communication.
	fromAddress common.Address    // Server wallet address.
	// List of contracts.
	volteContract VolteSessionHandler
}

func NewEthereumChainHandler() *EthereumContractHandler {
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
	ethChainID, err := strconv.ParseInt(*chainID, 10, 64)
	if err != nil {
		panic(err)
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	transactionOpsAuth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(ethChainID))
	if err != nil {
		slog.Error("Something went wrong while creating auth transaction sign.")
		panic(err)
	}
	volteContract, err := contracts.NewVolte(common.HexToAddress(*contractAddress), client)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to create contract. err : %s", err))
		panic(err)
	}

	return &EthereumContractHandler{
		client:      client,
		fromAddress: fromAddress,
		volteContract: &contracts.VolteSession{
			Contract: volteContract,
			CallOpts: bind.CallOpts{
				From:    fromAddress,
				Pending: true,
				Context: context.Background(),
			},
			TransactOpts: *transactionOpsAuth,
		},
	}
}

func (e *EthereumContractHandler) GetVolteContract() VolteSessionHandler {
	return e.volteContract
}

func (e *EthereumContractHandler) GetClient() *ethclient.Client {
	return e.client
}

func (e *EthereumContractHandler) GetFromAddress() common.Address {
	return e.fromAddress
}

func (e *EthereumContractHandler) VerifyAndSubmitVote(ballotProof *BallotProof,
	membershipProof *MembershipProof, proof *NullifierProof) error {
	return nil
}
