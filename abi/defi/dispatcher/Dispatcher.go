// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dispatcher

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
)

// DispatcherMetaData contains all meta data concerning the Dispatcher contract.
var DispatcherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dispatcher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawalAccount\",\"type\":\"address\"}],\"name\":\"chainBridgeToWithdrawalAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"}],\"name\":\"treasuryWithdrawAndDispatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DispatcherABI is the input ABI used to generate the binding from.
// Deprecated: Use DispatcherMetaData.ABI instead.
var DispatcherABI = DispatcherMetaData.ABI

// Dispatcher is an auto generated Go binding around an Ethereum contract.
type Dispatcher struct {
	DispatcherCaller     // Read-only binding to the contract
	DispatcherTransactor // Write-only binding to the contract
	DispatcherFilterer   // Log filterer for contract events
}

// DispatcherCaller is an auto generated read-only Go binding around an Ethereum contract.
type DispatcherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DispatcherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DispatcherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DispatcherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DispatcherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DispatcherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DispatcherSession struct {
	Contract     *Dispatcher       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DispatcherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DispatcherCallerSession struct {
	Contract *DispatcherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DispatcherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DispatcherTransactorSession struct {
	Contract     *DispatcherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DispatcherRaw is an auto generated low-level Go binding around an Ethereum contract.
type DispatcherRaw struct {
	Contract *Dispatcher // Generic contract binding to access the raw methods on
}

// DispatcherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DispatcherCallerRaw struct {
	Contract *DispatcherCaller // Generic read-only contract binding to access the raw methods on
}

// DispatcherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DispatcherTransactorRaw struct {
	Contract *DispatcherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDispatcher creates a new instance of Dispatcher, bound to a specific deployed contract.
func NewDispatcher(address common.Address, backend bind.ContractBackend) (*Dispatcher, error) {
	contract, err := bindDispatcher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dispatcher{DispatcherCaller: DispatcherCaller{contract: contract}, DispatcherTransactor: DispatcherTransactor{contract: contract}, DispatcherFilterer: DispatcherFilterer{contract: contract}}, nil
}

// NewDispatcherCaller creates a new read-only instance of Dispatcher, bound to a specific deployed contract.
func NewDispatcherCaller(address common.Address, caller bind.ContractCaller) (*DispatcherCaller, error) {
	contract, err := bindDispatcher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DispatcherCaller{contract: contract}, nil
}

// NewDispatcherTransactor creates a new write-only instance of Dispatcher, bound to a specific deployed contract.
func NewDispatcherTransactor(address common.Address, transactor bind.ContractTransactor) (*DispatcherTransactor, error) {
	contract, err := bindDispatcher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DispatcherTransactor{contract: contract}, nil
}

// NewDispatcherFilterer creates a new log filterer instance of Dispatcher, bound to a specific deployed contract.
func NewDispatcherFilterer(address common.Address, filterer bind.ContractFilterer) (*DispatcherFilterer, error) {
	contract, err := bindDispatcher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DispatcherFilterer{contract: contract}, nil
}

// bindDispatcher binds a generic wrapper to an already deployed contract.
func bindDispatcher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DispatcherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dispatcher *DispatcherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dispatcher.Contract.DispatcherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dispatcher *DispatcherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dispatcher.Contract.DispatcherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dispatcher *DispatcherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dispatcher.Contract.DispatcherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dispatcher *DispatcherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dispatcher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dispatcher *DispatcherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dispatcher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dispatcher *DispatcherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dispatcher.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dispatcher *DispatcherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dispatcher.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dispatcher *DispatcherSession) Owner() (common.Address, error) {
	return _Dispatcher.Contract.Owner(&_Dispatcher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dispatcher *DispatcherCallerSession) Owner() (common.Address, error) {
	return _Dispatcher.Contract.Owner(&_Dispatcher.CallOpts)
}

// ChainBridgeToWithdrawalAccount is a paid mutator transaction binding the contract method 0xd0090cdf.
//
// Solidity: function chainBridgeToWithdrawalAccount(uint256 pid, address token, address withdrawalAccount) returns()
func (_Dispatcher *DispatcherTransactor) ChainBridgeToWithdrawalAccount(opts *bind.TransactOpts, pid *big.Int, token common.Address, withdrawalAccount common.Address) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "chainBridgeToWithdrawalAccount", pid, token, withdrawalAccount)
}

// ChainBridgeToWithdrawalAccount is a paid mutator transaction binding the contract method 0xd0090cdf.
//
// Solidity: function chainBridgeToWithdrawalAccount(uint256 pid, address token, address withdrawalAccount) returns()
func (_Dispatcher *DispatcherSession) ChainBridgeToWithdrawalAccount(pid *big.Int, token common.Address, withdrawalAccount common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.ChainBridgeToWithdrawalAccount(&_Dispatcher.TransactOpts, pid, token, withdrawalAccount)
}

// ChainBridgeToWithdrawalAccount is a paid mutator transaction binding the contract method 0xd0090cdf.
//
// Solidity: function chainBridgeToWithdrawalAccount(uint256 pid, address token, address withdrawalAccount) returns()
func (_Dispatcher *DispatcherTransactorSession) ChainBridgeToWithdrawalAccount(pid *big.Int, token common.Address, withdrawalAccount common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.ChainBridgeToWithdrawalAccount(&_Dispatcher.TransactOpts, pid, token, withdrawalAccount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dispatcher *DispatcherTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dispatcher *DispatcherSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dispatcher.Contract.RenounceOwnership(&_Dispatcher.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dispatcher *DispatcherTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dispatcher.Contract.RenounceOwnership(&_Dispatcher.TransactOpts)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address user, bool allow) returns()
func (_Dispatcher *DispatcherTransactor) SetOperator(opts *bind.TransactOpts, user common.Address, allow bool) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "setOperator", user, allow)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address user, bool allow) returns()
func (_Dispatcher *DispatcherSession) SetOperator(user common.Address, allow bool) (*types.Transaction, error) {
	return _Dispatcher.Contract.SetOperator(&_Dispatcher.TransactOpts, user, allow)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address user, bool allow) returns()
func (_Dispatcher *DispatcherTransactorSession) SetOperator(user common.Address, allow bool) (*types.Transaction, error) {
	return _Dispatcher.Contract.SetOperator(&_Dispatcher.TransactOpts, user, allow)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dispatcher *DispatcherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dispatcher *DispatcherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.TransferOwnership(&_Dispatcher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dispatcher *DispatcherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.TransferOwnership(&_Dispatcher.TransactOpts, newOwner)
}

// TreasuryWithdrawAndDispatch is a paid mutator transaction binding the contract method 0x5d00a825.
//
// Solidity: function treasuryWithdrawAndDispatch(address from) returns()
func (_Dispatcher *DispatcherTransactor) TreasuryWithdrawAndDispatch(opts *bind.TransactOpts, from common.Address) (*types.Transaction, error) {
	return _Dispatcher.contract.Transact(opts, "treasuryWithdrawAndDispatch", from)
}

// TreasuryWithdrawAndDispatch is a paid mutator transaction binding the contract method 0x5d00a825.
//
// Solidity: function treasuryWithdrawAndDispatch(address from) returns()
func (_Dispatcher *DispatcherSession) TreasuryWithdrawAndDispatch(from common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.TreasuryWithdrawAndDispatch(&_Dispatcher.TransactOpts, from)
}

// TreasuryWithdrawAndDispatch is a paid mutator transaction binding the contract method 0x5d00a825.
//
// Solidity: function treasuryWithdrawAndDispatch(address from) returns()
func (_Dispatcher *DispatcherTransactorSession) TreasuryWithdrawAndDispatch(from common.Address) (*types.Transaction, error) {
	return _Dispatcher.Contract.TreasuryWithdrawAndDispatch(&_Dispatcher.TransactOpts, from)
}

// DispatcherOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dispatcher contract.
type DispatcherOwnershipTransferredIterator struct {
	Event *DispatcherOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DispatcherOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DispatcherOwnershipTransferred)
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
		it.Event = new(DispatcherOwnershipTransferred)
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
func (it *DispatcherOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DispatcherOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DispatcherOwnershipTransferred represents a OwnershipTransferred event raised by the Dispatcher contract.
type DispatcherOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dispatcher *DispatcherFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DispatcherOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dispatcher.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DispatcherOwnershipTransferredIterator{contract: _Dispatcher.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dispatcher *DispatcherFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DispatcherOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dispatcher.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DispatcherOwnershipTransferred)
				if err := _Dispatcher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dispatcher *DispatcherFilterer) ParseOwnershipTransferred(log types.Log) (*DispatcherOwnershipTransferred, error) {
	event := new(DispatcherOwnershipTransferred)
	if err := _Dispatcher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
