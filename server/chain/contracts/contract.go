// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// VolteContractBallotProof is an auto generated low-level Go binding around an user-defined struct.
type VolteContractBallotProof struct {
	Proof          VolteContractProof
	Input          [16]*big.Int
	CommitmentX    *big.Int
	CommitmentY    *big.Int
	CommitmentPokX *big.Int
	CommitmentPokY *big.Int
}

// VolteContractMembershipProof is an auto generated low-level Go binding around an user-defined struct.
type VolteContractMembershipProof struct {
	Proof VolteContractProof
	Input [2]*big.Int
}

// VolteContractNullifierProof is an auto generated low-level Go binding around an user-defined struct.
type VolteContractNullifierProof struct {
	Proof VolteContractProof
	Input [2]*big.Int
}

// VolteContractProof is an auto generated low-level Go binding around an user-defined struct.
type VolteContractProof struct {
	Arx  *big.Int
	Ary  *big.Int
	Brx0 *big.Int
	Brx1 *big.Int
	Bry0 *big.Int
	Bry1 *big.Int
	Cx   *big.Int
	Cy   *big.Int
}

// VolteContractProofs is an auto generated low-level Go binding around an user-defined struct.
type VolteContractProofs struct {
	Ballot     VolteContractBallotProof
	Membership VolteContractMembershipProof
	Nullifier  VolteContractNullifierProof
}

// VolteContractVoteSubmission is an auto generated low-level Go binding around an user-defined struct.
type VolteContractVoteSubmission struct {
	Sender  common.Address
	EventID string
	Proofs  VolteContractProofs
}

// VolteMetaData contains all meta data concerning the Volte contract.
var VolteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ballot\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_membership\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nullifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetEventHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetVoteMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetEventHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetVoteMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"Arx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cy\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.Proof\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[16]\",\"name\":\"Input\",\"type\":\"uint256[16]\"},{\"internalType\":\"uint256\",\"name\":\"CommitmentX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommitmentY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommitmentPokX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommitmentPokY\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.BallotProof\",\"name\":\"ballot\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"Arx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cy\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.Proof\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[2]\",\"name\":\"Input\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVolteContract.MembershipProof\",\"name\":\"membership\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"Arx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cy\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.Proof\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[2]\",\"name\":\"Input\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVolteContract.NullifierProof\",\"name\":\"nullifier\",\"type\":\"tuple\"}],\"internalType\":\"structVolteContract.Proofs\",\"name\":\"proofs\",\"type\":\"tuple\"}],\"internalType\":\"structVolteContract.VoteSubmission\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"C1\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"C2\",\"type\":\"uint256[2]\"}],\"name\":\"addCiphertexts\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"Csum\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballot\",\"outputs\":[{\"internalType\":\"contractBallotVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"eventHashes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"membership\",\"outputs\":[{\"internalType\":\"contractMembershipVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"membershipMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nullifier\",\"outputs\":[{\"internalType\":\"contractNullifierVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tallyScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"C1x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"C1y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"C2x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"C2y\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"voteMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051611804380380611804833981810160405281019061003191906101e2565b335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060065f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060405180608001604052805f81526020015f81526020015f81526020015f81525060015f820151815f0155602082015181600101556040820151816002015560608201518160030155905050505050610232565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101b182610188565b9050919050565b6101c1816101a7565b81146101cb575f5ffd5b50565b5f815190506101dc816101b8565b92915050565b5f5f5f606084860312156101f9576101f8610184565b5b5f610206868287016101ce565b9350506020610217868287016101ce565b9250506040610228868287016101ce565b9150509250925092565b6115c58061023f5f395ff3fe608060405234801561000f575f5ffd5b50600436106100e8575f3560e01c8063649a75ca1161008a578063ac3910a211610064578063ac3910a21461025d578063d59f76611461027b578063f851a440146102ab578063fdab8f23146102c9576100e8565b8063649a75ca146101f35780639b28fdf01461020f578063a2bf68b91461023f576100e8565b806343bf6cf0116100c657806343bf6cf01461016857806348a31121146101865780634c4ff130146101a75780635dd5ed28146101c3576100e8565b806337f5daab146100ec5780633aa82fee1461011c57806341260b3d14610138575b5f5ffd5b61010660048036038101906101019190610a32565b6102f9565b6040516101139190610aed565b60405180910390f35b61013660048036038101906101319190610b62565b6103aa565b005b610152600480360381019061014d9190610d08565b6103e0565b60405161015f9190610aed565b60405180910390f35b610170610493565b60405161017d9190610dc9565b60405180910390f35b61018e6104b8565b60405161019e9493929190610dfa565b60405180910390f35b6101c160048036038101906101bc9190610e60565b6104d5565b005b6101dd60048036038101906101d89190610a32565b61064d565b6040516101ea9190610aed565b60405180910390f35b61020d60048036038101906102089190610b62565b6106fe565b005b61022960048036038101906102249190610d08565b610734565b6040516102369190610aed565b60405180910390f35b6102476107e7565b6040516102549190610ec7565b60405180910390f35b61026561080c565b6040516102729190610f00565b60405180910390f35b61029560048036038101906102909190610d08565b610831565b6040516102a29190610aed565b60405180910390f35b6102b36108e4565b6040516102c09190610f39565b60405180910390f35b6102e360048036038101906102de919061102a565b610908565b6040516102f0919061110d565b60405180910390f35b60606009838360405161030d929190611154565b9081526020016040518091039020805461032690611199565b80601f016020809104026020016040519081016040528092919081815260200182805461035290611199565b801561039d5780601f106103745761010080835404028352916020019161039d565b820191905f5260205f20905b81548152906001019060200180831161038057829003601f168201915b5050505050905092915050565b8181600a86866040516103be929190611154565b908152602001604051809103902091826103d992919061136a565b5050505050565b600a818051602081018201805184825260208301602085012081835280955050505050505f91509050805461041490611199565b80601f016020809104026020016040519081016040528092919081815260200182805461044090611199565b801561048b5780601f106104625761010080835404028352916020019161048b565b820191905f5260205f20905b81548152906001019060200180831161046e57829003601f168201915b505050505081565b60065f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001805f0154908060010154908060020154908060030154905084565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dc6395e6604051806101000160405280846040015f015f015f01358152602001846040015f015f01602001358152602001846040015f015f01606001358152602001846040015f015f01604001358152602001846040015f015f0160a001358152602001846040015f015f01608001358152602001846040015f015f0160c001358152602001846040015f015f0160e001358152506040518060400160405280856040015f0161030001358152602001856040015f0161032001358152506040518060400160405280866040015f0161034001358152602001866040015f016103600135815250856040015f01610100016040518563ffffffff1660e01b815260040161061e94939291906114d0565b5f6040518083038186803b158015610634575f5ffd5b505afa158015610646573d5f5f3e3d5ffd5b5050505050565b6060600a8383604051610661929190611154565b9081526020016040518091039020805461067a90611199565b80601f01602080910402602001604051908101604052809291908181526020018280546106a690611199565b80156106f15780601f106106c8576101008083540402835291602001916106f1565b820191905f5260205f20905b8154815290600101906020018083116106d457829003601f168201915b5050505050905092915050565b818160098686604051610712929190611154565b9081526020016040518091039020918261072d92919061136a565b5050505050565b6009818051602081018201805184825260208301602085012081835280955050505050505f91509050805461076890611199565b80601f016020809104026020016040519081016040528092919081815260200182805461079490611199565b80156107df5780601f106107b6576101008083540402835291602001916107df565b820191905f5260205f20905b8154815290600101906020018083116107c257829003601f168201915b505050505081565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6008818051602081018201805184825260208301602085012081835280955050505050505f91509050805461086590611199565b80601f016020809104026020016040519081016040528092919081815260200182805461089190611199565b80156108dc5780601f106108b3576101008083540402835291602001916108dc565b820191905f5260205f20905b8154815290600101906020018083116108bf57829003601f168201915b505050505081565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b61091061099e565b5f6040518451815260208501516020820152835160408201526020840151606082015260408160808360065afa915082815181526020820151602082015260808201604052505080610997576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098e90611571565b60405180910390fd5b5092915050565b6040518060400160405280600290602082028036833780820191505090505090565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126109f2576109f16109d1565b5b8235905067ffffffffffffffff811115610a0f57610a0e6109d5565b5b602083019150836001820283011115610a2b57610a2a6109d9565b5b9250929050565b5f5f60208385031215610a4857610a476109c9565b5b5f83013567ffffffffffffffff811115610a6557610a646109cd565b5b610a71858286016109dd565b92509250509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610abf82610a7d565b610ac98185610a87565b9350610ad9818560208601610a97565b610ae281610aa5565b840191505092915050565b5f6020820190508181035f830152610b058184610ab5565b905092915050565b5f5f83601f840112610b2257610b216109d1565b5b8235905067ffffffffffffffff811115610b3f57610b3e6109d5565b5b602083019150836001820283011115610b5b57610b5a6109d9565b5b9250929050565b5f5f5f5f60408587031215610b7a57610b796109c9565b5b5f85013567ffffffffffffffff811115610b9757610b966109cd565b5b610ba3878288016109dd565b9450945050602085013567ffffffffffffffff811115610bc657610bc56109cd565b5b610bd287828801610b0d565b925092505092959194509250565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610c1a82610aa5565b810181811067ffffffffffffffff82111715610c3957610c38610be4565b5b80604052505050565b5f610c4b6109c0565b9050610c578282610c11565b919050565b5f67ffffffffffffffff821115610c7657610c75610be4565b5b610c7f82610aa5565b9050602081019050919050565b828183375f83830152505050565b5f610cac610ca784610c5c565b610c42565b905082815260208101848484011115610cc857610cc7610be0565b5b610cd3848285610c8c565b509392505050565b5f82601f830112610cef57610cee6109d1565b5b8135610cff848260208601610c9a565b91505092915050565b5f60208284031215610d1d57610d1c6109c9565b5b5f82013567ffffffffffffffff811115610d3a57610d396109cd565b5b610d4684828501610cdb565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610d91610d8c610d8784610d4f565b610d6e565b610d4f565b9050919050565b5f610da282610d77565b9050919050565b5f610db382610d98565b9050919050565b610dc381610da9565b82525050565b5f602082019050610ddc5f830184610dba565b92915050565b5f819050919050565b610df481610de2565b82525050565b5f608082019050610e0d5f830187610deb565b610e1a6020830186610deb565b610e276040830185610deb565b610e346060830184610deb565b95945050505050565b5f5ffd5b5f6106408284031215610e5757610e56610e3d565b5b81905092915050565b5f60208284031215610e7557610e746109c9565b5b5f82013567ffffffffffffffff811115610e9257610e916109cd565b5b610e9e84828501610e41565b91505092915050565b5f610eb182610d98565b9050919050565b610ec181610ea7565b82525050565b5f602082019050610eda5f830184610eb8565b92915050565b5f610eea82610d98565b9050919050565b610efa81610ee0565b82525050565b5f602082019050610f135f830184610ef1565b92915050565b5f610f2382610d4f565b9050919050565b610f3381610f19565b82525050565b5f602082019050610f4c5f830184610f2a565b92915050565b5f67ffffffffffffffff821115610f6c57610f6b610be4565b5b602082029050919050565b610f8081610de2565b8114610f8a575f5ffd5b50565b5f81359050610f9b81610f77565b92915050565b5f610fb3610fae84610f52565b610c42565b90508060208402830185811115610fcd57610fcc6109d9565b5b835b81811015610ff65780610fe28882610f8d565b845260208401935050602081019050610fcf565b5050509392505050565b5f82601f830112611014576110136109d1565b5b6002611021848285610fa1565b91505092915050565b5f5f608083850312156110405761103f6109c9565b5b5f61104d85828601611000565b925050604061105e85828601611000565b9150509250929050565b5f60029050919050565b5f81905092915050565b5f819050919050565b61108e81610de2565b82525050565b5f61109f8383611085565b60208301905092915050565b5f602082019050919050565b6110c081611068565b6110ca8184611072565b92506110d58261107c565b805f5b838110156111055781516110ec8782611094565b96506110f7836110ab565b9250506001810190506110d8565b505050505050565b5f6040820190506111205f8301846110b7565b92915050565b5f81905092915050565b5f61113b8385611126565b9350611148838584610c8c565b82840190509392505050565b5f611160828486611130565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806111b057607f821691505b6020821081036111c3576111c261116c565b5b50919050565b5f82905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261122f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826111f4565b61123986836111f4565b95508019841693508086168417925050509392505050565b5f61126b61126661126184610de2565b610d6e565b610de2565b9050919050565b5f819050919050565b61128483611251565b61129861129082611272565b848454611200565b825550505050565b5f5f905090565b6112af6112a0565b6112ba81848461127b565b505050565b5b818110156112dd576112d25f826112a7565b6001810190506112c0565b5050565b601f821115611322576112f3816111d3565b6112fc846111e5565b8101602085101561130b578190505b61131f611317856111e5565b8301826112bf565b50505b505050565b5f82821c905092915050565b5f6113425f1984600802611327565b1980831691505092915050565b5f61135a8383611333565b9150826002028217905092915050565b61137483836111c9565b67ffffffffffffffff81111561138d5761138c610be4565b5b6113978254611199565b6113a28282856112e1565b5f601f8311600181146113cf575f84156113bd578287013590505b6113c7858261134f565b86555061142e565b601f1984166113dd866111d3565b5f5b82811015611404578489013582556001820191506020850194506020810190506113df565b86831015611421578489013561141d601f891682611333565b8355505b6001600288020188555050505b50505050505050565b5f60089050919050565b5f81905092915050565b5f819050919050565b5f602082019050919050565b61146981611437565b6114738184611441565b925061147e8261144b565b805f5b838110156114ae5781516114958782611094565b96506114a083611454565b925050600181019050611481565b505050505050565b82818337505050565b6114cc61020083836114b6565b5050565b5f610380820190506114e45f830187611460565b6114f26101008301866110b7565b6115006101408301856110b7565b61150e6101808301846114bf565b95945050505050565b5f82825260208201905092915050565b7f4543414444206661696c656400000000000000000000000000000000000000005f82015250565b5f61155b600c83611517565b915061156682611527565b602082019050919050565b5f6020820190508181035f8301526115888161154f565b905091905056fea26469706673582212208254b8befebdcc18a16d9a4a968d9fd12469fff4129ecb31e1760f0f5ffffe7864736f6c634300081e0033",
}

// VolteABI is the input ABI used to generate the binding from.
// Deprecated: Use VolteMetaData.ABI instead.
var VolteABI = VolteMetaData.ABI

// VolteBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VolteMetaData.Bin instead.
var VolteBin = VolteMetaData.Bin

// DeployVolte deploys a new Ethereum contract, binding an instance of Volte to it.
func DeployVolte(auth *bind.TransactOpts, backend bind.ContractBackend, _ballot common.Address, _membership common.Address, _nullifier common.Address) (common.Address, *types.Transaction, *Volte, error) {
	parsed, err := VolteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VolteBin), backend, _ballot, _membership, _nullifier)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Volte{VolteCaller: VolteCaller{contract: contract}, VolteTransactor: VolteTransactor{contract: contract}, VolteFilterer: VolteFilterer{contract: contract}}, nil
}

// Volte is an auto generated Go binding around an Ethereum contract.
type Volte struct {
	VolteCaller     // Read-only binding to the contract
	VolteTransactor // Write-only binding to the contract
	VolteFilterer   // Log filterer for contract events
}

// VolteCaller is an auto generated read-only Go binding around an Ethereum contract.
type VolteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VolteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VolteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VolteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VolteSession struct {
	Contract     *Volte            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VolteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VolteCallerSession struct {
	Contract *VolteCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VolteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VolteTransactorSession struct {
	Contract     *VolteTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VolteRaw is an auto generated low-level Go binding around an Ethereum contract.
type VolteRaw struct {
	Contract *Volte // Generic contract binding to access the raw methods on
}

// VolteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VolteCallerRaw struct {
	Contract *VolteCaller // Generic read-only contract binding to access the raw methods on
}

// VolteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VolteTransactorRaw struct {
	Contract *VolteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVolte creates a new instance of Volte, bound to a specific deployed contract.
func NewVolte(address common.Address, backend bind.ContractBackend) (*Volte, error) {
	contract, err := bindVolte(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Volte{VolteCaller: VolteCaller{contract: contract}, VolteTransactor: VolteTransactor{contract: contract}, VolteFilterer: VolteFilterer{contract: contract}}, nil
}

// NewVolteCaller creates a new read-only instance of Volte, bound to a specific deployed contract.
func NewVolteCaller(address common.Address, caller bind.ContractCaller) (*VolteCaller, error) {
	contract, err := bindVolte(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VolteCaller{contract: contract}, nil
}

// NewVolteTransactor creates a new write-only instance of Volte, bound to a specific deployed contract.
func NewVolteTransactor(address common.Address, transactor bind.ContractTransactor) (*VolteTransactor, error) {
	contract, err := bindVolte(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VolteTransactor{contract: contract}, nil
}

// NewVolteFilterer creates a new log filterer instance of Volte, bound to a specific deployed contract.
func NewVolteFilterer(address common.Address, filterer bind.ContractFilterer) (*VolteFilterer, error) {
	contract, err := bindVolte(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VolteFilterer{contract: contract}, nil
}

// bindVolte binds a generic wrapper to an already deployed contract.
func bindVolte(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VolteMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Volte *VolteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Volte.Contract.VolteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Volte *VolteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Volte.Contract.VolteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Volte *VolteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Volte.Contract.VolteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Volte *VolteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Volte.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Volte *VolteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Volte.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Volte *VolteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Volte.Contract.contract.Transact(opts, method, params...)
}

// GetEventHash is a free data retrieval call binding the contract method 0x5dd5ed28.
//
// Solidity: function GetEventHash(string eventID) view returns(bytes)
func (_Volte *VolteCaller) GetEventHash(opts *bind.CallOpts, eventID string) ([]byte, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "GetEventHash", eventID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEventHash is a free data retrieval call binding the contract method 0x5dd5ed28.
//
// Solidity: function GetEventHash(string eventID) view returns(bytes)
func (_Volte *VolteSession) GetEventHash(eventID string) ([]byte, error) {
	return _Volte.Contract.GetEventHash(&_Volte.CallOpts, eventID)
}

// GetEventHash is a free data retrieval call binding the contract method 0x5dd5ed28.
//
// Solidity: function GetEventHash(string eventID) view returns(bytes)
func (_Volte *VolteCallerSession) GetEventHash(eventID string) ([]byte, error) {
	return _Volte.Contract.GetEventHash(&_Volte.CallOpts, eventID)
}

// GetVoteMerkleRoot is a free data retrieval call binding the contract method 0x37f5daab.
//
// Solidity: function GetVoteMerkleRoot(string eventID) view returns(bytes)
func (_Volte *VolteCaller) GetVoteMerkleRoot(opts *bind.CallOpts, eventID string) ([]byte, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "GetVoteMerkleRoot", eventID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetVoteMerkleRoot is a free data retrieval call binding the contract method 0x37f5daab.
//
// Solidity: function GetVoteMerkleRoot(string eventID) view returns(bytes)
func (_Volte *VolteSession) GetVoteMerkleRoot(eventID string) ([]byte, error) {
	return _Volte.Contract.GetVoteMerkleRoot(&_Volte.CallOpts, eventID)
}

// GetVoteMerkleRoot is a free data retrieval call binding the contract method 0x37f5daab.
//
// Solidity: function GetVoteMerkleRoot(string eventID) view returns(bytes)
func (_Volte *VolteCallerSession) GetVoteMerkleRoot(eventID string) ([]byte, error) {
	return _Volte.Contract.GetVoteMerkleRoot(&_Volte.CallOpts, eventID)
}

// Vote is a free data retrieval call binding the contract method 0x4c4ff130.
//
// Solidity: function Vote((address,string,(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[16],uint256,uint256,uint256,uint256),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]))) proof) view returns()
func (_Volte *VolteCaller) Vote(opts *bind.CallOpts, proof VolteContractVoteSubmission) error {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "Vote", proof)

	if err != nil {
		return err
	}

	return err

}

// Vote is a free data retrieval call binding the contract method 0x4c4ff130.
//
// Solidity: function Vote((address,string,(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[16],uint256,uint256,uint256,uint256),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]))) proof) view returns()
func (_Volte *VolteSession) Vote(proof VolteContractVoteSubmission) error {
	return _Volte.Contract.Vote(&_Volte.CallOpts, proof)
}

// Vote is a free data retrieval call binding the contract method 0x4c4ff130.
//
// Solidity: function Vote((address,string,(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[16],uint256,uint256,uint256,uint256),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]))) proof) view returns()
func (_Volte *VolteCallerSession) Vote(proof VolteContractVoteSubmission) error {
	return _Volte.Contract.Vote(&_Volte.CallOpts, proof)
}

// AddCiphertexts is a free data retrieval call binding the contract method 0xfdab8f23.
//
// Solidity: function addCiphertexts(uint256[2] C1, uint256[2] C2) view returns(uint256[2] Csum)
func (_Volte *VolteCaller) AddCiphertexts(opts *bind.CallOpts, C1 [2]*big.Int, C2 [2]*big.Int) ([2]*big.Int, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "addCiphertexts", C1, C2)

	if err != nil {
		return *new([2]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([2]*big.Int)).(*[2]*big.Int)

	return out0, err

}

// AddCiphertexts is a free data retrieval call binding the contract method 0xfdab8f23.
//
// Solidity: function addCiphertexts(uint256[2] C1, uint256[2] C2) view returns(uint256[2] Csum)
func (_Volte *VolteSession) AddCiphertexts(C1 [2]*big.Int, C2 [2]*big.Int) ([2]*big.Int, error) {
	return _Volte.Contract.AddCiphertexts(&_Volte.CallOpts, C1, C2)
}

// AddCiphertexts is a free data retrieval call binding the contract method 0xfdab8f23.
//
// Solidity: function addCiphertexts(uint256[2] C1, uint256[2] C2) view returns(uint256[2] Csum)
func (_Volte *VolteCallerSession) AddCiphertexts(C1 [2]*big.Int, C2 [2]*big.Int) ([2]*big.Int, error) {
	return _Volte.Contract.AddCiphertexts(&_Volte.CallOpts, C1, C2)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Volte *VolteCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Volte *VolteSession) Admin() (common.Address, error) {
	return _Volte.Contract.Admin(&_Volte.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Volte *VolteCallerSession) Admin() (common.Address, error) {
	return _Volte.Contract.Admin(&_Volte.CallOpts)
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_Volte *VolteCaller) Ballot(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "ballot")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_Volte *VolteSession) Ballot() (common.Address, error) {
	return _Volte.Contract.Ballot(&_Volte.CallOpts)
}

// Ballot is a free data retrieval call binding the contract method 0xac3910a2.
//
// Solidity: function ballot() view returns(address)
func (_Volte *VolteCallerSession) Ballot() (common.Address, error) {
	return _Volte.Contract.Ballot(&_Volte.CallOpts)
}

// EventHashes is a free data retrieval call binding the contract method 0x41260b3d.
//
// Solidity: function eventHashes(string ) view returns(bytes)
func (_Volte *VolteCaller) EventHashes(opts *bind.CallOpts, arg0 string) ([]byte, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "eventHashes", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EventHashes is a free data retrieval call binding the contract method 0x41260b3d.
//
// Solidity: function eventHashes(string ) view returns(bytes)
func (_Volte *VolteSession) EventHashes(arg0 string) ([]byte, error) {
	return _Volte.Contract.EventHashes(&_Volte.CallOpts, arg0)
}

// EventHashes is a free data retrieval call binding the contract method 0x41260b3d.
//
// Solidity: function eventHashes(string ) view returns(bytes)
func (_Volte *VolteCallerSession) EventHashes(arg0 string) ([]byte, error) {
	return _Volte.Contract.EventHashes(&_Volte.CallOpts, arg0)
}

// Membership is a free data retrieval call binding the contract method 0xa2bf68b9.
//
// Solidity: function membership() view returns(address)
func (_Volte *VolteCaller) Membership(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "membership")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Membership is a free data retrieval call binding the contract method 0xa2bf68b9.
//
// Solidity: function membership() view returns(address)
func (_Volte *VolteSession) Membership() (common.Address, error) {
	return _Volte.Contract.Membership(&_Volte.CallOpts)
}

// Membership is a free data retrieval call binding the contract method 0xa2bf68b9.
//
// Solidity: function membership() view returns(address)
func (_Volte *VolteCallerSession) Membership() (common.Address, error) {
	return _Volte.Contract.Membership(&_Volte.CallOpts)
}

// MembershipMerkleRoots is a free data retrieval call binding the contract method 0xd59f7661.
//
// Solidity: function membershipMerkleRoots(string ) view returns(bytes)
func (_Volte *VolteCaller) MembershipMerkleRoots(opts *bind.CallOpts, arg0 string) ([]byte, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "membershipMerkleRoots", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// MembershipMerkleRoots is a free data retrieval call binding the contract method 0xd59f7661.
//
// Solidity: function membershipMerkleRoots(string ) view returns(bytes)
func (_Volte *VolteSession) MembershipMerkleRoots(arg0 string) ([]byte, error) {
	return _Volte.Contract.MembershipMerkleRoots(&_Volte.CallOpts, arg0)
}

// MembershipMerkleRoots is a free data retrieval call binding the contract method 0xd59f7661.
//
// Solidity: function membershipMerkleRoots(string ) view returns(bytes)
func (_Volte *VolteCallerSession) MembershipMerkleRoots(arg0 string) ([]byte, error) {
	return _Volte.Contract.MembershipMerkleRoots(&_Volte.CallOpts, arg0)
}

// Nullifier is a free data retrieval call binding the contract method 0x43bf6cf0.
//
// Solidity: function nullifier() view returns(address)
func (_Volte *VolteCaller) Nullifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "nullifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nullifier is a free data retrieval call binding the contract method 0x43bf6cf0.
//
// Solidity: function nullifier() view returns(address)
func (_Volte *VolteSession) Nullifier() (common.Address, error) {
	return _Volte.Contract.Nullifier(&_Volte.CallOpts)
}

// Nullifier is a free data retrieval call binding the contract method 0x43bf6cf0.
//
// Solidity: function nullifier() view returns(address)
func (_Volte *VolteCallerSession) Nullifier() (common.Address, error) {
	return _Volte.Contract.Nullifier(&_Volte.CallOpts)
}

// TallyScore is a free data retrieval call binding the contract method 0x48a31121.
//
// Solidity: function tallyScore() view returns(uint256 C1x, uint256 C1y, uint256 C2x, uint256 C2y)
func (_Volte *VolteCaller) TallyScore(opts *bind.CallOpts) (struct {
	C1x *big.Int
	C1y *big.Int
	C2x *big.Int
	C2y *big.Int
}, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "tallyScore")

	outstruct := new(struct {
		C1x *big.Int
		C1y *big.Int
		C2x *big.Int
		C2y *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.C1x = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.C1y = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.C2x = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.C2y = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TallyScore is a free data retrieval call binding the contract method 0x48a31121.
//
// Solidity: function tallyScore() view returns(uint256 C1x, uint256 C1y, uint256 C2x, uint256 C2y)
func (_Volte *VolteSession) TallyScore() (struct {
	C1x *big.Int
	C1y *big.Int
	C2x *big.Int
	C2y *big.Int
}, error) {
	return _Volte.Contract.TallyScore(&_Volte.CallOpts)
}

// TallyScore is a free data retrieval call binding the contract method 0x48a31121.
//
// Solidity: function tallyScore() view returns(uint256 C1x, uint256 C1y, uint256 C2x, uint256 C2y)
func (_Volte *VolteCallerSession) TallyScore() (struct {
	C1x *big.Int
	C1y *big.Int
	C2x *big.Int
	C2y *big.Int
}, error) {
	return _Volte.Contract.TallyScore(&_Volte.CallOpts)
}

// VoteMerkleRoots is a free data retrieval call binding the contract method 0x9b28fdf0.
//
// Solidity: function voteMerkleRoots(string ) view returns(bytes)
func (_Volte *VolteCaller) VoteMerkleRoots(opts *bind.CallOpts, arg0 string) ([]byte, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "voteMerkleRoots", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// VoteMerkleRoots is a free data retrieval call binding the contract method 0x9b28fdf0.
//
// Solidity: function voteMerkleRoots(string ) view returns(bytes)
func (_Volte *VolteSession) VoteMerkleRoots(arg0 string) ([]byte, error) {
	return _Volte.Contract.VoteMerkleRoots(&_Volte.CallOpts, arg0)
}

// VoteMerkleRoots is a free data retrieval call binding the contract method 0x9b28fdf0.
//
// Solidity: function voteMerkleRoots(string ) view returns(bytes)
func (_Volte *VolteCallerSession) VoteMerkleRoots(arg0 string) ([]byte, error) {
	return _Volte.Contract.VoteMerkleRoots(&_Volte.CallOpts, arg0)
}

// SetEventHash is a paid mutator transaction binding the contract method 0x3aa82fee.
//
// Solidity: function SetEventHash(string eventID, bytes value) returns()
func (_Volte *VolteTransactor) SetEventHash(opts *bind.TransactOpts, eventID string, value []byte) (*types.Transaction, error) {
	return _Volte.contract.Transact(opts, "SetEventHash", eventID, value)
}

// SetEventHash is a paid mutator transaction binding the contract method 0x3aa82fee.
//
// Solidity: function SetEventHash(string eventID, bytes value) returns()
func (_Volte *VolteSession) SetEventHash(eventID string, value []byte) (*types.Transaction, error) {
	return _Volte.Contract.SetEventHash(&_Volte.TransactOpts, eventID, value)
}

// SetEventHash is a paid mutator transaction binding the contract method 0x3aa82fee.
//
// Solidity: function SetEventHash(string eventID, bytes value) returns()
func (_Volte *VolteTransactorSession) SetEventHash(eventID string, value []byte) (*types.Transaction, error) {
	return _Volte.Contract.SetEventHash(&_Volte.TransactOpts, eventID, value)
}

// SetVoteMerkleRoot is a paid mutator transaction binding the contract method 0x649a75ca.
//
// Solidity: function SetVoteMerkleRoot(string eventID, bytes value) returns()
func (_Volte *VolteTransactor) SetVoteMerkleRoot(opts *bind.TransactOpts, eventID string, value []byte) (*types.Transaction, error) {
	return _Volte.contract.Transact(opts, "SetVoteMerkleRoot", eventID, value)
}

// SetVoteMerkleRoot is a paid mutator transaction binding the contract method 0x649a75ca.
//
// Solidity: function SetVoteMerkleRoot(string eventID, bytes value) returns()
func (_Volte *VolteSession) SetVoteMerkleRoot(eventID string, value []byte) (*types.Transaction, error) {
	return _Volte.Contract.SetVoteMerkleRoot(&_Volte.TransactOpts, eventID, value)
}

// SetVoteMerkleRoot is a paid mutator transaction binding the contract method 0x649a75ca.
//
// Solidity: function SetVoteMerkleRoot(string eventID, bytes value) returns()
func (_Volte *VolteTransactorSession) SetVoteMerkleRoot(eventID string, value []byte) (*types.Transaction, error) {
	return _Volte.Contract.SetVoteMerkleRoot(&_Volte.TransactOpts, eventID, value)
}
