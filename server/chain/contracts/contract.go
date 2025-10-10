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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetEventHash\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetNullifierMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"GetVoteMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetEventHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetNullifierMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"SetVoteMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"eventID\",\"type\":\"string\"}],\"name\":\"Vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"eventHashes\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"nullifierMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"usedNullifiers\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"voteMerkleRoots\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b50335f5f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506112128061005b5f395ff3fe608060405234801561000f575f5ffd5b50600436106100b2575f3560e01c8063649a75ca1161006f578063649a75ca146101ae57806379488fdd146101ca5780639b28fdf0146101fa578063dd7a2ca21461022a578063e1a2b21f1461025a578063f851a4401461028a576100b2565b80632acceddd146100b6578063323082d7146100d257806337f5daab146101025780633aa82fee1461013257806341260b3d1461014e5780635dd5ed281461017e575b5f5ffd5b6100d060048036038101906100cb9190610ac8565b6102a8565b005b6100ec60048036038101906100e79190610b46565b61036c565b6040516100f99190610bab565b60405180910390f35b61011c60048036038101906101179190610b46565b610376565b6040516101299190610c34565b60405180910390f35b61014c60048036038101906101479190610ac8565b610427565b005b61016860048036038101906101639190610d7c565b6104eb565b6040516101759190610c34565b60405180910390f35b61019860048036038101906101939190610b46565b61059e565b6040516101a59190610c34565b60405180910390f35b6101c860048036038101906101c39190610ac8565b61064f565b005b6101e460048036038101906101df9190610d7c565b610713565b6040516101f19190610c34565b60405180910390f35b610214600480360381019061020f9190610d7c565b6107c6565b6040516102219190610c34565b60405180910390f35b610244600480360381019061023f9190610b46565b610879565b6040516102519190610c34565b60405180910390f35b610274600480360381019061026f9190610d7c565b61092a565b6040516102819190610c34565b60405180910390f35b6102926109dd565b60405161029f9190610e02565b60405180910390f35b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610336576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032d90610e9b565b60405180910390fd5b81816001868660405161034a929190610ee7565b9081526020016040518091039020918261036592919061110f565b5050505050565b5f5f905092915050565b60606002838360405161038a929190610ee7565b908152602001604051809103902080546103a390610f36565b80601f01602080910402602001604051908101604052809291908181526020018280546103cf90610f36565b801561041a5780601f106103f15761010080835404028352916020019161041a565b820191905f5260205f20905b8154815290600101906020018083116103fd57829003601f168201915b5050505050905092915050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104b5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ac90610e9b565b60405180910390fd5b8181600386866040516104c9929190610ee7565b908152602001604051809103902091826104e492919061110f565b5050505050565b6003818051602081018201805184825260208301602085012081835280955050505050505f91509050805461051f90610f36565b80601f016020809104026020016040519081016040528092919081815260200182805461054b90610f36565b80156105965780601f1061056d57610100808354040283529160200191610596565b820191905f5260205f20905b81548152906001019060200180831161057957829003601f168201915b505050505081565b6060600383836040516105b2929190610ee7565b908152602001604051809103902080546105cb90610f36565b80601f01602080910402602001604051908101604052809291908181526020018280546105f790610f36565b80156106425780601f1061061957610100808354040283529160200191610642565b820191905f5260205f20905b81548152906001019060200180831161062557829003601f168201915b5050505050905092915050565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d490610e9b565b60405180910390fd5b8181600286866040516106f1929190610ee7565b9081526020016040518091039020918261070c92919061110f565b5050505050565b6004818051602081018201805184825260208301602085012081835280955050505050505f91509050805461074790610f36565b80601f016020809104026020016040519081016040528092919081815260200182805461077390610f36565b80156107be5780601f10610795576101008083540402835291602001916107be565b820191905f5260205f20905b8154815290600101906020018083116107a157829003601f168201915b505050505081565b6002818051602081018201805184825260208301602085012081835280955050505050505f9150905080546107fa90610f36565b80601f016020809104026020016040519081016040528092919081815260200182805461082690610f36565b80156108715780601f1061084857610100808354040283529160200191610871565b820191905f5260205f20905b81548152906001019060200180831161085457829003601f168201915b505050505081565b60606001838360405161088d929190610ee7565b908152602001604051809103902080546108a690610f36565b80601f01602080910402602001604051908101604052809291908181526020018280546108d290610f36565b801561091d5780601f106108f45761010080835404028352916020019161091d565b820191905f5260205f20905b81548152906001019060200180831161090057829003601f168201915b5050505050905092915050565b6001818051602081018201805184825260208301602085012081835280955050505050505f91509050805461095e90610f36565b80601f016020809104026020016040519081016040528092919081815260200182805461098a90610f36565b80156109d55780601f106109ac576101008083540402835291602001916109d5565b820191905f5260205f20905b8154815290600101906020018083116109b857829003601f168201915b505050505081565b5f5f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610a3357610a32610a12565b5b8235905067ffffffffffffffff811115610a5057610a4f610a16565b5b602083019150836001820283011115610a6c57610a6b610a1a565b5b9250929050565b5f5f83601f840112610a8857610a87610a12565b5b8235905067ffffffffffffffff811115610aa557610aa4610a16565b5b602083019150836001820283011115610ac157610ac0610a1a565b5b9250929050565b5f5f5f5f60408587031215610ae057610adf610a0a565b5b5f85013567ffffffffffffffff811115610afd57610afc610a0e565b5b610b0987828801610a1e565b9450945050602085013567ffffffffffffffff811115610b2c57610b2b610a0e565b5b610b3887828801610a73565b925092505092959194509250565b5f5f60208385031215610b5c57610b5b610a0a565b5b5f83013567ffffffffffffffff811115610b7957610b78610a0e565b5b610b8585828601610a1e565b92509250509250929050565b5f8115159050919050565b610ba581610b91565b82525050565b5f602082019050610bbe5f830184610b9c565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610c0682610bc4565b610c108185610bce565b9350610c20818560208601610bde565b610c2981610bec565b840191505092915050565b5f6020820190508181035f830152610c4c8184610bfc565b905092915050565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610c8e82610bec565b810181811067ffffffffffffffff82111715610cad57610cac610c58565b5b80604052505050565b5f610cbf610a01565b9050610ccb8282610c85565b919050565b5f67ffffffffffffffff821115610cea57610ce9610c58565b5b610cf382610bec565b9050602081019050919050565b828183375f83830152505050565b5f610d20610d1b84610cd0565b610cb6565b905082815260208101848484011115610d3c57610d3b610c54565b5b610d47848285610d00565b509392505050565b5f82601f830112610d6357610d62610a12565b5b8135610d73848260208601610d0e565b91505092915050565b5f60208284031215610d9157610d90610a0a565b5b5f82013567ffffffffffffffff811115610dae57610dad610a0e565b5b610dba84828501610d4f565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610dec82610dc3565b9050919050565b610dfc81610de2565b82525050565b5f602082019050610e155f830184610df3565b92915050565b5f82825260208201905092915050565b7f4f6e6c79206f776e657220697320616c6c6f77656420746f20657865637574655f8201527f2074686973207472616e73616374696f6e2e0000000000000000000000000000602082015250565b5f610e85603283610e1b565b9150610e9082610e2b565b604082019050919050565b5f6020820190508181035f830152610eb281610e79565b9050919050565b5f81905092915050565b5f610ece8385610eb9565b9350610edb838584610d00565b82840190509392505050565b5f610ef3828486610ec3565b91508190509392505050565b5f82905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610f4d57607f821691505b602082108103610f6057610f5f610f09565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610fc27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610f87565b610fcc8683610f87565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61101061100b61100684610fe4565b610fed565b610fe4565b9050919050565b5f819050919050565b61102983610ff6565b61103d61103582611017565b848454610f93565b825550505050565b5f5f905090565b611054611045565b61105f818484611020565b505050565b5b81811015611082576110775f8261104c565b600181019050611065565b5050565b601f8211156110c75761109881610f66565b6110a184610f78565b810160208510156110b0578190505b6110c46110bc85610f78565b830182611064565b50505b505050565b5f82821c905092915050565b5f6110e75f19846008026110cc565b1980831691505092915050565b5f6110ff83836110d8565b9150826002028217905092915050565b6111198383610eff565b67ffffffffffffffff81111561113257611131610c58565b5b61113c8254610f36565b611147828285611086565b5f601f831160018114611174575f8415611162578287013590505b61116c85826110f4565b8655506111d3565b601f19841661118286610f66565b5f5b828110156111a957848901358255600182019150602085019450602081019050611184565b868310156111c657848901356111c2601f8916826110d8565b8355505b6001600288020188555050505b5050505050505056fea264697066735822122068dde24b4015b26ff30dc203a8ec13c7affdb871142b42f32500df8416fc0ea364736f6c634300081e0033",
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

// Vote is a paid mutator transaction binding the contract method 0x323082d7.
//
// Solidity: function Vote(string eventID) returns(bool)
func (_Volte *VolteTransactor) Vote(opts *bind.TransactOpts, eventID string) (*types.Transaction, error) {
	return _Volte.contract.Transact(opts, "Vote", eventID)
}

// Vote is a paid mutator transaction binding the contract method 0x323082d7.
//
// Solidity: function Vote(string eventID) returns(bool)
func (_Volte *VolteSession) Vote(eventID string) (*types.Transaction, error) {
	return _Volte.Contract.Vote(&_Volte.TransactOpts, eventID)
}

// Vote is a paid mutator transaction binding the contract method 0x323082d7.
//
// Solidity: function Vote(string eventID) returns(bool)
func (_Volte *VolteTransactorSession) Vote(eventID string) (*types.Transaction, error) {
	return _Volte.Contract.Vote(&_Volte.TransactOpts, eventID)
}
