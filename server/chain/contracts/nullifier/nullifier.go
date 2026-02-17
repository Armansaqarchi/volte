// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nullifier

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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"_pA\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"_pB\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"_pC\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[3]\",\"name\":\"_pubSignals\",\"type\":\"uint256[3]\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b5061069d8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610029575f3560e01c806311479fea1461002d575b5f5ffd5b610047600480360381019061004291906105ce565b61005d565b604051610054919061064e565b60405180910390f35b5f61051d565b7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000018110610092575f5f5260205ff35b50565b5f60405183815284602082015285604082015260408160608360076107d05a03fa9150816100c5575f5f5260205ff35b825160408201526020830151606082015260408360808360066107d05a03fa9150816100f3575f5f5260205ff35b505050505050565b5f608086015f87017f1302fca42f734644fb973c48ef884338fab0a5bc0cea8c1cfc3dc7471ecceba381527f2713a42ca4f5f6dffacdf9ade1bfdc96feabee5379a12d542b3e31f23ba59559602082015261019b5f8801357f1ddfa8c25e7dd8ae45137a17c9a8675bf96ac492bab0e0c30eab209153fa162f7f0bdbb403ccc424bcb38ff725ea78e2067e3b7807943b9fb5c053c6d219c308f284610095565b6101eb60208801357f2cdcc5c8326dc4547f08f47ecfcd64d3fcb77b4ec74617104e89c627e3aeaeb37f0bd42bd6bc0ab37b5a264cad11a4bc0e1a423f429fb025443f2ebfd8b73f827784610095565b61023b60408801357f0cadbdfc4d81a70df52b7aeac6ff35755630517f9a51911dedcb75270728e5f07f15b5fb9ebcb5a536a5d360d9ef05e6b3686a88839cb735cf90260d9eb2d0b02a84610095565b833582527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760208501357f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4703066020830152843560408301526020850135606083015260408501356080830152606085013560a08301527f2130b8c9661d6f75c92158aff45c4aed604bb13a523a397af5c56f9753ded94960c08301527f2a1c5b233df16707a4a18b4cedcff32daf316bf9215bccd1351c6a958dd1952a60e08301527f0251a0626d478c29cdb60f044d49a7e82729518d153c11fd6ec0da37345ed3f56101008301527f0d261e37e3b02c30bf5f87046260e72fec8250822d4084e42bf6fbd9badd6b406101208301527f15cf832f3f7c1b7950befab926ecc6bfc0dbb784e5f13b4572edb425f5f19fab6101408301527f28990dcd55f60b11799b320da5dc3c06c48afdfc19c86f13cb51d99dcd7d48156101608301525f88015161018083015260205f018801516101a08301527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c26101c08301527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6101e08301527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b6102008301527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa610220830152853561024083015260208601356102608301527f2b676132d45ae309ffd1ab9938be043d42a371925a56a7d5f7d64667b80157c66102808301527f08c195782636e7cf6ff2ff9243ac34a307da1cb3d40ad11b7366fcdb958604156102a08301527f11bc66342246e2d4150866adeaa82f12d31bbe816a4f484073e2fbc1b5f9401c6102c08301527f0ac43b1b95450b548b4eec2bcbd098f3dfb6b6b1faf1bbacbe215e4d15afeaa66102e08301526020826103008460086107d05a03fa82518116935050505095945050505050565b60405161038081016040526105345f840135610063565b6105416020840135610063565b61054e6040840135610063565b61055b818486888a6100fb565b805f5260205ff35b5f5ffd5b5f5ffd5b5f8190508260206002028201111561058657610585610567565b5b92915050565b5f819050826040600202820111156105a7576105a6610567565b5b92915050565b5f819050826020600302820111156105c8576105c7610567565b5b92915050565b5f5f5f5f61016085870312156105e7576105e6610563565b5b5f6105f48782880161056b565b94505060406106058782880161058c565b93505060c06106168782880161056b565b925050610100610628878288016105ad565b91505092959194509250565b5f8115159050919050565b61064881610634565b82525050565b5f6020820190506106615f83018461063f565b9291505056fea2646970667358221220a41fe9166acf5da6842aee82da19a578027eaddd83bdc40804ac81da0b9757f964736f6c634300081e0033",
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

// VerifyProof is a free data retrieval call binding the contract method 0x11479fea.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[3] _pubSignals) view returns(bool)
func (_Volte *VolteCaller) VerifyProof(opts *bind.CallOpts, _pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [3]*big.Int) (bool, error) {
	var out []interface{}
	err := _Volte.contract.Call(opts, &out, "verifyProof", _pA, _pB, _pC, _pubSignals)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x11479fea.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[3] _pubSignals) view returns(bool)
func (_Volte *VolteSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [3]*big.Int) (bool, error) {
	return _Volte.Contract.VerifyProof(&_Volte.CallOpts, _pA, _pB, _pC, _pubSignals)
}

// VerifyProof is a free data retrieval call binding the contract method 0x11479fea.
//
// Solidity: function verifyProof(uint256[2] _pA, uint256[2][2] _pB, uint256[2] _pC, uint256[3] _pubSignals) view returns(bool)
func (_Volte *VolteCallerSession) VerifyProof(_pA [2]*big.Int, _pB [2][2]*big.Int, _pC [2]*big.Int, _pubSignals [3]*big.Int) (bool, error) {
	return _Volte.Contract.VerifyProof(&_Volte.CallOpts, _pA, _pB, _pC, _pubSignals)
}
