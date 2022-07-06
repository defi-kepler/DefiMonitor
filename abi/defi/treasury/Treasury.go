// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package treasury

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

// TreasuryMetaData contains all meta data concerning the Treasury contract.
var TreasuryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Sweep\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stoken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TreasuryABI is the input ABI used to generate the binding from.
// Deprecated: Use TreasuryMetaData.ABI instead.
var TreasuryABI = TreasuryMetaData.ABI

// Treasury is an auto generated Go binding around an Ethereum contract.
type Treasury struct {
	TreasuryCaller     // Read-only binding to the contract
	TreasuryTransactor // Write-only binding to the contract
	TreasuryFilterer   // Log filterer for contract events
}

// TreasuryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TreasuryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TreasuryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TreasuryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasurySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TreasurySession struct {
	Contract     *Treasury         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TreasuryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TreasuryCallerSession struct {
	Contract *TreasuryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TreasuryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TreasuryTransactorSession struct {
	Contract     *TreasuryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TreasuryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TreasuryRaw struct {
	Contract *Treasury // Generic contract binding to access the raw methods on
}

// TreasuryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TreasuryCallerRaw struct {
	Contract *TreasuryCaller // Generic read-only contract binding to access the raw methods on
}

// TreasuryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TreasuryTransactorRaw struct {
	Contract *TreasuryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTreasury creates a new instance of Treasury, bound to a specific deployed contract.
func NewTreasury(address common.Address, backend bind.ContractBackend) (*Treasury, error) {
	contract, err := bindTreasury(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Treasury{TreasuryCaller: TreasuryCaller{contract: contract}, TreasuryTransactor: TreasuryTransactor{contract: contract}, TreasuryFilterer: TreasuryFilterer{contract: contract}}, nil
}

// NewTreasuryCaller creates a new read-only instance of Treasury, bound to a specific deployed contract.
func NewTreasuryCaller(address common.Address, caller bind.ContractCaller) (*TreasuryCaller, error) {
	contract, err := bindTreasury(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryCaller{contract: contract}, nil
}

// NewTreasuryTransactor creates a new write-only instance of Treasury, bound to a specific deployed contract.
func NewTreasuryTransactor(address common.Address, transactor bind.ContractTransactor) (*TreasuryTransactor, error) {
	contract, err := bindTreasury(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryTransactor{contract: contract}, nil
}

// NewTreasuryFilterer creates a new log filterer instance of Treasury, bound to a specific deployed contract.
func NewTreasuryFilterer(address common.Address, filterer bind.ContractFilterer) (*TreasuryFilterer, error) {
	contract, err := bindTreasury(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TreasuryFilterer{contract: contract}, nil
}

// bindTreasury binds a generic wrapper to an already deployed contract.
func bindTreasury(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TreasuryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Treasury *TreasuryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Treasury.Contract.TreasuryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Treasury *TreasuryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.Contract.TreasuryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Treasury *TreasuryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Treasury.Contract.TreasuryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Treasury *TreasuryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Treasury.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Treasury *TreasuryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Treasury *TreasuryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Treasury.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Treasury *TreasuryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Treasury *TreasurySession) Owner() (common.Address, error) {
	return _Treasury.Contract.Owner(&_Treasury.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Treasury *TreasuryCallerSession) Owner() (common.Address, error) {
	return _Treasury.Contract.Owner(&_Treasury.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Treasury *TreasuryCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Treasury.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Treasury *TreasurySession) Token() (common.Address, error) {
	return _Treasury.Contract.Token(&_Treasury.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Treasury *TreasuryCallerSession) Token() (common.Address, error) {
	return _Treasury.Contract.Token(&_Treasury.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Treasury *TreasuryTransactor) Deposit(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "deposit", amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Treasury *TreasurySession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.Deposit(&_Treasury.TransactOpts, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xb6b55f25.
//
// Solidity: function deposit(uint256 amount) returns()
func (_Treasury *TreasuryTransactorSession) Deposit(amount *big.Int) (*types.Transaction, error) {
	return _Treasury.Contract.Deposit(&_Treasury.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Treasury *TreasuryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Treasury *TreasurySession) RenounceOwnership() (*types.Transaction, error) {
	return _Treasury.Contract.RenounceOwnership(&_Treasury.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Treasury *TreasuryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Treasury.Contract.RenounceOwnership(&_Treasury.TransactOpts)
}

// Sweep is a paid mutator transaction binding the contract method 0xb8dc491b.
//
// Solidity: function sweep(address stoken, address recipient) returns()
func (_Treasury *TreasuryTransactor) Sweep(opts *bind.TransactOpts, stoken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "sweep", stoken, recipient)
}

// Sweep is a paid mutator transaction binding the contract method 0xb8dc491b.
//
// Solidity: function sweep(address stoken, address recipient) returns()
func (_Treasury *TreasurySession) Sweep(stoken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.Sweep(&_Treasury.TransactOpts, stoken, recipient)
}

// Sweep is a paid mutator transaction binding the contract method 0xb8dc491b.
//
// Solidity: function sweep(address stoken, address recipient) returns()
func (_Treasury *TreasuryTransactorSession) Sweep(stoken common.Address, recipient common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.Sweep(&_Treasury.TransactOpts, stoken, recipient)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Treasury *TreasuryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Treasury *TreasurySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.TransferOwnership(&_Treasury.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Treasury *TreasuryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.TransferOwnership(&_Treasury.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Treasury *TreasuryTransactor) Withdraw(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "withdraw", recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Treasury *TreasurySession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.Withdraw(&_Treasury.TransactOpts, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_Treasury *TreasuryTransactorSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.Withdraw(&_Treasury.TransactOpts, recipient)
}

// TreasuryDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Treasury contract.
type TreasuryDepositIterator struct {
	Event *TreasuryDeposit // Event containing the contract specifics and raw log

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
func (it *TreasuryDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryDeposit)
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
		it.Event = new(TreasuryDeposit)
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
func (it *TreasuryDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryDeposit represents a Deposit event raised by the Treasury contract.
type TreasuryDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address user, uint256 amount)
func (_Treasury *TreasuryFilterer) FilterDeposit(opts *bind.FilterOpts) (*TreasuryDepositIterator, error) {

	logs, sub, err := _Treasury.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &TreasuryDepositIterator{contract: _Treasury.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address user, uint256 amount)
func (_Treasury *TreasuryFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *TreasuryDeposit) (event.Subscription, error) {

	logs, sub, err := _Treasury.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryDeposit)
				if err := _Treasury.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address user, uint256 amount)
func (_Treasury *TreasuryFilterer) ParseDeposit(log types.Log) (*TreasuryDeposit, error) {
	event := new(TreasuryDeposit)
	if err := _Treasury.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Treasury contract.
type TreasuryOwnershipTransferredIterator struct {
	Event *TreasuryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TreasuryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryOwnershipTransferred)
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
		it.Event = new(TreasuryOwnershipTransferred)
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
func (it *TreasuryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryOwnershipTransferred represents a OwnershipTransferred event raised by the Treasury contract.
type TreasuryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Treasury *TreasuryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TreasuryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Treasury.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TreasuryOwnershipTransferredIterator{contract: _Treasury.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Treasury *TreasuryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TreasuryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Treasury.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryOwnershipTransferred)
				if err := _Treasury.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Treasury *TreasuryFilterer) ParseOwnershipTransferred(log types.Log) (*TreasuryOwnershipTransferred, error) {
	event := new(TreasuryOwnershipTransferred)
	if err := _Treasury.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasurySweepIterator is returned from FilterSweep and is used to iterate over the raw logs and unpacked data for Sweep events raised by the Treasury contract.
type TreasurySweepIterator struct {
	Event *TreasurySweep // Event containing the contract specifics and raw log

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
func (it *TreasurySweepIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasurySweep)
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
		it.Event = new(TreasurySweep)
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
func (it *TreasurySweepIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasurySweepIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasurySweep represents a Sweep event raised by the Treasury contract.
type TreasurySweep struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSweep is a free log retrieval operation binding the contract event 0xed679328aebf74ede77ae09efcf36e90244f83643dadac1c2d9f0b21a46f6ab7.
//
// Solidity: event Sweep(address token, address recipient, uint256 amount)
func (_Treasury *TreasuryFilterer) FilterSweep(opts *bind.FilterOpts) (*TreasurySweepIterator, error) {

	logs, sub, err := _Treasury.contract.FilterLogs(opts, "Sweep")
	if err != nil {
		return nil, err
	}
	return &TreasurySweepIterator{contract: _Treasury.contract, event: "Sweep", logs: logs, sub: sub}, nil
}

// WatchSweep is a free log subscription operation binding the contract event 0xed679328aebf74ede77ae09efcf36e90244f83643dadac1c2d9f0b21a46f6ab7.
//
// Solidity: event Sweep(address token, address recipient, uint256 amount)
func (_Treasury *TreasuryFilterer) WatchSweep(opts *bind.WatchOpts, sink chan<- *TreasurySweep) (event.Subscription, error) {

	logs, sub, err := _Treasury.contract.WatchLogs(opts, "Sweep")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasurySweep)
				if err := _Treasury.contract.UnpackLog(event, "Sweep", log); err != nil {
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

// ParseSweep is a log parse operation binding the contract event 0xed679328aebf74ede77ae09efcf36e90244f83643dadac1c2d9f0b21a46f6ab7.
//
// Solidity: event Sweep(address token, address recipient, uint256 amount)
func (_Treasury *TreasuryFilterer) ParseSweep(log types.Log) (*TreasurySweep, error) {
	event := new(TreasurySweep)
	if err := _Treasury.contract.UnpackLog(event, "Sweep", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TreasuryWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Treasury contract.
type TreasuryWithdrawIterator struct {
	Event *TreasuryWithdraw // Event containing the contract specifics and raw log

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
func (it *TreasuryWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TreasuryWithdraw)
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
		it.Event = new(TreasuryWithdraw)
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
func (it *TreasuryWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TreasuryWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TreasuryWithdraw represents a Withdraw event raised by the Treasury contract.
type TreasuryWithdraw struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address recipient, uint256 amount)
func (_Treasury *TreasuryFilterer) FilterWithdraw(opts *bind.FilterOpts) (*TreasuryWithdrawIterator, error) {

	logs, sub, err := _Treasury.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &TreasuryWithdrawIterator{contract: _Treasury.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address recipient, uint256 amount)
func (_Treasury *TreasuryFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TreasuryWithdraw) (event.Subscription, error) {

	logs, sub, err := _Treasury.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TreasuryWithdraw)
				if err := _Treasury.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address recipient, uint256 amount)
func (_Treasury *TreasuryFilterer) ParseWithdraw(log types.Log) (*TreasuryWithdraw, error) {
	event := new(TreasuryWithdraw)
	if err := _Treasury.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
