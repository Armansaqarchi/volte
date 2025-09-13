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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"}],\"name\":\"GetEventHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"}],\"name\":\"GetNullifierMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"}],\"name\":\"GetVoteMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetEventHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetNullifierMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"eventID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetVoteMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"eventHashes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nullifierMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"voteMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610e078061005b5f395ff3fe608060405234801561000f575f5ffd5b506004361061009c575f3560e01c8063a873ecee11610064578063a873ecee1461017c578063b0e2eccf14610198578063f30436f7146101b4578063f4c9d1e1146101e4578063f851a440146102145761009c565b806307acde36146100a05780632be94db7146100d057806355a67f8d146100ec57806365d108341461011c5780639d6f47031461014c575b5f5ffd5b6100ba60048036038101906100b59190610861565b610232565b6040516100c791906108fc565b60405180910390f35b6100ea60048036038101906100e5919061097d565b6102d3565b005b61010660048036038101906101019190610861565b610387565b60405161011391906108fc565b60405180910390f35b61013660048036038101906101319190610861565b610422565b60405161014391906108fc565b60405180910390f35b61016660048036038101906101619190610861565b6104bd565b60405161017391906108fc565b60405180910390f35b6101966004803603810190610191919061097d565b610558565b005b6101b260048036038101906101ad919061097d565b61060c565b005b6101ce60048036038101906101c99190610861565b6106c0565b6040516101db91906108fc565b60405180910390f35b6101fe60048036038101906101f99190610861565b610761565b60405161020b91906108fc565b60405180910390f35b61021c610802565b6040516102299190610a19565b60405180910390f35b606060025f8381526020019081526020015f20805461025090610a5f565b80601f016020809104026020016040519081016040528092919081815260200182805461027c90610a5f565b80156102c75780601f1061029e576101008083540402835291602001916102c7565b820191905f5260205f20905b8154815290600101906020018083116102aa57829003601f168201915b50505050509050919050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610361576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035890610b0f565b60405180910390fd5b818160015f8681526020019081526020015f209182610381929190610d04565b50505050565b6002602052805f5260405f205f9150905080546103a390610a5f565b80601f01602080910402602001604051908101604052809291908181526020018280546103cf90610a5f565b801561041a5780601f106103f15761010080835404028352916020019161041a565b820191905f5260205f20905b8154815290600101906020018083116103fd57829003601f168201915b505050505081565b6001602052805f5260405f205f91509050805461043e90610a5f565b80601f016020809104026020016040519081016040528092919081815260200182805461046a90610a5f565b80156104b55780601f1061048c576101008083540402835291602001916104b5565b820191905f5260205f20905b81548152906001019060200180831161049857829003601f168201915b505050505081565b6003602052805f5260405f205f9150905080546104d990610a5f565b80601f016020809104026020016040519081016040528092919081815260200182805461050590610a5f565b80156105505780601f1061052757610100808354040283529160200191610550565b820191905f5260205f20905b81548152906001019060200180831161053357829003601f168201915b505050505081565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105dd90610b0f565b60405180910390fd5b818160035f8681526020019081526020015f209182610606929190610d04565b50505050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461069a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069190610b0f565b60405180910390fd5b818160025f8681526020019081526020015f2091826106ba929190610d04565b50505050565b606060035f8381526020019081526020015f2080546106de90610a5f565b80601f016020809104026020016040519081016040528092919081815260200182805461070a90610a5f565b80156107555780601f1061072c57610100808354040283529160200191610755565b820191905f5260205f20905b81548152906001019060200180831161073857829003601f168201915b50505050509050919050565b606060015f8381526020019081526020015f20805461077f90610a5f565b80601f01602080910402602001604051908101604052809291908181526020018280546107ab90610a5f565b80156107f65780601f106107cd576101008083540402835291602001916107f6565b820191905f5260205f20905b8154815290600101906020018083116107d957829003601f168201915b50505050509050919050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f5ffd5b5f5ffd5b5f819050919050565b6108408161082e565b811461084a575f5ffd5b50565b5f8135905061085b81610837565b92915050565b5f6020828403121561087657610875610826565b5b5f6108838482850161084d565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6108ce8261088c565b6108d88185610896565b93506108e88185602086016108a6565b6108f1816108b4565b840191505092915050565b5f6020820190508181035f83015261091481846108c4565b905092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261093d5761093c61091c565b5b8235905067ffffffffffffffff81111561095a57610959610920565b5b60208301915083600182028301111561097657610975610924565b5b9250929050565b5f5f5f6040848603121561099457610993610826565b5b5f6109a18682870161084d565b935050602084013567ffffffffffffffff8111156109c2576109c161082a565b5b6109ce86828701610928565b92509250509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610a03826109da565b9050919050565b610a13816109f9565b82525050565b5f602082019050610a2c5f830184610a0a565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610a7657607f821691505b602082108103610a8957610a88610a32565b5b50919050565b5f82825260208201905092915050565b7f4f6e6c79206f776e657220697320616c6c6f77656420746f20657865637574655f8201527f2074686973207472616e73616374696f6e2e0000000000000000000000000000602082015250565b5f610af9603283610a8f565b9150610b0482610a9f565b604082019050919050565b5f6020820190508181035f830152610b2681610aed565b9050919050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610bc07fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610b85565b610bca8683610b85565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610c05610c00610bfb8461082e565b610be2565b61082e565b9050919050565b5f819050919050565b610c1e83610beb565b610c32610c2a82610c0c565b848454610b91565b825550505050565b5f5f905090565b610c49610c3a565b610c54818484610c15565b505050565b5b81811015610c7757610c6c5f82610c41565b600181019050610c5a565b5050565b601f821115610cbc57610c8d81610b64565b610c9684610b76565b81016020851015610ca5578190505b610cb9610cb185610b76565b830182610c59565b50505b505050565b5f82821c905092915050565b5f610cdc5f1984600802610cc1565b1980831691505092915050565b5f610cf48383610ccd565b9150826002028217905092915050565b610d0e8383610b2d565b67ffffffffffffffff811115610d2757610d26610b37565b5b610d318254610a5f565b610d3c828285610c7b565b5f601f831160018114610d69575f8415610d57578287013590505b610d618582610ce9565b865550610dc8565b601f198416610d7786610b64565b5f5b82811015610d9e57848901358255600182019150602085019450602081019050610d79565b86831015610dbb5784890135610db7601f891682610ccd565b8355505b6001600288020188555050505b5050505050505056fea2646970667358221220ca4018ca1fe98557aa65d54160382f9b94ed90e88f96d332793bba7420925e7b64736f6c634300081e0033",
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

// GetEventHash is a free data retrieval call binding the contract method 0xf30436f7.
//
// Solidity: function GetEventHash(uint256 eventID) view returns(bytes)
func (_Volte *VolteCaller) GetEventHash(opts *bind.CallOpts, eventID *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "GetEventHash", eventID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetEventHash is a free data retrieval call binding the contract method 0xf30436f7.
//
// Solidity: function GetEventHash(uint256 eventID) view returns(bytes)
func (_Volte *VolteSession) GetEventHash(eventID *big.Int) ([]byte, error) {
	return _Volte.Contract.GetEventHash(&_Volte.CallOpts, eventID)
}

// GetEventHash is a free data retrieval call binding the contract method 0xf30436f7.
//
// Solidity: function GetEventHash(uint256 eventID) view returns(bytes)
func (_Volte *VolteCallerSession) GetEventHash(eventID *big.Int) ([]byte, error) {
	return _Volte.Contract.GetEventHash(&_Volte.CallOpts, eventID)
}

// GetNullifierMerkleRoot is a free data retrieval call binding the contract method 0xf4c9d1e1.
//
// Solidity: function GetNullifierMerkleRoot(uint256 eventID) view returns(bytes)
func (_Volte *VolteCaller) GetNullifierMerkleRoot(opts *bind.CallOpts, eventID *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "GetNullifierMerkleRoot", eventID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetNullifierMerkleRoot is a free data retrieval call binding the contract method 0xf4c9d1e1.
//
// Solidity: function GetNullifierMerkleRoot(uint256 eventID) view returns(bytes)
func (_Volte *VolteSession) GetNullifierMerkleRoot(eventID *big.Int) ([]byte, error) {
	return _Volte.Contract.GetNullifierMerkleRoot(&_Volte.CallOpts, eventID)
}

// GetNullifierMerkleRoot is a free data retrieval call binding the contract method 0xf4c9d1e1.
//
// Solidity: function GetNullifierMerkleRoot(uint256 eventID) view returns(bytes)
func (_Volte *VolteCallerSession) GetNullifierMerkleRoot(eventID *big.Int) ([]byte, error) {
	return _Volte.Contract.GetNullifierMerkleRoot(&_Volte.CallOpts, eventID)
}

// GetVoteMerkleRoot is a free data retrieval call binding the contract method 0x07acde36.
//
// Solidity: function GetVoteMerkleRoot(uint256 eventID) view returns(bytes)
func (_Volte *VolteCaller) GetVoteMerkleRoot(opts *bind.CallOpts, eventID *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "GetVoteMerkleRoot", eventID)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetVoteMerkleRoot is a free data retrieval call binding the contract method 0x07acde36.
//
// Solidity: function GetVoteMerkleRoot(uint256 eventID) view returns(bytes)
func (_Volte *VolteSession) GetVoteMerkleRoot(eventID *big.Int) ([]byte, error) {
	return _Volte.Contract.GetVoteMerkleRoot(&_Volte.CallOpts, eventID)
}

// GetVoteMerkleRoot is a free data retrieval call binding the contract method 0x07acde36.
//
// Solidity: function GetVoteMerkleRoot(uint256 eventID) view returns(bytes)
func (_Volte *VolteCallerSession) GetVoteMerkleRoot(eventID *big.Int) ([]byte, error) {
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

// EventHashes is a free data retrieval call binding the contract method 0x9d6f4703.
//
// Solidity: function eventHashes(uint256 ) view returns(bytes)
func (_Volte *VolteCaller) EventHashes(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "eventHashes", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// EventHashes is a free data retrieval call binding the contract method 0x9d6f4703.
//
// Solidity: function eventHashes(uint256 ) view returns(bytes)
func (_Volte *VolteSession) EventHashes(arg0 *big.Int) ([]byte, error) {
	return _Volte.Contract.EventHashes(&_Volte.CallOpts, arg0)
}

// EventHashes is a free data retrieval call binding the contract method 0x9d6f4703.
//
// Solidity: function eventHashes(uint256 ) view returns(bytes)
func (_Volte *VolteCallerSession) EventHashes(arg0 *big.Int) ([]byte, error) {
	return _Volte.Contract.EventHashes(&_Volte.CallOpts, arg0)
}

// NullifierMerkleRoots is a free data retrieval call binding the contract method 0x65d10834.
//
// Solidity: function nullifierMerkleRoots(uint256 ) view returns(bytes)
func (_Volte *VolteCaller) NullifierMerkleRoots(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "nullifierMerkleRoots", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// NullifierMerkleRoots is a free data retrieval call binding the contract method 0x65d10834.
//
// Solidity: function nullifierMerkleRoots(uint256 ) view returns(bytes)
func (_Volte *VolteSession) NullifierMerkleRoots(arg0 *big.Int) ([]byte, error) {
	return _Volte.Contract.NullifierMerkleRoots(&_Volte.CallOpts, arg0)
}

// NullifierMerkleRoots is a free data retrieval call binding the contract method 0x65d10834.
//
// Solidity: function nullifierMerkleRoots(uint256 ) view returns(bytes)
func (_Volte *VolteCallerSession) NullifierMerkleRoots(arg0 *big.Int) ([]byte, error) {
	return _Volte.Contract.NullifierMerkleRoots(&_Volte.CallOpts, arg0)
}

// VoteMerkleRoots is a free data retrieval call binding the contract method 0x55a67f8d.
//
// Solidity: function voteMerkleRoots(uint256 ) view returns(bytes)
func (_Volte *VolteCaller) VoteMerkleRoots(opts *bind.CallOpts, arg0 *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "voteMerkleRoots", arg0)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// VoteMerkleRoots is a free data retrieval call binding the contract method 0x55a67f8d.
//
// Solidity: function voteMerkleRoots(uint256 ) view returns(bytes)
func (_Volte *VolteSession) VoteMerkleRoots(arg0 *big.Int) ([]byte, error) {
	return _Volte.Contract.VoteMerkleRoots(&_Volte.CallOpts, arg0)
}

// VoteMerkleRoots is a free data retrieval call binding the contract method 0x55a67f8d.
//
// Solidity: function voteMerkleRoots(uint256 ) view returns(bytes)
func (_Volte *VolteCallerSession) VoteMerkleRoots(arg0 *big.Int) ([]byte, error) {
	return _Volte.Contract.VoteMerkleRoots(&_Volte.CallOpts, arg0)
}

// SetEventHash is a paid mutator transaction binding the contract method 0xa873ecee.
//
// Solidity: function SetEventHash(uint256 eventID, bytes value) returns()
func (_Volte *VolteTransactor) SetEventHash(opts *bind.TransactOpts, eventID *big.Int, value []byte) (*types.Transaction, error) {
	return _Volte.contract.Transact(opts, "SetEventHash", eventID, value)
}

// SetEventHash is a paid mutator transaction binding the contract method 0xa873ecee.
//
// Solidity: function SetEventHash(uint256 eventID, bytes value) returns()
func (_Volte *VolteSession) SetEventHash(eventID *big.Int, value []byte) (*types.Transaction, error) {
	return _Volte.Contract.SetEventHash(&_Volte.TransactOpts, eventID, value)
}

// SetEventHash is a paid mutator transaction binding the contract method 0xa873ecee.
//
// Solidity: function SetEventHash(uint256 eventID, bytes value) returns()
func (_Volte *VolteTransactorSession) SetEventHash(eventID *big.Int, value []byte) (*types.Transaction, error) {
	return _Volte.Contract.SetEventHash(&_Volte.TransactOpts, eventID, value)
}

// SetNullifierMerkleRoot is a paid mutator transaction binding the contract method 0x2be94db7.
//
// Solidity: function SetNullifierMerkleRoot(uint256 eventID, bytes value) returns()
func (_Volte *VolteTransactor) SetNullifierMerkleRoot(opts *bind.TransactOpts, eventID *big.Int, value []byte) (*types.Transaction, error) {
	return _Volte.contract.Transact(opts, "SetNullifierMerkleRoot", eventID, value)
}

// SetNullifierMerkleRoot is a paid mutator transaction binding the contract method 0x2be94db7.
//
// Solidity: function SetNullifierMerkleRoot(uint256 eventID, bytes value) returns()
func (_Volte *VolteSession) SetNullifierMerkleRoot(eventID *big.Int, value []byte) (*types.Transaction, error) {
	return _Volte.Contract.SetNullifierMerkleRoot(&_Volte.TransactOpts, eventID, value)
}

// SetNullifierMerkleRoot is a paid mutator transaction binding the contract method 0x2be94db7.
//
// Solidity: function SetNullifierMerkleRoot(uint256 eventID, bytes value) returns()
func (_Volte *VolteTransactorSession) SetNullifierMerkleRoot(eventID *big.Int, value []byte) (*types.Transaction, error) {
	return _Volte.Contract.SetNullifierMerkleRoot(&_Volte.TransactOpts, eventID, value)
}

// SetVoteMerkleRoot is a paid mutator transaction binding the contract method 0xb0e2eccf.
//
// Solidity: function SetVoteMerkleRoot(uint256 eventID, bytes value) returns()
func (_Volte *VolteTransactor) SetVoteMerkleRoot(opts *bind.TransactOpts, eventID *big.Int, value []byte) (*types.Transaction, error) {
	return _Volte.contract.Transact(opts, "SetVoteMerkleRoot", eventID, value)
}

// SetVoteMerkleRoot is a paid mutator transaction binding the contract method 0xb0e2eccf.
//
// Solidity: function SetVoteMerkleRoot(uint256 eventID, bytes value) returns()
func (_Volte *VolteSession) SetVoteMerkleRoot(eventID *big.Int, value []byte) (*types.Transaction, error) {
	return _Volte.Contract.SetVoteMerkleRoot(&_Volte.TransactOpts, eventID, value)
}

// SetVoteMerkleRoot is a paid mutator transaction binding the contract method 0xb0e2eccf.
//
// Solidity: function SetVoteMerkleRoot(uint256 eventID, bytes value) returns()
func (_Volte *VolteTransactorSession) SetVoteMerkleRoot(eventID *big.Int, value []byte) (*types.Transaction, error) {
	return _Volte.Contract.SetVoteMerkleRoot(&_Volte.TransactOpts, eventID, value)
}
