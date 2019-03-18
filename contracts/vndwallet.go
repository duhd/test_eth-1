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

// ContractReceiverABI is the input ABI used to generate the binding from.
const ContractReceiverABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractReceiverBin is the compiled bytecode used for deploying new contracts.
const ContractReceiverBin = `0x`

// DeployContractReceiver deploys a new Ethereum contract, binding an instance of ContractReceiver to it.
func DeployContractReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ContractReceiver, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractReceiverABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractReceiver{ContractReceiverCaller: ContractReceiverCaller{contract: contract}, ContractReceiverTransactor: ContractReceiverTransactor{contract: contract}, ContractReceiverFilterer: ContractReceiverFilterer{contract: contract}}, nil
}

// ContractReceiver is an auto generated Go binding around an Ethereum contract.
type ContractReceiver struct {
	ContractReceiverCaller     // Read-only binding to the contract
	ContractReceiverTransactor // Write-only binding to the contract
	ContractReceiverFilterer   // Log filterer for contract events
}

// ContractReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractReceiverSession struct {
	Contract     *ContractReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractReceiverCallerSession struct {
	Contract *ContractReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ContractReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractReceiverTransactorSession struct {
	Contract     *ContractReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ContractReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractReceiverRaw struct {
	Contract *ContractReceiver // Generic contract binding to access the raw methods on
}

// ContractReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractReceiverCallerRaw struct {
	Contract *ContractReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// ContractReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractReceiverTransactorRaw struct {
	Contract *ContractReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractReceiver creates a new instance of ContractReceiver, bound to a specific deployed contract.
func NewContractReceiver(address common.Address, backend bind.ContractBackend) (*ContractReceiver, error) {
	contract, err := bindContractReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractReceiver{ContractReceiverCaller: ContractReceiverCaller{contract: contract}, ContractReceiverTransactor: ContractReceiverTransactor{contract: contract}, ContractReceiverFilterer: ContractReceiverFilterer{contract: contract}}, nil
}

// NewContractReceiverCaller creates a new read-only instance of ContractReceiver, bound to a specific deployed contract.
func NewContractReceiverCaller(address common.Address, caller bind.ContractCaller) (*ContractReceiverCaller, error) {
	contract, err := bindContractReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractReceiverCaller{contract: contract}, nil
}

// NewContractReceiverTransactor creates a new write-only instance of ContractReceiver, bound to a specific deployed contract.
func NewContractReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractReceiverTransactor, error) {
	contract, err := bindContractReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractReceiverTransactor{contract: contract}, nil
}

// NewContractReceiverFilterer creates a new log filterer instance of ContractReceiver, bound to a specific deployed contract.
func NewContractReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractReceiverFilterer, error) {
	contract, err := bindContractReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractReceiverFilterer{contract: contract}, nil
}

// bindContractReceiver binds a generic wrapper to an already deployed contract.
func bindContractReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractReceiverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractReceiver *ContractReceiverRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractReceiver.Contract.ContractReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractReceiver *ContractReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractReceiver.Contract.ContractReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractReceiver *ContractReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractReceiver.Contract.ContractReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractReceiver *ContractReceiverCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ContractReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractReceiver *ContractReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractReceiver *ContractReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractReceiver.Contract.contract.Transact(opts, method, params...)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address from, uint256 value, bytes data) returns()
func (_ContractReceiver *ContractReceiverTransactor) TokenFallback(opts *bind.TransactOpts, from common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ContractReceiver.contract.Transact(opts, "tokenFallback", from, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address from, uint256 value, bytes data) returns()
func (_ContractReceiver *ContractReceiverSession) TokenFallback(from common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ContractReceiver.Contract.TokenFallback(&_ContractReceiver.TransactOpts, from, value, data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address from, uint256 value, bytes data) returns()
func (_ContractReceiver *ContractReceiverTransactorSession) TokenFallback(from common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ContractReceiver.Contract.TokenFallback(&_ContractReceiver.TransactOpts, from, value, data)
}

// ERC223TokenBasicABI is the input ABI used to generate the binding from.
const ERC223TokenBasicABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"typ\",\"type\":\"uint256\"}],\"name\":\"changeAccountType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferFromChild\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"blockByMasterKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"checkAccountType\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"unblockByMasterKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"checkOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"selfBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fromMasterKey\",\"type\":\"address\"},{\"name\":\"toMasterKey\",\"type\":\"address\"}],\"name\":\"changeMasterKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC223TokenBasicBin is the compiled bytecode used for deploying new contracts.
const ERC223TokenBasicBin = `0x`

// DeployERC223TokenBasic deploys a new Ethereum contract, binding an instance of ERC223TokenBasic to it.
func DeployERC223TokenBasic(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC223TokenBasic, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC223TokenBasicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC223TokenBasicBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC223TokenBasic{ERC223TokenBasicCaller: ERC223TokenBasicCaller{contract: contract}, ERC223TokenBasicTransactor: ERC223TokenBasicTransactor{contract: contract}, ERC223TokenBasicFilterer: ERC223TokenBasicFilterer{contract: contract}}, nil
}

// ERC223TokenBasic is an auto generated Go binding around an Ethereum contract.
type ERC223TokenBasic struct {
	ERC223TokenBasicCaller     // Read-only binding to the contract
	ERC223TokenBasicTransactor // Write-only binding to the contract
	ERC223TokenBasicFilterer   // Log filterer for contract events
}

// ERC223TokenBasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC223TokenBasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223TokenBasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC223TokenBasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223TokenBasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC223TokenBasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC223TokenBasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC223TokenBasicSession struct {
	Contract     *ERC223TokenBasic // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC223TokenBasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC223TokenBasicCallerSession struct {
	Contract *ERC223TokenBasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC223TokenBasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC223TokenBasicTransactorSession struct {
	Contract     *ERC223TokenBasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC223TokenBasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC223TokenBasicRaw struct {
	Contract *ERC223TokenBasic // Generic contract binding to access the raw methods on
}

// ERC223TokenBasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC223TokenBasicCallerRaw struct {
	Contract *ERC223TokenBasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC223TokenBasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC223TokenBasicTransactorRaw struct {
	Contract *ERC223TokenBasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC223TokenBasic creates a new instance of ERC223TokenBasic, bound to a specific deployed contract.
func NewERC223TokenBasic(address common.Address, backend bind.ContractBackend) (*ERC223TokenBasic, error) {
	contract, err := bindERC223TokenBasic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC223TokenBasic{ERC223TokenBasicCaller: ERC223TokenBasicCaller{contract: contract}, ERC223TokenBasicTransactor: ERC223TokenBasicTransactor{contract: contract}, ERC223TokenBasicFilterer: ERC223TokenBasicFilterer{contract: contract}}, nil
}

// NewERC223TokenBasicCaller creates a new read-only instance of ERC223TokenBasic, bound to a specific deployed contract.
func NewERC223TokenBasicCaller(address common.Address, caller bind.ContractCaller) (*ERC223TokenBasicCaller, error) {
	contract, err := bindERC223TokenBasic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC223TokenBasicCaller{contract: contract}, nil
}

// NewERC223TokenBasicTransactor creates a new write-only instance of ERC223TokenBasic, bound to a specific deployed contract.
func NewERC223TokenBasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC223TokenBasicTransactor, error) {
	contract, err := bindERC223TokenBasic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC223TokenBasicTransactor{contract: contract}, nil
}

// NewERC223TokenBasicFilterer creates a new log filterer instance of ERC223TokenBasic, bound to a specific deployed contract.
func NewERC223TokenBasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC223TokenBasicFilterer, error) {
	contract, err := bindERC223TokenBasic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC223TokenBasicFilterer{contract: contract}, nil
}

// bindERC223TokenBasic binds a generic wrapper to an already deployed contract.
func bindERC223TokenBasic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC223TokenBasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC223TokenBasic *ERC223TokenBasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC223TokenBasic.Contract.ERC223TokenBasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC223TokenBasic *ERC223TokenBasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.ERC223TokenBasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC223TokenBasic *ERC223TokenBasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.ERC223TokenBasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC223TokenBasic *ERC223TokenBasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC223TokenBasic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC223TokenBasic *ERC223TokenBasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC223TokenBasic *ERC223TokenBasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC223TokenBasic *ERC223TokenBasicCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC223TokenBasic.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC223TokenBasic *ERC223TokenBasicSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC223TokenBasic.Contract.BalanceOf(&_ERC223TokenBasic.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_ERC223TokenBasic *ERC223TokenBasicCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ERC223TokenBasic.Contract.BalanceOf(&_ERC223TokenBasic.CallOpts, owner)
}

// CheckOwners is a free data retrieval call binding the contract method 0xc14ee5a4.
//
// Solidity: function checkOwners(address addr, address owner) constant returns(bool)
func (_ERC223TokenBasic *ERC223TokenBasicCaller) CheckOwners(opts *bind.CallOpts, addr common.Address, owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ERC223TokenBasic.contract.Call(opts, out, "checkOwners", addr, owner)
	return *ret0, err
}

// CheckOwners is a free data retrieval call binding the contract method 0xc14ee5a4.
//
// Solidity: function checkOwners(address addr, address owner) constant returns(bool)
func (_ERC223TokenBasic *ERC223TokenBasicSession) CheckOwners(addr common.Address, owner common.Address) (bool, error) {
	return _ERC223TokenBasic.Contract.CheckOwners(&_ERC223TokenBasic.CallOpts, addr, owner)
}

// CheckOwners is a free data retrieval call binding the contract method 0xc14ee5a4.
//
// Solidity: function checkOwners(address addr, address owner) constant returns(bool)
func (_ERC223TokenBasic *ERC223TokenBasicCallerSession) CheckOwners(addr common.Address, owner common.Address) (bool, error) {
	return _ERC223TokenBasic.Contract.CheckOwners(&_ERC223TokenBasic.CallOpts, addr, owner)
}

// BlockByMasterKey is a paid mutator transaction binding the contract method 0x87200737.
//
// Solidity: function blockByMasterKey(address addr) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactor) BlockByMasterKey(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.contract.Transact(opts, "blockByMasterKey", addr)
}

// BlockByMasterKey is a paid mutator transaction binding the contract method 0x87200737.
//
// Solidity: function blockByMasterKey(address addr) returns()
func (_ERC223TokenBasic *ERC223TokenBasicSession) BlockByMasterKey(addr common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.BlockByMasterKey(&_ERC223TokenBasic.TransactOpts, addr)
}

// BlockByMasterKey is a paid mutator transaction binding the contract method 0x87200737.
//
// Solidity: function blockByMasterKey(address addr) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactorSession) BlockByMasterKey(addr common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.BlockByMasterKey(&_ERC223TokenBasic.TransactOpts, addr)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _ERC223TokenBasic.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ERC223TokenBasic *ERC223TokenBasicSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.Burn(&_ERC223TokenBasic.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.Burn(&_ERC223TokenBasic.TransactOpts, value)
}

// ChangeAccountType is a paid mutator transaction binding the contract method 0x0c14ad67.
//
// Solidity: function changeAccountType(address addr, uint256 typ) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactor) ChangeAccountType(opts *bind.TransactOpts, addr common.Address, typ *big.Int) (*types.Transaction, error) {
	return _ERC223TokenBasic.contract.Transact(opts, "changeAccountType", addr, typ)
}

// ChangeAccountType is a paid mutator transaction binding the contract method 0x0c14ad67.
//
// Solidity: function changeAccountType(address addr, uint256 typ) returns()
func (_ERC223TokenBasic *ERC223TokenBasicSession) ChangeAccountType(addr common.Address, typ *big.Int) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.ChangeAccountType(&_ERC223TokenBasic.TransactOpts, addr, typ)
}

// ChangeAccountType is a paid mutator transaction binding the contract method 0x0c14ad67.
//
// Solidity: function changeAccountType(address addr, uint256 typ) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactorSession) ChangeAccountType(addr common.Address, typ *big.Int) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.ChangeAccountType(&_ERC223TokenBasic.TransactOpts, addr, typ)
}

// ChangeMasterKey is a paid mutator transaction binding the contract method 0xdedf13a3.
//
// Solidity: function changeMasterKey(address fromMasterKey, address toMasterKey) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactor) ChangeMasterKey(opts *bind.TransactOpts, fromMasterKey common.Address, toMasterKey common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.contract.Transact(opts, "changeMasterKey", fromMasterKey, toMasterKey)
}

// ChangeMasterKey is a paid mutator transaction binding the contract method 0xdedf13a3.
//
// Solidity: function changeMasterKey(address fromMasterKey, address toMasterKey) returns()
func (_ERC223TokenBasic *ERC223TokenBasicSession) ChangeMasterKey(fromMasterKey common.Address, toMasterKey common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.ChangeMasterKey(&_ERC223TokenBasic.TransactOpts, fromMasterKey, toMasterKey)
}

// ChangeMasterKey is a paid mutator transaction binding the contract method 0xdedf13a3.
//
// Solidity: function changeMasterKey(address fromMasterKey, address toMasterKey) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactorSession) ChangeMasterKey(fromMasterKey common.Address, toMasterKey common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.ChangeMasterKey(&_ERC223TokenBasic.TransactOpts, fromMasterKey, toMasterKey)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address addr, address owner) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactor) ChangeOwner(opts *bind.TransactOpts, addr common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.contract.Transact(opts, "changeOwner", addr, owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address addr, address owner) returns()
func (_ERC223TokenBasic *ERC223TokenBasicSession) ChangeOwner(addr common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.ChangeOwner(&_ERC223TokenBasic.TransactOpts, addr, owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address addr, address owner) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactorSession) ChangeOwner(addr common.Address, owner common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.ChangeOwner(&_ERC223TokenBasic.TransactOpts, addr, owner)
}

// CheckAccountType is a paid mutator transaction binding the contract method 0x8b50c9d9.
//
// Solidity: function checkAccountType(address _addr) returns(uint256)
func (_ERC223TokenBasic *ERC223TokenBasicTransactor) CheckAccountType(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.contract.Transact(opts, "checkAccountType", _addr)
}

// CheckAccountType is a paid mutator transaction binding the contract method 0x8b50c9d9.
//
// Solidity: function checkAccountType(address _addr) returns(uint256)
func (_ERC223TokenBasic *ERC223TokenBasicSession) CheckAccountType(_addr common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.CheckAccountType(&_ERC223TokenBasic.TransactOpts, _addr)
}

// CheckAccountType is a paid mutator transaction binding the contract method 0x8b50c9d9.
//
// Solidity: function checkAccountType(address _addr) returns(uint256)
func (_ERC223TokenBasic *ERC223TokenBasicTransactorSession) CheckAccountType(_addr common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.CheckAccountType(&_ERC223TokenBasic.TransactOpts, _addr)
}

// IsBlock is a paid mutator transaction binding the contract method 0x007d6fff.
//
// Solidity: function isBlock(address _addr) returns(bool)
func (_ERC223TokenBasic *ERC223TokenBasicTransactor) IsBlock(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.contract.Transact(opts, "isBlock", _addr)
}

// IsBlock is a paid mutator transaction binding the contract method 0x007d6fff.
//
// Solidity: function isBlock(address _addr) returns(bool)
func (_ERC223TokenBasic *ERC223TokenBasicSession) IsBlock(_addr common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.IsBlock(&_ERC223TokenBasic.TransactOpts, _addr)
}

// IsBlock is a paid mutator transaction binding the contract method 0x007d6fff.
//
// Solidity: function isBlock(address _addr) returns(bool)
func (_ERC223TokenBasic *ERC223TokenBasicTransactorSession) IsBlock(_addr common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.IsBlock(&_ERC223TokenBasic.TransactOpts, _addr)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactor) Mint(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _ERC223TokenBasic.contract.Transact(opts, "mint", value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_ERC223TokenBasic *ERC223TokenBasicSession) Mint(value *big.Int) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.Mint(&_ERC223TokenBasic.TransactOpts, value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactorSession) Mint(value *big.Int) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.Mint(&_ERC223TokenBasic.TransactOpts, value)
}

// Refund is a paid mutator transaction binding the contract method 0x363d6a05.
//
// Solidity: function refund(address from, address to, uint256 value, bytes data) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactor) Refund(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC223TokenBasic.contract.Transact(opts, "refund", from, to, value, data)
}

// Refund is a paid mutator transaction binding the contract method 0x363d6a05.
//
// Solidity: function refund(address from, address to, uint256 value, bytes data) returns()
func (_ERC223TokenBasic *ERC223TokenBasicSession) Refund(from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.Refund(&_ERC223TokenBasic.TransactOpts, from, to, value, data)
}

// Refund is a paid mutator transaction binding the contract method 0x363d6a05.
//
// Solidity: function refund(address from, address to, uint256 value, bytes data) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactorSession) Refund(from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.Refund(&_ERC223TokenBasic.TransactOpts, from, to, value, data)
}

// SelfBlock is a paid mutator transaction binding the contract method 0xc233de85.
//
// Solidity: function selfBlock() returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactor) SelfBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC223TokenBasic.contract.Transact(opts, "selfBlock")
}

// SelfBlock is a paid mutator transaction binding the contract method 0xc233de85.
//
// Solidity: function selfBlock() returns()
func (_ERC223TokenBasic *ERC223TokenBasicSession) SelfBlock() (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.SelfBlock(&_ERC223TokenBasic.TransactOpts)
}

// SelfBlock is a paid mutator transaction binding the contract method 0xc233de85.
//
// Solidity: function selfBlock() returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactorSession) SelfBlock() (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.SelfBlock(&_ERC223TokenBasic.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address receiver, uint256 amount, bytes data) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC223TokenBasic.contract.Transact(opts, "transfer", receiver, amount, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address receiver, uint256 amount, bytes data) returns()
func (_ERC223TokenBasic *ERC223TokenBasicSession) Transfer(receiver common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.Transfer(&_ERC223TokenBasic.TransactOpts, receiver, amount, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address receiver, uint256 amount, bytes data) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactorSession) Transfer(receiver common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.Transfer(&_ERC223TokenBasic.TransactOpts, receiver, amount, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC223TokenBasic.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns()
func (_ERC223TokenBasic *ERC223TokenBasicSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.TransferFrom(&_ERC223TokenBasic.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.TransferFrom(&_ERC223TokenBasic.TransactOpts, from, to, value)
}

// TransferFromChild is a paid mutator transaction binding the contract method 0x49bb3e8f.
//
// Solidity: function transferFromChild(address from, address to, uint256 value, bytes data) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactor) TransferFromChild(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC223TokenBasic.contract.Transact(opts, "transferFromChild", from, to, value, data)
}

// TransferFromChild is a paid mutator transaction binding the contract method 0x49bb3e8f.
//
// Solidity: function transferFromChild(address from, address to, uint256 value, bytes data) returns()
func (_ERC223TokenBasic *ERC223TokenBasicSession) TransferFromChild(from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.TransferFromChild(&_ERC223TokenBasic.TransactOpts, from, to, value, data)
}

// TransferFromChild is a paid mutator transaction binding the contract method 0x49bb3e8f.
//
// Solidity: function transferFromChild(address from, address to, uint256 value, bytes data) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactorSession) TransferFromChild(from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.TransferFromChild(&_ERC223TokenBasic.TransactOpts, from, to, value, data)
}

// UnblockByMasterKey is a paid mutator transaction binding the contract method 0xc02e3654.
//
// Solidity: function unblockByMasterKey(address addr) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactor) UnblockByMasterKey(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.contract.Transact(opts, "unblockByMasterKey", addr)
}

// UnblockByMasterKey is a paid mutator transaction binding the contract method 0xc02e3654.
//
// Solidity: function unblockByMasterKey(address addr) returns()
func (_ERC223TokenBasic *ERC223TokenBasicSession) UnblockByMasterKey(addr common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.UnblockByMasterKey(&_ERC223TokenBasic.TransactOpts, addr)
}

// UnblockByMasterKey is a paid mutator transaction binding the contract method 0xc02e3654.
//
// Solidity: function unblockByMasterKey(address addr) returns()
func (_ERC223TokenBasic *ERC223TokenBasicTransactorSession) UnblockByMasterKey(addr common.Address) (*types.Transaction, error) {
	return _ERC223TokenBasic.Contract.UnblockByMasterKey(&_ERC223TokenBasic.TransactOpts, addr)
}

// TokenRecipientABI is the input ABI used to generate the binding from.
const TokenRecipientABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"receiveApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenRecipientBin is the compiled bytecode used for deploying new contracts.
const TokenRecipientBin = `0x`

// DeployTokenRecipient deploys a new Ethereum contract, binding an instance of TokenRecipient to it.
func DeployTokenRecipient(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenRecipient, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenRecipientABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenRecipientBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenRecipient{TokenRecipientCaller: TokenRecipientCaller{contract: contract}, TokenRecipientTransactor: TokenRecipientTransactor{contract: contract}, TokenRecipientFilterer: TokenRecipientFilterer{contract: contract}}, nil
}

// TokenRecipient is an auto generated Go binding around an Ethereum contract.
type TokenRecipient struct {
	TokenRecipientCaller     // Read-only binding to the contract
	TokenRecipientTransactor // Write-only binding to the contract
	TokenRecipientFilterer   // Log filterer for contract events
}

// TokenRecipientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenRecipientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRecipientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenRecipientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRecipientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenRecipientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenRecipientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenRecipientSession struct {
	Contract     *TokenRecipient   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRecipientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenRecipientCallerSession struct {
	Contract *TokenRecipientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TokenRecipientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenRecipientTransactorSession struct {
	Contract     *TokenRecipientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TokenRecipientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRecipientRaw struct {
	Contract *TokenRecipient // Generic contract binding to access the raw methods on
}

// TokenRecipientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenRecipientCallerRaw struct {
	Contract *TokenRecipientCaller // Generic read-only contract binding to access the raw methods on
}

// TokenRecipientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenRecipientTransactorRaw struct {
	Contract *TokenRecipientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenRecipient creates a new instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipient(address common.Address, backend bind.ContractBackend) (*TokenRecipient, error) {
	contract, err := bindTokenRecipient(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenRecipient{TokenRecipientCaller: TokenRecipientCaller{contract: contract}, TokenRecipientTransactor: TokenRecipientTransactor{contract: contract}, TokenRecipientFilterer: TokenRecipientFilterer{contract: contract}}, nil
}

// NewTokenRecipientCaller creates a new read-only instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipientCaller(address common.Address, caller bind.ContractCaller) (*TokenRecipientCaller, error) {
	contract, err := bindTokenRecipient(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRecipientCaller{contract: contract}, nil
}

// NewTokenRecipientTransactor creates a new write-only instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipientTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenRecipientTransactor, error) {
	contract, err := bindTokenRecipient(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenRecipientTransactor{contract: contract}, nil
}

// NewTokenRecipientFilterer creates a new log filterer instance of TokenRecipient, bound to a specific deployed contract.
func NewTokenRecipientFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenRecipientFilterer, error) {
	contract, err := bindTokenRecipient(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenRecipientFilterer{contract: contract}, nil
}

// bindTokenRecipient binds a generic wrapper to an already deployed contract.
func bindTokenRecipient(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenRecipientABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRecipient *TokenRecipientRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenRecipient.Contract.TokenRecipientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRecipient *TokenRecipientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRecipient.Contract.TokenRecipientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRecipient *TokenRecipientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRecipient.Contract.TokenRecipientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenRecipient *TokenRecipientCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenRecipient.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenRecipient *TokenRecipientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenRecipient.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenRecipient *TokenRecipientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenRecipient.Contract.contract.Transact(opts, method, params...)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0xa2d57853.
//
// Solidity: function receiveApproval(address from, uint256 value, bytes data) returns()
func (_TokenRecipient *TokenRecipientTransactor) ReceiveApproval(opts *bind.TransactOpts, from common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenRecipient.contract.Transact(opts, "receiveApproval", from, value, data)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0xa2d57853.
//
// Solidity: function receiveApproval(address from, uint256 value, bytes data) returns()
func (_TokenRecipient *TokenRecipientSession) ReceiveApproval(from common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenRecipient.Contract.ReceiveApproval(&_TokenRecipient.TransactOpts, from, value, data)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0xa2d57853.
//
// Solidity: function receiveApproval(address from, uint256 value, bytes data) returns()
func (_TokenRecipient *TokenRecipientTransactorSession) ReceiveApproval(from common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _TokenRecipient.Contract.ReceiveApproval(&_TokenRecipient.TransactOpts, from, value, data)
}

// VNDWalletABI is the input ABI used to generate the binding from.
const VNDWalletABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"isBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_typ\",\"type\":\"uint256\"}],\"name\":\"changeAccountType\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferFromChild\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"custom_fallback\",\"type\":\"string\"}],\"name\":\"transferInvokeCallBack\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"masterKey1\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"masterKey2\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"blockByMasterKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"checkAccountType\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"unblockByMasterKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"checkOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"selfBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"context\",\"type\":\"bytes\"}],\"name\":\"approveAndCall\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"fromMasterKey\",\"type\":\"address\"},{\"name\":\"toMasterKey\",\"type\":\"address\"}],\"name\":\"changeMasterKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_currentValue\",\"type\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"safeApprove\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\"},{\"name\":\"tokenName\",\"type\":\"string\"},{\"name\":\"decimalUnits\",\"type\":\"uint8\"},{\"name\":\"tokenSymbol\",\"type\":\"string\"},{\"name\":\"initMasterKey1\",\"type\":\"address\"},{\"name\":\"initMasterKey2\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"typ\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"ownerTransactionId\",\"type\":\"uint256\"}],\"name\":\"ChangeOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"by\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"whichMasterKey\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"ChangeMasterKey\",\"type\":\"event\"}]"

// VNDWalletBin is the compiled bytecode used for deploying new contracts.
const VNDWalletBin = `0x60806040523480156200001157600080fd5b506040516200214138038062002141833981018060405260c08110156200003757600080fd5b8151602083018051919392830192916401000000008111156200005957600080fd5b820160208101848111156200006d57600080fd5b81516401000000008111828201871017156200008857600080fd5b50506020820151604090920180519194929391640100000000811115620000ae57600080fd5b82016020810184811115620000c257600080fd5b8151640100000000811182820187101715620000dd57600080fd5b505060208083015160409384015160ff8816600a0a8a026003819055336000908152600885529586205588519396509094509262000120929091880190620001ff565b506002805460ff191660ff8616179055825162000145906001906020860190620001ff565b5060048054600160a060020a03199081163390811790925560058054600160a060020a03868116918416919091179091556006805491851691909216179055600380546040805191825260208201819052818101929092527f3078300000000000000000000000000000000000000000000000000000000000606082015290516000917fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16919081900360800190a3505050505050620002a4565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200024257805160ff191683800117855562000272565b8280016001018555821562000272579182015b828111156200027257825182559160200191906001019062000255565b506200028092915062000284565b5090565b620002a191905b808211156200028057600081556001016200028b565b90565b611e8d80620002b46000396000f3fe6080604052600436106101de576000357c01000000000000000000000000000000000000000000000000000000009004806379cc679011610114578063c02e3654116100b2578063dd62ed3e11610081578063dd62ed3e146109b9578063dedf13a3146109f4578063f00d4b5d14610a2f578063f650366214610a6a576101de565b8063c02e36541461086e578063c14ee5a4146108a1578063c233de85146108dc578063cae9ca51146108f1576101de565b80638b50c9d9116100ee5780638b50c9d91461076a57806395d89b411461079d578063a0712d68146107b2578063be45fd62146107dc576101de565b806379cc6790146106e95780637e7899f9146107225780638720073714610737576101de565b8063313ce5671161018157806349bb3e8f1161015b57806349bb3e8f146104b55780635a0906af146105525780636756dd22146106a157806370a08231146106b6576101de565b8063313ce567146103c3578063363d6a05146103ee57806342966c681461048b576101de565b80630c14ad67116101bd5780630c14ad67146102ed57806318160ddd146103285780631d1438481461034f57806323b872dd14610380576101de565b80627d6fff146101e357806306fdde031461022a578063095ea7b3146102b4575b600080fd5b3480156101ef57600080fd5b506102166004803603602081101561020657600080fd5b5035600160a060020a0316610aa9565b604080519115158252519081900360200190f35b34801561023657600080fd5b5061023f610ae0565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610279578181015183820152602001610261565b50505050905090810190601f1680156102a65780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102c057600080fd5b50610216600480360360408110156102d757600080fd5b50600160a060020a038135169060200135610b6e565b3480156102f957600080fd5b506103266004803603604081101561031057600080fd5b50600160a060020a038135169060200135610bd4565b005b34801561033457600080fd5b5061033d610d0c565b60408051918252519081900360200190f35b34801561035b57600080fd5b50610364610d12565b60408051600160a060020a039092168252519081900360200190f35b34801561038c57600080fd5b50610326600480360360608110156103a357600080fd5b50600160a060020a03813581169160208101359091169060400135610d21565b3480156103cf57600080fd5b506103d8610db5565b6040805160ff9092168252519081900360200190f35b3480156103fa57600080fd5b506103266004803603608081101561041157600080fd5b600160a060020a0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561044c57600080fd5b82018360208201111561045e57600080fd5b8035906020019184600183028401116401000000008311171561048057600080fd5b509092509050610dbe565b34801561049757600080fd5b50610326600480360360208110156104ae57600080fd5b5035610e4a565b3480156104c157600080fd5b50610326600480360360808110156104d857600080fd5b600160a060020a0382358116926020810135909116916040820135919081019060808101606082013564010000000081111561051357600080fd5b82018360208201111561052557600080fd5b8035906020019184600183028401116401000000008311171561054757600080fd5b509092509050610ebd565b34801561055e57600080fd5b506102166004803603608081101561057557600080fd5b600160a060020a03823516916020810135918101906060810160408201356401000000008111156105a557600080fd5b8201836020820111156105b757600080fd5b803590602001918460018302840111640100000000831117156105d957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929594936020810193503591505064010000000081111561062c57600080fd5b82018360208201111561063e57600080fd5b8035906020019184600183028401116401000000008311171561066057600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610f29945050505050565b3480156106ad57600080fd5b50610364611147565b3480156106c257600080fd5b5061033d600480360360208110156106d957600080fd5b5035600160a060020a0316611156565b3480156106f557600080fd5b506103266004803603604081101561070c57600080fd5b50600160a060020a038135169060200135611171565b34801561072e57600080fd5b5061036461123d565b34801561074357600080fd5b506103266004803603602081101561075a57600080fd5b5035600160a060020a031661124c565b34801561077657600080fd5b5061033d6004803603602081101561078d57600080fd5b5035600160a060020a0316611334565b3480156107a957600080fd5b5061023f61136a565b3480156107be57600080fd5b50610326600480360360208110156107d557600080fd5b50356113c4565b3480156107e857600080fd5b50610326600480360360608110156107ff57600080fd5b600160a060020a038235169160208101359181019060608101604082013564010000000081111561082f57600080fd5b82018360208201111561084157600080fd5b8035906020019184600183028401116401000000008311171561086357600080fd5b509092509050611419565b34801561087a57600080fd5b506103266004803603602081101561089157600080fd5b5035600160a060020a03166114cd565b3480156108ad57600080fd5b50610216600480360360408110156108c457600080fd5b50600160a060020a038135811691602001351661154f565b3480156108e857600080fd5b50610326611574565b3480156108fd57600080fd5b506102166004803603606081101561091457600080fd5b600160a060020a038235169160208101359181019060608101604082013564010000000081111561094457600080fd5b82018360208201111561095657600080fd5b8035906020019184600183028401116401000000008311171561097857600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506115a6945050505050565b3480156109c557600080fd5b5061033d600480360360408110156109dc57600080fd5b50600160a060020a03813581169160200135166116b9565b348015610a0057600080fd5b5061032660048036036040811015610a1757600080fd5b50600160a060020a03813581169160200135166116e4565b348015610a3b57600080fd5b5061032660048036036040811015610a5257600080fd5b50600160a060020a0381358116916020013516611a55565b348015610a7657600080fd5b5061021660048036036060811015610a8d57600080fd5b50600160a060020a038135169060208101359060400135611bde565b600160a060020a0381166000908152600760205260408120546001811415610ad5576001915050610adb565b60009150505b919050565b6000805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610b665780601f10610b3b57610100808354040283529160200191610b66565b820191906000526020600020905b815481529060010190602001808311610b4957829003601f168201915b505050505081565b336000818152600960209081526040808320600160a060020a038716808552908352818420869055815186815291519394909390927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925928290030190a350600192915050565b600554600160a060020a0316331480610bf75750600654600160a060020a031633145b1515610c6457604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f596f7520617265206e6f74206d6173746572206b657921000000000000000000604482015290519081900360640190fd5b600160a060020a0382161515610c7957600080fd5b600160a060020a03821660009081526007602052604090205460011415610c9f57600080fd5b600160a060020a0382166000908152600760205260409020600101541580610ce25750600160a060020a0382166000908152600760205260409020600101548114155b1515610ced57600080fd5b600160a060020a03909116600090815260076020526040902060010155565b60035481565b600454600160a060020a031681565b600160a060020a03831660009081526007602052604090205460011415610d4757600080fd5b600160a060020a0383166000908152600960209081526040808320338452909152902054811115610d7757600080fd5b600160a060020a03831660009081526009602090815260408083203384529091529020805482900390556060610daf84848484611c1a565b50505050565b60025460ff1681565b600554600160a060020a0316331480610de15750600654600160a060020a031633145b1515610dec57600080fd5b600160a060020a0385161515610e0157600080fd5b610e4385858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611c1a92505050565b5050505050565b33600090815260086020526040902054811115610e6657600080fd5b3360008181526008602090815260409182902080548590039055600380548590039055815184815291517fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca59281900390910190a250565b3360009081526007602052604090205460011415610eda57600080fd5b600160a060020a03851660009081526007602052604090205460011415610f0057600080fd5b600160a060020a03858116600090815260076020526040902060020154163314610e0157600080fd5b6000610f3733868686611c1a565b610f4085611d53565b1561113c57600085905080600160a060020a03166000843388886040516024018084600160a060020a0316600160a060020a0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610fb8578181015183820152602001610fa0565b50505050905090810190601f168015610fe55780820380516001836020036101000a031916815260200191505b50945050505050604051602081830303815290604052906040518082805190602001908083835b6020831061102b5780518252601f19909201916020918201910161100c565b51815160001960209485036101000a01908116901991909116179052604080519490920184900390932092860180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090941693909317835251855190945084935090508083835b602083106110d05780518252601f1990920191602091820191016110b1565b6001836020036101000a03801982511681845116808217855250505050505090500191505060006040518083038185875af1925050503d8060008114611132576040519150601f19603f3d011682016040523d82523d6000602084013e611137565b606091505b505050505b506001949350505050565b600554600160a060020a031681565b600160a060020a031660009081526008602052604090205490565b600160a060020a03821660009081526008602052604090205481111561119657600080fd5b600160a060020a03821660009081526009602090815260408083203384529091529020548111156111c657600080fd5b600160a060020a0382166000818152600860209081526040808320805486900390556009825280832033845282529182902080548590039055600380548590039055815184815291517fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca59281900390910190a25050565b600654600160a060020a031681565b600554600160a060020a031633148061126f5750600654600160a060020a031633145b15156112dc57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f596f7520617265206e6f74206d6173746572206b657921000000000000000000604482015290519081900360640190fd5b600160a060020a03811615156112f157600080fd5b600160a060020a0381166000908152600760205260409020546001141561131757600080fd5b600160a060020a0316600090815260076020526040902060019055565b6000600160a060020a038216151561134b57600080fd5b50600160a060020a031660009081526007602052604090206001015490565b60018054604080516020600284861615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610b665780601f10610b3b57610100808354040283529160200191610b66565b336000818152600860209081526040918290208054850190556003805485019055815184815291517f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968859281900390910190a250565b336000908152600760205260409020546001141561143657600080fd5b61143f84611d53565b1561148b57611485848484848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611d5b92505050565b50610daf565b610daf33858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611c1a92505050565b600554600160a060020a03163314806114f05750600654600160a060020a031633145b15156114fb57600080fd5b600160a060020a038116151561151057600080fd5b600160a060020a03811660009081526007602052604090205460011461153557600080fd5b600160a060020a0316600090815260076020526040812055565b600160a060020a03918216600090815260076020526040902060020154821691161490565b336000908152600760205260409020546001141561159157600080fd5b33600090815260076020526040902060019055565b60006115b28484610b6e565b156116ae576040517fa2d578530000000000000000000000000000000000000000000000000000000081523360048201818152602483018690526060604484019081528551606485015285518894600160a060020a0386169463a2d578539490938a938a9360840190602085019080838360005b8381101561163e578181015183820152602001611626565b50505050905090810190601f16801561166b5780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b15801561168c57600080fd5b505af11580156116a0573d6000803e3d6000fd5b5050505060019150506116b2565b5060005b9392505050565b600160a060020a03918216600090815260096020908152604080832093909416825291909152205490565b600160a060020a03821615156116f957600080fd5b600160a060020a038116151561170e57600080fd5b600554600160a060020a03163314806117315750600654600160a060020a031633145b151561179e57604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600760248201527f64756f6e67687400000000000000000000000000000000000000000000000000604482015290519081900360640190fd5b600160a060020a0382811690821614156117b757600080fd5b600554600160a060020a0383811691161480156117e25750600654600160a060020a03828116911614155b806118135750600654600160a060020a0383811691161480156118135750600654600160a060020a03828116911614155b151561181e57600080fd5b600160a060020a038281166000908152600a6020908152604080832085851684529091529020541615801561187d5750600160a060020a038281166000908152600a602090815260408083208585168452909152902060010154163314155b156118bc57600160a060020a038083166000908152600a602090815260408083209385168352929052208054600160a060020a03191633179055611959565b600160a060020a038281166000908152600a6020908152604080832085851684529091529020600101541615801561191b5750600160a060020a038281166000908152600a602090815260408083208585168452909152902054163314155b1561195957600160a060020a038083166000908152600a602090815260408083209385168352929052206001018054600160a060020a031916331790555b600160a060020a038281166000908152600a60209081526040808320858516845290915290205416158015906119b85750600160a060020a038281166000908152600a6020908152604080832085851684529091529020600101541615155b15611a5157600554600160a060020a03838116911614156119f35760058054600160a060020a031916600160a060020a038316179055611a0f565b60068054600160a060020a031916600160a060020a0383161790555b600160a060020a038281166000908152600a602090815260408083209385168352929052208054600160a060020a031990811682556001909101805490911690555b5050565b600160a060020a0382161515611a6a57600080fd5b600160a060020a038281169082161415611a8357600080fd5b600160a060020a0382811660009081526007602052604090206002015481169082161415611ab057600080fd5b33600160a060020a038316148015611ae35750600160a060020a0382811660009081526007602052604090206002015416155b80611b0a5750600160a060020a038281166000908152600760205260409020600201541633145b1515611b1557600080fd5b33600160a060020a038316148015611b485750600160a060020a0382811660009081526007602052604090206002015416155b15611b8357600160a060020a0382811660009081526007602052604090206002018054600160a060020a031916918316919091179055611a51565b600160a060020a0382811660009081526007602052604090206002015416331415611a5157600160a060020a0382811660009081526007602052604090206002018054600160a060020a031916918316919091179055611a51565b336000908152600960209081526040808320600160a060020a03871684529091528120548314156116ae57611c138483610b6e565b90506116b2565b600160a060020a0383161515611c2f57600080fd5b600160a060020a038416600090815260086020526040902054821115611c5457600080fd5b600160a060020a03831660009081526008602052604090205482810111611c7a57600080fd5b600160a060020a038085166000818152600860209081526040808320805488900390559387168083528483208054880190558451878152808301868152875196820196909652865191957fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1694899489946060850192918601918190849084905b83811015611d12578181015183820152602001611cfa565b50505050905090810190601f168015611d3f5780820380516001836020036101000a031916815260200191505b50935050505060405180910390a350505050565b6000903b1190565b6000611d6933858585611c1a565b6040517fc0ee0b8a0000000000000000000000000000000000000000000000000000000081523360048201818152602483018690526060604484019081528551606485015285518894600160a060020a0386169463c0ee0b8a9490938a938a9360840190602085019080838360005b83811015611df0578181015183820152602001611dd8565b50505050905090810190601f168015611e1d5780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b158015611e3e57600080fd5b505af1158015611e52573d6000803e3d6000fd5b5060019897505050505050505056fea165627a7a7230582004d6364b9bcc21d0d2824bfe10ac7b97f8f1a1f4ea644d5963a45ddeb1410e040029`

// DeployVNDWallet deploys a new Ethereum contract, binding an instance of VNDWallet to it.
func DeployVNDWallet(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, tokenName string, decimalUnits uint8, tokenSymbol string, initMasterKey1 common.Address, initMasterKey2 common.Address) (common.Address, *types.Transaction, *VNDWallet, error) {
	parsed, err := abi.JSON(strings.NewReader(VNDWalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VNDWalletBin), backend, initialSupply, tokenName, decimalUnits, tokenSymbol, initMasterKey1, initMasterKey2)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VNDWallet{VNDWalletCaller: VNDWalletCaller{contract: contract}, VNDWalletTransactor: VNDWalletTransactor{contract: contract}, VNDWalletFilterer: VNDWalletFilterer{contract: contract}}, nil
}

// VNDWallet is an auto generated Go binding around an Ethereum contract.
type VNDWallet struct {
	VNDWalletCaller     // Read-only binding to the contract
	VNDWalletTransactor // Write-only binding to the contract
	VNDWalletFilterer   // Log filterer for contract events
}

// VNDWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type VNDWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VNDWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VNDWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VNDWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VNDWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VNDWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VNDWalletSession struct {
	Contract     *VNDWallet        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VNDWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VNDWalletCallerSession struct {
	Contract *VNDWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// VNDWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VNDWalletTransactorSession struct {
	Contract     *VNDWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// VNDWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type VNDWalletRaw struct {
	Contract *VNDWallet // Generic contract binding to access the raw methods on
}

// VNDWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VNDWalletCallerRaw struct {
	Contract *VNDWalletCaller // Generic read-only contract binding to access the raw methods on
}

// VNDWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VNDWalletTransactorRaw struct {
	Contract *VNDWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVNDWallet creates a new instance of VNDWallet, bound to a specific deployed contract.
func NewVNDWallet(address common.Address, backend bind.ContractBackend) (*VNDWallet, error) {
	contract, err := bindVNDWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VNDWallet{VNDWalletCaller: VNDWalletCaller{contract: contract}, VNDWalletTransactor: VNDWalletTransactor{contract: contract}, VNDWalletFilterer: VNDWalletFilterer{contract: contract}}, nil
}

// NewVNDWalletCaller creates a new read-only instance of VNDWallet, bound to a specific deployed contract.
func NewVNDWalletCaller(address common.Address, caller bind.ContractCaller) (*VNDWalletCaller, error) {
	contract, err := bindVNDWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VNDWalletCaller{contract: contract}, nil
}

// NewVNDWalletTransactor creates a new write-only instance of VNDWallet, bound to a specific deployed contract.
func NewVNDWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*VNDWalletTransactor, error) {
	contract, err := bindVNDWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VNDWalletTransactor{contract: contract}, nil
}

// NewVNDWalletFilterer creates a new log filterer instance of VNDWallet, bound to a specific deployed contract.
func NewVNDWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*VNDWalletFilterer, error) {
	contract, err := bindVNDWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VNDWalletFilterer{contract: contract}, nil
}

// bindVNDWallet binds a generic wrapper to an already deployed contract.
func bindVNDWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VNDWalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VNDWallet *VNDWalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VNDWallet.Contract.VNDWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VNDWallet *VNDWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VNDWallet.Contract.VNDWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VNDWallet *VNDWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VNDWallet.Contract.VNDWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VNDWallet *VNDWalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VNDWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VNDWallet *VNDWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VNDWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VNDWallet *VNDWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VNDWallet.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256 remaining)
func (_VNDWallet *VNDWalletCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VNDWallet.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256 remaining)
func (_VNDWallet *VNDWalletSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VNDWallet.Contract.Allowance(&_VNDWallet.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) constant returns(uint256 remaining)
func (_VNDWallet *VNDWalletCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _VNDWallet.Contract.Allowance(&_VNDWallet.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_VNDWallet *VNDWalletCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VNDWallet.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_VNDWallet *VNDWalletSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VNDWallet.Contract.BalanceOf(&_VNDWallet.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) constant returns(uint256)
func (_VNDWallet *VNDWalletCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _VNDWallet.Contract.BalanceOf(&_VNDWallet.CallOpts, owner)
}

// CheckOwners is a free data retrieval call binding the contract method 0xc14ee5a4.
//
// Solidity: function checkOwners(address _addr, address _owner) constant returns(bool)
func (_VNDWallet *VNDWalletCaller) CheckOwners(opts *bind.CallOpts, _addr common.Address, _owner common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _VNDWallet.contract.Call(opts, out, "checkOwners", _addr, _owner)
	return *ret0, err
}

// CheckOwners is a free data retrieval call binding the contract method 0xc14ee5a4.
//
// Solidity: function checkOwners(address _addr, address _owner) constant returns(bool)
func (_VNDWallet *VNDWalletSession) CheckOwners(_addr common.Address, _owner common.Address) (bool, error) {
	return _VNDWallet.Contract.CheckOwners(&_VNDWallet.CallOpts, _addr, _owner)
}

// CheckOwners is a free data retrieval call binding the contract method 0xc14ee5a4.
//
// Solidity: function checkOwners(address _addr, address _owner) constant returns(bool)
func (_VNDWallet *VNDWalletCallerSession) CheckOwners(_addr common.Address, _owner common.Address) (bool, error) {
	return _VNDWallet.Contract.CheckOwners(&_VNDWallet.CallOpts, _addr, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_VNDWallet *VNDWalletCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _VNDWallet.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_VNDWallet *VNDWalletSession) Decimals() (uint8, error) {
	return _VNDWallet.Contract.Decimals(&_VNDWallet.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_VNDWallet *VNDWalletCallerSession) Decimals() (uint8, error) {
	return _VNDWallet.Contract.Decimals(&_VNDWallet.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_VNDWallet *VNDWalletCaller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VNDWallet.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_VNDWallet *VNDWalletSession) Issuer() (common.Address, error) {
	return _VNDWallet.Contract.Issuer(&_VNDWallet.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_VNDWallet *VNDWalletCallerSession) Issuer() (common.Address, error) {
	return _VNDWallet.Contract.Issuer(&_VNDWallet.CallOpts)
}

// MasterKey1 is a free data retrieval call binding the contract method 0x6756dd22.
//
// Solidity: function masterKey1() constant returns(address)
func (_VNDWallet *VNDWalletCaller) MasterKey1(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VNDWallet.contract.Call(opts, out, "masterKey1")
	return *ret0, err
}

// MasterKey1 is a free data retrieval call binding the contract method 0x6756dd22.
//
// Solidity: function masterKey1() constant returns(address)
func (_VNDWallet *VNDWalletSession) MasterKey1() (common.Address, error) {
	return _VNDWallet.Contract.MasterKey1(&_VNDWallet.CallOpts)
}

// MasterKey1 is a free data retrieval call binding the contract method 0x6756dd22.
//
// Solidity: function masterKey1() constant returns(address)
func (_VNDWallet *VNDWalletCallerSession) MasterKey1() (common.Address, error) {
	return _VNDWallet.Contract.MasterKey1(&_VNDWallet.CallOpts)
}

// MasterKey2 is a free data retrieval call binding the contract method 0x7e7899f9.
//
// Solidity: function masterKey2() constant returns(address)
func (_VNDWallet *VNDWalletCaller) MasterKey2(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VNDWallet.contract.Call(opts, out, "masterKey2")
	return *ret0, err
}

// MasterKey2 is a free data retrieval call binding the contract method 0x7e7899f9.
//
// Solidity: function masterKey2() constant returns(address)
func (_VNDWallet *VNDWalletSession) MasterKey2() (common.Address, error) {
	return _VNDWallet.Contract.MasterKey2(&_VNDWallet.CallOpts)
}

// MasterKey2 is a free data retrieval call binding the contract method 0x7e7899f9.
//
// Solidity: function masterKey2() constant returns(address)
func (_VNDWallet *VNDWalletCallerSession) MasterKey2() (common.Address, error) {
	return _VNDWallet.Contract.MasterKey2(&_VNDWallet.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_VNDWallet *VNDWalletCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _VNDWallet.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_VNDWallet *VNDWalletSession) Name() (string, error) {
	return _VNDWallet.Contract.Name(&_VNDWallet.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_VNDWallet *VNDWalletCallerSession) Name() (string, error) {
	return _VNDWallet.Contract.Name(&_VNDWallet.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_VNDWallet *VNDWalletCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _VNDWallet.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_VNDWallet *VNDWalletSession) Symbol() (string, error) {
	return _VNDWallet.Contract.Symbol(&_VNDWallet.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_VNDWallet *VNDWalletCallerSession) Symbol() (string, error) {
	return _VNDWallet.Contract.Symbol(&_VNDWallet.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_VNDWallet *VNDWalletCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VNDWallet.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_VNDWallet *VNDWalletSession) TotalSupply() (*big.Int, error) {
	return _VNDWallet.Contract.TotalSupply(&_VNDWallet.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_VNDWallet *VNDWalletCallerSession) TotalSupply() (*big.Int, error) {
	return _VNDWallet.Contract.TotalSupply(&_VNDWallet.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_VNDWallet *VNDWalletTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_VNDWallet *VNDWalletSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.Contract.Approve(&_VNDWallet.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool success)
func (_VNDWallet *VNDWalletTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.Contract.Approve(&_VNDWallet.TransactOpts, spender, value)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes context) returns(bool success)
func (_VNDWallet *VNDWalletTransactor) ApproveAndCall(opts *bind.TransactOpts, spender common.Address, value *big.Int, context []byte) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "approveAndCall", spender, value, context)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes context) returns(bool success)
func (_VNDWallet *VNDWalletSession) ApproveAndCall(spender common.Address, value *big.Int, context []byte) (*types.Transaction, error) {
	return _VNDWallet.Contract.ApproveAndCall(&_VNDWallet.TransactOpts, spender, value, context)
}

// ApproveAndCall is a paid mutator transaction binding the contract method 0xcae9ca51.
//
// Solidity: function approveAndCall(address spender, uint256 value, bytes context) returns(bool success)
func (_VNDWallet *VNDWalletTransactorSession) ApproveAndCall(spender common.Address, value *big.Int, context []byte) (*types.Transaction, error) {
	return _VNDWallet.Contract.ApproveAndCall(&_VNDWallet.TransactOpts, spender, value, context)
}

// BlockByMasterKey is a paid mutator transaction binding the contract method 0x87200737.
//
// Solidity: function blockByMasterKey(address _addr) returns()
func (_VNDWallet *VNDWalletTransactor) BlockByMasterKey(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "blockByMasterKey", _addr)
}

// BlockByMasterKey is a paid mutator transaction binding the contract method 0x87200737.
//
// Solidity: function blockByMasterKey(address _addr) returns()
func (_VNDWallet *VNDWalletSession) BlockByMasterKey(_addr common.Address) (*types.Transaction, error) {
	return _VNDWallet.Contract.BlockByMasterKey(&_VNDWallet.TransactOpts, _addr)
}

// BlockByMasterKey is a paid mutator transaction binding the contract method 0x87200737.
//
// Solidity: function blockByMasterKey(address _addr) returns()
func (_VNDWallet *VNDWalletTransactorSession) BlockByMasterKey(_addr common.Address) (*types.Transaction, error) {
	return _VNDWallet.Contract.BlockByMasterKey(&_VNDWallet.TransactOpts, _addr)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_VNDWallet *VNDWalletTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_VNDWallet *VNDWalletSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.Contract.Burn(&_VNDWallet.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_VNDWallet *VNDWalletTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.Contract.Burn(&_VNDWallet.TransactOpts, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 value) returns()
func (_VNDWallet *VNDWalletTransactor) BurnFrom(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "burnFrom", from, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 value) returns()
func (_VNDWallet *VNDWalletSession) BurnFrom(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.Contract.BurnFrom(&_VNDWallet.TransactOpts, from, value)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address from, uint256 value) returns()
func (_VNDWallet *VNDWalletTransactorSession) BurnFrom(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.Contract.BurnFrom(&_VNDWallet.TransactOpts, from, value)
}

// ChangeAccountType is a paid mutator transaction binding the contract method 0x0c14ad67.
//
// Solidity: function changeAccountType(address _addr, uint256 _typ) returns()
func (_VNDWallet *VNDWalletTransactor) ChangeAccountType(opts *bind.TransactOpts, _addr common.Address, _typ *big.Int) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "changeAccountType", _addr, _typ)
}

// ChangeAccountType is a paid mutator transaction binding the contract method 0x0c14ad67.
//
// Solidity: function changeAccountType(address _addr, uint256 _typ) returns()
func (_VNDWallet *VNDWalletSession) ChangeAccountType(_addr common.Address, _typ *big.Int) (*types.Transaction, error) {
	return _VNDWallet.Contract.ChangeAccountType(&_VNDWallet.TransactOpts, _addr, _typ)
}

// ChangeAccountType is a paid mutator transaction binding the contract method 0x0c14ad67.
//
// Solidity: function changeAccountType(address _addr, uint256 _typ) returns()
func (_VNDWallet *VNDWalletTransactorSession) ChangeAccountType(_addr common.Address, _typ *big.Int) (*types.Transaction, error) {
	return _VNDWallet.Contract.ChangeAccountType(&_VNDWallet.TransactOpts, _addr, _typ)
}

// ChangeMasterKey is a paid mutator transaction binding the contract method 0xdedf13a3.
//
// Solidity: function changeMasterKey(address fromMasterKey, address toMasterKey) returns()
func (_VNDWallet *VNDWalletTransactor) ChangeMasterKey(opts *bind.TransactOpts, fromMasterKey common.Address, toMasterKey common.Address) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "changeMasterKey", fromMasterKey, toMasterKey)
}

// ChangeMasterKey is a paid mutator transaction binding the contract method 0xdedf13a3.
//
// Solidity: function changeMasterKey(address fromMasterKey, address toMasterKey) returns()
func (_VNDWallet *VNDWalletSession) ChangeMasterKey(fromMasterKey common.Address, toMasterKey common.Address) (*types.Transaction, error) {
	return _VNDWallet.Contract.ChangeMasterKey(&_VNDWallet.TransactOpts, fromMasterKey, toMasterKey)
}

// ChangeMasterKey is a paid mutator transaction binding the contract method 0xdedf13a3.
//
// Solidity: function changeMasterKey(address fromMasterKey, address toMasterKey) returns()
func (_VNDWallet *VNDWalletTransactorSession) ChangeMasterKey(fromMasterKey common.Address, toMasterKey common.Address) (*types.Transaction, error) {
	return _VNDWallet.Contract.ChangeMasterKey(&_VNDWallet.TransactOpts, fromMasterKey, toMasterKey)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address _addr, address _owner) returns()
func (_VNDWallet *VNDWalletTransactor) ChangeOwner(opts *bind.TransactOpts, _addr common.Address, _owner common.Address) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "changeOwner", _addr, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address _addr, address _owner) returns()
func (_VNDWallet *VNDWalletSession) ChangeOwner(_addr common.Address, _owner common.Address) (*types.Transaction, error) {
	return _VNDWallet.Contract.ChangeOwner(&_VNDWallet.TransactOpts, _addr, _owner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xf00d4b5d.
//
// Solidity: function changeOwner(address _addr, address _owner) returns()
func (_VNDWallet *VNDWalletTransactorSession) ChangeOwner(_addr common.Address, _owner common.Address) (*types.Transaction, error) {
	return _VNDWallet.Contract.ChangeOwner(&_VNDWallet.TransactOpts, _addr, _owner)
}

// CheckAccountType is a paid mutator transaction binding the contract method 0x8b50c9d9.
//
// Solidity: function checkAccountType(address _addr) returns(uint256)
func (_VNDWallet *VNDWalletTransactor) CheckAccountType(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "checkAccountType", _addr)
}

// CheckAccountType is a paid mutator transaction binding the contract method 0x8b50c9d9.
//
// Solidity: function checkAccountType(address _addr) returns(uint256)
func (_VNDWallet *VNDWalletSession) CheckAccountType(_addr common.Address) (*types.Transaction, error) {
	return _VNDWallet.Contract.CheckAccountType(&_VNDWallet.TransactOpts, _addr)
}

// CheckAccountType is a paid mutator transaction binding the contract method 0x8b50c9d9.
//
// Solidity: function checkAccountType(address _addr) returns(uint256)
func (_VNDWallet *VNDWalletTransactorSession) CheckAccountType(_addr common.Address) (*types.Transaction, error) {
	return _VNDWallet.Contract.CheckAccountType(&_VNDWallet.TransactOpts, _addr)
}

// IsBlock is a paid mutator transaction binding the contract method 0x007d6fff.
//
// Solidity: function isBlock(address _addr) returns(bool)
func (_VNDWallet *VNDWalletTransactor) IsBlock(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "isBlock", _addr)
}

// IsBlock is a paid mutator transaction binding the contract method 0x007d6fff.
//
// Solidity: function isBlock(address _addr) returns(bool)
func (_VNDWallet *VNDWalletSession) IsBlock(_addr common.Address) (*types.Transaction, error) {
	return _VNDWallet.Contract.IsBlock(&_VNDWallet.TransactOpts, _addr)
}

// IsBlock is a paid mutator transaction binding the contract method 0x007d6fff.
//
// Solidity: function isBlock(address _addr) returns(bool)
func (_VNDWallet *VNDWalletTransactorSession) IsBlock(_addr common.Address) (*types.Transaction, error) {
	return _VNDWallet.Contract.IsBlock(&_VNDWallet.TransactOpts, _addr)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_VNDWallet *VNDWalletTransactor) Mint(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "mint", value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_VNDWallet *VNDWalletSession) Mint(value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.Contract.Mint(&_VNDWallet.TransactOpts, value)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 value) returns()
func (_VNDWallet *VNDWalletTransactorSession) Mint(value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.Contract.Mint(&_VNDWallet.TransactOpts, value)
}

// Refund is a paid mutator transaction binding the contract method 0x363d6a05.
//
// Solidity: function refund(address from, address to, uint256 value, bytes data) returns()
func (_VNDWallet *VNDWalletTransactor) Refund(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "refund", from, to, value, data)
}

// Refund is a paid mutator transaction binding the contract method 0x363d6a05.
//
// Solidity: function refund(address from, address to, uint256 value, bytes data) returns()
func (_VNDWallet *VNDWalletSession) Refund(from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VNDWallet.Contract.Refund(&_VNDWallet.TransactOpts, from, to, value, data)
}

// Refund is a paid mutator transaction binding the contract method 0x363d6a05.
//
// Solidity: function refund(address from, address to, uint256 value, bytes data) returns()
func (_VNDWallet *VNDWalletTransactorSession) Refund(from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VNDWallet.Contract.Refund(&_VNDWallet.TransactOpts, from, to, value, data)
}

// SafeApprove is a paid mutator transaction binding the contract method 0xf6503662.
//
// Solidity: function safeApprove(address _spender, uint256 _currentValue, uint256 _value) returns(bool success)
func (_VNDWallet *VNDWalletTransactor) SafeApprove(opts *bind.TransactOpts, _spender common.Address, _currentValue *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "safeApprove", _spender, _currentValue, _value)
}

// SafeApprove is a paid mutator transaction binding the contract method 0xf6503662.
//
// Solidity: function safeApprove(address _spender, uint256 _currentValue, uint256 _value) returns(bool success)
func (_VNDWallet *VNDWalletSession) SafeApprove(_spender common.Address, _currentValue *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.Contract.SafeApprove(&_VNDWallet.TransactOpts, _spender, _currentValue, _value)
}

// SafeApprove is a paid mutator transaction binding the contract method 0xf6503662.
//
// Solidity: function safeApprove(address _spender, uint256 _currentValue, uint256 _value) returns(bool success)
func (_VNDWallet *VNDWalletTransactorSession) SafeApprove(_spender common.Address, _currentValue *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.Contract.SafeApprove(&_VNDWallet.TransactOpts, _spender, _currentValue, _value)
}

// SelfBlock is a paid mutator transaction binding the contract method 0xc233de85.
//
// Solidity: function selfBlock() returns()
func (_VNDWallet *VNDWalletTransactor) SelfBlock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "selfBlock")
}

// SelfBlock is a paid mutator transaction binding the contract method 0xc233de85.
//
// Solidity: function selfBlock() returns()
func (_VNDWallet *VNDWalletSession) SelfBlock() (*types.Transaction, error) {
	return _VNDWallet.Contract.SelfBlock(&_VNDWallet.TransactOpts)
}

// SelfBlock is a paid mutator transaction binding the contract method 0xc233de85.
//
// Solidity: function selfBlock() returns()
func (_VNDWallet *VNDWalletTransactorSession) SelfBlock() (*types.Transaction, error) {
	return _VNDWallet.Contract.SelfBlock(&_VNDWallet.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address to, uint256 value, bytes data) returns()
func (_VNDWallet *VNDWalletTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "transfer", to, value, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address to, uint256 value, bytes data) returns()
func (_VNDWallet *VNDWalletSession) Transfer(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VNDWallet.Contract.Transfer(&_VNDWallet.TransactOpts, to, value, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xbe45fd62.
//
// Solidity: function transfer(address to, uint256 value, bytes data) returns()
func (_VNDWallet *VNDWalletTransactorSession) Transfer(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VNDWallet.Contract.Transfer(&_VNDWallet.TransactOpts, to, value, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns()
func (_VNDWallet *VNDWalletTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns()
func (_VNDWallet *VNDWalletSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.Contract.TransferFrom(&_VNDWallet.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns()
func (_VNDWallet *VNDWalletTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _VNDWallet.Contract.TransferFrom(&_VNDWallet.TransactOpts, from, to, value)
}

// TransferFromChild is a paid mutator transaction binding the contract method 0x49bb3e8f.
//
// Solidity: function transferFromChild(address from, address to, uint256 value, bytes data) returns()
func (_VNDWallet *VNDWalletTransactor) TransferFromChild(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "transferFromChild", from, to, value, data)
}

// TransferFromChild is a paid mutator transaction binding the contract method 0x49bb3e8f.
//
// Solidity: function transferFromChild(address from, address to, uint256 value, bytes data) returns()
func (_VNDWallet *VNDWalletSession) TransferFromChild(from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VNDWallet.Contract.TransferFromChild(&_VNDWallet.TransactOpts, from, to, value, data)
}

// TransferFromChild is a paid mutator transaction binding the contract method 0x49bb3e8f.
//
// Solidity: function transferFromChild(address from, address to, uint256 value, bytes data) returns()
func (_VNDWallet *VNDWalletTransactorSession) TransferFromChild(from common.Address, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _VNDWallet.Contract.TransferFromChild(&_VNDWallet.TransactOpts, from, to, value, data)
}

// TransferInvokeCallBack is a paid mutator transaction binding the contract method 0x5a0906af.
//
// Solidity: function transferInvokeCallBack(address to, uint256 value, bytes data, string custom_fallback) returns(bool success)
func (_VNDWallet *VNDWalletTransactor) TransferInvokeCallBack(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte, custom_fallback string) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "transferInvokeCallBack", to, value, data, custom_fallback)
}

// TransferInvokeCallBack is a paid mutator transaction binding the contract method 0x5a0906af.
//
// Solidity: function transferInvokeCallBack(address to, uint256 value, bytes data, string custom_fallback) returns(bool success)
func (_VNDWallet *VNDWalletSession) TransferInvokeCallBack(to common.Address, value *big.Int, data []byte, custom_fallback string) (*types.Transaction, error) {
	return _VNDWallet.Contract.TransferInvokeCallBack(&_VNDWallet.TransactOpts, to, value, data, custom_fallback)
}

// TransferInvokeCallBack is a paid mutator transaction binding the contract method 0x5a0906af.
//
// Solidity: function transferInvokeCallBack(address to, uint256 value, bytes data, string custom_fallback) returns(bool success)
func (_VNDWallet *VNDWalletTransactorSession) TransferInvokeCallBack(to common.Address, value *big.Int, data []byte, custom_fallback string) (*types.Transaction, error) {
	return _VNDWallet.Contract.TransferInvokeCallBack(&_VNDWallet.TransactOpts, to, value, data, custom_fallback)
}

// UnblockByMasterKey is a paid mutator transaction binding the contract method 0xc02e3654.
//
// Solidity: function unblockByMasterKey(address _addr) returns()
func (_VNDWallet *VNDWalletTransactor) UnblockByMasterKey(opts *bind.TransactOpts, _addr common.Address) (*types.Transaction, error) {
	return _VNDWallet.contract.Transact(opts, "unblockByMasterKey", _addr)
}

// UnblockByMasterKey is a paid mutator transaction binding the contract method 0xc02e3654.
//
// Solidity: function unblockByMasterKey(address _addr) returns()
func (_VNDWallet *VNDWalletSession) UnblockByMasterKey(_addr common.Address) (*types.Transaction, error) {
	return _VNDWallet.Contract.UnblockByMasterKey(&_VNDWallet.TransactOpts, _addr)
}

// UnblockByMasterKey is a paid mutator transaction binding the contract method 0xc02e3654.
//
// Solidity: function unblockByMasterKey(address _addr) returns()
func (_VNDWallet *VNDWalletTransactorSession) UnblockByMasterKey(_addr common.Address) (*types.Transaction, error) {
	return _VNDWallet.Contract.UnblockByMasterKey(&_VNDWallet.TransactOpts, _addr)
}

// VNDWalletApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the VNDWallet contract.
type VNDWalletApprovalIterator struct {
	Event *VNDWalletApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VNDWalletApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VNDWalletApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VNDWalletApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VNDWalletApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VNDWalletApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VNDWalletApproval represents a Approval event raised by the VNDWallet contract.
type VNDWalletApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VNDWallet *VNDWalletFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*VNDWalletApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VNDWallet.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VNDWalletApprovalIterator{contract: _VNDWallet.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_VNDWallet *VNDWalletFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VNDWalletApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _VNDWallet.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VNDWalletApproval)
				if err := _VNDWallet.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// VNDWalletBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the VNDWallet contract.
type VNDWalletBurnIterator struct {
	Event *VNDWalletBurn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VNDWalletBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VNDWalletBurn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VNDWalletBurn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VNDWalletBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VNDWalletBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VNDWalletBurn represents a Burn event raised by the VNDWallet contract.
type VNDWalletBurn struct {
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_VNDWallet *VNDWalletFilterer) FilterBurn(opts *bind.FilterOpts, from []common.Address) (*VNDWalletBurnIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _VNDWallet.contract.FilterLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return &VNDWalletBurnIterator{contract: _VNDWallet.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed from, uint256 value)
func (_VNDWallet *VNDWalletFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *VNDWalletBurn, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _VNDWallet.contract.WatchLogs(opts, "Burn", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VNDWalletBurn)
				if err := _VNDWallet.contract.UnpackLog(event, "Burn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// VNDWalletChangeMasterKeyIterator is returned from FilterChangeMasterKey and is used to iterate over the raw logs and unpacked data for ChangeMasterKey events raised by the VNDWallet contract.
type VNDWalletChangeMasterKeyIterator struct {
	Event *VNDWalletChangeMasterKey // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VNDWalletChangeMasterKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VNDWalletChangeMasterKey)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VNDWalletChangeMasterKey)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VNDWalletChangeMasterKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VNDWalletChangeMasterKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VNDWalletChangeMasterKey represents a ChangeMasterKey event raised by the VNDWallet contract.
type VNDWalletChangeMasterKey struct {
	WhichMasterKey *big.Int
	Data           []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterChangeMasterKey is a free log retrieval operation binding the contract event 0x47031ce53f6639aac67d13de394fce91019d012a92b4254ed9cd2dcceecf3b86.
//
// Solidity: event ChangeMasterKey(uint256 whichMasterKey, bytes data)
func (_VNDWallet *VNDWalletFilterer) FilterChangeMasterKey(opts *bind.FilterOpts) (*VNDWalletChangeMasterKeyIterator, error) {

	logs, sub, err := _VNDWallet.contract.FilterLogs(opts, "ChangeMasterKey")
	if err != nil {
		return nil, err
	}
	return &VNDWalletChangeMasterKeyIterator{contract: _VNDWallet.contract, event: "ChangeMasterKey", logs: logs, sub: sub}, nil
}

// WatchChangeMasterKey is a free log subscription operation binding the contract event 0x47031ce53f6639aac67d13de394fce91019d012a92b4254ed9cd2dcceecf3b86.
//
// Solidity: event ChangeMasterKey(uint256 whichMasterKey, bytes data)
func (_VNDWallet *VNDWalletFilterer) WatchChangeMasterKey(opts *bind.WatchOpts, sink chan<- *VNDWalletChangeMasterKey) (event.Subscription, error) {

	logs, sub, err := _VNDWallet.contract.WatchLogs(opts, "ChangeMasterKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VNDWalletChangeMasterKey)
				if err := _VNDWallet.contract.UnpackLog(event, "ChangeMasterKey", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// VNDWalletChangeOwnerIterator is returned from FilterChangeOwner and is used to iterate over the raw logs and unpacked data for ChangeOwner events raised by the VNDWallet contract.
type VNDWalletChangeOwnerIterator struct {
	Event *VNDWalletChangeOwner // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VNDWalletChangeOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VNDWalletChangeOwner)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VNDWalletChangeOwner)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VNDWalletChangeOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VNDWalletChangeOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VNDWalletChangeOwner represents a ChangeOwner event raised by the VNDWallet contract.
type VNDWalletChangeOwner struct {
	Addr               common.Address
	Owner              common.Address
	Typ                *big.Int
	OwnerTransactionId *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterChangeOwner is a free log retrieval operation binding the contract event 0xe38c516ca079e79d39ba4f754c78096855cce4d441f985694e95d2498ed8f419.
//
// Solidity: event ChangeOwner(address indexed addr, address indexed owner, uint256 typ, uint256 ownerTransactionId)
func (_VNDWallet *VNDWalletFilterer) FilterChangeOwner(opts *bind.FilterOpts, addr []common.Address, owner []common.Address) (*VNDWalletChangeOwnerIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VNDWallet.contract.FilterLogs(opts, "ChangeOwner", addrRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &VNDWalletChangeOwnerIterator{contract: _VNDWallet.contract, event: "ChangeOwner", logs: logs, sub: sub}, nil
}

// WatchChangeOwner is a free log subscription operation binding the contract event 0xe38c516ca079e79d39ba4f754c78096855cce4d441f985694e95d2498ed8f419.
//
// Solidity: event ChangeOwner(address indexed addr, address indexed owner, uint256 typ, uint256 ownerTransactionId)
func (_VNDWallet *VNDWalletFilterer) WatchChangeOwner(opts *bind.WatchOpts, sink chan<- *VNDWalletChangeOwner, addr []common.Address, owner []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _VNDWallet.contract.WatchLogs(opts, "ChangeOwner", addrRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VNDWalletChangeOwner)
				if err := _VNDWallet.contract.UnpackLog(event, "ChangeOwner", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// VNDWalletMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the VNDWallet contract.
type VNDWalletMintIterator struct {
	Event *VNDWalletMint // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VNDWalletMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VNDWalletMint)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VNDWalletMint)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VNDWalletMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VNDWalletMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VNDWalletMint represents a Mint event raised by the VNDWallet contract.
type VNDWalletMint struct {
	By    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed by, uint256 value)
func (_VNDWallet *VNDWalletFilterer) FilterMint(opts *bind.FilterOpts, by []common.Address) (*VNDWalletMintIterator, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _VNDWallet.contract.FilterLogs(opts, "Mint", byRule)
	if err != nil {
		return nil, err
	}
	return &VNDWalletMintIterator{contract: _VNDWallet.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed by, uint256 value)
func (_VNDWallet *VNDWalletFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *VNDWalletMint, by []common.Address) (event.Subscription, error) {

	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _VNDWallet.contract.WatchLogs(opts, "Mint", byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VNDWalletMint)
				if err := _VNDWallet.contract.UnpackLog(event, "Mint", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// VNDWalletTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the VNDWallet contract.
type VNDWalletTransferIterator struct {
	Event *VNDWalletTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VNDWalletTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VNDWalletTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VNDWalletTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VNDWalletTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VNDWalletTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VNDWalletTransfer represents a Transfer event raised by the VNDWallet contract.
type VNDWalletTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_VNDWallet *VNDWalletFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VNDWalletTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VNDWallet.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VNDWalletTransferIterator{contract: _VNDWallet.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_VNDWallet *VNDWalletFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VNDWalletTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VNDWallet.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VNDWalletTransfer)
				if err := _VNDWallet.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
