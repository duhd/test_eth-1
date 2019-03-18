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

// ExistingWithoutABIABI is the input ABI used to generate the binding from.
const ExistingWithoutABIABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_val\",\"type\":\"uint256\"}],\"name\":\"setA_ASM\",\"outputs\":[{\"name\":\"answer\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_val\",\"type\":\"uint256\"}],\"name\":\"setA_Signature\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_t\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ExistingWithoutABIBin is the compiled bytecode used for deploying new contracts.
const ExistingWithoutABIBin = `0x608060405234801561001057600080fd5b506040516020806102da8339810180604052602081101561003057600080fd5b505160008054600160a060020a03909216600160a060020a0319909216919091179055610278806100626000396000f3fe608060405234801561001057600080fd5b5060043610610052577c010000000000000000000000000000000000000000000000000000000060003504637a8b011481146100575780638976762d14610086575b600080fd5b6100746004803603602081101561006d57600080fd5b50356100b7565b60408051918252519081900360200190f35b6100a36004803603602081101561009c57600080fd5b503561012a565b604080519115158252519081900360200190f35b60008060405180807f736574412875696e743235362900000000000000000000000000000000000000815250600d0190506040518091039020905060405181815283600482015260208160248360008054613a98f180151561011857600080fd5b50805160249091016040529392505050565b6000805460408051602480820186905282518083039091018152604490910182526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fee919d50000000000000000000000000000000000000000000000000000000001781529151815173ffffffffffffffffffffffffffffffffffffffff90941693919290918291908083835b602083106101da5780518252601f1990920191602091820191016101bb565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d806000811461023c576040519150601f19603f3d011682016040523d82523d6000602084013e610241565b606091505b50600194935050505056fea165627a7a7230582083eb20d96cf5697b38a18e7d93513627e9964f5ee148b8a70b7960b0d0e4f5050029`

// DeployExistingWithoutABI deploys a new Ethereum contract, binding an instance of ExistingWithoutABI to it.
func DeployExistingWithoutABI(auth *bind.TransactOpts, backend bind.ContractBackend, _t common.Address) (common.Address, *types.Transaction, *ExistingWithoutABI, error) {
	parsed, err := abi.JSON(strings.NewReader(ExistingWithoutABIABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExistingWithoutABIBin), backend, _t)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ExistingWithoutABI{ExistingWithoutABICaller: ExistingWithoutABICaller{contract: contract}, ExistingWithoutABITransactor: ExistingWithoutABITransactor{contract: contract}, ExistingWithoutABIFilterer: ExistingWithoutABIFilterer{contract: contract}}, nil
}

// ExistingWithoutABI is an auto generated Go binding around an Ethereum contract.
type ExistingWithoutABI struct {
	ExistingWithoutABICaller     // Read-only binding to the contract
	ExistingWithoutABITransactor // Write-only binding to the contract
	ExistingWithoutABIFilterer   // Log filterer for contract events
}

// ExistingWithoutABICaller is an auto generated read-only Go binding around an Ethereum contract.
type ExistingWithoutABICaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExistingWithoutABITransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExistingWithoutABITransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExistingWithoutABIFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExistingWithoutABIFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExistingWithoutABISession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExistingWithoutABISession struct {
	Contract     *ExistingWithoutABI // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExistingWithoutABICallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExistingWithoutABICallerSession struct {
	Contract *ExistingWithoutABICaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ExistingWithoutABITransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExistingWithoutABITransactorSession struct {
	Contract     *ExistingWithoutABITransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ExistingWithoutABIRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExistingWithoutABIRaw struct {
	Contract *ExistingWithoutABI // Generic contract binding to access the raw methods on
}

// ExistingWithoutABICallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExistingWithoutABICallerRaw struct {
	Contract *ExistingWithoutABICaller // Generic read-only contract binding to access the raw methods on
}

// ExistingWithoutABITransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExistingWithoutABITransactorRaw struct {
	Contract *ExistingWithoutABITransactor // Generic write-only contract binding to access the raw methods on
}

// NewExistingWithoutABI creates a new instance of ExistingWithoutABI, bound to a specific deployed contract.
func NewExistingWithoutABI(address common.Address, backend bind.ContractBackend) (*ExistingWithoutABI, error) {
	contract, err := bindExistingWithoutABI(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExistingWithoutABI{ExistingWithoutABICaller: ExistingWithoutABICaller{contract: contract}, ExistingWithoutABITransactor: ExistingWithoutABITransactor{contract: contract}, ExistingWithoutABIFilterer: ExistingWithoutABIFilterer{contract: contract}}, nil
}

// NewExistingWithoutABICaller creates a new read-only instance of ExistingWithoutABI, bound to a specific deployed contract.
func NewExistingWithoutABICaller(address common.Address, caller bind.ContractCaller) (*ExistingWithoutABICaller, error) {
	contract, err := bindExistingWithoutABI(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExistingWithoutABICaller{contract: contract}, nil
}

// NewExistingWithoutABITransactor creates a new write-only instance of ExistingWithoutABI, bound to a specific deployed contract.
func NewExistingWithoutABITransactor(address common.Address, transactor bind.ContractTransactor) (*ExistingWithoutABITransactor, error) {
	contract, err := bindExistingWithoutABI(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExistingWithoutABITransactor{contract: contract}, nil
}

// NewExistingWithoutABIFilterer creates a new log filterer instance of ExistingWithoutABI, bound to a specific deployed contract.
func NewExistingWithoutABIFilterer(address common.Address, filterer bind.ContractFilterer) (*ExistingWithoutABIFilterer, error) {
	contract, err := bindExistingWithoutABI(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExistingWithoutABIFilterer{contract: contract}, nil
}

// bindExistingWithoutABI binds a generic wrapper to an already deployed contract.
func bindExistingWithoutABI(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExistingWithoutABIABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExistingWithoutABI *ExistingWithoutABIRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExistingWithoutABI.Contract.ExistingWithoutABICaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExistingWithoutABI *ExistingWithoutABIRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExistingWithoutABI.Contract.ExistingWithoutABITransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExistingWithoutABI *ExistingWithoutABIRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExistingWithoutABI.Contract.ExistingWithoutABITransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExistingWithoutABI *ExistingWithoutABICallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExistingWithoutABI.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExistingWithoutABI *ExistingWithoutABITransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExistingWithoutABI.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExistingWithoutABI *ExistingWithoutABITransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExistingWithoutABI.Contract.contract.Transact(opts, method, params...)
}

// SetAASM is a paid mutator transaction binding the contract method 0x7a8b0114.
//
// Solidity: function setA_ASM(uint256 _val) returns(uint256 answer)
func (_ExistingWithoutABI *ExistingWithoutABITransactor) SetAASM(opts *bind.TransactOpts, _val *big.Int) (*types.Transaction, error) {
	return _ExistingWithoutABI.contract.Transact(opts, "setA_ASM", _val)
}

// SetAASM is a paid mutator transaction binding the contract method 0x7a8b0114.
//
// Solidity: function setA_ASM(uint256 _val) returns(uint256 answer)
func (_ExistingWithoutABI *ExistingWithoutABISession) SetAASM(_val *big.Int) (*types.Transaction, error) {
	return _ExistingWithoutABI.Contract.SetAASM(&_ExistingWithoutABI.TransactOpts, _val)
}

// SetAASM is a paid mutator transaction binding the contract method 0x7a8b0114.
//
// Solidity: function setA_ASM(uint256 _val) returns(uint256 answer)
func (_ExistingWithoutABI *ExistingWithoutABITransactorSession) SetAASM(_val *big.Int) (*types.Transaction, error) {
	return _ExistingWithoutABI.Contract.SetAASM(&_ExistingWithoutABI.TransactOpts, _val)
}

// SetASignature is a paid mutator transaction binding the contract method 0x8976762d.
//
// Solidity: function setA_Signature(uint256 _val) returns(bool success)
func (_ExistingWithoutABI *ExistingWithoutABITransactor) SetASignature(opts *bind.TransactOpts, _val *big.Int) (*types.Transaction, error) {
	return _ExistingWithoutABI.contract.Transact(opts, "setA_Signature", _val)
}

// SetASignature is a paid mutator transaction binding the contract method 0x8976762d.
//
// Solidity: function setA_Signature(uint256 _val) returns(bool success)
func (_ExistingWithoutABI *ExistingWithoutABISession) SetASignature(_val *big.Int) (*types.Transaction, error) {
	return _ExistingWithoutABI.Contract.SetASignature(&_ExistingWithoutABI.TransactOpts, _val)
}

// SetASignature is a paid mutator transaction binding the contract method 0x8976762d.
//
// Solidity: function setA_Signature(uint256 _val) returns(bool success)
func (_ExistingWithoutABI *ExistingWithoutABITransactorSession) SetASignature(_val *big.Int) (*types.Transaction, error) {
	return _ExistingWithoutABI.Contract.SetASignature(&_ExistingWithoutABI.TransactOpts, _val)
}
