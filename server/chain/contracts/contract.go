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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ballot\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_membership\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nullifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetEventHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"}],\"name\":\"GetTallyScore\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetVoteMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetEventHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetVoteMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"Arx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cy\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.Proof\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[16]\",\"name\":\"Input\",\"type\":\"uint256[16]\"},{\"internalType\":\"uint256\",\"name\":\"CommitmentX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommitmentY\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommitmentPokX\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommitmentPokY\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.BallotProof\",\"name\":\"ballot\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"Arx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cy\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.Proof\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[2]\",\"name\":\"Input\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVolteContract.MembershipProof\",\"name\":\"membership\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"Arx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cy\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.Proof\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[2]\",\"name\":\"Input\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVolteContract.NullifierProof\",\"name\":\"nullifier\",\"type\":\"tuple\"}],\"internalType\":\"structVolteContract.Proofs\",\"name\":\"proofs\",\"type\":\"tuple\"}],\"internalType\":\"structVolteContract.VoteSubmission\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"C1\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"C2\",\"type\":\"uint256[2]\"}],\"name\":\"addCiphertexts\",\"outputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"Csum\",\"type\":\"uint256[2]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballot\",\"outputs\":[{\"internalType\":\"contractBallotVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"eventHashes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"membership\",\"outputs\":[{\"internalType\":\"contractMembershipVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"membershipMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nullifier\",\"outputs\":[{\"internalType\":\"contractNullifierVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tallyScores\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"C1x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"C1y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"C2x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"C2y\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"voteMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161226b38038061226b83398181016040528101906100319190610196565b335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506101e6565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101658261013c565b9050919050565b6101758161015b565b811461017f575f5ffd5b50565b5f815190506101908161016c565b92915050565b5f5f5f606084860312156101ad576101ac610138565b5b5f6101ba86828701610182565b93505060206101cb86828701610182565b92505060406101dc86828701610182565b9150509250925092565b612078806101f35f395ff3fe608060405234801561000f575f5ffd5b50600436106100f3575f3560e01c80639b28fdf011610095578063d59f766111610064578063d59f766114610298578063f3c6fe65146102c8578063f851a440146102f8578063fdab8f2314610316576100f3565b80639b28fdf0146101f9578063a2bf68b914610229578063ac3910a214610247578063c5126a4814610265576100f3565b806343bf6cf0116100d157806343bf6cf0146101735780634c4ff130146101915780635dd5ed28146101ad578063649a75ca146101dd576100f3565b806337f5daab146100f75780633aa82fee1461012757806341260b3d14610143575b5f5ffd5b610111600480360381019061010c919061138f565b610346565b60405161011e919061144a565b60405180910390f35b610141600480360381019061013c91906114bf565b6103f7565b005b61015d60048036038101906101589190611665565b61042d565b60405161016a919061144a565b60405180910390f35b61017b6104e0565b6040516101889190611726565b60405180910390f35b6101ab60048036038101906101a69190611762565b610505565b005b6101c760048036038101906101c2919061138f565b610e32565b6040516101d4919061144a565b60405180910390f35b6101f760048036038101906101f291906114bf565b610ee3565b005b610213600480360381019061020e9190611665565b610f19565b604051610220919061144a565b60405180910390f35b610231610fcc565b60405161023e91906117c9565b60405180910390f35b61024f610ff1565b60405161025c9190611802565b60405180910390f35b61027f600480360381019061027a919061184e565b611016565b60405161028f9493929190611888565b60405180910390f35b6102b260048036038101906102ad9190611665565b611042565b6040516102bf919061144a565b60405180910390f35b6102e260048036038101906102dd919061184e565b6110f5565b6040516102ef9190611970565b60405180910390f35b610300611149565b60405161030d91906119a9565b60405180910390f35b610330600480360381019061032b9190611a70565b61116d565b60405161033d9190611b2d565b60405180910390f35b60606005838360405161035a929190611b74565b9081526020016040518091039020805461037390611bb9565b80601f016020809104026020016040519081016040528092919081815260200182805461039f90611bb9565b80156103ea5780601f106103c1576101008083540402835291602001916103ea565b820191905f5260205f20905b8154815290600101906020018083116103cd57829003601f168201915b5050505050905092915050565b81816006868660405161040b929190611b74565b90815260200160405180910390209182610426929190611d8a565b5050505050565b6006818051602081018201805184825260208301602085012081835280955050505050505f91509050805461046190611bb9565b80601f016020809104026020016040519081016040528092919081815260200182805461048d90611bb9565b80156104d85780601f106104af576101008083540402835291602001916104d8565b820191905f5260205f20905b8154815290600101906020018083116104bb57829003601f168201915b505050505081565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635fe24f23604051806101000160405280846040016104c0015f015f01358152602001846040016104c0015f01602001358152602001846040016104c0015f01606001358152602001846040016104c0015f01604001358152602001846040016104c0015f0160a001358152602001846040016104c0015f01608001358152602001846040016104c0015f0160c001358152602001846040016104c0015f0160e00135815250836040016104c001610100016040518363ffffffff1660e01b815260040161060c929190611eef565b5f6040518083038186803b158015610622575f5ffd5b505afa158015610634573d5f5f3e3d5ffd5b5050505060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635fe24f2360405180610100016040528084604001610380015f015f0135815260200184604001610380015f0160200135815260200184604001610380015f0160600135815260200184604001610380015f0160400135815260200184604001610380015f0160a00135815260200184604001610380015f0160800135815260200184604001610380015f0160c00135815260200184604001610380015f0160e001358152508360400161038001610100016040518363ffffffff1660e01b815260040161073f929190611eef565b5f6040518083038186803b158015610755575f5ffd5b505afa158015610767573d5f5f3e3d5ffd5b5050505060015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dc6395e6604051806101000160405280846040015f015f015f01358152602001846040015f015f01602001358152602001846040015f015f01606001358152602001846040015f015f01604001358152602001846040015f015f0160a001358152602001846040015f015f01608001358152602001846040015f015f0160c001358152602001846040015f015f0160e001358152506040518060400160405280856040015f0161030001358152602001856040015f0161032001358152506040518060400160405280866040015f0161034001358152602001866040015f016103600135815250856040015f01610100016040518563ffffffff1660e01b81526004016108b49493929190611f29565b5f6040518083038186803b1580156108ca575f5ffd5b505afa1580156108dc573d5f5f3e3d5ffd5b505050505f6109dd6040518060800160405280846040015f01610100015f6010811061090b5761090a611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001846040015f016101000160016010811061094657610945611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001846040015f016101000160026010811061098157610980611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001846040015f01610100016003601081106109bc576109bb611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff16815250611203565b90505f610add6040518060800160405280856040015f0161010001600460108110610a0b57610a0a611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001856040015f0161010001600560108110610a4657610a45611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001856040015f0161010001600660108110610a8157610a80611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001856040015f0161010001600760108110610abc57610abb611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff16815250611203565b90505f610bdd6040518060800160405280866040015f0161010001600860108110610b0b57610b0a611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001866040015f0161010001600960108110610b4657610b45611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001866040015f0161010001600a60108110610b8157610b80611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001866040015f0161010001600b60108110610bbc57610bbb611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff16815250611203565b90505f610cdd6040518060800160405280876040015f0161010001600c60108110610c0b57610c0a611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001876040015f0161010001600d60108110610c4657610c45611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001876040015f0161010001600e60108110610c8157610c80611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff168152602001876040015f0161010001600f60108110610cbc57610cbb611f70565b5b602002013567ffffffffffffffff1667ffffffffffffffff16815250611203565b90505f856040016104c001610100015f60028110610cfe57610cfd611f70565b5b602002013590505f60075f8381526020019081526020015f2090505f60405180604001604052808881526020018781525090505f60405180604001604052808781526020018681525090505f6040518060400160405280855f01548152602001856001015481525090505f604051806040016040528086600201548152602001866003015481525090505f610d93858461116d565b90505f610da0858461116d565b9050815f60028110610db557610db4611f70565b5b6020020151875f018190555081600160028110610dd557610dd4611f70565b5b60200201518760010181905550805f60028110610df557610df4611f70565b5b6020020151876002018190555080600160028110610e1657610e15611f70565b5b6020020151876003018190555050505050505050505050505050565b606060068383604051610e46929190611b74565b90815260200160405180910390208054610e5f90611bb9565b80601f0160208091040260200160405190810160405280929190818152602001828054610e8b90611bb9565b8015610ed65780601f10610ead57610100808354040283529160200191610ed6565b820191905f5260205f20905b815481529060010190602001808311610eb957829003601f168201915b5050505050905092915050565b818160058686604051610ef7929190611b74565b90815260200160405180910390209182610f12929190611d8a565b5050505050565b6005818051602081018201805184825260208301602085012081835280955050505050505f915090508054610f4d90611bb9565b80601f0160208091040260200160405190810160405280929190818152602001828054610f7990611bb9565b8015610fc45780601f10610f9b57610100808354040283529160200191610fc4565b820191905f5260205f20905b815481529060010190602001808311610fa757829003601f168201915b505050505081565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6007602052805f5260405f205f91509050805f0154908060010154908060020154908060030154905084565b6004818051602081018201805184825260208301602085012081835280955050505050505f91509050805461107690611bb9565b80601f01602080910402602001604051908101604052809291908181526020018280546110a290611bb9565b80156110ed5780601f106110c4576101008083540402835291602001916110ed565b820191905f5260205f20905b8154815290600101906020018083116110d057829003601f168201915b505050505081565b6110fd6112d9565b5f60075f8481526020019081526020015f2090506040518060800160405280825f0154815260200182600101548152602001826002015481526020018260030154815250915050919050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6111756112fb565b5f6040518451815260208501516020820152835160408201526020840151606082015260408160808360065afa9150828151815260208201516020820152608082016040525050806111fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111f390611ff7565b60405180910390fd5b5092915050565b5f60c08260036004811061121a57611219611f70565b5b602002015167ffffffffffffffff16901b60808360026004811061124157611240611f70565b5b602002015167ffffffffffffffff16901b60408460016004811061126857611267611f70565b5b602002015167ffffffffffffffff16901b845f6004811061128c5761128b611f70565b5b602002015167ffffffffffffffff1617171790507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001806112cf576112ce612015565b5b5f82089050919050565b6040518060800160405280600490602082028036833780820191505090505090565b6040518060400160405280600290602082028036833780820191505090505090565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261134f5761134e61132e565b5b8235905067ffffffffffffffff81111561136c5761136b611332565b5b60208301915083600182028301111561138857611387611336565b5b9250929050565b5f5f602083850312156113a5576113a4611326565b5b5f83013567ffffffffffffffff8111156113c2576113c161132a565b5b6113ce8582860161133a565b92509250509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61141c826113da565b61142681856113e4565b93506114368185602086016113f4565b61143f81611402565b840191505092915050565b5f6020820190508181035f8301526114628184611412565b905092915050565b5f5f83601f84011261147f5761147e61132e565b5b8235905067ffffffffffffffff81111561149c5761149b611332565b5b6020830191508360018202830111156114b8576114b7611336565b5b9250929050565b5f5f5f5f604085870312156114d7576114d6611326565b5b5f85013567ffffffffffffffff8111156114f4576114f361132a565b5b6115008782880161133a565b9450945050602085013567ffffffffffffffff8111156115235761152261132a565b5b61152f8782880161146a565b925092505092959194509250565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61157782611402565b810181811067ffffffffffffffff8211171561159657611595611541565b5b80604052505050565b5f6115a861131d565b90506115b4828261156e565b919050565b5f67ffffffffffffffff8211156115d3576115d2611541565b5b6115dc82611402565b9050602081019050919050565b828183375f83830152505050565b5f611609611604846115b9565b61159f565b9050828152602081018484840111156116255761162461153d565b5b6116308482856115e9565b509392505050565b5f82601f83011261164c5761164b61132e565b5b813561165c8482602086016115f7565b91505092915050565b5f6020828403121561167a57611679611326565b5b5f82013567ffffffffffffffff8111156116975761169661132a565b5b6116a384828501611638565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f6116ee6116e96116e4846116ac565b6116cb565b6116ac565b9050919050565b5f6116ff826116d4565b9050919050565b5f611710826116f5565b9050919050565b61172081611706565b82525050565b5f6020820190506117395f830184611717565b92915050565b5f5ffd5b5f61064082840312156117595761175861173f565b5b81905092915050565b5f6020828403121561177757611776611326565b5b5f82013567ffffffffffffffff8111156117945761179361132a565b5b6117a084828501611743565b91505092915050565b5f6117b3826116f5565b9050919050565b6117c3816117a9565b82525050565b5f6020820190506117dc5f8301846117ba565b92915050565b5f6117ec826116f5565b9050919050565b6117fc816117e2565b82525050565b5f6020820190506118155f8301846117f3565b92915050565b5f819050919050565b61182d8161181b565b8114611837575f5ffd5b50565b5f8135905061184881611824565b92915050565b5f6020828403121561186357611862611326565b5b5f6118708482850161183a565b91505092915050565b6118828161181b565b82525050565b5f60808201905061189b5f830187611879565b6118a86020830186611879565b6118b56040830185611879565b6118c26060830184611879565b95945050505050565b5f60049050919050565b5f81905092915050565b5f819050919050565b6118f18161181b565b82525050565b5f61190283836118e8565b60208301905092915050565b5f602082019050919050565b611923816118cb565b61192d81846118d5565b9250611938826118df565b805f5b8381101561196857815161194f87826118f7565b965061195a8361190e565b92505060018101905061193b565b505050505050565b5f6080820190506119835f83018461191a565b92915050565b5f611993826116ac565b9050919050565b6119a381611989565b82525050565b5f6020820190506119bc5f83018461199a565b92915050565b5f67ffffffffffffffff8211156119dc576119db611541565b5b602082029050919050565b5f6119f96119f4846119c2565b61159f565b90508060208402830185811115611a1357611a12611336565b5b835b81811015611a3c5780611a28888261183a565b845260208401935050602081019050611a15565b5050509392505050565b5f82601f830112611a5a57611a5961132e565b5b6002611a678482856119e7565b91505092915050565b5f5f60808385031215611a8657611a85611326565b5b5f611a9385828601611a46565b9250506040611aa485828601611a46565b9150509250929050565b5f60029050919050565b5f81905092915050565b5f819050919050565b5f602082019050919050565b611ae081611aae565b611aea8184611ab8565b9250611af582611ac2565b805f5b83811015611b25578151611b0c87826118f7565b9650611b1783611acb565b925050600181019050611af8565b505050505050565b5f604082019050611b405f830184611ad7565b92915050565b5f81905092915050565b5f611b5b8385611b46565b9350611b688385846115e9565b82840190509392505050565b5f611b80828486611b50565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611bd057607f821691505b602082108103611be357611be2611b8c565b5b50919050565b5f82905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611c4f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611c14565b611c598683611c14565b95508019841693508086168417925050509392505050565b5f611c8b611c86611c818461181b565b6116cb565b61181b565b9050919050565b5f819050919050565b611ca483611c71565b611cb8611cb082611c92565b848454611c20565b825550505050565b5f5f905090565b611ccf611cc0565b611cda818484611c9b565b505050565b5b81811015611cfd57611cf25f82611cc7565b600181019050611ce0565b5050565b601f821115611d4257611d1381611bf3565b611d1c84611c05565b81016020851015611d2b578190505b611d3f611d3785611c05565b830182611cdf565b50505b505050565b5f82821c905092915050565b5f611d625f1984600802611d47565b1980831691505092915050565b5f611d7a8383611d53565b9150826002028217905092915050565b611d948383611be9565b67ffffffffffffffff811115611dad57611dac611541565b5b611db78254611bb9565b611dc2828285611d01565b5f601f831160018114611def575f8415611ddd578287013590505b611de78582611d6f565b865550611e4e565b601f198416611dfd86611bf3565b5f5b82811015611e2457848901358255600182019150602085019450602081019050611dff565b86831015611e415784890135611e3d601f891682611d53565b8355505b6001600288020188555050505b50505050505050565b5f60089050919050565b5f81905092915050565b5f819050919050565b5f602082019050919050565b611e8981611e57565b611e938184611e61565b9250611e9e82611e6b565b805f5b83811015611ece578151611eb587826118f7565b9650611ec083611e74565b925050600181019050611ea1565b505050505050565b82818337505050565b611eeb60408383611ed6565b5050565b5f61014082019050611f035f830185611e80565b611f11610100830184611edf565b9392505050565b611f256102008383611ed6565b5050565b5f61038082019050611f3d5f830187611e80565b611f4b610100830186611ad7565b611f59610140830185611ad7565b611f67610180830184611f18565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f82825260208201905092915050565b7f4543414444206661696c656400000000000000000000000000000000000000005f82015250565b5f611fe1600c83611f9d565b9150611fec82611fad565b602082019050919050565b5f6020820190508181035f83015261200e81611fd5565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffdfea2646970667358221220767af2b9b4dce98b355ee51619f3d6d59db0b4e416265316bbdd2dcae034678a64736f6c634300081e0033",
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

// GetTallyScore is a free data retrieval call binding the contract method 0xf3c6fe65.
//
// Solidity: function GetTallyScore(uint256 eventID) view returns(uint256[4])
func (_Volte *VolteCaller) GetTallyScore(opts *bind.CallOpts, eventID *big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "GetTallyScore", eventID)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetTallyScore is a free data retrieval call binding the contract method 0xf3c6fe65.
//
// Solidity: function GetTallyScore(uint256 eventID) view returns(uint256[4])
func (_Volte *VolteSession) GetTallyScore(eventID *big.Int) ([4]*big.Int, error) {
	return _Volte.Contract.GetTallyScore(&_Volte.CallOpts, eventID)
}

// GetTallyScore is a free data retrieval call binding the contract method 0xf3c6fe65.
//
// Solidity: function GetTallyScore(uint256 eventID) view returns(uint256[4])
func (_Volte *VolteCallerSession) GetTallyScore(eventID *big.Int) ([4]*big.Int, error) {
	return _Volte.Contract.GetTallyScore(&_Volte.CallOpts, eventID)
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

// TallyScores is a free data retrieval call binding the contract method 0xc5126a48.
//
// Solidity: function tallyScores(uint256 ) view returns(uint256 C1x, uint256 C1y, uint256 C2x, uint256 C2y)
func (_Volte *VolteCaller) TallyScores(opts *bind.CallOpts, arg0 *big.Int) (struct {
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

// TallyScores is a free data retrieval call binding the contract method 0xc5126a48.
//
// Solidity: function tallyScores(uint256 ) view returns(uint256 C1x, uint256 C1y, uint256 C2x, uint256 C2y)
func (_Volte *VolteSession) TallyScores(arg0 *big.Int) (struct {
	C1x *big.Int
	C1y *big.Int
	C2x *big.Int
	C2y *big.Int
}, error) {
	return _Volte.Contract.TallyScores(&_Volte.CallOpts, arg0)
}

// TallyScores is a free data retrieval call binding the contract method 0xc5126a48.
//
// Solidity: function tallyScores(uint256 ) view returns(uint256 C1x, uint256 C1y, uint256 C2x, uint256 C2y)
func (_Volte *VolteCallerSession) TallyScores(arg0 *big.Int) (struct {
	C1x *big.Int
	C1y *big.Int
	C2x *big.Int
	C2y *big.Int
}, error) {
	return _Volte.Contract.TallyScores(&_Volte.CallOpts, arg0)
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

// Vote is a paid mutator transaction binding the contract method 0x4c4ff130.
//
// Solidity: function Vote((address,string,(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[16],uint256,uint256,uint256,uint256),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]))) proof) returns()
func (_Volte *VolteTransactor) Vote(opts *bind.TransactOpts, proof VolteContractVoteSubmission) (*types.Transaction, error) {
	return _Volte.contract.Transact(opts, "Vote", proof)
}

// Vote is a paid mutator transaction binding the contract method 0x4c4ff130.
//
// Solidity: function Vote((address,string,(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[16],uint256,uint256,uint256,uint256),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]))) proof) returns()
func (_Volte *VolteSession) Vote(proof VolteContractVoteSubmission) (*types.Transaction, error) {
	return _Volte.Contract.Vote(&_Volte.TransactOpts, proof)
}

// Vote is a paid mutator transaction binding the contract method 0x4c4ff130.
//
// Solidity: function Vote((address,string,(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[16],uint256,uint256,uint256,uint256),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]))) proof) returns()
func (_Volte *VolteTransactorSession) Vote(proof VolteContractVoteSubmission) (*types.Transaction, error) {
	return _Volte.Contract.Vote(&_Volte.TransactOpts, proof)
}
