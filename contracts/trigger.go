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

// TriggerABI is the input ABI used to generate the binding from.
const TriggerABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_trigger\",\"type\":\"uint256\"}],\"name\":\"trigger\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_trigger\",\"type\":\"uint256\"}],\"name\":\"TriggerEvt\",\"type\":\"event\"}]"

// TriggerBin is the compiled bytecode used for deploying new contracts.
const TriggerBin = `0x608060405234801561001057600080fd5b5060008054600160a060020a03191633179055610124806100326000396000f3fe6080604052348015600f57600080fd5b5060043610604e577c01000000000000000000000000000000000000000000000000000000006000350463893d20e881146053578063ed684cc6146082575b600080fd5b6059609e565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b609c60048036036020811015609657600080fd5b503560ba565b005b60005473ffffffffffffffffffffffffffffffffffffffff1690565b604080513381526020810183905281517f7453df022b3c775a1d8aad3cd61495415e1799d0e8fb0462baf8ef58e6797a4b929181900390910190a15056fea165627a7a72305820b07033f55a05c36040ed0047935e52750e1fb4d3a692ea37e86d1e37b93f247d0029`

// DeployTrigger deploys a new Ethereum contract, binding an instance of Trigger to it.
func DeployTrigger(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Trigger, error) {
	parsed, err := abi.JSON(strings.NewReader(TriggerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TriggerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Trigger{TriggerCaller: TriggerCaller{contract: contract}, TriggerTransactor: TriggerTransactor{contract: contract}, TriggerFilterer: TriggerFilterer{contract: contract}}, nil
}

// Trigger is an auto generated Go binding around an Ethereum contract.
type Trigger struct {
	TriggerCaller     // Read-only binding to the contract
	TriggerTransactor // Write-only binding to the contract
	TriggerFilterer   // Log filterer for contract events
}

// TriggerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TriggerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TriggerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TriggerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TriggerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TriggerSession struct {
	Contract     *Trigger          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TriggerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TriggerCallerSession struct {
	Contract *TriggerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TriggerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TriggerTransactorSession struct {
	Contract     *TriggerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TriggerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TriggerRaw struct {
	Contract *Trigger // Generic contract binding to access the raw methods on
}

// TriggerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TriggerCallerRaw struct {
	Contract *TriggerCaller // Generic read-only contract binding to access the raw methods on
}

// TriggerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TriggerTransactorRaw struct {
	Contract *TriggerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrigger creates a new instance of Trigger, bound to a specific deployed contract.
func NewTrigger(address common.Address, backend bind.ContractBackend) (*Trigger, error) {
	contract, err := bindTrigger(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trigger{TriggerCaller: TriggerCaller{contract: contract}, TriggerTransactor: TriggerTransactor{contract: contract}, TriggerFilterer: TriggerFilterer{contract: contract}}, nil
}

// NewTriggerCaller creates a new read-only instance of Trigger, bound to a specific deployed contract.
func NewTriggerCaller(address common.Address, caller bind.ContractCaller) (*TriggerCaller, error) {
	contract, err := bindTrigger(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TriggerCaller{contract: contract}, nil
}

// NewTriggerTransactor creates a new write-only instance of Trigger, bound to a specific deployed contract.
func NewTriggerTransactor(address common.Address, transactor bind.ContractTransactor) (*TriggerTransactor, error) {
	contract, err := bindTrigger(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TriggerTransactor{contract: contract}, nil
}

// NewTriggerFilterer creates a new log filterer instance of Trigger, bound to a specific deployed contract.
func NewTriggerFilterer(address common.Address, filterer bind.ContractFilterer) (*TriggerFilterer, error) {
	contract, err := bindTrigger(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TriggerFilterer{contract: contract}, nil
}

// bindTrigger binds a generic wrapper to an already deployed contract.
func bindTrigger(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TriggerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trigger *TriggerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Trigger.Contract.TriggerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trigger *TriggerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trigger.Contract.TriggerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trigger *TriggerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trigger.Contract.TriggerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trigger *TriggerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Trigger.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trigger *TriggerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trigger.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trigger *TriggerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trigger.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_Trigger *TriggerTransactor) GetOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trigger.contract.Transact(opts, "getOwner")
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_Trigger *TriggerSession) GetOwner() (*types.Transaction, error) {
	return _Trigger.Contract.GetOwner(&_Trigger.TransactOpts)
}

// GetOwner is a paid mutator transaction binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() returns(address)
func (_Trigger *TriggerTransactorSession) GetOwner() (*types.Transaction, error) {
	return _Trigger.Contract.GetOwner(&_Trigger.TransactOpts)
}

// Trigger is a paid mutator transaction binding the contract method 0xed684cc6.
//
// Solidity: function trigger(uint256 _trigger) returns()
func (_Trigger *TriggerTransactor) Trigger(opts *bind.TransactOpts, _trigger *big.Int) (*types.Transaction, error) {
	return _Trigger.contract.Transact(opts, "trigger", _trigger)
}

// Trigger is a paid mutator transaction binding the contract method 0xed684cc6.
//
// Solidity: function trigger(uint256 _trigger) returns()
func (_Trigger *TriggerSession) Trigger(_trigger *big.Int) (*types.Transaction, error) {
	return _Trigger.Contract.Trigger(&_Trigger.TransactOpts, _trigger)
}

// Trigger is a paid mutator transaction binding the contract method 0xed684cc6.
//
// Solidity: function trigger(uint256 _trigger) returns()
func (_Trigger *TriggerTransactorSession) Trigger(_trigger *big.Int) (*types.Transaction, error) {
	return _Trigger.Contract.Trigger(&_Trigger.TransactOpts, _trigger)
}

// TriggerTriggerEvtIterator is returned from FilterTriggerEvt and is used to iterate over the raw logs and unpacked data for TriggerEvt events raised by the Trigger contract.
type TriggerTriggerEvtIterator struct {
	Event *TriggerTriggerEvt // Event containing the contract specifics and raw log

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
func (it *TriggerTriggerEvtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TriggerTriggerEvt)
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
		it.Event = new(TriggerTriggerEvt)
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
func (it *TriggerTriggerEvtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TriggerTriggerEvtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TriggerTriggerEvt represents a TriggerEvt event raised by the Trigger contract.
type TriggerTriggerEvt struct {
	Sender  common.Address
	Trigger *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTriggerEvt is a free log retrieval operation binding the contract event 0x7453df022b3c775a1d8aad3cd61495415e1799d0e8fb0462baf8ef58e6797a4b.
//
// Solidity: event TriggerEvt(address _sender, uint256 _trigger)
func (_Trigger *TriggerFilterer) FilterTriggerEvt(opts *bind.FilterOpts) (*TriggerTriggerEvtIterator, error) {

	logs, sub, err := _Trigger.contract.FilterLogs(opts, "TriggerEvt")
	if err != nil {
		return nil, err
	}
	return &TriggerTriggerEvtIterator{contract: _Trigger.contract, event: "TriggerEvt", logs: logs, sub: sub}, nil
}

// WatchTriggerEvt is a free log subscription operation binding the contract event 0x7453df022b3c775a1d8aad3cd61495415e1799d0e8fb0462baf8ef58e6797a4b.
//
// Solidity: event TriggerEvt(address _sender, uint256 _trigger)
func (_Trigger *TriggerFilterer) WatchTriggerEvt(opts *bind.WatchOpts, sink chan<- *TriggerTriggerEvt) (event.Subscription, error) {

	logs, sub, err := _Trigger.contract.WatchLogs(opts, "TriggerEvt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TriggerTriggerEvt)
				if err := _Trigger.contract.UnpackLog(event, "TriggerEvt", log); err != nil {
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
