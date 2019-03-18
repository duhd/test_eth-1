// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package metacoin

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

// CoinCallerABI is the input ABI used to generate the binding from.
const CoinCallerABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"coinContractAddress\",\"type\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendCoin\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"txstatus\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"sendCoinEvt\",\"type\":\"event\"}]"

// CoinCallerBin is the compiled bytecode used for deploying new contracts.
const CoinCallerBin = `0x608060405234801561001057600080fd5b50610322806100206000396000f3fe608060405234801561001057600080fd5b5060043610610047577c010000000000000000000000000000000000000000000000000000000060003504630b40bd88811461004c575b600080fd5b61008f6004803603606081101561006257600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160208101359091169060400135610091565b005b60018054600090815260208181526040808320805473ffffffffffffffffffffffffffffffffffffffff808a1673ffffffffffffffffffffffffffffffffffffffff19928316178084559683018890556002830180548a831693168317905583517f412664ae000000000000000000000000000000000000000000000000000000008152600481019290925260248201889052925191959092169363412664ae93604480850194919392918390030190829087803b15801561015257600080fd5b505af1158015610166573d6000803e3d6000fd5b505050506040513d602081101561017c57600080fd5b5051600282018054911515740100000000000000000000000000000000000000000274ff0000000000000000000000000000000000000000199092169190911790558054604080517f27e235e3000000000000000000000000000000000000000000000000000000008152326004820152905173ffffffffffffffffffffffffffffffffffffffff909216916327e235e391602480820192602092909190829003018186803b15801561022e57600080fd5b505afa158015610242573d6000803e3d6000fd5b505050506040513d602081101561025857600080fd5b5051600382019081556001805481019055600282015490546040805132815273ffffffffffffffffffffffffffffffffffffffff871660208201528082018690527401000000000000000000000000000000000000000090930460ff16151560608401526080830191909152517fa120daf44c15b264b9cd106ea0ba86d890166c59a770ae1965a1ab2af642c62c9160a0908290030190a15050505056fea165627a7a72305820f6542ee4e9aec9f47c27859f66825f8bb293c40d6a1da699eefc5b725803b5910029`

// DeployCoinCaller deploys a new Ethereum contract, binding an instance of CoinCaller to it.
func DeployCoinCaller(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CoinCaller, error) {
	parsed, err := abi.JSON(strings.NewReader(CoinCallerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CoinCallerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CoinCaller{CoinCallerCaller: CoinCallerCaller{contract: contract}, CoinCallerTransactor: CoinCallerTransactor{contract: contract}, CoinCallerFilterer: CoinCallerFilterer{contract: contract}}, nil
}

// CoinCaller is an auto generated Go binding around an Ethereum contract.
type CoinCaller struct {
	CoinCallerCaller     // Read-only binding to the contract
	CoinCallerTransactor // Write-only binding to the contract
	CoinCallerFilterer   // Log filterer for contract events
}

// CoinCallerCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoinCallerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinCallerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoinCallerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinCallerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoinCallerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinCallerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoinCallerSession struct {
	Contract     *CoinCaller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoinCallerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoinCallerCallerSession struct {
	Contract *CoinCallerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// CoinCallerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoinCallerTransactorSession struct {
	Contract     *CoinCallerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// CoinCallerRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoinCallerRaw struct {
	Contract *CoinCaller // Generic contract binding to access the raw methods on
}

// CoinCallerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoinCallerCallerRaw struct {
	Contract *CoinCallerCaller // Generic read-only contract binding to access the raw methods on
}

// CoinCallerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoinCallerTransactorRaw struct {
	Contract *CoinCallerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoinCaller creates a new instance of CoinCaller, bound to a specific deployed contract.
func NewCoinCaller(address common.Address, backend bind.ContractBackend) (*CoinCaller, error) {
	contract, err := bindCoinCaller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CoinCaller{CoinCallerCaller: CoinCallerCaller{contract: contract}, CoinCallerTransactor: CoinCallerTransactor{contract: contract}, CoinCallerFilterer: CoinCallerFilterer{contract: contract}}, nil
}

// NewCoinCallerCaller creates a new read-only instance of CoinCaller, bound to a specific deployed contract.
func NewCoinCallerCaller(address common.Address, caller bind.ContractCaller) (*CoinCallerCaller, error) {
	contract, err := bindCoinCaller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoinCallerCaller{contract: contract}, nil
}

// NewCoinCallerTransactor creates a new write-only instance of CoinCaller, bound to a specific deployed contract.
func NewCoinCallerTransactor(address common.Address, transactor bind.ContractTransactor) (*CoinCallerTransactor, error) {
	contract, err := bindCoinCaller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoinCallerTransactor{contract: contract}, nil
}

// NewCoinCallerFilterer creates a new log filterer instance of CoinCaller, bound to a specific deployed contract.
func NewCoinCallerFilterer(address common.Address, filterer bind.ContractFilterer) (*CoinCallerFilterer, error) {
	contract, err := bindCoinCaller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoinCallerFilterer{contract: contract}, nil
}

// bindCoinCaller binds a generic wrapper to an already deployed contract.
func bindCoinCaller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoinCallerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoinCaller *CoinCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CoinCaller.Contract.CoinCallerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoinCaller *CoinCallerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinCaller.Contract.CoinCallerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoinCaller *CoinCallerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoinCaller.Contract.CoinCallerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoinCaller *CoinCallerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CoinCaller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoinCaller *CoinCallerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinCaller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoinCaller *CoinCallerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoinCaller.Contract.contract.Transact(opts, method, params...)
}

// SendCoin is a paid mutator transaction binding the contract method 0x0b40bd88.
//
// Solidity: function sendCoin(address coinContractAddress, address receiver, uint256 amount) returns()
func (_CoinCaller *CoinCallerTransactor) SendCoin(opts *bind.TransactOpts, coinContractAddress common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoinCaller.contract.Transact(opts, "sendCoin", coinContractAddress, receiver, amount)
}

// SendCoin is a paid mutator transaction binding the contract method 0x0b40bd88.
//
// Solidity: function sendCoin(address coinContractAddress, address receiver, uint256 amount) returns()
func (_CoinCaller *CoinCallerSession) SendCoin(coinContractAddress common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoinCaller.Contract.SendCoin(&_CoinCaller.TransactOpts, coinContractAddress, receiver, amount)
}

// SendCoin is a paid mutator transaction binding the contract method 0x0b40bd88.
//
// Solidity: function sendCoin(address coinContractAddress, address receiver, uint256 amount) returns()
func (_CoinCaller *CoinCallerTransactorSession) SendCoin(coinContractAddress common.Address, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CoinCaller.Contract.SendCoin(&_CoinCaller.TransactOpts, coinContractAddress, receiver, amount)
}

// CoinCallerSendCoinEvtIterator is returned from FilterSendCoinEvt and is used to iterate over the raw logs and unpacked data for SendCoinEvt events raised by the CoinCaller contract.
type CoinCallerSendCoinEvtIterator struct {
	Event *CoinCallerSendCoinEvt // Event containing the contract specifics and raw log

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
func (it *CoinCallerSendCoinEvtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinCallerSendCoinEvt)
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
		it.Event = new(CoinCallerSendCoinEvt)
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
func (it *CoinCallerSendCoinEvtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinCallerSendCoinEvtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinCallerSendCoinEvt represents a SendCoinEvt event raised by the CoinCaller contract.
type CoinCallerSendCoinEvt struct {
	From     common.Address
	To       common.Address
	Amount   *big.Int
	Txstatus bool
	Balance  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSendCoinEvt is a free log retrieval operation binding the contract event 0xa120daf44c15b264b9cd106ea0ba86d890166c59a770ae1965a1ab2af642c62c.
//
// Solidity: event sendCoinEvt(address from, address to, uint256 amount, bool txstatus, uint256 balance)
func (_CoinCaller *CoinCallerFilterer) FilterSendCoinEvt(opts *bind.FilterOpts) (*CoinCallerSendCoinEvtIterator, error) {

	logs, sub, err := _CoinCaller.contract.FilterLogs(opts, "sendCoinEvt")
	if err != nil {
		return nil, err
	}
	return &CoinCallerSendCoinEvtIterator{contract: _CoinCaller.contract, event: "sendCoinEvt", logs: logs, sub: sub}, nil
}

// WatchSendCoinEvt is a free log subscription operation binding the contract event 0xa120daf44c15b264b9cd106ea0ba86d890166c59a770ae1965a1ab2af642c62c.
//
// Solidity: event sendCoinEvt(address from, address to, uint256 amount, bool txstatus, uint256 balance)
func (_CoinCaller *CoinCallerFilterer) WatchSendCoinEvt(opts *bind.WatchOpts, sink chan<- *CoinCallerSendCoinEvt) (event.Subscription, error) {

	logs, sub, err := _CoinCaller.contract.WatchLogs(opts, "sendCoinEvt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinCallerSendCoinEvt)
				if err := _CoinCaller.contract.UnpackLog(event, "sendCoinEvt", log); err != nil {
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

// CoinSpawnABI is the input ABI used to generate the binding from.
const CoinSpawnABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"name\":\"createCoin\",\"outputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deployedContracts\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"name\":\"CreateCoinEvt\",\"type\":\"event\"}]"

// CoinSpawnBin is the compiled bytecode used for deploying new contracts.
const CoinSpawnBin = `0x6080604052600060015534801561001557600080fd5b50610420806100256000396000f3fe608060405234801561001057600080fd5b5060043610610052577c0100000000000000000000000000000000000000000000000000000000600035046380f7884d81146100575780639ad1ee101461009d575b600080fd5b6100746004803603602081101561006d57600080fd5b50356100ba565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b610074600480360360208110156100b357600080fd5b5035610180565b600080826040516100ca906101a8565b90815260405190819003602001906000f0801580156100ed573d6000803e3d6000fd5b5060015460009081526020818152604091829020805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff8516908117909155825190815290810186905281519293507fe1c5c24fe88d7c3d89714b1e5ec744610dc7f3ea260e47074caaa3a7d3d9a68a929081900390910190a1600180548101905592915050565b60006020819052908152604090205473ffffffffffffffffffffffffffffffffffffffff1681565b61023f806101b68339019056fe608060405234801561001057600080fd5b5060405160208061023f8339810180604052602081101561003057600080fd5b5051326000908152602081905260409020556101ee806100516000396000f3fe608060405234801561001057600080fd5b506004361061005d577c0100000000000000000000000000000000000000000000000000000000600035046327e235e38114610062578063412664ae146100a757806370a08231146100f4575b600080fd5b6100956004803603602081101561007857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610127565b60408051918252519081900360200190f35b6100e0600480360360408110156100bd57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610139565b604080519115158252519081900360200190f35b6100956004803603602081101561010a57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661019a565b60006020819052908152604090205481565b3260009081526020819052604081205482111561015857506000610194565b50326000908152602081905260408082208054849003905573ffffffffffffffffffffffffffffffffffffffff84168252902080548201905560015b92915050565b73ffffffffffffffffffffffffffffffffffffffff166000908152602081905260409020549056fea165627a7a72305820dbdcc05decef5374d12e5683c41fb2d094fa08ab7abce7105657fa0e1d317e6a0029a165627a7a72305820e957e6a2ef811d3d69344c86e9248e4b8bb0dee2a091d7d74fb7ff23476473230029`

// DeployCoinSpawn deploys a new Ethereum contract, binding an instance of CoinSpawn to it.
func DeployCoinSpawn(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CoinSpawn, error) {
	parsed, err := abi.JSON(strings.NewReader(CoinSpawnABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CoinSpawnBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CoinSpawn{CoinSpawnCaller: CoinSpawnCaller{contract: contract}, CoinSpawnTransactor: CoinSpawnTransactor{contract: contract}, CoinSpawnFilterer: CoinSpawnFilterer{contract: contract}}, nil
}

// CoinSpawn is an auto generated Go binding around an Ethereum contract.
type CoinSpawn struct {
	CoinSpawnCaller     // Read-only binding to the contract
	CoinSpawnTransactor // Write-only binding to the contract
	CoinSpawnFilterer   // Log filterer for contract events
}

// CoinSpawnCaller is an auto generated read-only Go binding around an Ethereum contract.
type CoinSpawnCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinSpawnTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CoinSpawnTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinSpawnFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CoinSpawnFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CoinSpawnSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CoinSpawnSession struct {
	Contract     *CoinSpawn        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CoinSpawnCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CoinSpawnCallerSession struct {
	Contract *CoinSpawnCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CoinSpawnTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CoinSpawnTransactorSession struct {
	Contract     *CoinSpawnTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CoinSpawnRaw is an auto generated low-level Go binding around an Ethereum contract.
type CoinSpawnRaw struct {
	Contract *CoinSpawn // Generic contract binding to access the raw methods on
}

// CoinSpawnCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CoinSpawnCallerRaw struct {
	Contract *CoinSpawnCaller // Generic read-only contract binding to access the raw methods on
}

// CoinSpawnTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CoinSpawnTransactorRaw struct {
	Contract *CoinSpawnTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCoinSpawn creates a new instance of CoinSpawn, bound to a specific deployed contract.
func NewCoinSpawn(address common.Address, backend bind.ContractBackend) (*CoinSpawn, error) {
	contract, err := bindCoinSpawn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CoinSpawn{CoinSpawnCaller: CoinSpawnCaller{contract: contract}, CoinSpawnTransactor: CoinSpawnTransactor{contract: contract}, CoinSpawnFilterer: CoinSpawnFilterer{contract: contract}}, nil
}

// NewCoinSpawnCaller creates a new read-only instance of CoinSpawn, bound to a specific deployed contract.
func NewCoinSpawnCaller(address common.Address, caller bind.ContractCaller) (*CoinSpawnCaller, error) {
	contract, err := bindCoinSpawn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CoinSpawnCaller{contract: contract}, nil
}

// NewCoinSpawnTransactor creates a new write-only instance of CoinSpawn, bound to a specific deployed contract.
func NewCoinSpawnTransactor(address common.Address, transactor bind.ContractTransactor) (*CoinSpawnTransactor, error) {
	contract, err := bindCoinSpawn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CoinSpawnTransactor{contract: contract}, nil
}

// NewCoinSpawnFilterer creates a new log filterer instance of CoinSpawn, bound to a specific deployed contract.
func NewCoinSpawnFilterer(address common.Address, filterer bind.ContractFilterer) (*CoinSpawnFilterer, error) {
	contract, err := bindCoinSpawn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CoinSpawnFilterer{contract: contract}, nil
}

// bindCoinSpawn binds a generic wrapper to an already deployed contract.
func bindCoinSpawn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CoinSpawnABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoinSpawn *CoinSpawnRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CoinSpawn.Contract.CoinSpawnCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoinSpawn *CoinSpawnRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinSpawn.Contract.CoinSpawnTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoinSpawn *CoinSpawnRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoinSpawn.Contract.CoinSpawnTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CoinSpawn *CoinSpawnCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _CoinSpawn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CoinSpawn *CoinSpawnTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CoinSpawn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CoinSpawn *CoinSpawnTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CoinSpawn.Contract.contract.Transact(opts, method, params...)
}

// DeployedContracts is a free data retrieval call binding the contract method 0x9ad1ee10.
//
// Solidity: function deployedContracts(uint256 ) constant returns(address)
func (_CoinSpawn *CoinSpawnCaller) DeployedContracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _CoinSpawn.contract.Call(opts, out, "deployedContracts", arg0)
	return *ret0, err
}

// DeployedContracts is a free data retrieval call binding the contract method 0x9ad1ee10.
//
// Solidity: function deployedContracts(uint256 ) constant returns(address)
func (_CoinSpawn *CoinSpawnSession) DeployedContracts(arg0 *big.Int) (common.Address, error) {
	return _CoinSpawn.Contract.DeployedContracts(&_CoinSpawn.CallOpts, arg0)
}

// DeployedContracts is a free data retrieval call binding the contract method 0x9ad1ee10.
//
// Solidity: function deployedContracts(uint256 ) constant returns(address)
func (_CoinSpawn *CoinSpawnCallerSession) DeployedContracts(arg0 *big.Int) (common.Address, error) {
	return _CoinSpawn.Contract.DeployedContracts(&_CoinSpawn.CallOpts, arg0)
}

// CreateCoin is a paid mutator transaction binding the contract method 0x80f7884d.
//
// Solidity: function createCoin(uint256 initialBalance) returns(address a)
func (_CoinSpawn *CoinSpawnTransactor) CreateCoin(opts *bind.TransactOpts, initialBalance *big.Int) (*types.Transaction, error) {
	return _CoinSpawn.contract.Transact(opts, "createCoin", initialBalance)
}

// CreateCoin is a paid mutator transaction binding the contract method 0x80f7884d.
//
// Solidity: function createCoin(uint256 initialBalance) returns(address a)
func (_CoinSpawn *CoinSpawnSession) CreateCoin(initialBalance *big.Int) (*types.Transaction, error) {
	return _CoinSpawn.Contract.CreateCoin(&_CoinSpawn.TransactOpts, initialBalance)
}

// CreateCoin is a paid mutator transaction binding the contract method 0x80f7884d.
//
// Solidity: function createCoin(uint256 initialBalance) returns(address a)
func (_CoinSpawn *CoinSpawnTransactorSession) CreateCoin(initialBalance *big.Int) (*types.Transaction, error) {
	return _CoinSpawn.Contract.CreateCoin(&_CoinSpawn.TransactOpts, initialBalance)
}

// CoinSpawnCreateCoinEvtIterator is returned from FilterCreateCoinEvt and is used to iterate over the raw logs and unpacked data for CreateCoinEvt events raised by the CoinSpawn contract.
type CoinSpawnCreateCoinEvtIterator struct {
	Event *CoinSpawnCreateCoinEvt // Event containing the contract specifics and raw log

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
func (it *CoinSpawnCreateCoinEvtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CoinSpawnCreateCoinEvt)
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
		it.Event = new(CoinSpawnCreateCoinEvt)
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
func (it *CoinSpawnCreateCoinEvtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CoinSpawnCreateCoinEvtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CoinSpawnCreateCoinEvt represents a CreateCoinEvt event raised by the CoinSpawn contract.
type CoinSpawnCreateCoinEvt struct {
	Addr           common.Address
	InitialBalance *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCreateCoinEvt is a free log retrieval operation binding the contract event 0xe1c5c24fe88d7c3d89714b1e5ec744610dc7f3ea260e47074caaa3a7d3d9a68a.
//
// Solidity: event CreateCoinEvt(address addr, uint256 initialBalance)
func (_CoinSpawn *CoinSpawnFilterer) FilterCreateCoinEvt(opts *bind.FilterOpts) (*CoinSpawnCreateCoinEvtIterator, error) {

	logs, sub, err := _CoinSpawn.contract.FilterLogs(opts, "CreateCoinEvt")
	if err != nil {
		return nil, err
	}
	return &CoinSpawnCreateCoinEvtIterator{contract: _CoinSpawn.contract, event: "CreateCoinEvt", logs: logs, sub: sub}, nil
}

// WatchCreateCoinEvt is a free log subscription operation binding the contract event 0xe1c5c24fe88d7c3d89714b1e5ec744610dc7f3ea260e47074caaa3a7d3d9a68a.
//
// Solidity: event CreateCoinEvt(address addr, uint256 initialBalance)
func (_CoinSpawn *CoinSpawnFilterer) WatchCreateCoinEvt(opts *bind.WatchOpts, sink chan<- *CoinSpawnCreateCoinEvt) (event.Subscription, error) {

	logs, sub, err := _CoinSpawn.contract.WatchLogs(opts, "CreateCoinEvt")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CoinSpawnCreateCoinEvt)
				if err := _CoinSpawn.contract.UnpackLog(event, "CreateCoinEvt", log); err != nil {
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

// MetaCoinABI is the input ABI used to generate the binding from.
const MetaCoinABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendToken\",\"outputs\":[{\"name\":\"successful\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"initialBalance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// MetaCoinBin is the compiled bytecode used for deploying new contracts.
const MetaCoinBin = `0x608060405234801561001057600080fd5b5060405160208061023f8339810180604052602081101561003057600080fd5b5051326000908152602081905260409020556101ee806100516000396000f3fe608060405234801561001057600080fd5b506004361061005d577c0100000000000000000000000000000000000000000000000000000000600035046327e235e38114610062578063412664ae146100a757806370a08231146100f4575b600080fd5b6100956004803603602081101561007857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610127565b60408051918252519081900360200190f35b6100e0600480360360408110156100bd57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610139565b604080519115158252519081900360200190f35b6100956004803603602081101561010a57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661019a565b60006020819052908152604090205481565b3260009081526020819052604081205482111561015857506000610194565b50326000908152602081905260408082208054849003905573ffffffffffffffffffffffffffffffffffffffff84168252902080548201905560015b92915050565b73ffffffffffffffffffffffffffffffffffffffff166000908152602081905260409020549056fea165627a7a72305820dbdcc05decef5374d12e5683c41fb2d094fa08ab7abce7105657fa0e1d317e6a0029`

// DeployMetaCoin deploys a new Ethereum contract, binding an instance of MetaCoin to it.
func DeployMetaCoin(auth *bind.TransactOpts, backend bind.ContractBackend, initialBalance *big.Int) (common.Address, *types.Transaction, *MetaCoin, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaCoinABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MetaCoinBin), backend, initialBalance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MetaCoin{MetaCoinCaller: MetaCoinCaller{contract: contract}, MetaCoinTransactor: MetaCoinTransactor{contract: contract}, MetaCoinFilterer: MetaCoinFilterer{contract: contract}}, nil
}

// MetaCoin is an auto generated Go binding around an Ethereum contract.
type MetaCoin struct {
	MetaCoinCaller     // Read-only binding to the contract
	MetaCoinTransactor // Write-only binding to the contract
	MetaCoinFilterer   // Log filterer for contract events
}

// MetaCoinCaller is an auto generated read-only Go binding around an Ethereum contract.
type MetaCoinCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaCoinTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MetaCoinTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaCoinFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MetaCoinFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MetaCoinSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MetaCoinSession struct {
	Contract     *MetaCoin         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MetaCoinCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MetaCoinCallerSession struct {
	Contract *MetaCoinCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MetaCoinTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MetaCoinTransactorSession struct {
	Contract     *MetaCoinTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MetaCoinRaw is an auto generated low-level Go binding around an Ethereum contract.
type MetaCoinRaw struct {
	Contract *MetaCoin // Generic contract binding to access the raw methods on
}

// MetaCoinCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MetaCoinCallerRaw struct {
	Contract *MetaCoinCaller // Generic read-only contract binding to access the raw methods on
}

// MetaCoinTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MetaCoinTransactorRaw struct {
	Contract *MetaCoinTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMetaCoin creates a new instance of MetaCoin, bound to a specific deployed contract.
func NewMetaCoin(address common.Address, backend bind.ContractBackend) (*MetaCoin, error) {
	contract, err := bindMetaCoin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MetaCoin{MetaCoinCaller: MetaCoinCaller{contract: contract}, MetaCoinTransactor: MetaCoinTransactor{contract: contract}, MetaCoinFilterer: MetaCoinFilterer{contract: contract}}, nil
}

// NewMetaCoinCaller creates a new read-only instance of MetaCoin, bound to a specific deployed contract.
func NewMetaCoinCaller(address common.Address, caller bind.ContractCaller) (*MetaCoinCaller, error) {
	contract, err := bindMetaCoin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MetaCoinCaller{contract: contract}, nil
}

// NewMetaCoinTransactor creates a new write-only instance of MetaCoin, bound to a specific deployed contract.
func NewMetaCoinTransactor(address common.Address, transactor bind.ContractTransactor) (*MetaCoinTransactor, error) {
	contract, err := bindMetaCoin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MetaCoinTransactor{contract: contract}, nil
}

// NewMetaCoinFilterer creates a new log filterer instance of MetaCoin, bound to a specific deployed contract.
func NewMetaCoinFilterer(address common.Address, filterer bind.ContractFilterer) (*MetaCoinFilterer, error) {
	contract, err := bindMetaCoin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MetaCoinFilterer{contract: contract}, nil
}

// bindMetaCoin binds a generic wrapper to an already deployed contract.
func bindMetaCoin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MetaCoinABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaCoin *MetaCoinRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MetaCoin.Contract.MetaCoinCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaCoin *MetaCoinRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaCoin.Contract.MetaCoinTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaCoin *MetaCoinRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaCoin.Contract.MetaCoinTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MetaCoin *MetaCoinCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MetaCoin.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MetaCoin *MetaCoinTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MetaCoin.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MetaCoin *MetaCoinTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MetaCoin.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) constant returns(uint256)
func (_MetaCoin *MetaCoinCaller) BalanceOf(opts *bind.CallOpts, _account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MetaCoin.contract.Call(opts, out, "balanceOf", _account)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) constant returns(uint256)
func (_MetaCoin *MetaCoinSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _MetaCoin.Contract.BalanceOf(&_MetaCoin.CallOpts, _account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _account) constant returns(uint256)
func (_MetaCoin *MetaCoinCallerSession) BalanceOf(_account common.Address) (*big.Int, error) {
	return _MetaCoin.Contract.BalanceOf(&_MetaCoin.CallOpts, _account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_MetaCoin *MetaCoinCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MetaCoin.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_MetaCoin *MetaCoinSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _MetaCoin.Contract.Balances(&_MetaCoin.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) constant returns(uint256)
func (_MetaCoin *MetaCoinCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _MetaCoin.Contract.Balances(&_MetaCoin.CallOpts, arg0)
}

// SendToken is a paid mutator transaction binding the contract method 0x412664ae.
//
// Solidity: function sendToken(address receiver, uint256 amount) returns(bool successful)
func (_MetaCoin *MetaCoinTransactor) SendToken(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaCoin.contract.Transact(opts, "sendToken", receiver, amount)
}

// SendToken is a paid mutator transaction binding the contract method 0x412664ae.
//
// Solidity: function sendToken(address receiver, uint256 amount) returns(bool successful)
func (_MetaCoin *MetaCoinSession) SendToken(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaCoin.Contract.SendToken(&_MetaCoin.TransactOpts, receiver, amount)
}

// SendToken is a paid mutator transaction binding the contract method 0x412664ae.
//
// Solidity: function sendToken(address receiver, uint256 amount) returns(bool successful)
func (_MetaCoin *MetaCoinTransactorSession) SendToken(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MetaCoin.Contract.SendToken(&_MetaCoin.TransactOpts, receiver, amount)
}
