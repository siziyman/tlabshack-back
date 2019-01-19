// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Ownable *OwnableCallerSession) IsOwner() (bool, error) {
	return _Ownable.Contract.IsOwner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// SmartCityRideShareABI is the input ABI used to generate the binding from.
const SmartCityRideShareABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"payOut\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"parkPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initiateParkingPayment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"driver\",\"type\":\"address\"}],\"name\":\"startRide\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ridePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// SmartCityRideShareBin is the compiled bytecode used for deploying new contracts.
const SmartCityRideShareBin = `0x608060408190526001600381905560045560008054600160a060020a0319163317808255600160a060020a0316917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a36103c0806100616000396000f3fe608060405234801561001057600080fd5b50600436106100b0576000357c0100000000000000000000000000000000000000000000000000000000900480638da5cb5b116100835780638da5cb5b1461010d5780638f32d59b14610131578063910bc1401461014d578063f2fde38b14610173578063fda46d7d14610199576100b0565b806305061616146100b5578063073d33c0146100e35780632444217e146100fd578063715018a614610105575b600080fd5b6100e1600480360360408110156100cb57600080fd5b5080359060200135600160a060020a03166101a1565b005b6100eb6101d4565b60408051918252519081900360200190f35b6100e16101da565b6100e1610213565b61011561027d565b60408051600160a060020a039092168252519081900360200190f35b61013961028c565b604080519115158252519081900360200190f35b6100e16004803603602081101561016357600080fd5b5035600160a060020a031661029d565b6100e16004803603602081101561018957600080fd5b5035600160a060020a03166102f2565b6100eb610311565b6101a961028c565b15156101b457600080fd5b600160a060020a0316600090815260026020526040902080549091019055565b60045481565b60045433600090815260026020526040902054116101f757600080fd5b6004543360009081526002602052604090208054919091039055565b61021b61028c565b151561022657600080fd5b60008054604051600160a060020a03909116907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a36000805473ffffffffffffffffffffffffffffffffffffffff19169055565b600054600160a060020a031690565b600054600160a060020a0316331490565b60035433600090815260026020526040902054116102ba57600080fd5b60038054600160a060020a03909216600090815260026020526040808220805490940190935590543382529190208054919091039055565b6102fa61028c565b151561030557600080fd5b61030e81610317565b50565b60035481565b600160a060020a038116151561032c57600080fd5b60008054604051600160a060020a03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a36000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905556fea165627a7a723058205e5e0c01ad942628d93225ffaeef3c009d64474479dd21155fef9bb1312df8fc0029`

// DeploySmartCityRideShare deploys a new Ethereum contract, binding an instance of SmartCityRideShare to it.
func DeploySmartCityRideShare(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SmartCityRideShare, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartCityRideShareABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SmartCityRideShareBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SmartCityRideShare{SmartCityRideShareCaller: SmartCityRideShareCaller{contract: contract}, SmartCityRideShareTransactor: SmartCityRideShareTransactor{contract: contract}, SmartCityRideShareFilterer: SmartCityRideShareFilterer{contract: contract}}, nil
}

// SmartCityRideShare is an auto generated Go binding around an Ethereum contract.
type SmartCityRideShare struct {
	SmartCityRideShareCaller     // Read-only binding to the contract
	SmartCityRideShareTransactor // Write-only binding to the contract
	SmartCityRideShareFilterer   // Log filterer for contract events
}

// SmartCityRideShareCaller is an auto generated read-only Go binding around an Ethereum contract.
type SmartCityRideShareCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartCityRideShareTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SmartCityRideShareTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartCityRideShareFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SmartCityRideShareFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SmartCityRideShareSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SmartCityRideShareSession struct {
	Contract     *SmartCityRideShare // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SmartCityRideShareCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SmartCityRideShareCallerSession struct {
	Contract *SmartCityRideShareCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SmartCityRideShareTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SmartCityRideShareTransactorSession struct {
	Contract     *SmartCityRideShareTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SmartCityRideShareRaw is an auto generated low-level Go binding around an Ethereum contract.
type SmartCityRideShareRaw struct {
	Contract *SmartCityRideShare // Generic contract binding to access the raw methods on
}

// SmartCityRideShareCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SmartCityRideShareCallerRaw struct {
	Contract *SmartCityRideShareCaller // Generic read-only contract binding to access the raw methods on
}

// SmartCityRideShareTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SmartCityRideShareTransactorRaw struct {
	Contract *SmartCityRideShareTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSmartCityRideShare creates a new instance of SmartCityRideShare, bound to a specific deployed contract.
func NewSmartCityRideShare(address common.Address, backend bind.ContractBackend) (*SmartCityRideShare, error) {
	contract, err := bindSmartCityRideShare(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SmartCityRideShare{SmartCityRideShareCaller: SmartCityRideShareCaller{contract: contract}, SmartCityRideShareTransactor: SmartCityRideShareTransactor{contract: contract}, SmartCityRideShareFilterer: SmartCityRideShareFilterer{contract: contract}}, nil
}

// NewSmartCityRideShareCaller creates a new read-only instance of SmartCityRideShare, bound to a specific deployed contract.
func NewSmartCityRideShareCaller(address common.Address, caller bind.ContractCaller) (*SmartCityRideShareCaller, error) {
	contract, err := bindSmartCityRideShare(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SmartCityRideShareCaller{contract: contract}, nil
}

// NewSmartCityRideShareTransactor creates a new write-only instance of SmartCityRideShare, bound to a specific deployed contract.
func NewSmartCityRideShareTransactor(address common.Address, transactor bind.ContractTransactor) (*SmartCityRideShareTransactor, error) {
	contract, err := bindSmartCityRideShare(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SmartCityRideShareTransactor{contract: contract}, nil
}

// NewSmartCityRideShareFilterer creates a new log filterer instance of SmartCityRideShare, bound to a specific deployed contract.
func NewSmartCityRideShareFilterer(address common.Address, filterer bind.ContractFilterer) (*SmartCityRideShareFilterer, error) {
	contract, err := bindSmartCityRideShare(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SmartCityRideShareFilterer{contract: contract}, nil
}

// bindSmartCityRideShare binds a generic wrapper to an already deployed contract.
func bindSmartCityRideShare(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SmartCityRideShareABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartCityRideShare *SmartCityRideShareRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartCityRideShare.Contract.SmartCityRideShareCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartCityRideShare *SmartCityRideShareRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartCityRideShare.Contract.SmartCityRideShareTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartCityRideShare *SmartCityRideShareRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartCityRideShare.Contract.SmartCityRideShareTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SmartCityRideShare *SmartCityRideShareCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SmartCityRideShare.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SmartCityRideShare *SmartCityRideShareTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartCityRideShare.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SmartCityRideShare *SmartCityRideShareTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SmartCityRideShare.Contract.contract.Transact(opts, method, params...)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SmartCityRideShare *SmartCityRideShareCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SmartCityRideShare.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SmartCityRideShare *SmartCityRideShareSession) IsOwner() (bool, error) {
	return _SmartCityRideShare.Contract.IsOwner(&_SmartCityRideShare.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_SmartCityRideShare *SmartCityRideShareCallerSession) IsOwner() (bool, error) {
	return _SmartCityRideShare.Contract.IsOwner(&_SmartCityRideShare.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SmartCityRideShare *SmartCityRideShareCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SmartCityRideShare.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SmartCityRideShare *SmartCityRideShareSession) Owner() (common.Address, error) {
	return _SmartCityRideShare.Contract.Owner(&_SmartCityRideShare.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SmartCityRideShare *SmartCityRideShareCallerSession) Owner() (common.Address, error) {
	return _SmartCityRideShare.Contract.Owner(&_SmartCityRideShare.CallOpts)
}

// ParkPrice is a free data retrieval call binding the contract method 0x073d33c0.
//
// Solidity: function parkPrice() constant returns(uint256)
func (_SmartCityRideShare *SmartCityRideShareCaller) ParkPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartCityRideShare.contract.Call(opts, out, "parkPrice")
	return *ret0, err
}

// ParkPrice is a free data retrieval call binding the contract method 0x073d33c0.
//
// Solidity: function parkPrice() constant returns(uint256)
func (_SmartCityRideShare *SmartCityRideShareSession) ParkPrice() (*big.Int, error) {
	return _SmartCityRideShare.Contract.ParkPrice(&_SmartCityRideShare.CallOpts)
}

// ParkPrice is a free data retrieval call binding the contract method 0x073d33c0.
//
// Solidity: function parkPrice() constant returns(uint256)
func (_SmartCityRideShare *SmartCityRideShareCallerSession) ParkPrice() (*big.Int, error) {
	return _SmartCityRideShare.Contract.ParkPrice(&_SmartCityRideShare.CallOpts)
}

// RidePrice is a free data retrieval call binding the contract method 0xfda46d7d.
//
// Solidity: function ridePrice() constant returns(uint256)
func (_SmartCityRideShare *SmartCityRideShareCaller) RidePrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SmartCityRideShare.contract.Call(opts, out, "ridePrice")
	return *ret0, err
}

// RidePrice is a free data retrieval call binding the contract method 0xfda46d7d.
//
// Solidity: function ridePrice() constant returns(uint256)
func (_SmartCityRideShare *SmartCityRideShareSession) RidePrice() (*big.Int, error) {
	return _SmartCityRideShare.Contract.RidePrice(&_SmartCityRideShare.CallOpts)
}

// RidePrice is a free data retrieval call binding the contract method 0xfda46d7d.
//
// Solidity: function ridePrice() constant returns(uint256)
func (_SmartCityRideShare *SmartCityRideShareCallerSession) RidePrice() (*big.Int, error) {
	return _SmartCityRideShare.Contract.RidePrice(&_SmartCityRideShare.CallOpts)
}

// InitiateParkingPayment is a paid mutator transaction binding the contract method 0x2444217e.
//
// Solidity: function initiateParkingPayment() returns()
func (_SmartCityRideShare *SmartCityRideShareTransactor) InitiateParkingPayment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartCityRideShare.contract.Transact(opts, "initiateParkingPayment")
}

// InitiateParkingPayment is a paid mutator transaction binding the contract method 0x2444217e.
//
// Solidity: function initiateParkingPayment() returns()
func (_SmartCityRideShare *SmartCityRideShareSession) InitiateParkingPayment() (*types.Transaction, error) {
	return _SmartCityRideShare.Contract.InitiateParkingPayment(&_SmartCityRideShare.TransactOpts)
}

// InitiateParkingPayment is a paid mutator transaction binding the contract method 0x2444217e.
//
// Solidity: function initiateParkingPayment() returns()
func (_SmartCityRideShare *SmartCityRideShareTransactorSession) InitiateParkingPayment() (*types.Transaction, error) {
	return _SmartCityRideShare.Contract.InitiateParkingPayment(&_SmartCityRideShare.TransactOpts)
}

// PayOut is a paid mutator transaction binding the contract method 0x05061616.
//
// Solidity: function payOut(uint256 amount, address user) returns()
func (_SmartCityRideShare *SmartCityRideShareTransactor) PayOut(opts *bind.TransactOpts, amount *big.Int, user common.Address) (*types.Transaction, error) {
	return _SmartCityRideShare.contract.Transact(opts, "payOut", amount, user)
}

// PayOut is a paid mutator transaction binding the contract method 0x05061616.
//
// Solidity: function payOut(uint256 amount, address user) returns()
func (_SmartCityRideShare *SmartCityRideShareSession) PayOut(amount *big.Int, user common.Address) (*types.Transaction, error) {
	return _SmartCityRideShare.Contract.PayOut(&_SmartCityRideShare.TransactOpts, amount, user)
}

// PayOut is a paid mutator transaction binding the contract method 0x05061616.
//
// Solidity: function payOut(uint256 amount, address user) returns()
func (_SmartCityRideShare *SmartCityRideShareTransactorSession) PayOut(amount *big.Int, user common.Address) (*types.Transaction, error) {
	return _SmartCityRideShare.Contract.PayOut(&_SmartCityRideShare.TransactOpts, amount, user)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SmartCityRideShare *SmartCityRideShareTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SmartCityRideShare.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SmartCityRideShare *SmartCityRideShareSession) RenounceOwnership() (*types.Transaction, error) {
	return _SmartCityRideShare.Contract.RenounceOwnership(&_SmartCityRideShare.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SmartCityRideShare *SmartCityRideShareTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SmartCityRideShare.Contract.RenounceOwnership(&_SmartCityRideShare.TransactOpts)
}

// StartRide is a paid mutator transaction binding the contract method 0x910bc140.
//
// Solidity: function startRide(address driver) returns()
func (_SmartCityRideShare *SmartCityRideShareTransactor) StartRide(opts *bind.TransactOpts, driver common.Address) (*types.Transaction, error) {
	return _SmartCityRideShare.contract.Transact(opts, "startRide", driver)
}

// StartRide is a paid mutator transaction binding the contract method 0x910bc140.
//
// Solidity: function startRide(address driver) returns()
func (_SmartCityRideShare *SmartCityRideShareSession) StartRide(driver common.Address) (*types.Transaction, error) {
	return _SmartCityRideShare.Contract.StartRide(&_SmartCityRideShare.TransactOpts, driver)
}

// StartRide is a paid mutator transaction binding the contract method 0x910bc140.
//
// Solidity: function startRide(address driver) returns()
func (_SmartCityRideShare *SmartCityRideShareTransactorSession) StartRide(driver common.Address) (*types.Transaction, error) {
	return _SmartCityRideShare.Contract.StartRide(&_SmartCityRideShare.TransactOpts, driver)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SmartCityRideShare *SmartCityRideShareTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SmartCityRideShare.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SmartCityRideShare *SmartCityRideShareSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SmartCityRideShare.Contract.TransferOwnership(&_SmartCityRideShare.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SmartCityRideShare *SmartCityRideShareTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SmartCityRideShare.Contract.TransferOwnership(&_SmartCityRideShare.TransactOpts, newOwner)
}

// SmartCityRideShareOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SmartCityRideShare contract.
type SmartCityRideShareOwnershipTransferredIterator struct {
	Event *SmartCityRideShareOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SmartCityRideShareOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SmartCityRideShareOwnershipTransferred)
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
		it.Event = new(SmartCityRideShareOwnershipTransferred)
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
func (it *SmartCityRideShareOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SmartCityRideShareOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SmartCityRideShareOwnershipTransferred represents a OwnershipTransferred event raised by the SmartCityRideShare contract.
type SmartCityRideShareOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SmartCityRideShare *SmartCityRideShareFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SmartCityRideShareOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SmartCityRideShare.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SmartCityRideShareOwnershipTransferredIterator{contract: _SmartCityRideShare.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SmartCityRideShare *SmartCityRideShareFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SmartCityRideShareOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SmartCityRideShare.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SmartCityRideShareOwnershipTransferred)
				if err := _SmartCityRideShare.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

