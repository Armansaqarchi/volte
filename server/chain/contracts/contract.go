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
	Proof VolteContractProof
	Input [4]*big.Int
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ballot\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_membership\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nullifier\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"BallotProofInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gotRoot\",\"type\":\"uint256\"}],\"name\":\"EventRootMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MembershipProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NullifierProofInvalid\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetEventHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetTallyScore\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetTotalEventVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetVoteMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetEventHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SetVoteMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"Arx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cy\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.Proof\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[4]\",\"name\":\"Input\",\"type\":\"uint256[4]\"}],\"internalType\":\"structVolteContract.BallotProof\",\"name\":\"ballot\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"Arx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cy\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.Proof\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[2]\",\"name\":\"Input\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVolteContract.MembershipProof\",\"name\":\"membership\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"Arx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Ary\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Brx1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Bry1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cx\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Cy\",\"type\":\"uint256\"}],\"internalType\":\"structVolteContract.Proof\",\"name\":\"Proof\",\"type\":\"tuple\"},{\"internalType\":\"uint256[3]\",\"name\":\"Input\",\"type\":\"uint256[3]\"}],\"internalType\":\"structVolteContract.NullifierProof\",\"name\":\"nullifier\",\"type\":\"tuple\"}],\"internalType\":\"structVolteContract.Proofs\",\"name\":\"proofs\",\"type\":\"tuple\"}],\"internalType\":\"structVolteContract.VoteSubmission\",\"name\":\"proof\",\"type\":\"tuple\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ballot\",\"outputs\":[{\"internalType\":\"contractBallotVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"eventHashes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"membership\",\"outputs\":[{\"internalType\":\"contractMembershipVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nullifier\",\"outputs\":[{\"internalType\":\"contractNullifierVerifier\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"tallyScores\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"C1x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"C1y\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"C2x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"C2y\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"totalEventVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"voteMerkleRoots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b5060405161278a38038061278a83398181016040528101906100319190610196565b335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506101e6565b5f5ffd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6101658261013c565b9050919050565b6101758161015b565b811461017f575f5ffd5b50565b5f815190506101908161016c565b92915050565b5f5f5f606084860312156101ad576101ac610138565b5b5f6101ba86828701610182565b93505060206101cb86828701610182565b92505060406101dc86828701610182565b9150509250925092565b612597806101f35f395ff3fe608060405234801561000f575f5ffd5b50600436106100f3575f3560e01c80635dd5ed2811610095578063a68965c811610064578063a68965c8146102a7578063ac3910a2146102da578063f018026c146102f8578063f851a44014610328576100f3565b80635dd5ed281461020d5780638b441a3f1461023d5780639b28fdf014610259578063a2bf68b914610289576100f3565b806337f5daab116100d157806337f5daab146101735780633aa82fee146101a357806341260b3d146101bf57806343bf6cf0146101ef576100f3565b80630d9a3825146100f75780630f00659d1461012757806312b3f7b314610143575b5f5ffd5b610111600480360381019061010c9190611650565b610346565b60405161011e9190611745565b60405180910390f35b610141600480360381019061013c9190611781565b6103a7565b005b61015d60048036038101906101589190611650565b610ca9565b60405161016a91906117d7565b60405180910390f35b61018d6004803603810190610188919061184d565b610cd0565b60405161019a91906117d7565b60405180910390f35b6101bd60048036038101906101b891906118ed565b610cfa565b005b6101d960048036038101906101d49190611650565b610d30565b6040516101e691906119cb565b60405180910390f35b6101f7610de3565b6040516102049190611a65565b60405180910390f35b6102276004803603810190610222919061184d565b610e08565b60405161023491906119cb565b60405180910390f35b61025760048036038101906102529190611aa8565b610eb9565b005b610273600480360381019061026e9190611650565b610ee3565b60405161028091906117d7565b60405180910390f35b610291610f10565b60405161029e9190611b25565b60405180910390f35b6102c160048036038101906102bc9190611650565b610f35565b6040516102d19493929190611b3e565b60405180910390f35b6102e2610f79565b6040516102ef9190611ba1565b60405180910390f35b610312600480360381019061030d9190611650565b610f9e565b60405161031f91906117d7565b60405180910390f35b610330610fcb565b60405161033d9190611bda565b60405180910390f35b61034e6114e1565b5f60068360405161035f9190611c37565b908152602001604051809103902090506040518060800160405280825f0154815260200182600101548152602001826002015481526020018260030154815250915050919050565b600481805f01906103b89190611c59565b6040516103c6929190611cdf565b9081526020016040518091039020548160200161018001610100015f600281106103f3576103f2611cf7565b5b60200201351461048b57600481805f019061040e9190611c59565b60405161041c929190611cdf565b9081526020016040518091039020548160200161018001610100015f6002811061044957610448611cf7565b5b60200201356040517f1f0189e2000000000000000000000000000000000000000000000000000000008152600401610482929190611d24565b60405180910390fd5b5f60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166311479fea6040518060400160405280856020016102c0015f015f01358152602001856020016102c0015f016020013581525060405180604001604052806040518060400160405280886020016102c0015f01606001358152602001886020016102c0015f016040013581525081526020016040518060400160405280886020016102c0015f0160a001358152602001886020016102c0015f01608001358152508152506040518060400160405280876020016102c0015f0160c001358152602001876020016102c0015f0160e00135815250866020016102c001610100016040518563ffffffff1660e01b81526004016105c29493929190611ed9565b602060405180830381865afa1580156105dd573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106019190611f53565b90508061063a576040517f0d83a54d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f5c9d69e604051806040016040528086602001610180015f015f0135815260200186602001610180015f01602001358152506040518060400160405280604051806040016040528089602001610180015f0160600135815260200189602001610180015f01604001358152508152602001604051806040016040528089602001610180015f0160a00135815260200189602001610180015f0160800135815250815250604051806040016040528088602001610180015f0160c00135815260200188602001610180015f0160e001358152508760200161018001610100016040518563ffffffff1660e01b81526004016107719493929190611f8e565b602060405180830381865afa15801561078c573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107b09190611f53565b9050806107e9576040517f986d47fa00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635fe8c13b6040518060400160405280876020015f015f015f01358152602001876020015f015f0160200135815250604051806040016040528060405180604001604052808a6020015f015f016060013581526020018a6020015f015f0160400135815250815260200160405180604001604052808a6020015f015f0160a0013581526020018a6020015f015f01608001358152508152506040518060400160405280896020015f015f0160c001358152602001896020015f015f0160e00135815250886020015f01610100016040518563ffffffff1660e01b815260040161090e9493929190611fe3565b602060405180830381865afa158015610929573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061094d9190611f53565b905080610986576040517fbf2c074200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f84805f01906109969190611c59565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505090505f6006826040516109eb9190611c37565b908152602001604051809103902090505f6040518060400160405280886020015f01610100015f60048110610a2357610a22611cf7565b5b60200201358152602001886020015f0161010001600160048110610a4a57610a49611cf7565b5b602002013581525090505f6040518060400160405280896020015f0161010001600260048110610a7d57610a7c611cf7565b5b60200201358152602001896020015f0161010001600360048110610aa457610aa3611cf7565b5b602002013581525090505f835f0154148015610ac357505f8360010154145b15610adb575f835f0181905550600183600101819055505b5f8360020154148015610af157505f8360030154145b15610b0a575f8360020181905550600183600301819055505b5f6040518060400160405280855f01548152602001856001015481525090505f604051806040016040528086600201548152602001866003015481525090505f5f610bb5865f60028110610b6157610b60611cf7565b5b602002015187600160028110610b7a57610b79611cf7565b5b6020020151865f60028110610b9257610b91611cf7565b5b602002015187600160028110610bab57610baa611cf7565b5b6020020151610fef565b915091505f5f610c25875f60028110610bd157610bd0611cf7565b5b602002015188600160028110610bea57610be9611cf7565b5b6020020151875f60028110610c0257610c01611cf7565b5b602002015188600160028110610c1b57610c1a611cf7565b5b6020020151610fef565b9150915083895f0181905550828960010181905550818960020181905550808960030181905550600160078b604051610c5e9190611c37565b908152602001604051809103902054610c779190612055565b60078b604051610c879190611c37565b9081526020016040518091039020819055505050505050505050505050505050565b5f600782604051610cba9190611c37565b9081526020016040518091039020549050919050565b5f60048383604051610ce3929190611cdf565b908152602001604051809103902054905092915050565b818160058686604051610d0e929190611cdf565b90815260200160405180910390209182610d29929190612286565b5050505050565b6005818051602081018201805184825260208301602085012081835280955050505050505f915090508054610d64906120bf565b80601f0160208091040260200160405190810160405280929190818152602001828054610d90906120bf565b8015610ddb5780601f10610db257610100808354040283529160200191610ddb565b820191905f5260205f20905b815481529060010190602001808311610dbe57829003601f168201915b505050505081565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060058383604051610e1c929190611cdf565b90815260200160405180910390208054610e35906120bf565b80601f0160208091040260200160405190810160405280929190818152602001828054610e61906120bf565b8015610eac5780601f10610e8357610100808354040283529160200191610eac565b820191905f5260205f20905b815481529060010190602001808311610e8f57829003601f168201915b5050505050905092915050565b8060048484604051610ecc929190611cdf565b908152602001604051809103902081905550505050565b6004818051602081018201805184825260208301602085012081835280955050505050505f915090505481565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6006818051602081018201805184825260208301602085012081835280955050505050505f91509050805f0154908060010154908060020154908060030154905084565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6007818051602081018201805184825260208301602085012081835280955050505050505f915090505481565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5f5f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018061102157611020612353565b5b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001806110505761104f612353565b5b8688097f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018061108257611081612353565b5b868a090890505f6111297f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001806110bb576110ba612353565b5b8689097f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001806110ed576110ec612353565b5b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018061111c5761111b612353565b5b898c09620292fc096112c0565b90505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018061115b5761115a612353565b5b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018061118a57611189612353565b5b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001806111b9576111b8612353565b5b888b097f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001806111eb576111ea612353565b5b8a8d0909620292f80990505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018061122657611225612353565b5b8260010890505f6112386001846112c0565b90507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018061126957611268612353565b5b6112728361134d565b860996507f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001806112a5576112a4612353565b5b6112ae8261134d565b85099550505050505094509492505050565b5f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001806112f0576112ef612353565b5b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000183816113205761131f612353565b5b067f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f0000001038408905092915050565b5f5f8203611390576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611387906123da565b60405180910390fd5b6113e78260027f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016113c191906123f8565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000016113ee565b9050919050565b5f5f602080602087878760405160200161140d9695949392919061244b565b60405160208183030381529060405290505f602067ffffffffffffffff81111561143a5761143961152c565b5b6040519080825280601f01601f19166020018201604052801561146c5781602001600182028036833780820191505090505b5090505f602080830184516020860160055afa9050806114c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114b890612504565b60405180910390fd5b818060200190518101906114d59190612536565b93505050509392505050565b6040518060800160405280600490602082028036833780820191505090505090565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6115628261151c565b810181811067ffffffffffffffff821117156115815761158061152c565b5b80604052505050565b5f611593611503565b905061159f8282611559565b919050565b5f67ffffffffffffffff8211156115be576115bd61152c565b5b6115c78261151c565b9050602081019050919050565b828183375f83830152505050565b5f6115f46115ef846115a4565b61158a565b9050828152602081018484840111156116105761160f611518565b5b61161b8482856115d4565b509392505050565b5f82601f83011261163757611636611514565b5b81356116478482602086016115e2565b91505092915050565b5f602082840312156116655761166461150c565b5b5f82013567ffffffffffffffff81111561168257611681611510565b5b61168e84828501611623565b91505092915050565b5f60049050919050565b5f81905092915050565b5f819050919050565b5f819050919050565b6116c6816116b4565b82525050565b5f6116d783836116bd565b60208301905092915050565b5f602082019050919050565b6116f881611697565b61170281846116a1565b925061170d826116ab565b805f5b8381101561173d57815161172487826116cc565b965061172f836116e3565b925050600181019050611710565b505050505050565b5f6080820190506117585f8301846116ef565b92915050565b5f5ffd5b5f61044082840312156117785761177761175e565b5b81905092915050565b5f602082840312156117965761179561150c565b5b5f82013567ffffffffffffffff8111156117b3576117b2611510565b5b6117bf84828501611762565b91505092915050565b6117d1816116b4565b82525050565b5f6020820190506117ea5f8301846117c8565b92915050565b5f5ffd5b5f5ffd5b5f5f83601f84011261180d5761180c611514565b5b8235905067ffffffffffffffff81111561182a576118296117f0565b5b602083019150836001820283011115611846576118456117f4565b5b9250929050565b5f5f602083850312156118635761186261150c565b5b5f83013567ffffffffffffffff8111156118805761187f611510565b5b61188c858286016117f8565b92509250509250929050565b5f5f83601f8401126118ad576118ac611514565b5b8235905067ffffffffffffffff8111156118ca576118c96117f0565b5b6020830191508360018202830111156118e6576118e56117f4565b5b9250929050565b5f5f5f5f604085870312156119055761190461150c565b5b5f85013567ffffffffffffffff81111561192257611921611510565b5b61192e878288016117f8565b9450945050602085013567ffffffffffffffff81111561195157611950611510565b5b61195d87828801611898565b925092505092959194509250565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f61199d8261196b565b6119a78185611975565b93506119b7818560208601611985565b6119c08161151c565b840191505092915050565b5f6020820190508181035f8301526119e38184611993565b905092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f611a2d611a28611a23846119eb565b611a0a565b6119eb565b9050919050565b5f611a3e82611a13565b9050919050565b5f611a4f82611a34565b9050919050565b611a5f81611a45565b82525050565b5f602082019050611a785f830184611a56565b92915050565b611a87816116b4565b8114611a91575f5ffd5b50565b5f81359050611aa281611a7e565b92915050565b5f5f5f60408486031215611abf57611abe61150c565b5b5f84013567ffffffffffffffff811115611adc57611adb611510565b5b611ae8868287016117f8565b93509350506020611afb86828701611a94565b9150509250925092565b5f611b0f82611a34565b9050919050565b611b1f81611b05565b82525050565b5f602082019050611b385f830184611b16565b92915050565b5f608082019050611b515f8301876117c8565b611b5e60208301866117c8565b611b6b60408301856117c8565b611b7860608301846117c8565b95945050505050565b5f611b8b82611a34565b9050919050565b611b9b81611b81565b82525050565b5f602082019050611bb45f830184611b92565b92915050565b5f611bc4826119eb565b9050919050565b611bd481611bba565b82525050565b5f602082019050611bed5f830184611bcb565b92915050565b5f81519050919050565b5f81905092915050565b5f611c1182611bf3565b611c1b8185611bfd565b9350611c2b818560208601611985565b80840191505092915050565b5f611c428284611c07565b915081905092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83356001602003843603038112611c7557611c74611c4d565b5b80840192508235915067ffffffffffffffff821115611c9757611c96611c51565b5b602083019250600182023603831315611cb357611cb2611c55565b5b509250929050565b5f611cc68385611bfd565b9350611cd38385846115d4565b82840190509392505050565b5f611ceb828486611cbb565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f604082019050611d375f8301856117c8565b611d4460208301846117c8565b9392505050565b5f60029050919050565b5f81905092915050565b5f819050919050565b5f602082019050919050565b611d7d81611d4b565b611d878184611d55565b9250611d9282611d5f565b805f5b83811015611dc2578151611da987826116cc565b9650611db483611d68565b925050600181019050611d95565b505050505050565b5f60029050919050565b5f81905092915050565b5f819050919050565b5f81905092915050565b611dfa81611d4b565b611e048184611de7565b9250611e0f82611d5f565b805f5b83811015611e3f578151611e2687826116cc565b9650611e3183611d68565b925050600181019050611e12565b505050505050565b5f611e528383611df1565b60408301905092915050565b5f602082019050919050565b611e7381611dca565b611e7d8184611dd4565b9250611e8882611dde565b805f5b83811015611eb8578151611e9f8782611e47565b9650611eaa83611e5e565b925050600181019050611e8b565b505050505050565b82818337505050565b611ed560608383611ec0565b5050565b5f61016082019050611eed5f830187611d74565b611efa6040830186611e6a565b611f0760c0830185611d74565b611f15610100830184611ec9565b95945050505050565b5f8115159050919050565b611f3281611f1e565b8114611f3c575f5ffd5b50565b5f81519050611f4d81611f29565b92915050565b5f60208284031215611f6857611f6761150c565b5b5f611f7584828501611f3f565b91505092915050565b611f8a60408383611ec0565b5050565b5f61014082019050611fa25f830187611d74565b611faf6040830186611e6a565b611fbc60c0830185611d74565b611fca610100830184611f7e565b95945050505050565b611fdf60808383611ec0565b5050565b5f61018082019050611ff75f830187611d74565b6120046040830186611e6a565b61201160c0830185611d74565b61201f610100830184611fd3565b95945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61205f826116b4565b915061206a836116b4565b925082820190508082111561208257612081612028565b5b92915050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806120d657607f821691505b6020821081036120e9576120e8612092565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261214b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82612110565b6121558683612110565b95508019841693508086168417925050509392505050565b5f61218761218261217d846116b4565b611a0a565b6116b4565b9050919050565b5f819050919050565b6121a08361216d565b6121b46121ac8261218e565b84845461211c565b825550505050565b5f5f905090565b6121cb6121bc565b6121d6818484612197565b505050565b5b818110156121f9576121ee5f826121c3565b6001810190506121dc565b5050565b601f82111561223e5761220f816120ef565b61221884612101565b81016020851015612227578190505b61223b61223385612101565b8301826121db565b50505b505050565b5f82821c905092915050565b5f61225e5f1984600802612243565b1980831691505092915050565b5f612276838361224f565b9150826002028217905092915050565b6122908383612088565b67ffffffffffffffff8111156122a9576122a861152c565b5b6122b382546120bf565b6122be8282856121fd565b5f601f8311600181146122eb575f84156122d9578287013590505b6122e3858261226b565b86555061234a565b601f1984166122f9866120ef565b5f5b82811015612320578489013582556001820191506020850194506020810190506122fb565b8683101561233d5784890135612339601f89168261224f565b8355505b6001600288020188555050505b50505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f82825260208201905092915050565b7f696e7628302900000000000000000000000000000000000000000000000000005f82015250565b5f6123c4600683612380565b91506123cf82612390565b602082019050919050565b5f6020820190508181035f8301526123f1816123b8565b9050919050565b5f612402826116b4565b915061240d836116b4565b925082820390508181111561242557612424612028565b5b92915050565b5f819050919050565b612445612440826116b4565b61242b565b82525050565b5f6124568289612434565b6020820191506124668288612434565b6020820191506124768287612434565b6020820191506124868286612434565b6020820191506124968285612434565b6020820191506124a68284612434565b602082019150819050979650505050505050565b7f6d6f64657870206661696c6564000000000000000000000000000000000000005f82015250565b5f6124ee600d83612380565b91506124f9826124ba565b602082019050919050565b5f6020820190508181035f83015261251b816124e2565b9050919050565b5f8151905061253081611a7e565b92915050565b5f6020828403121561254b5761254a61150c565b5b5f61255884828501612522565b9150509291505056fea26469706673582212204dcf80e58d8c12d213a69b123683ff3617c0b5429781ebbaac2574e1c01e03fc64736f6c634300081e0033",
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

// Vote is a paid mutator transaction binding the contract method 0x0f00659d.
//
// Solidity: function Vote((string,(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[4]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[3]))) proof) returns()
func (_Volte *VolteTransactor) Vote(opts *bind.TransactOpts, proof VolteContractVoteSubmission) (*types.Transaction, error) {
	return _Volte.contract.Transact(opts, "Vote", proof)
}

// Vote is a paid mutator transaction binding the contract method 0x0f00659d.
//
// Solidity: function Vote((string,(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[4]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[3]))) proof) returns()
func (_Volte *VolteSession) Vote(proof VolteContractVoteSubmission) (*types.Transaction, error) {
	return _Volte.Contract.Vote(&_Volte.TransactOpts, proof)
}

// Vote is a paid mutator transaction binding the contract method 0x0f00659d.
//
// Solidity: function Vote((string,(((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[4]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[2]),((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256),uint256[3]))) proof) returns()
func (_Volte *VolteTransactorSession) Vote(proof VolteContractVoteSubmission) (*types.Transaction, error) {
	return _Volte.Contract.Vote(&_Volte.TransactOpts, proof)
}
