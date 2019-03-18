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
const DeployedABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"a\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_a\",\"type\":\"uint256\"}],\"name\":\"setA\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DeployedBin is the compiled bytecode used for deploying new contracts.
const DeployedBin = `0x6080604052600160005534801561001557600080fd5b5060bf806100246000396000f3fe6080604052348015600f57600080fd5b5060043610604e577c010000000000000000000000000000000000000000000000000000000060003504630dbe671f81146053578063ee919d5014606b575b600080fd5b60596085565b60408051918252519081900360200190f35b605960048036036020811015607f57600080fd5b5035608b565b60005481565b60008190559056fea165627a7a7230582060ed592477fc020ada7e718a17e519a3d552fe807fe2292bbe59a02372654b640029`

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
// Solidity: function setA(uint256 _a) returns(uint256)
func (_Deployed *DeployedTransactor) SetA(opts *bind.TransactOpts, _a *big.Int) (*types.Transaction, error) {
	return _Deployed.contract.Transact(opts, "setA", _a)
}

// SetA is a paid mutator transaction binding the contract method 0xee919d50.
//
// Solidity: function setA(uint256 _a) returns(uint256)
func (_Deployed *DeployedSession) SetA(_a *big.Int) (*types.Transaction, error) {
	return _Deployed.Contract.SetA(&_Deployed.TransactOpts, _a)
}

// SetA is a paid mutator transaction binding the contract method 0xee919d50.
//
// Solidity: function setA(uint256 _a) returns(uint256)
func (_Deployed *DeployedTransactorSession) SetA(_a *big.Int) (*types.Transaction, error) {
	return _Deployed.Contract.SetA(&_Deployed.TransactOpts, _a)
}
