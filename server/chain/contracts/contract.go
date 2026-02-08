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
	Input [3]*big.Int
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
	EventID string
	Proofs  VolteContractProofs
}

// VolteMetaData contains all meta data concerning the Volte contract.
var VolteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ballot\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_membership\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nullifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gotRoot\",\"type\":\"uint256\"}],\"name\":\"EventRootMismatch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetEventHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetTallyScore\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetTotalEventVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetVoteMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetEventHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SetVoteMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"Arx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cy\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.Proof\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[16]\",\"name\":\"Input\",\"type\":\"uint256[16]\"},{\"internalType\":\"uint256\",\"name\":\"CommitmentX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommitmentY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommitmentPokX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommitmentPokY\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.BallotProof\",\"name\":\"ballot\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"Arx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cy\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.Proof\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[2]\",\"name\":\"Input\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVolteContract.MembershipProof\",\"name\":\"membership\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"Arx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cy\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.Proof\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[3]\",\"name\":\"Input\",\"type\":\"uint256[3]\"}],\"internalType\":\"structVolteContract.NullifierProof\",\"name\":\"nullifier\",\"type\":\"tuple\"}],\"internalType\":\"structVolteContract.Proofs\",\"name\":\"proofs\",\"type\":\"tuple\"}],\"internalType\":\"structVolteContract.VoteSubmission\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"C1\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"C2\",\"type\":\"uint256[2]\"}],\"name\":\"addCiphertexts\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"Csum\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballot\",\"outputs\":[{\"internalType\":\"contractBallotVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"eventHashes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"membership\",\"outputs\":[{\"internalType\":\"contractMembershipVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nullifier\",\"outputs\":[{\"internalType\":\"contractNullifierVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tallyScores\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"C1x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"C1y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"C2x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"C2y\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"totalEventVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"voteMerkleRoots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161248838038061248883398181016040528101906100319190610196565b335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506101e6565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101658261013c565b9050919050565b6101758161015b565b811461017f575f5ffd5b50565b5f815190506101908161016c565b92915050565b5f5f5f606084860312156101ad576101ac610138565b5b5f6101ba86828701610182565b93505060206101cb86828701610182565b92505060406101dc86828701610182565b9150509250925092565b612295806101f35f395ff3fe608060405234801561000f575f5ffd5b50600436106100fe575f3560e01c80638b441a3f11610095578063ac3910a211610064578063ac3910a2146102e5578063f018026c14610303578063f851a44014610333578063fdab8f2314610351576100fe565b80638b441a3f146102485780639b28fdf014610264578063a2bf68b914610294578063a68965c8146102b2576100fe565b806341260b3d116100d157806341260b3d146101ae57806343bf6cf0146101de5780635dd5ed28146101fc5780637c277c9e1461022c576100fe565b80630d9a38251461010257806312b3f7b31461013257806337f5daab146101625780633aa82fee14610192575b5f5ffd5b61011c600480360381019061011791906114be565b610381565b60405161012991906115b3565b60405180910390f35b61014c600480360381019061014791906114be565b6103e2565b60405161015991906115db565b60405180910390f35b61017c60048036038101906101779190611651565b610409565b60405161018991906115db565b60405180910390f35b6101ac60048036038101906101a791906116f1565b610433565b005b6101c860048036038101906101c391906114be565b610469565b6040516101d591906117cf565b60405180910390f35b6101e661051c565b6040516101f39190611869565b60405180910390f35b61021660048036038101906102119190611651565b610541565b60405161022391906117cf565b60405180910390f35b610246600480360381019061024191906118a5565b6105f2565b005b610262600480360381019061025d9190611916565b61108b565b005b61027e600480360381019061027991906114be565b6110b5565b60405161028b91906115db565b60405180910390f35b61029c6110e2565b6040516102a99190611993565b60405180910390f35b6102cc60048036038101906102c791906114be565b611107565b6040516102dc94939291906119ac565b60405180910390f35b6102ed61114b565b6040516102fa9190611a0f565b60405180910390f35b61031d600480360381019061031891906114be565b611170565b60405161032a91906115db565b60405180910390f35b61033b61119d565b6040516103489190611a48565b60405180910390f35b61036b60048036038101906103669190611b0f565b6111c1565b6040516103789190611bcc565b60405180910390f35b61038961132d565b5f60068360405161039a9190611c29565b908152602001604051809103902090506040518060800160405280825f0154815260200182600101548152602001826002015481526020018260030154815250915050919050565b5f6007826040516103f39190611c29565b9081526020016040518091039020549050919050565b5f6004838360405161041c929190611c63565b908152602001604051809103902054905092915050565b818160058686604051610447929190611c63565b90815260200160405180910390209182610462929190611e79565b5050505050565b6005818051602081018201805184825260208301602085012081835280955050505050505f91509050805461049d90611cb2565b80601f01602080910402602001604051908101604052809291908181526020018280546104c990611cb2565b80156105145780601f106104eb57610100808354040283529160200191610514565b820191905f5260205f20905b8154815290600101906020018083116104f757829003601f168201915b505050505081565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060058383604051610555929190611c63565b9081526020016040518091039020805461056e90611cb2565b80601f016020809104026020016040519081016040528092919081815260200182805461059a90611cb2565b80156105e55780601f106105bc576101008083540402835291602001916105e5565b820191905f5260205f20905b8154815290600101906020018083116105c857829003601f168201915b5050505050905092915050565b600481805f01906106039190611f52565b604051610611929190611c63565b9081526020016040518091039020548160200161038001610100015f6002811061063e5761063d611fb4565b5b6020020135146106d657600481805f01906106599190611f52565b604051610667929190611c63565b9081526020016040518091039020548160200161038001610100015f6002811061069457610693611fb4565b5b60200201356040517f1f0189e20000000000000000000000000000000000000000000000000000000081526004016106cd929190611fe1565b60405180910390fd5b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166365c03259604051806101000160405280846020016104c0015f015f01358152602001846020016104c0015f01602001358152602001846020016104c0015f01606001358152602001846020016104c0015f01604001358152602001846020016104c0015f0160a001358152602001846020016104c0015f01608001358152602001846020016104c0015f0160c001358152602001846020016104c0015f0160e00135815250836020016104c001610100016040518363ffffffff1660e01b81526004016107dd9291906120a0565b5f6040518083038186803b1580156107f3575f5ffd5b505afa158015610805573d5f5f3e3d5ffd5b5050505060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635fe24f2360405180610100016040528084602001610380015f015f0135815260200184602001610380015f0160200135815260200184602001610380015f0160600135815260200184602001610380015f0160400135815260200184602001610380015f0160a00135815260200184602001610380015f0160800135815260200184602001610380015f0160c00135815260200184602001610380015f0160e001358152508360200161038001610100016040518363ffffffff1660e01b81526004016109109291906120d9565b5f6040518083038186803b158015610926575f5ffd5b505afa158015610938573d5f5f3e3d5ffd5b5050505060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dc6395e6604051806101000160405280846020015f015f015f01358152602001846020015f015f01602001358152602001846020015f015f01606001358152602001846020015f015f01604001358152602001846020015f015f0160a001358152602001846020015f015f01608001358152602001846020015f015f0160c001358152602001846020015f015f0160e001358152506040518060400160405280856020015f0161030001358152602001856020015f0161032001358152506040518060400160405280866020015f0161034001358152602001866020015f016103600135815250856020015f01610100016040518563ffffffff1660e01b8152600401610a859493929190612113565b5f6040518083038186803b158015610a9b575f5ffd5b505afa158015610aad573d5f5f3e3d5ffd5b505050505f610bae6040518060800160405280846020015f01610100015f60108110610adc57610adb611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001846020015f0161010001600160108110610b1757610b16611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001846020015f0161010001600260108110610b5257610b51611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001846020015f0161010001600360108110610b8d57610b8c611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff16815250611257565b90505f610cae6040518060800160405280856020015f0161010001600460108110610bdc57610bdb611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001856020015f0161010001600560108110610c1757610c16611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001856020015f0161010001600660108110610c5257610c51611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001856020015f0161010001600760108110610c8d57610c8c611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff16815250611257565b90505f610dae6040518060800160405280866020015f0161010001600860108110610cdc57610cdb611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001866020015f0161010001600960108110610d1757610d16611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001866020015f0161010001600a60108110610d5257610d51611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001866020015f0161010001600b60108110610d8d57610d8c611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff16815250611257565b90505f610eae6040518060800160405280876020015f0161010001600c60108110610ddc57610ddb611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001876020015f0161010001600d60108110610e1757610e16611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001876020015f0161010001600e60108110610e5257610e51611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001876020015f0161010001600f60108110610e8d57610e8c611fb4565b5b602002013567ffffffffffffffff1667ffffffffffffffff16815250611257565b90505f85805f0190610ec09190611f52565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505090505f600682604051610f159190611c29565b908152602001604051809103902090505f60405180604001604052808881526020018781525090505f60405180604001604052808781526020018681525090505f6040518060400160405280855f01548152602001856001015481525090505f604051806040016040528086600201548152602001866003015481525090505f610f9f85846111c1565b90505f610fac85846111c1565b9050815f60028110610fc157610fc0611fb4565b5b6020020151875f018190555081600160028110610fe157610fe0611fb4565b5b60200201518760010181905550805f6002811061100157611000611fb4565b5b602002015187600201819055508060016002811061102257611021611fb4565b5b6020020151876003018190555060016007896040516110419190611c29565b90815260200160405180910390205461105a9190612187565b60078960405161106a9190611c29565b90815260200160405180910390208190555050505050505050505050505050565b806004848460405161109e929190611c63565b908152602001604051809103902081905550505050565b6004818051602081018201805184825260208301602085012081835280955050505050505f915090505481565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6006818051602081018201805184825260208301602085012081835280955050505050505f91509050805f0154908060010154908060020154908060030154905084565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6007818051602081018201805184825260208301602085012081835280955050505050505f915090505481565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6111c961134f565b5f6040518451815260208501516020820152835160408201526020840151606082015260408160808360065afa915082815181526020820151602082015260808201604052505080611250576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161124790612214565b60405180910390fd5b5092915050565b5f60c08260036004811061126e5761126d611fb4565b5b602002015167ffffffffffffffff16901b60808360026004811061129557611294611fb4565b5b602002015167ffffffffffffffff16901b6040846001600481106112bc576112bb611fb4565b5b602002015167ffffffffffffffff16901b845f600481106112e0576112df611fb4565b5b602002015167ffffffffffffffff1617171790507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018061132357611322612232565b5b5f82089050919050565b6040518060800160405280600490602082028036833780820191505090505090565b6040518060400160405280600290602082028036833780820191505090505090565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6113d08261138a565b810181811067ffffffffffffffff821117156113ef576113ee61139a565b5b80604052505050565b5f611401611371565b905061140d82826113c7565b919050565b5f67ffffffffffffffff82111561142c5761142b61139a565b5b6114358261138a565b9050602081019050919050565b828183375f83830152505050565b5f61146261145d84611412565b6113f8565b90508281526020810184848401111561147e5761147d611386565b5b611489848285611442565b509392505050565b5f82601f8301126114a5576114a4611382565b5b81356114b5848260208601611450565b91505092915050565b5f602082840312156114d3576114d261137a565b5b5f82013567ffffffffffffffff8111156114f0576114ef61137e565b5b6114fc84828501611491565b91505092915050565b5f60049050919050565b5f81905092915050565b5f819050919050565b5f819050919050565b61153481611522565b82525050565b5f611545838361152b565b60208301905092915050565b5f602082019050919050565b61156681611505565b611570818461150f565b925061157b82611519565b805f5b838110156115ab578151611592878261153a565b965061159d83611551565b92505060018101905061157e565b505050505050565b5f6080820190506115c65f83018461155d565b92915050565b6115d581611522565b82525050565b5f6020820190506115ee5f8301846115cc565b92915050565b5f5ffd5b5f5ffd5b5f5f83601f84011261161157611610611382565b5b8235905067ffffffffffffffff81111561162e5761162d6115f4565b5b60208301915083600182028301111561164a576116496115f8565b5b9250929050565b5f5f602083850312156116675761166661137a565b5b5f83013567ffffffffffffffff8111156116845761168361137e565b5b611690858286016115fc565b92509250509250929050565b5f5f83601f8401126116b1576116b0611382565b5b8235905067ffffffffffffffff8111156116ce576116cd6115f4565b5b6020830191508360018202830111156116ea576116e96115f8565b5b9250929050565b5f5f5f5f604085870312156117095761170861137a565b5b5f85013567ffffffffffffffff8111156117265761172561137e565b5b611732878288016115fc565b9450945050602085013567ffffffffffffffff8111156117555761175461137e565b5b6117618782880161169c565b925092505092959194509250565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6117a18261176f565b6117ab8185611779565b93506117bb818560208601611789565b6117c48161138a565b840191505092915050565b5f6020820190508181035f8301526117e78184611797565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f61183161182c611827846117ef565b61180e565b6117ef565b9050919050565b5f61184282611817565b9050919050565b5f61185382611838565b9050919050565b61186381611849565b82525050565b5f60208201905061187c5f83018461185a565b92915050565b5f5ffd5b5f610640828403121561189c5761189b611882565b5b81905092915050565b5f602082840312156118ba576118b961137a565b5b5f82013567ffffffffffffffff8111156118d7576118d661137e565b5b6118e384828501611886565b91505092915050565b6118f581611522565b81146118ff575f5ffd5b50565b5f81359050611910816118ec565b92915050565b5f5f5f6040848603121561192d5761192c61137a565b5b5f84013567ffffffffffffffff81111561194a5761194961137e565b5b611956868287016115fc565b9350935050602061196986828701611902565b9150509250925092565b5f61197d82611838565b9050919050565b61198d81611973565b82525050565b5f6020820190506119a65f830184611984565b92915050565b5f6080820190506119bf5f8301876115cc565b6119cc60208301866115cc565b6119d960408301856115cc565b6119e660608301846115cc565b95945050505050565b5f6119f982611838565b9050919050565b611a09816119ef565b82525050565b5f602082019050611a225f830184611a00565b92915050565b5f611a32826117ef565b9050919050565b611a4281611a28565b82525050565b5f602082019050611a5b5f830184611a39565b92915050565b5f67ffffffffffffffff821115611a7b57611a7a61139a565b5b602082029050919050565b5f611a98611a9384611a61565b6113f8565b90508060208402830185811115611ab257611ab16115f8565b5b835b81811015611adb5780611ac78882611902565b845260208401935050602081019050611ab4565b5050509392505050565b5f82601f830112611af957611af8611382565b5b6002611b06848285611a86565b91505092915050565b5f5f60808385031215611b2557611b2461137a565b5b5f611b3285828601611ae5565b9250506040611b4385828601611ae5565b9150509250929050565b5f60029050919050565b5f81905092915050565b5f819050919050565b5f602082019050919050565b611b7f81611b4d565b611b898184611b57565b9250611b9482611b61565b805f5b83811015611bc4578151611bab878261153a565b9650611bb683611b6a565b925050600181019050611b97565b505050505050565b5f604082019050611bdf5f830184611b76565b92915050565b5f81519050919050565b5f81905092915050565b5f611c0382611be5565b611c0d8185611bef565b9350611c1d818560208601611789565b80840191505092915050565b5f611c348284611bf9565b915081905092915050565b5f611c4a8385611bef565b9350611c57838584611442565b82840190509392505050565b5f611c6f828486611c3f565b91508190509392505050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611cc957607f821691505b602082108103611cdc57611cdb611c85565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611d3e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611d03565b611d488683611d03565b95508019841693508086168417925050509392505050565b5f611d7a611d75611d7084611522565b61180e565b611522565b9050919050565b5f819050919050565b611d9383611d60565b611da7611d9f82611d81565b848454611d0f565b825550505050565b5f5f905090565b611dbe611daf565b611dc9818484611d8a565b505050565b5b81811015611dec57611de15f82611db6565b600181019050611dcf565b5050565b601f821115611e3157611e0281611ce2565b611e0b84611cf4565b81016020851015611e1a578190505b611e2e611e2685611cf4565b830182611dce565b50505b505050565b5f82821c905092915050565b5f611e515f1984600802611e36565b1980831691505092915050565b5f611e698383611e42565b9150826002028217905092915050565b611e838383611c7b565b67ffffffffffffffff811115611e9c57611e9b61139a565b5b611ea68254611cb2565b611eb1828285611df0565b5f601f831160018114611ede575f8415611ecc578287013590505b611ed68582611e5e565b865550611f3d565b601f198416611eec86611ce2565b5f5b82811015611f1357848901358255600182019150602085019450602081019050611eee565b86831015611f305784890135611f2c601f891682611e42565b8355505b6001600288020188555050505b50505050505050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83356001602003843603038112611f6e57611f6d611f46565b5b80840192508235915067ffffffffffffffff821115611f9057611f8f611f4a565b5b602083019250600182023603831315611fac57611fab611f4e565b5b509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f604082019050611ff45f8301856115cc565b61200160208301846115cc565b9392505050565b5f60089050919050565b5f81905092915050565b5f819050919050565b5f602082019050919050565b61203a81612008565b6120448184612012565b925061204f8261201c565b805f5b8381101561207f578151612066878261153a565b965061207183612025565b925050600181019050612052565b505050505050565b82818337505050565b61209c60608383612087565b5050565b5f610160820190506120b45f830185612031565b6120c2610100830184612090565b9392505050565b6120d560408383612087565b5050565b5f610140820190506120ed5f830185612031565b6120fb6101008301846120c9565b9392505050565b61210f6102008383612087565b5050565b5f610380820190506121275f830187612031565b612135610100830186611b76565b612143610140830185611b76565b612151610180830184612102565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61219182611522565b915061219c83611522565b92508282019050808211156121b4576121b361215a565b5b92915050565b5f82825260208201905092915050565b7f4543414444206661696c656400000000000000000000000000000000000000005f82015250565b5f6121fe600c836121ba565b9150612209826121ca565b602082019050919050565b5f6020820190508181035f83015261222b816121f2565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffdfea2646970667358221220e35d8ccde94d39b5a7fa0c74fe6f93a1510feef4b21fc8439c5ebd1ef879bccf64736f6c634300081e0033",
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

// GetTallyScore is a free data retrieval call binding the contract method 0x0d9a3825.
//
// Solidity: function GetTallyScore(string eventID) view returns(uint256[4])
func (_Volte *VolteCaller) GetTallyScore(opts *bind.CallOpts, eventID string) ([4]*big.Int, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "GetTallyScore", eventID)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetTallyScore is a free data retrieval call binding the contract method 0x0d9a3825.
//
// Solidity: function GetTallyScore(string eventID) view returns(uint256[4])
func (_Volte *VolteSession) GetTallyScore(eventID string) ([4]*big.Int, error) {
	return _Volte.Contract.GetTallyScore(&_Volte.CallOpts, eventID)
}

// GetTallyScore is a free data retrieval call binding the contract method 0x0d9a3825.
//
// Solidity: function GetTallyScore(string eventID) view returns(uint256[4])
func (_Volte *VolteCallerSession) GetTallyScore(eventID string) ([4]*big.Int, error) {
	return _Volte.Contract.GetTallyScore(&_Volte.CallOpts, eventID)
}

// GetTotalEventVotes is a free data retrieval call binding the contract method 0x12b3f7b3.
//
// Solidity: function GetTotalEventVotes(string eventID) view returns(uint256)
func (_Volte *VolteCaller) GetTotalEventVotes(opts *bind.CallOpts, eventID string) (*big.Int, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "GetTotalEventVotes", eventID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalEventVotes is a free data retrieval call binding the contract method 0x12b3f7b3.
//
// Solidity: function GetTotalEventVotes(string eventID) view returns(uint256)
func (_Volte *VolteSession) GetTotalEventVotes(eventID string) (*big.Int, error) {
	return _Volte.Contract.GetTotalEventVotes(&_Volte.CallOpts, eventID)
}

// GetTotalEventVotes is a free data retrieval call binding the contract method 0x12b3f7b3.
//
// Solidity: function GetTotalEventVotes(string eventID) view returns(uint256)
func (_Volte *VolteCallerSession) GetTotalEventVotes(eventID string) (*big.Int, error) {
	return _Volte.Contract.GetTotalEventVotes(&_Volte.CallOpts, eventID)
}

// GetVoteMerkleRoot is a free data retrieval call binding the contract method 0x37f5daab.
//
// Solidity: function GetVoteMerkleRoot(string eventID) view returns(uint256)
func (_Volte *VolteCaller) GetVoteMerkleRoot(opts *bind.CallOpts, eventID string) (*big.Int, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "GetVoteMerkleRoot", eventID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVoteMerkleRoot is a free data retrieval call binding the contract method 0x37f5daab.
//
// Solidity: function GetVoteMerkleRoot(string eventID) view returns(uint256)
func (_Volte *VolteSession) GetVoteMerkleRoot(eventID string) (*big.Int, error) {
	return _Volte.Contract.GetVoteMerkleRoot(&_Volte.CallOpts, eventID)
}

// GetVoteMerkleRoot is a free data retrieval call binding the contract method 0x37f5daab.
//
// Solidity: function GetVoteMerkleRoot(string eventID) view returns(uint256)
func (_Volte *VolteCallerSession) GetVoteMerkleRoot(eventID string) (*big.Int, error) {
	return _Volte.Contract.GetVoteMerkleRoot(&_Volte.CallOpts, eventID)
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

// TallyScores is a free data retrieval call binding the contract method 0xa68965c8.
//
// Solidity: function tallyScores(string ) view returns(uint256 C1x, uint256 C1y, uint256 C2x, uint256 C2y)
func (_Volte *VolteCaller) TallyScores(opts *bind.CallOpts, arg0 string) (struct {
	C1x *big.Int
	C1y *big.Int
	C2x *big.Int
	C2y *big.Int
}, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "tallyScores", arg0)

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

// TallyScores is a free data retrieval call binding the contract method 0xa68965c8.
//
// Solidity: function tallyScores(string ) view returns(uint256 C1x, uint256 C1y, uint256 C2x, uint256 C2y)
func (_Volte *VolteSession) TallyScores(arg0 string) (struct {
	C1x *big.Int
	C1y *big.Int
	C2x *big.Int
	C2y *big.Int
}, error) {
	return _Volte.Contract.TallyScores(&_Volte.CallOpts, arg0)
}

// TallyScores is a free data retrieval call binding the contract method 0xa68965c8.
//
// Solidity: function tallyScores(string ) view returns(uint256 C1x, uint256 C1y, uint256 C2x, uint256 C2y)
func (_Volte *VolteCallerSession) TallyScores(arg0 string) (struct {
	C1x *big.Int
	C1y *big.Int
	C2x *big.Int
	C2y *big.Int
}, error) {
	return _Volte.Contract.TallyScores(&_Volte.CallOpts, arg0)
}

// TotalEventVotes is a free data retrieval call binding the contract method 0xf018026c.
//
// Solidity: function totalEventVotes(string ) view returns(uint256)
func (_Volte *VolteCaller) TotalEventVotes(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "totalEventVotes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEventVotes is a free data retrieval call binding the contract method 0xf018026c.
//
// Solidity: function totalEventVotes(string ) view returns(uint256)
func (_Volte *VolteSession) TotalEventVotes(arg0 string) (*big.Int, error) {
	return _Volte.Contract.TotalEventVotes(&_Volte.CallOpts, arg0)
}

// TotalEventVotes is a free data retrieval call binding the contract method 0xf018026c.
//
// Solidity: function totalEventVotes(string ) view returns(uint256)
func (_Volte *VolteCallerSession) TotalEventVotes(arg0 string) (*big.Int, error) {
	return _Volte.Contract.TotalEventVotes(&_Volte.CallOpts, arg0)
}

// VoteMerkleRoots is a free data retrieval call binding the contract method 0x9b28fdf0.
//
// Solidity: function voteMerkleRoots(string ) view returns(uint256)
func (_Volte *VolteCaller) VoteMerkleRoots(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "voteMerkleRoots", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteMerkleRoots is a free data retrieval call binding the contract method 0x9b28fdf0.
//
// Solidity: function voteMerkleRoots(string ) view returns(uint256)
func (_Volte *VolteSession) VoteMerkleRoots(arg0 string) (*big.Int, error) {
	return _Volte.Contract.VoteMerkleRoots(&_Volte.CallOpts, arg0)
}

// VoteMerkleRoots is a free data retrieval call binding the contract method 0x9b28fdf0.
//
// Solidity: function voteMerkleRoots(string ) view returns(uint256)
func (_Volte *VolteCallerSession) VoteMerkleRoots(arg0 string) (*big.Int, error) {
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

// SetVoteMerkleRoot is a paid mutator transaction binding the contract method 0x8b441a3f.
//
// Solidity: function SetVoteMerkleRoot(string eventID, uint256 value) returns()
func (_Volte *VolteTransactor) SetVoteMerkleRoot(opts *bind.TransactOpts, eventID string, value *big.Int) (*types.Transaction, error) {
	return _Volte.contract.Transact(opts, "SetVoteMerkleRoot", eventID, value)
}

// SetVoteMerkleRoot is a paid mutator transaction binding the contract method 0x8b441a3f.
//
// Solidity: function SetVoteMerkleRoot(string eventID, uint256 value) returns()
func (_Volte *VolteSession) SetVoteMerkleRoot(eventID string, value *big.Int) (*types.Transaction, error) {
	return _Volte.Contract.SetVoteMerkleRoot(&_Volte.TransactOpts, eventID, value)
}

// SetVoteMerkleRoot is a paid mutator transaction binding the contract method 0x8b441a3f.
//
// Solidity: function SetVoteMerkleRoot(string eventID, uint256 value) returns()
func (_Volte *VolteTransactorSession) SetVoteMerkleRoot(eventID string, value *big.Int) (*types.Transaction, error) {
	return _Volte.Contract.SetVoteMerkleRoot(&_Volte.TransactOpts, eventID, value)
}

// Vote is a paid mutator transaction binding the contract method 0x7c277c9e.
//
// Solidity: function Vote((string,(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[16],uint256,uint256,uint256,uint256),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[3]))) proof) returns()
func (_Volte *VolteTransactor) Vote(opts *bind.TransactOpts, proof VolteContractVoteSubmission) (*types.Transaction, error) {
	return _Volte.contract.Transact(opts, "Vote", proof)
}

// Vote is a paid mutator transaction binding the contract method 0x7c277c9e.
//
// Solidity: function Vote((string,(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[16],uint256,uint256,uint256,uint256),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[3]))) proof) returns()
func (_Volte *VolteSession) Vote(proof VolteContractVoteSubmission) (*types.Transaction, error) {
	return _Volte.Contract.Vote(&_Volte.TransactOpts, proof)
}

// Vote is a paid mutator transaction binding the contract method 0x7c277c9e.
//
// Solidity: function Vote((string,(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[16],uint256,uint256,uint256,uint256),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[3]))) proof) returns()
func (_Volte *VolteTransactorSession) Vote(proof VolteContractVoteSubmission) (*types.Transaction, error) {
	return _Volte.Contract.Vote(&_Volte.TransactOpts, proof)
}
