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

// VolteMetaData contains all meta data concerning the Volte contract.
var VolteMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetEventHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetNullifierMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetVoteMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetEventHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetNullifierMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetVoteMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"eventHashes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"nullifierMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"usedNullifiers\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"voteMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061119a8061005b5f395ff3fe608060405234801561000f575f5ffd5b50600436106100a7575f3560e01c8063649a75ca1161006f578063649a75ca1461017357806379488fdd1461018f5780639b28fdf0146101bf578063dd7a2ca2146101ef578063e1a2b21f1461021f578063f851a4401461024f576100a7565b80632acceddd146100ab57806337f5daab146100c75780633aa82fee146100f757806341260b3d146101135780635dd5ed2814610143575b5f5ffd5b6100c560048036038101906100c09190610a83565b61026d565b005b6100e160048036038101906100dc9190610b01565b610331565b6040516100ee9190610bbc565b60405180910390f35b610111600480360381019061010c9190610a83565b6103e2565b005b61012d60048036038101906101289190610d04565b6104a6565b60405161013a9190610bbc565b60405180910390f35b61015d60048036038101906101589190610b01565b610559565b60405161016a9190610bbc565b60405180910390f35b61018d60048036038101906101889190610a83565b61060a565b005b6101a960048036038101906101a49190610d04565b6106ce565b6040516101b69190610bbc565b60405180910390f35b6101d960048036038101906101d49190610d04565b610781565b6040516101e69190610bbc565b60405180910390f35b61020960048036038101906102049190610b01565b610834565b6040516102169190610bbc565b60405180910390f35b61023960048036038101906102349190610d04565b6108e5565b6040516102469190610bbc565b60405180910390f35b610257610998565b6040516102649190610d8a565b60405180910390f35b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146102fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102f290610e23565b60405180910390fd5b81816001868660405161030f929190610e6f565b9081526020016040518091039020918261032a929190611097565b5050505050565b606060028383604051610345929190610e6f565b9081526020016040518091039020805461035e90610ebe565b80601f016020809104026020016040519081016040528092919081815260200182805461038a90610ebe565b80156103d55780601f106103ac576101008083540402835291602001916103d5565b820191905f5260205f20905b8154815290600101906020018083116103b857829003601f168201915b5050505050905092915050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610470576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161046790610e23565b60405180910390fd5b818160038686604051610484929190610e6f565b9081526020016040518091039020918261049f929190611097565b5050505050565b6003818051602081018201805184825260208301602085012081835280955050505050505f9150905080546104da90610ebe565b80601f016020809104026020016040519081016040528092919081815260200182805461050690610ebe565b80156105515780601f1061052857610100808354040283529160200191610551565b820191905f5260205f20905b81548152906001019060200180831161053457829003601f168201915b505050505081565b60606003838360405161056d929190610e6f565b9081526020016040518091039020805461058690610ebe565b80601f01602080910402602001604051908101604052809291908181526020018280546105b290610ebe565b80156105fd5780601f106105d4576101008083540402835291602001916105fd565b820191905f5260205f20905b8154815290600101906020018083116105e057829003601f168201915b5050505050905092915050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610698576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161068f90610e23565b60405180910390fd5b8181600286866040516106ac929190610e6f565b908152602001604051809103902091826106c7929190611097565b5050505050565b6004818051602081018201805184825260208301602085012081835280955050505050505f91509050805461070290610ebe565b80601f016020809104026020016040519081016040528092919081815260200182805461072e90610ebe565b80156107795780601f1061075057610100808354040283529160200191610779565b820191905f5260205f20905b81548152906001019060200180831161075c57829003601f168201915b505050505081565b6002818051602081018201805184825260208301602085012081835280955050505050505f9150905080546107b590610ebe565b80601f01602080910402602001604051908101604052809291908181526020018280546107e190610ebe565b801561082c5780601f106108035761010080835404028352916020019161082c565b820191905f5260205f20905b81548152906001019060200180831161080f57829003601f168201915b505050505081565b606060018383604051610848929190610e6f565b9081526020016040518091039020805461086190610ebe565b80601f016020809104026020016040519081016040528092919081815260200182805461088d90610ebe565b80156108d85780601f106108af576101008083540402835291602001916108d8565b820191905f5260205f20905b8154815290600101906020018083116108bb57829003601f168201915b5050505050905092915050565b6001818051602081018201805184825260208301602085012081835280955050505050505f91509050805461091990610ebe565b80601f016020809104026020016040519081016040528092919081815260200182805461094590610ebe565b80156109905780601f1061096757610100808354040283529160200191610990565b820191905f5260205f20905b81548152906001019060200180831161097357829003601f168201915b505050505081565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126109ee576109ed6109cd565b5b8235905067ffffffffffffffff811115610a0b57610a0a6109d1565b5b602083019150836001820283011115610a2757610a266109d5565b5b9250929050565b5f5f83601f840112610a4357610a426109cd565b5b8235905067ffffffffffffffff811115610a6057610a5f6109d1565b5b602083019150836001820283011115610a7c57610a7b6109d5565b5b9250929050565b5f5f5f5f60408587031215610a9b57610a9a6109c5565b5b5f85013567ffffffffffffffff811115610ab857610ab76109c9565b5b610ac4878288016109d9565b9450945050602085013567ffffffffffffffff811115610ae757610ae66109c9565b5b610af387828801610a2e565b925092505092959194509250565b5f5f60208385031215610b1757610b166109c5565b5b5f83013567ffffffffffffffff811115610b3457610b336109c9565b5b610b40858286016109d9565b92509250509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610b8e82610b4c565b610b988185610b56565b9350610ba8818560208601610b66565b610bb181610b74565b840191505092915050565b5f6020820190508181035f830152610bd48184610b84565b905092915050565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610c1682610b74565b810181811067ffffffffffffffff82111715610c3557610c34610be0565b5b80604052505050565b5f610c476109bc565b9050610c538282610c0d565b919050565b5f67ffffffffffffffff821115610c7257610c71610be0565b5b610c7b82610b74565b9050602081019050919050565b828183375f83830152505050565b5f610ca8610ca384610c58565b610c3e565b905082815260208101848484011115610cc457610cc3610bdc565b5b610ccf848285610c88565b509392505050565b5f82601f830112610ceb57610cea6109cd565b5b8135610cfb848260208601610c96565b91505092915050565b5f60208284031215610d1957610d186109c5565b5b5f82013567ffffffffffffffff811115610d3657610d356109c9565b5b610d4284828501610cd7565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610d7482610d4b565b9050919050565b610d8481610d6a565b82525050565b5f602082019050610d9d5f830184610d7b565b92915050565b5f82825260208201905092915050565b7f4f6e6c79206f776e657220697320616c6c6f77656420746f20657865637574655f8201527f2074686973207472616e73616374696f6e2e0000000000000000000000000000602082015250565b5f610e0d603283610da3565b9150610e1882610db3565b604082019050919050565b5f6020820190508181035f830152610e3a81610e01565b9050919050565b5f81905092915050565b5f610e568385610e41565b9350610e63838584610c88565b82840190509392505050565b5f610e7b828486610e4b565b91508190509392505050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610ed557607f821691505b602082108103610ee857610ee7610e91565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610f4a7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610f0f565b610f548683610f0f565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f610f98610f93610f8e84610f6c565b610f75565b610f6c565b9050919050565b5f819050919050565b610fb183610f7e565b610fc5610fbd82610f9f565b848454610f1b565b825550505050565b5f5f905090565b610fdc610fcd565b610fe7818484610fa8565b505050565b5b8181101561100a57610fff5f82610fd4565b600181019050610fed565b5050565b601f82111561104f5761102081610eee565b61102984610f00565b81016020851015611038578190505b61104c61104485610f00565b830182610fec565b50505b505050565b5f82821c905092915050565b5f61106f5f1984600802611054565b1980831691505092915050565b5f6110878383611060565b9150826002028217905092915050565b6110a18383610e87565b67ffffffffffffffff8111156110ba576110b9610be0565b5b6110c48254610ebe565b6110cf82828561100e565b5f601f8311600181146110fc575f84156110ea578287013590505b6110f4858261107c565b86555061115b565b601f19841661110a86610eee565b5f5b828110156111315784890135825560018201915060208501945060208101905061110c565b8683101561114e578489013561114a601f891682611060565b8355505b6001600288020188555050505b5050505050505056fea2646970667358221220ad9c84878199ab76d04837b984d9b800b51314dea62f2ccf4f9abb4a45a7a3ed64736f6c634300081e0033",
}

// VolteABI is the input ABI used to generate the binding from.
// Deprecated: Use VolteMetaData.ABI instead.
var VolteABI = VolteMetaData.ABI

// VolteBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VolteMetaData.Bin instead.
var VolteBin = VolteMetaData.Bin

// DeployVolte deploys a new Ethereum contract, binding an instance of Volte to it.
func DeployVolte(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Volte, error) {
	parsed, err := VolteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VolteBin), backend)
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

// GetNullifierMerkleRoot is a free data retrieval call binding the contract method 0xdd7a2ca2.
//
// Solidity: function GetNullifierMerkleRoot(string eventID) view returns(bytes)
func (_Volte *VolteCaller) GetNullifierMerkleRoot(opts *bind.CallOpts, eventID string) ([]byte, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "GetNullifierMerkleRoot", eventID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNullifierMerkleRoot is a free data retrieval call binding the contract method 0xdd7a2ca2.
//
// Solidity: function GetNullifierMerkleRoot(string eventID) view returns(bytes)
func (_Volte *VolteSession) GetNullifierMerkleRoot(eventID string) ([]byte, error) {
	return _Volte.Contract.GetNullifierMerkleRoot(&_Volte.CallOpts, eventID)
}

// GetNullifierMerkleRoot is a free data retrieval call binding the contract method 0xdd7a2ca2.
//
// Solidity: function GetNullifierMerkleRoot(string eventID) view returns(bytes)
func (_Volte *VolteCallerSession) GetNullifierMerkleRoot(eventID string) ([]byte, error) {
	return _Volte.Contract.GetNullifierMerkleRoot(&_Volte.CallOpts, eventID)
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

// NullifierMerkleRoots is a free data retrieval call binding the contract method 0xe1a2b21f.
//
// Solidity: function nullifierMerkleRoots(string ) view returns(bytes)
func (_Volte *VolteCaller) NullifierMerkleRoots(opts *bind.CallOpts, arg0 string) ([]byte, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "nullifierMerkleRoots", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// NullifierMerkleRoots is a free data retrieval call binding the contract method 0xe1a2b21f.
//
// Solidity: function nullifierMerkleRoots(string ) view returns(bytes)
func (_Volte *VolteSession) NullifierMerkleRoots(arg0 string) ([]byte, error) {
	return _Volte.Contract.NullifierMerkleRoots(&_Volte.CallOpts, arg0)
}

// NullifierMerkleRoots is a free data retrieval call binding the contract method 0xe1a2b21f.
//
// Solidity: function nullifierMerkleRoots(string ) view returns(bytes)
func (_Volte *VolteCallerSession) NullifierMerkleRoots(arg0 string) ([]byte, error) {
	return _Volte.Contract.NullifierMerkleRoots(&_Volte.CallOpts, arg0)
}

// UsedNullifiers is a free data retrieval call binding the contract method 0x79488fdd.
//
// Solidity: function usedNullifiers(string ) view returns(bytes)
func (_Volte *VolteCaller) UsedNullifiers(opts *bind.CallOpts, arg0 string) ([]byte, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "usedNullifiers", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// UsedNullifiers is a free data retrieval call binding the contract method 0x79488fdd.
//
// Solidity: function usedNullifiers(string ) view returns(bytes)
func (_Volte *VolteSession) UsedNullifiers(arg0 string) ([]byte, error) {
	return _Volte.Contract.UsedNullifiers(&_Volte.CallOpts, arg0)
}

// UsedNullifiers is a free data retrieval call binding the contract method 0x79488fdd.
//
// Solidity: function usedNullifiers(string ) view returns(bytes)
func (_Volte *VolteCallerSession) UsedNullifiers(arg0 string) ([]byte, error) {
	return _Volte.Contract.UsedNullifiers(&_Volte.CallOpts, arg0)
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

// SetNullifierMerkleRoot is a paid mutator transaction binding the contract method 0x2acceddd.
//
// Solidity: function SetNullifierMerkleRoot(string eventID, bytes value) returns()
func (_Volte *VolteTransactor) SetNullifierMerkleRoot(opts *bind.TransactOpts, eventID string, value []byte) (*types.Transaction, error) {
	return _Volte.contract.Transact(opts, "SetNullifierMerkleRoot", eventID, value)
}

// SetNullifierMerkleRoot is a paid mutator transaction binding the contract method 0x2acceddd.
//
// Solidity: function SetNullifierMerkleRoot(string eventID, bytes value) returns()
func (_Volte *VolteSession) SetNullifierMerkleRoot(eventID string, value []byte) (*types.Transaction, error) {
	return _Volte.Contract.SetNullifierMerkleRoot(&_Volte.TransactOpts, eventID, value)
}

// SetNullifierMerkleRoot is a paid mutator transaction binding the contract method 0x2acceddd.
//
// Solidity: function SetNullifierMerkleRoot(string eventID, bytes value) returns()
func (_Volte *VolteTransactorSession) SetNullifierMerkleRoot(eventID string, value []byte) (*types.Transaction, error) {
	return _Volte.Contract.SetNullifierMerkleRoot(&_Volte.TransactOpts, eventID, value)
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
