// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// DeployedABI is the input ABI used to generate the binding from.
const DeployedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"a\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"setA\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DeployedBin is the compiled bytecode used for deploying new contracts.
const DeployedBin = `0x608060405234801561001057600080fd5b5060bc8061001f6000396000f3fe6080604052348015600f57600080fd5b5060043610604e577c010000000000000000000000000000000000000000000000000000000060003504630dbe671f81146053578063ee919d5014606b575b600080fd5b60596085565b60408051918252519081900360200190f35b605960048036036020811015607f57600080fd5b5035608a565b600090565b5060009056fea165627a7a72305820949a524de0f22528aba44cd79fd15397231334cabe0a31928e53656ebc496b7f0029`

// DeployDeployed deploys a new Ethereum contract, binding an instance of Deployed to it.
func DeployDeployed(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Deployed, error) {
	parsed, err := abi.JSON(strings.NewReader(DeployedABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DeployedBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Deployed{DeployedCaller: DeployedCaller{contract: contract}, DeployedTransactor: DeployedTransactor{contract: contract}, DeployedFilterer: DeployedFilterer{contract: contract}}, nil
}

// Deployed is an auto generated Go binding around an Ethereum contract.
type Deployed struct {
	DeployedCaller     // Read-only binding to the contract
	DeployedTransactor // Write-only binding to the contract
	DeployedFilterer   // Log filterer for contract events
}

// DeployedCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeployedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeployedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeployedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeployedSession struct {
	Contract     *Deployed         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeployedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeployedCallerSession struct {
	Contract *DeployedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DeployedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeployedTransactorSession struct {
	Contract     *DeployedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DeployedRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeployedRaw struct {
	Contract *Deployed // Generic contract binding to access the raw methods on
}

// DeployedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeployedCallerRaw struct {
	Contract *DeployedCaller // Generic read-only contract binding to access the raw methods on
}

// DeployedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeployedTransactorRaw struct {
	Contract *DeployedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeployed creates a new instance of Deployed, bound to a specific deployed contract.
func NewDeployed(address common.Address, backend bind.ContractBackend) (*Deployed, error) {
	contract, err := bindDeployed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deployed{DeployedCaller: DeployedCaller{contract: contract}, DeployedTransactor: DeployedTransactor{contract: contract}, DeployedFilterer: DeployedFilterer{contract: contract}}, nil
}

// NewDeployedCaller creates a new read-only instance of Deployed, bound to a specific deployed contract.
func NewDeployedCaller(address common.Address, caller bind.ContractCaller) (*DeployedCaller, error) {
	contract, err := bindDeployed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeployedCaller{contract: contract}, nil
}

// NewDeployedTransactor creates a new write-only instance of Deployed, bound to a specific deployed contract.
func NewDeployedTransactor(address common.Address, transactor bind.ContractTransactor) (*DeployedTransactor, error) {
	contract, err := bindDeployed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeployedTransactor{contract: contract}, nil
}

// NewDeployedFilterer creates a new log filterer instance of Deployed, bound to a specific deployed contract.
func NewDeployedFilterer(address common.Address, filterer bind.ContractFilterer) (*DeployedFilterer, error) {
	contract, err := bindDeployed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeployedFilterer{contract: contract}, nil
}

// bindDeployed binds a generic wrapper to an already deployed contract.
func bindDeployed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeployedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deployed *DeployedRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Deployed.Contract.DeployedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deployed *DeployedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployed.Contract.DeployedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deployed *DeployedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deployed.Contract.DeployedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deployed *DeployedCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Deployed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deployed *DeployedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deployed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deployed *DeployedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deployed.Contract.contract.Transact(opts, method, params...)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() constant returns(uint256)
func (_Deployed *DeployedCaller) A(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Deployed.contract.Call(opts, out, "a")
	return *ret0, err
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() constant returns(uint256)
func (_Deployed *DeployedSession) A() (*big.Int, error) {
	return _Deployed.Contract.A(&_Deployed.CallOpts)
}

// A is a free data retrieval call binding the contract method 0x0dbe671f.
//
// Solidity: function a() constant returns(uint256)
func (_Deployed *DeployedCallerSession) A() (*big.Int, error) {
	return _Deployed.Contract.A(&_Deployed.CallOpts)
}

// SetA is a paid mutator transaction binding the contract method 0xee919d50.
//
// Solidity: function setA(uint256 ) returns(uint256)
func (_Deployed *DeployedTransactor) SetA(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Deployed.contract.Transact(opts, "setA", arg0)
}

// SetA is a paid mutator transaction binding the contract method 0xee919d50.
//
// Solidity: function setA(uint256 ) returns(uint256)
func (_Deployed *DeployedSession) SetA(arg0 *big.Int) (*types.Transaction, error) {
	return _Deployed.Contract.SetA(&_Deployed.TransactOpts, arg0)
}

// SetA is a paid mutator transaction binding the contract method 0xee919d50.
//
// Solidity: function setA(uint256 ) returns(uint256)
func (_Deployed *DeployedTransactorSession) SetA(arg0 *big.Int) (*types.Transaction, error) {
	return _Deployed.Contract.SetA(&_Deployed.TransactOpts, arg0)
}

// ExistingABI is the input ABI used to generate the binding from.
const ExistingABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getA\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_val\",\"type\":\"uint256\"}],\"name\":\"setA\",\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_t\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ExistingBin is the compiled bytecode used for deploying new contracts.
const ExistingBin = `0x608060405234801561001057600080fd5b506040516020806102788339810180604052602081101561003057600080fd5b505160008054600160a060020a03909216600160a060020a0319909216919091179055610216806100626000396000f3fe608060405234801561001057600080fd5b5060043610610052577c01000000000000000000000000000000000000000000000000000000006000350463d46300fd8114610057578063ee919d5014610071575b600080fd5b61005f61008e565b60408051918252519081900360200190f35b61005f6004803603602081101561008757600080fd5b5035610144565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630dbe671f6040518163ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040160206040518083038186803b15801561011357600080fd5b505afa158015610127573d6000803e3d6000fd5b505050506040513d602081101561013d57600080fd5b5051905090565b60008054604080517fee919d5000000000000000000000000000000000000000000000000000000000815260048101859052905173ffffffffffffffffffffffffffffffffffffffff9092169163ee919d509160248082019260209290919082900301818787803b1580156101b857600080fd5b505af11580156101cc573d6000803e3d6000fd5b505050506040513d60208110156101e257600080fd5b50919291505056fea165627a7a72305820868c443a5ed7b9f841d4439be69f605dbae566ee81960e7741a1387ba6925c280029`

// DeployExisting deploys a new Ethereum contract, binding an instance of Existing to it.
func DeployExisting(auth *bind.TransactOpts, backend bind.ContractBackend, _t common.Address) (common.Address, *types.Transaction, *Existing, error) {
	parsed, err := abi.JSON(strings.NewReader(ExistingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExistingBin), backend, _t)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Existing{ExistingCaller: ExistingCaller{contract: contract}, ExistingTransactor: ExistingTransactor{contract: contract}, ExistingFilterer: ExistingFilterer{contract: contract}}, nil
}

// Existing is an auto generated Go binding around an Ethereum contract.
type Existing struct {
	ExistingCaller     // Read-only binding to the contract
	ExistingTransactor // Write-only binding to the contract
	ExistingFilterer   // Log filterer for contract events
}

// ExistingCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExistingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExistingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExistingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExistingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExistingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExistingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExistingSession struct {
	Contract     *Existing         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExistingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExistingCallerSession struct {
	Contract *ExistingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExistingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExistingTransactorSession struct {
	Contract     *ExistingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExistingRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExistingRaw struct {
	Contract *Existing // Generic contract binding to access the raw methods on
}

// ExistingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExistingCallerRaw struct {
	Contract *ExistingCaller // Generic read-only contract binding to access the raw methods on
}

// ExistingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExistingTransactorRaw struct {
	Contract *ExistingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExisting creates a new instance of Existing, bound to a specific deployed contract.
func NewExisting(address common.Address, backend bind.ContractBackend) (*Existing, error) {
	contract, err := bindExisting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Existing{ExistingCaller: ExistingCaller{contract: contract}, ExistingTransactor: ExistingTransactor{contract: contract}, ExistingFilterer: ExistingFilterer{contract: contract}}, nil
}

// NewExistingCaller creates a new read-only instance of Existing, bound to a specific deployed contract.
func NewExistingCaller(address common.Address, caller bind.ContractCaller) (*ExistingCaller, error) {
	contract, err := bindExisting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExistingCaller{contract: contract}, nil
}

// NewExistingTransactor creates a new write-only instance of Existing, bound to a specific deployed contract.
func NewExistingTransactor(address common.Address, transactor bind.ContractTransactor) (*ExistingTransactor, error) {
	contract, err := bindExisting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExistingTransactor{contract: contract}, nil
}

// NewExistingFilterer creates a new log filterer instance of Existing, bound to a specific deployed contract.
func NewExistingFilterer(address common.Address, filterer bind.ContractFilterer) (*ExistingFilterer, error) {
	contract, err := bindExisting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExistingFilterer{contract: contract}, nil
}

// bindExisting binds a generic wrapper to an already deployed contract.
func bindExisting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExistingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Existing *ExistingRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Existing.Contract.ExistingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Existing *ExistingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Existing.Contract.ExistingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Existing *ExistingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Existing.Contract.ExistingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Existing *ExistingCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Existing.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Existing *ExistingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Existing.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Existing *ExistingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Existing.Contract.contract.Transact(opts, method, params...)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() constant returns(uint256 result)
func (_Existing *ExistingCaller) GetA(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Existing.contract.Call(opts, out, "getA")
	return *ret0, err
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() constant returns(uint256 result)
func (_Existing *ExistingSession) GetA() (*big.Int, error) {
	return _Existing.Contract.GetA(&_Existing.CallOpts)
}

// GetA is a free data retrieval call binding the contract method 0xd46300fd.
//
// Solidity: function getA() constant returns(uint256 result)
func (_Existing *ExistingCallerSession) GetA() (*big.Int, error) {
	return _Existing.Contract.GetA(&_Existing.CallOpts)
}

// SetA is a paid mutator transaction binding the contract method 0xee919d50.
//
// Solidity: function setA(uint256 _val) returns(uint256 result)
func (_Existing *ExistingTransactor) SetA(opts *bind.TransactOpts, _val *big.Int) (*types.Transaction, error) {
	return _Existing.contract.Transact(opts, "setA", _val)
}

// SetA is a paid mutator transaction binding the contract method 0xee919d50.
//
// Solidity: function setA(uint256 _val) returns(uint256 result)
func (_Existing *ExistingSession) SetA(_val *big.Int) (*types.Transaction, error) {
	return _Existing.Contract.SetA(&_Existing.TransactOpts, _val)
}

// SetA is a paid mutator transaction binding the contract method 0xee919d50.
//
// Solidity: function setA(uint256 _val) returns(uint256 result)
func (_Existing *ExistingTransactorSession) SetA(_val *big.Int) (*types.Transaction, error) {
	return _Existing.Contract.SetA(&_Existing.TransactOpts, _val)
}
