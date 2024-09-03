// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
	"errors"
	"math/big"
	"strings"
	"fmt"

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
	_ = abi.ConvertType
)

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"filePath\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"FileSaved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"filePath\",\"type\":\"string\"}],\"name\":\"get\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"filePath\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"save\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "6080604052348015600e575f5ffd5b5061090e8061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c8063693ec85e14610038578063962939b814610068575b5f5ffd5b610052600480360381019061004d9190610380565b610084565b60405161005f9190610427565b60405180910390f35b610082600480360381019061007d9190610447565b610131565b005b60605f8260405161009591906104f7565b908152602001604051809103902080546100ae9061053a565b80601f01602080910402602001604051908101604052809291908181526020018280546100da9061053a565b80156101255780601f106100fc57610100808354040283529160200191610125565b820191905f5260205f20905b81548152906001019060200180831161010857829003601f168201915b50505050509050919050565b5f825111610174576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161016b906105b4565b60405180910390fd5b5f8151116101b7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ae9061061c565b60405180910390fd5b805f836040516101c791906104f7565b908152602001604051809103902090816101e191906107e3565b50816040516101f091906104f7565b60405180910390207fcadc7184c55de53d424a8e73df016947523f46250bb8957192d0d084403dfd25826040516102279190610427565b60405180910390a25050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6102928261024c565b810181811067ffffffffffffffff821117156102b1576102b061025c565b5b80604052505050565b5f6102c3610233565b90506102cf8282610289565b919050565b5f67ffffffffffffffff8211156102ee576102ed61025c565b5b6102f78261024c565b9050602081019050919050565b828183375f83830152505050565b5f61032461031f846102d4565b6102ba565b9050828152602081018484840111156103405761033f610248565b5b61034b848285610304565b509392505050565b5f82601f83011261036757610366610244565b5b8135610377848260208601610312565b91505092915050565b5f602082840312156103955761039461023c565b5b5f82013567ffffffffffffffff8111156103b2576103b1610240565b5b6103be84828501610353565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f6103f9826103c7565b61040381856103d1565b93506104138185602086016103e1565b61041c8161024c565b840191505092915050565b5f6020820190508181035f83015261043f81846103ef565b905092915050565b5f5f6040838503121561045d5761045c61023c565b5b5f83013567ffffffffffffffff81111561047a57610479610240565b5b61048685828601610353565b925050602083013567ffffffffffffffff8111156104a7576104a6610240565b5b6104b385828601610353565b9150509250929050565b5f81905092915050565b5f6104d1826103c7565b6104db81856104bd565b93506104eb8185602086016103e1565b80840191505092915050565b5f61050282846104c7565b915081905092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061055157607f821691505b6020821081036105645761056361050d565b5b50919050565b7f46696c6520706174682063616e6e6f7420626520656d707479000000000000005f82015250565b5f61059e6019836103d1565b91506105a98261056a565b602082019050919050565b5f6020820190508181035f8301526105cb81610592565b9050919050565b7f4349442063616e6e6f7420626520656d707479000000000000000000000000005f82015250565b5f6106066013836103d1565b9150610611826105d2565b602082019050919050565b5f6020820190508181035f830152610633816105fa565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026106967fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261065b565b6106a0868361065b565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6106e46106df6106da846106b8565b6106c1565b6106b8565b9050919050565b5f819050919050565b6106fd836106ca565b610711610709826106eb565b848454610667565b825550505050565b5f5f905090565b610728610719565b6107338184846106f4565b505050565b5b818110156107565761074b5f82610720565b600181019050610739565b5050565b601f82111561079b5761076c8161063a565b6107758461064c565b81016020851015610784578190505b6107986107908561064c565b830182610738565b50505b505050565b5f82821c905092915050565b5f6107bb5f19846008026107a0565b1980831691505092915050565b5f6107d383836107ac565b9150826002028217905092915050565b6107ec826103c7565b67ffffffffffffffff8111156108055761080461025c565b5b61080f825461053a565b61081a82828561075a565b5f60209050601f83116001811461084b575f8415610839578287015190505b61084385826107c8565b8655506108aa565b601f1984166108598661063a565b5f5b828110156108805784890151825560018201915060208501945060208101905061085b565b8683101561089d5784890151610899601f8916826107ac565b8355505b6001600288020188555050505b50505050505056fea264697066735822122080e4d051fe12dc773d8df2bf0f9fb569e0e82263c2a619429d3002a40d47bcf164736f6c637828302e382e32372d646576656c6f702e323032342e382e31322b636f6d6d69742e34656361626263660059",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	fmt.Println(string(common.FromHex(StoreBin)))

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string filePath) view returns(string)
func (_Store *StoreCaller) Get(opts *bind.CallOpts, filePath string) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "get", filePath)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string filePath) view returns(string)
func (_Store *StoreSession) Get(filePath string) (string, error) {
	return _Store.Contract.Get(&_Store.CallOpts, filePath)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string filePath) view returns(string)
func (_Store *StoreCallerSession) Get(filePath string) (string, error) {
	return _Store.Contract.Get(&_Store.CallOpts, filePath)
}

// Save is a paid mutator transaction binding the contract method 0x962939b8.
//
// Solidity: function save(string filePath, string cid) returns()
func (_Store *StoreTransactor) Save(opts *bind.TransactOpts, filePath string, cid string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "save", filePath, cid)
}

// Save is a paid mutator transaction binding the contract method 0x962939b8.
//
// Solidity: function save(string filePath, string cid) returns()
func (_Store *StoreSession) Save(filePath string, cid string) (*types.Transaction, error) {
	return _Store.Contract.Save(&_Store.TransactOpts, filePath, cid)
}

// Save is a paid mutator transaction binding the contract method 0x962939b8.
//
// Solidity: function save(string filePath, string cid) returns()
func (_Store *StoreTransactorSession) Save(filePath string, cid string) (*types.Transaction, error) {
	return _Store.Contract.Save(&_Store.TransactOpts, filePath, cid)
}

// StoreFileSavedIterator is returned from FilterFileSaved and is used to iterate over the raw logs and unpacked data for FileSaved events raised by the Store contract.
type StoreFileSavedIterator struct {
	Event *StoreFileSaved // Event containing the contract specifics and raw log

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
func (it *StoreFileSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreFileSaved)
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
		it.Event = new(StoreFileSaved)
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
func (it *StoreFileSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreFileSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreFileSaved represents a FileSaved event raised by the Store contract.
type StoreFileSaved struct {
	FilePath common.Hash
	Cid      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFileSaved is a free log retrieval operation binding the contract event 0xcadc7184c55de53d424a8e73df016947523f46250bb8957192d0d084403dfd25.
//
// Solidity: event FileSaved(string indexed filePath, string cid)
func (_Store *StoreFilterer) FilterFileSaved(opts *bind.FilterOpts, filePath []string) (*StoreFileSavedIterator, error) {

	var filePathRule []interface{}
	for _, filePathItem := range filePath {
		filePathRule = append(filePathRule, filePathItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "FileSaved", filePathRule)
	if err != nil {
		return nil, err
	}
	return &StoreFileSavedIterator{contract: _Store.contract, event: "FileSaved", logs: logs, sub: sub}, nil
}

// WatchFileSaved is a free log subscription operation binding the contract event 0xcadc7184c55de53d424a8e73df016947523f46250bb8957192d0d084403dfd25.
//
// Solidity: event FileSaved(string indexed filePath, string cid)
func (_Store *StoreFilterer) WatchFileSaved(opts *bind.WatchOpts, sink chan<- *StoreFileSaved, filePath []string) (event.Subscription, error) {

	var filePathRule []interface{}
	for _, filePathItem := range filePath {
		filePathRule = append(filePathRule, filePathItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "FileSaved", filePathRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreFileSaved)
				if err := _Store.contract.UnpackLog(event, "FileSaved", log); err != nil {
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

// ParseFileSaved is a log parse operation binding the contract event 0xcadc7184c55de53d424a8e73df016947523f46250bb8957192d0d084403dfd25.
//
// Solidity: event FileSaved(string indexed filePath, string cid)
func (_Store *StoreFilterer) ParseFileSaved(log types.Log) (*StoreFileSaved, error) {
	event := new(StoreFileSaved)
	if err := _Store.contract.UnpackLog(event, "FileSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
