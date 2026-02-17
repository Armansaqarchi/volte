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
	"github.com/iden3/go-iden3-crypto/v2/babyjub"
	"log/slog"
	"math"
	"math/big"
	"strconv"
	"volte/backend/chain/contracts"
	"volte/backend/crypto/utils"
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
	// ChainID, unique for each ethereum network, default set to sepolia network.
	chainID           = flag.String("eth_based_network_chain_id", "11155111", "Chain Commitment")
	ElgamalPrivateKey = flag.String("elgamal_private_key", "20", "elgamal encryption private key")
)

func NewIntFromString(str string) *big.Int {
	key := big.Int{}
	key.SetString(str, 10)
	return &key
}

type VolteSessionHandler interface {
	Vote(proof contracts.VolteContractVoteSubmission) (*types.Transaction, error)
	GetTallyScore(eventID string) ([4]*big.Int, error)
	SetVoteMerkleRoot(eventID string, value *big.Int) (*types.Transaction, error)
	SetEventHash(eventID string, value []byte) (*types.Transaction, error)
	GetVoteMerkleRoot(eventID string) (*big.Int, error)
	GetTotalEventVotes(eventID string) (*big.Int, error)
	GetEventHash(eventID string) ([]byte, error)
}

type ContractHandler interface {
	GetClient() *ethclient.Client
	GetFromAddress() common.Address
	GetVolteContract() VolteSessionHandler
	GetTallyScore(eventID string) (*TallyScore, error)
}

type TallyScore struct {
	Score uint64
	Total uint64
}

type EthereumContractHandler struct {
	client      *ethclient.Client // Ethereum client for RPC communication.
	fromAddress common.Address    // Server wallet address.
	// List of contracts.
	volteContract   VolteSessionHandler
	tallyPrivateKey *big.Int
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
	// Load elgamal encryption private key
	tallyPrivateKey := &big.Int{}
	tallyPrivateKey, ok := tallyPrivateKey.SetString(*ElgamalPrivateKey, 10)
	if !ok {
		panic("Failed to parse elgamal private key")
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
		tallyPrivateKey: tallyPrivateKey,
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

func (e *EthereumContractHandler) GetTallyScore(eventID string) (*TallyScore, error) {
	scores, err := e.GetVolteContract().GetTallyScore(eventID)
	if err != nil {
		slog.Error("Failed to get tally score")
		return nil, fmt.Errorf("failed to get tally score, %w", err)
	}

	votes, err := e.volteContract.GetTotalEventVotes(eventID)
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
	C1.X = NewIntFromString(scores[0].String())
	C1.Y = NewIntFromString(scores[1].String())
	C2.X = NewIntFromString(scores[2].String())
	C2.Y = NewIntFromString(scores[3].String())
	score, err := utils.DecryptM_BSGS(C1, C2, NewIntFromString(*ElgamalPrivateKey), uint64(math.Pow(2, 32)))
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to decrypt tally score, err : %s", err.Error()))
		return nil, fmt.Errorf("failed to decrypt tally score, %w", err)
	}

	return &TallyScore{
		Score: score,
		Total: votes.Uint64(),
	}, nil
}
