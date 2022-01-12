// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package zeroxtoken

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

// ZeroxtokenMetaData contains all meta data concerning the Zeroxtoken contract.
var ZeroxtokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]",
}

// ZeroxtokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ZeroxtokenMetaData.ABI instead.
var ZeroxtokenABI = ZeroxtokenMetaData.ABI

// Zeroxtoken is an auto generated Go binding around an Ethereum contract.
type Zeroxtoken struct {
	ZeroxtokenCaller     // Read-only binding to the contract
	ZeroxtokenTransactor // Write-only binding to the contract
	ZeroxtokenFilterer   // Log filterer for contract events
}

// ZeroxtokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZeroxtokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroxtokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZeroxtokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroxtokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZeroxtokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZeroxtokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZeroxtokenSession struct {
	Contract     *Zeroxtoken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZeroxtokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZeroxtokenCallerSession struct {
	Contract *ZeroxtokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ZeroxtokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZeroxtokenTransactorSession struct {
	Contract     *ZeroxtokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ZeroxtokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZeroxtokenRaw struct {
	Contract *Zeroxtoken // Generic contract binding to access the raw methods on
}

// ZeroxtokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZeroxtokenCallerRaw struct {
	Contract *ZeroxtokenCaller // Generic read-only contract binding to access the raw methods on
}

// ZeroxtokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZeroxtokenTransactorRaw struct {
	Contract *ZeroxtokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZeroxtoken creates a new instance of Zeroxtoken, bound to a specific deployed contract.
func NewZeroxtoken(address common.Address, backend bind.ContractBackend) (*Zeroxtoken, error) {
	contract, err := bindZeroxtoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Zeroxtoken{ZeroxtokenCaller: ZeroxtokenCaller{contract: contract}, ZeroxtokenTransactor: ZeroxtokenTransactor{contract: contract}, ZeroxtokenFilterer: ZeroxtokenFilterer{contract: contract}}, nil
}

// NewZeroxtokenCaller creates a new read-only instance of Zeroxtoken, bound to a specific deployed contract.
func NewZeroxtokenCaller(address common.Address, caller bind.ContractCaller) (*ZeroxtokenCaller, error) {
	contract, err := bindZeroxtoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroxtokenCaller{contract: contract}, nil
}

// NewZeroxtokenTransactor creates a new write-only instance of Zeroxtoken, bound to a specific deployed contract.
func NewZeroxtokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ZeroxtokenTransactor, error) {
	contract, err := bindZeroxtoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZeroxtokenTransactor{contract: contract}, nil
}

// NewZeroxtokenFilterer creates a new log filterer instance of Zeroxtoken, bound to a specific deployed contract.
func NewZeroxtokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ZeroxtokenFilterer, error) {
	contract, err := bindZeroxtoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZeroxtokenFilterer{contract: contract}, nil
}

// bindZeroxtoken binds a generic wrapper to an already deployed contract.
func bindZeroxtoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ZeroxtokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zeroxtoken *ZeroxtokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zeroxtoken.Contract.ZeroxtokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zeroxtoken *ZeroxtokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zeroxtoken.Contract.ZeroxtokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zeroxtoken *ZeroxtokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zeroxtoken.Contract.ZeroxtokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Zeroxtoken *ZeroxtokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Zeroxtoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Zeroxtoken *ZeroxtokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Zeroxtoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Zeroxtoken *ZeroxtokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Zeroxtoken.Contract.contract.Transact(opts, method, params...)
}

// ZeroxtokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Zeroxtoken contract.
type ZeroxtokenApprovalIterator struct {
	Event *ZeroxtokenApproval // Event containing the contract specifics and raw log

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
func (it *ZeroxtokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroxtokenApproval)
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
		it.Event = new(ZeroxtokenApproval)
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
func (it *ZeroxtokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroxtokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroxtokenApproval represents a Approval event raised by the Zeroxtoken contract.
type ZeroxtokenApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Zeroxtoken *ZeroxtokenFilterer) FilterApproval(opts *bind.FilterOpts, tokenOwner []common.Address, spender []common.Address) (*ZeroxtokenApprovalIterator, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Zeroxtoken.contract.FilterLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ZeroxtokenApprovalIterator{contract: _Zeroxtoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Zeroxtoken *ZeroxtokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ZeroxtokenApproval, tokenOwner []common.Address, spender []common.Address) (event.Subscription, error) {

	var tokenOwnerRule []interface{}
	for _, tokenOwnerItem := range tokenOwner {
		tokenOwnerRule = append(tokenOwnerRule, tokenOwnerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Zeroxtoken.contract.WatchLogs(opts, "Approval", tokenOwnerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroxtokenApproval)
				if err := _Zeroxtoken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed tokenOwner, address indexed spender, uint256 tokens)
func (_Zeroxtoken *ZeroxtokenFilterer) ParseApproval(log types.Log) (*ZeroxtokenApproval, error) {
	event := new(ZeroxtokenApproval)
	if err := _Zeroxtoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZeroxtokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Zeroxtoken contract.
type ZeroxtokenTransferIterator struct {
	Event *ZeroxtokenTransfer // Event containing the contract specifics and raw log

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
func (it *ZeroxtokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZeroxtokenTransfer)
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
		it.Event = new(ZeroxtokenTransfer)
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
func (it *ZeroxtokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZeroxtokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZeroxtokenTransfer represents a Transfer event raised by the Zeroxtoken contract.
type ZeroxtokenTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Zeroxtoken *ZeroxtokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ZeroxtokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Zeroxtoken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ZeroxtokenTransferIterator{contract: _Zeroxtoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Zeroxtoken *ZeroxtokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ZeroxtokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Zeroxtoken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZeroxtokenTransfer)
				if err := _Zeroxtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 tokens)
func (_Zeroxtoken *ZeroxtokenFilterer) ParseTransfer(log types.Log) (*ZeroxtokenTransfer, error) {
	event := new(ZeroxtokenTransfer)
	if err := _Zeroxtoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
