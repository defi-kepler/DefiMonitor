// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package withdrawal

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

// DefiMetaData contains all meta data concerning the Defi contract.
var DefiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_dispatcher\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"dispatcher\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_dispatcher\",\"type\":\"address\"}],\"name\":\"setDispatcher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allow\",\"type\":\"bool\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DefiABI is the input ABI used to generate the binding from.
// Deprecated: Use DefiMetaData.ABI instead.
var DefiABI = DefiMetaData.ABI

// Defi is an auto generated Go binding around an Ethereum contract.
type Defi struct {
	DefiCaller     // Read-only binding to the contract
	DefiTransactor // Write-only binding to the contract
	DefiFilterer   // Log filterer for contract events
}

// DefiCaller is an auto generated read-only Go binding around an Ethereum contract.
type DefiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DefiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DefiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DefiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DefiSession struct {
	Contract     *Defi             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DefiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DefiCallerSession struct {
	Contract *DefiCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DefiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DefiTransactorSession struct {
	Contract     *DefiTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DefiRaw is an auto generated low-level Go binding around an Ethereum contract.
type DefiRaw struct {
	Contract *Defi // Generic contract binding to access the raw methods on
}

// DefiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DefiCallerRaw struct {
	Contract *DefiCaller // Generic read-only contract binding to access the raw methods on
}

// DefiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DefiTransactorRaw struct {
	Contract *DefiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDefi creates a new instance of Defi, bound to a specific deployed contract.
func NewDefi(address common.Address, backend bind.ContractBackend) (*Defi, error) {
	contract, err := bindDefi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Defi{DefiCaller: DefiCaller{contract: contract}, DefiTransactor: DefiTransactor{contract: contract}, DefiFilterer: DefiFilterer{contract: contract}}, nil
}

// NewDefiCaller creates a new read-only instance of Defi, bound to a specific deployed contract.
func NewDefiCaller(address common.Address, caller bind.ContractCaller) (*DefiCaller, error) {
	contract, err := bindDefi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DefiCaller{contract: contract}, nil
}

// NewDefiTransactor creates a new write-only instance of Defi, bound to a specific deployed contract.
func NewDefiTransactor(address common.Address, transactor bind.ContractTransactor) (*DefiTransactor, error) {
	contract, err := bindDefi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DefiTransactor{contract: contract}, nil
}

// NewDefiFilterer creates a new log filterer instance of Defi, bound to a specific deployed contract.
func NewDefiFilterer(address common.Address, filterer bind.ContractFilterer) (*DefiFilterer, error) {
	contract, err := bindDefi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DefiFilterer{contract: contract}, nil
}

// bindDefi binds a generic wrapper to an already deployed contract.
func bindDefi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DefiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Defi *DefiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Defi.Contract.DefiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Defi *DefiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Defi.Contract.DefiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Defi *DefiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Defi.Contract.DefiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Defi *DefiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Defi.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Defi *DefiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Defi.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Defi *DefiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Defi.Contract.contract.Transact(opts, method, params...)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_Defi *DefiCaller) Dispatcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "dispatcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_Defi *DefiSession) Dispatcher() (common.Address, error) {
	return _Defi.Contract.Dispatcher(&_Defi.CallOpts)
}

// Dispatcher is a free data retrieval call binding the contract method 0xcb7e9057.
//
// Solidity: function dispatcher() view returns(address)
func (_Defi *DefiCallerSession) Dispatcher() (common.Address, error) {
	return _Defi.Contract.Dispatcher(&_Defi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Defi *DefiCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Defi *DefiSession) Owner() (common.Address, error) {
	return _Defi.Contract.Owner(&_Defi.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Defi *DefiCallerSession) Owner() (common.Address, error) {
	return _Defi.Contract.Owner(&_Defi.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Defi *DefiCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Defi.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Defi *DefiSession) Token() (common.Address, error) {
	return _Defi.Contract.Token(&_Defi.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Defi *DefiCallerSession) Token() (common.Address, error) {
	return _Defi.Contract.Token(&_Defi.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Defi *DefiTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Defi *DefiSession) RenounceOwnership() (*types.Transaction, error) {
	return _Defi.Contract.RenounceOwnership(&_Defi.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Defi *DefiTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Defi.Contract.RenounceOwnership(&_Defi.TransactOpts)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address _dispatcher) returns()
func (_Defi *DefiTransactor) SetDispatcher(opts *bind.TransactOpts, _dispatcher common.Address) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "setDispatcher", _dispatcher)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address _dispatcher) returns()
func (_Defi *DefiSession) SetDispatcher(_dispatcher common.Address) (*types.Transaction, error) {
	return _Defi.Contract.SetDispatcher(&_Defi.TransactOpts, _dispatcher)
}

// SetDispatcher is a paid mutator transaction binding the contract method 0xba22bd76.
//
// Solidity: function setDispatcher(address _dispatcher) returns()
func (_Defi *DefiTransactorSession) SetDispatcher(_dispatcher common.Address) (*types.Transaction, error) {
	return _Defi.Contract.SetDispatcher(&_Defi.TransactOpts, _dispatcher)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address user, bool allow) returns()
func (_Defi *DefiTransactor) SetOperator(opts *bind.TransactOpts, user common.Address, allow bool) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "setOperator", user, allow)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address user, bool allow) returns()
func (_Defi *DefiSession) SetOperator(user common.Address, allow bool) (*types.Transaction, error) {
	return _Defi.Contract.SetOperator(&_Defi.TransactOpts, user, allow)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address user, bool allow) returns()
func (_Defi *DefiTransactorSession) SetOperator(user common.Address, allow bool) (*types.Transaction, error) {
	return _Defi.Contract.SetOperator(&_Defi.TransactOpts, user, allow)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Defi *DefiTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Defi *DefiSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Defi.Contract.TransferOwnership(&_Defi.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Defi *DefiTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Defi.Contract.TransferOwnership(&_Defi.TransactOpts, newOwner)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x5a6b26ba.
//
// Solidity: function withdrawal(address user, uint256 amount) returns()
func (_Defi *DefiTransactor) Withdrawal(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Defi.contract.Transact(opts, "withdrawal", user, amount)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x5a6b26ba.
//
// Solidity: function withdrawal(address user, uint256 amount) returns()
func (_Defi *DefiSession) Withdrawal(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.Withdrawal(&_Defi.TransactOpts, user, amount)
}

// Withdrawal is a paid mutator transaction binding the contract method 0x5a6b26ba.
//
// Solidity: function withdrawal(address user, uint256 amount) returns()
func (_Defi *DefiTransactorSession) Withdrawal(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Defi.Contract.Withdrawal(&_Defi.TransactOpts, user, amount)
}

// DefiOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Defi contract.
type DefiOwnershipTransferredIterator struct {
	Event *DefiOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DefiOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiOwnershipTransferred)
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
		it.Event = new(DefiOwnershipTransferred)
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
func (it *DefiOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiOwnershipTransferred represents a OwnershipTransferred event raised by the Defi contract.
type DefiOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Defi *DefiFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DefiOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Defi.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DefiOwnershipTransferredIterator{contract: _Defi.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Defi *DefiFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DefiOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Defi.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiOwnershipTransferred)
				if err := _Defi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Defi *DefiFilterer) ParseOwnershipTransferred(log types.Log) (*DefiOwnershipTransferred, error) {
	event := new(DefiOwnershipTransferred)
	if err := _Defi.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DefiWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Defi contract.
type DefiWithdrawalIterator struct {
	Event *DefiWithdrawal // Event containing the contract specifics and raw log

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
func (it *DefiWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DefiWithdrawal)
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
		it.Event = new(DefiWithdrawal)
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
func (it *DefiWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DefiWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DefiWithdrawal represents a Withdrawal event raised by the Defi contract.
type DefiWithdrawal struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address user, uint256 amount)
func (_Defi *DefiFilterer) FilterWithdrawal(opts *bind.FilterOpts) (*DefiWithdrawalIterator, error) {

	logs, sub, err := _Defi.contract.FilterLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return &DefiWithdrawalIterator{contract: _Defi.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address user, uint256 amount)
func (_Defi *DefiFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *DefiWithdrawal) (event.Subscription, error) {

	logs, sub, err := _Defi.contract.WatchLogs(opts, "Withdrawal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DefiWithdrawal)
				if err := _Defi.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address user, uint256 amount)
func (_Defi *DefiFilterer) ParseWithdrawal(log types.Log) (*DefiWithdrawal, error) {
	event := new(DefiWithdrawal)
	if err := _Defi.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
