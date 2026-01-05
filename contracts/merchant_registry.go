// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// IMerchantRegistryMerchant is an auto generated low-level Go binding around an user-defined struct.
type IMerchantRegistryMerchant struct {
	Owner              common.Address
	PayoutWallet       common.Address
	MetadataURI        string
	Exists             bool
	VerificationStatus uint8
	CreatedAt          *big.Int
}

// MerchantRegistryMetaData contains all meta data concerning the MerchantRegistry contract.
var MerchantRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMerchantCreatedAt\",\"inputs\":[{\"name\":\"_merchantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMerchantInfo\",\"inputs\":[{\"name\":\"_merchantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIMerchantRegistry.Merchant\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"payoutWallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"metadataURI\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"verificationStatus\",\"type\":\"uint8\",\"internalType\":\"enumIMerchantRegistry.VerificationStatus\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMerchantVerificationStatus\",\"inputs\":[{\"name\":\"_merchantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIMerchantRegistry.VerificationStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isMerchantVerified\",\"inputs\":[{\"name\":\"_merchantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"isVerified\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerMerchant\",\"inputs\":[{\"name\":\"_payoutWallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_metadataUri\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"merchantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMerchant\",\"inputs\":[{\"name\":\"_merchantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_payoutWallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_metadataUri\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMerchantVerificationStatus\",\"inputs\":[{\"name\":\"_merchantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_newStatus\",\"type\":\"uint8\",\"internalType\":\"enumIMerchantRegistry.VerificationStatus\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MerchantRegistered\",\"inputs\":[{\"name\":\"merchantId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"payoutWallet\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MerchantUpdated\",\"inputs\":[{\"name\":\"merchantId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MerchantVerificationStatusUpdated\",\"inputs\":[{\"name\":\"merchantId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"oldStatus\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIMerchantRegistry.VerificationStatus\"},{\"name\":\"newStatus\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIMerchantRegistry.VerificationStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerchantRegistry__InvalidMetadataUri\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerchantRegistry__InvalidVerificationStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerchantRegistry__MerchantAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerchantRegistry__MerchantNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerchantRegistry__MerchantNotVerified\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerchantRegistry__ThrowZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerchantRegistry__UnauthorizedMerchant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerchantRegistry__UnverifiedMerchant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// MerchantRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use MerchantRegistryMetaData.ABI instead.
var MerchantRegistryABI = MerchantRegistryMetaData.ABI

// MerchantRegistry is an auto generated Go binding around an Ethereum contract.
type MerchantRegistry struct {
	MerchantRegistryCaller     // Read-only binding to the contract
	MerchantRegistryTransactor // Write-only binding to the contract
	MerchantRegistryFilterer   // Log filterer for contract events
}

// MerchantRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerchantRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerchantRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerchantRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerchantRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerchantRegistrySession struct {
	Contract     *MerchantRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerchantRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerchantRegistryCallerSession struct {
	Contract *MerchantRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// MerchantRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerchantRegistryTransactorSession struct {
	Contract     *MerchantRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// MerchantRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerchantRegistryRaw struct {
	Contract *MerchantRegistry // Generic contract binding to access the raw methods on
}

// MerchantRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerchantRegistryCallerRaw struct {
	Contract *MerchantRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// MerchantRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerchantRegistryTransactorRaw struct {
	Contract *MerchantRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerchantRegistry creates a new instance of MerchantRegistry, bound to a specific deployed contract.
func NewMerchantRegistry(address common.Address, backend bind.ContractBackend) (*MerchantRegistry, error) {
	contract, err := bindMerchantRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistry{MerchantRegistryCaller: MerchantRegistryCaller{contract: contract}, MerchantRegistryTransactor: MerchantRegistryTransactor{contract: contract}, MerchantRegistryFilterer: MerchantRegistryFilterer{contract: contract}}, nil
}

// NewMerchantRegistryCaller creates a new read-only instance of MerchantRegistry, bound to a specific deployed contract.
func NewMerchantRegistryCaller(address common.Address, caller bind.ContractCaller) (*MerchantRegistryCaller, error) {
	contract, err := bindMerchantRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryCaller{contract: contract}, nil
}

// NewMerchantRegistryTransactor creates a new write-only instance of MerchantRegistry, bound to a specific deployed contract.
func NewMerchantRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*MerchantRegistryTransactor, error) {
	contract, err := bindMerchantRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryTransactor{contract: contract}, nil
}

// NewMerchantRegistryFilterer creates a new log filterer instance of MerchantRegistry, bound to a specific deployed contract.
func NewMerchantRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*MerchantRegistryFilterer, error) {
	contract, err := bindMerchantRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryFilterer{contract: contract}, nil
}

// bindMerchantRegistry binds a generic wrapper to an already deployed contract.
func bindMerchantRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MerchantRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerchantRegistry *MerchantRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerchantRegistry.Contract.MerchantRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerchantRegistry *MerchantRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.MerchantRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerchantRegistry *MerchantRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.MerchantRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerchantRegistry *MerchantRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerchantRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerchantRegistry *MerchantRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerchantRegistry *MerchantRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetMerchantCreatedAt is a free data retrieval call binding the contract method 0x5ea3550f.
//
// Solidity: function getMerchantCreatedAt(bytes32 _merchantId) view returns(uint256 createdAt)
func (_MerchantRegistry *MerchantRegistryCaller) GetMerchantCreatedAt(opts *bind.CallOpts, _merchantId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _MerchantRegistry.contract.Call(opts, &out, "getMerchantCreatedAt", _merchantId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMerchantCreatedAt is a free data retrieval call binding the contract method 0x5ea3550f.
//
// Solidity: function getMerchantCreatedAt(bytes32 _merchantId) view returns(uint256 createdAt)
func (_MerchantRegistry *MerchantRegistrySession) GetMerchantCreatedAt(_merchantId [32]byte) (*big.Int, error) {
	return _MerchantRegistry.Contract.GetMerchantCreatedAt(&_MerchantRegistry.CallOpts, _merchantId)
}

// GetMerchantCreatedAt is a free data retrieval call binding the contract method 0x5ea3550f.
//
// Solidity: function getMerchantCreatedAt(bytes32 _merchantId) view returns(uint256 createdAt)
func (_MerchantRegistry *MerchantRegistryCallerSession) GetMerchantCreatedAt(_merchantId [32]byte) (*big.Int, error) {
	return _MerchantRegistry.Contract.GetMerchantCreatedAt(&_MerchantRegistry.CallOpts, _merchantId)
}

// GetMerchantInfo is a free data retrieval call binding the contract method 0x2f389922.
//
// Solidity: function getMerchantInfo(bytes32 _merchantId) view returns((address,address,string,bool,uint8,uint256))
func (_MerchantRegistry *MerchantRegistryCaller) GetMerchantInfo(opts *bind.CallOpts, _merchantId [32]byte) (IMerchantRegistryMerchant, error) {
	var out []interface{}
	err := _MerchantRegistry.contract.Call(opts, &out, "getMerchantInfo", _merchantId)

	if err != nil {
		return *new(IMerchantRegistryMerchant), err
	}

	out0 := *abi.ConvertType(out[0], new(IMerchantRegistryMerchant)).(*IMerchantRegistryMerchant)

	return out0, err

}

// GetMerchantInfo is a free data retrieval call binding the contract method 0x2f389922.
//
// Solidity: function getMerchantInfo(bytes32 _merchantId) view returns((address,address,string,bool,uint8,uint256))
func (_MerchantRegistry *MerchantRegistrySession) GetMerchantInfo(_merchantId [32]byte) (IMerchantRegistryMerchant, error) {
	return _MerchantRegistry.Contract.GetMerchantInfo(&_MerchantRegistry.CallOpts, _merchantId)
}

// GetMerchantInfo is a free data retrieval call binding the contract method 0x2f389922.
//
// Solidity: function getMerchantInfo(bytes32 _merchantId) view returns((address,address,string,bool,uint8,uint256))
func (_MerchantRegistry *MerchantRegistryCallerSession) GetMerchantInfo(_merchantId [32]byte) (IMerchantRegistryMerchant, error) {
	return _MerchantRegistry.Contract.GetMerchantInfo(&_MerchantRegistry.CallOpts, _merchantId)
}

// GetMerchantVerificationStatus is a free data retrieval call binding the contract method 0x18b981fc.
//
// Solidity: function getMerchantVerificationStatus(bytes32 _merchantId) view returns(uint8 status)
func (_MerchantRegistry *MerchantRegistryCaller) GetMerchantVerificationStatus(opts *bind.CallOpts, _merchantId [32]byte) (uint8, error) {
	var out []interface{}
	err := _MerchantRegistry.contract.Call(opts, &out, "getMerchantVerificationStatus", _merchantId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetMerchantVerificationStatus is a free data retrieval call binding the contract method 0x18b981fc.
//
// Solidity: function getMerchantVerificationStatus(bytes32 _merchantId) view returns(uint8 status)
func (_MerchantRegistry *MerchantRegistrySession) GetMerchantVerificationStatus(_merchantId [32]byte) (uint8, error) {
	return _MerchantRegistry.Contract.GetMerchantVerificationStatus(&_MerchantRegistry.CallOpts, _merchantId)
}

// GetMerchantVerificationStatus is a free data retrieval call binding the contract method 0x18b981fc.
//
// Solidity: function getMerchantVerificationStatus(bytes32 _merchantId) view returns(uint8 status)
func (_MerchantRegistry *MerchantRegistryCallerSession) GetMerchantVerificationStatus(_merchantId [32]byte) (uint8, error) {
	return _MerchantRegistry.Contract.GetMerchantVerificationStatus(&_MerchantRegistry.CallOpts, _merchantId)
}

// IsMerchantVerified is a free data retrieval call binding the contract method 0x24e922aa.
//
// Solidity: function isMerchantVerified(bytes32 _merchantId) view returns(bool isVerified)
func (_MerchantRegistry *MerchantRegistryCaller) IsMerchantVerified(opts *bind.CallOpts, _merchantId [32]byte) (bool, error) {
	var out []interface{}
	err := _MerchantRegistry.contract.Call(opts, &out, "isMerchantVerified", _merchantId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMerchantVerified is a free data retrieval call binding the contract method 0x24e922aa.
//
// Solidity: function isMerchantVerified(bytes32 _merchantId) view returns(bool isVerified)
func (_MerchantRegistry *MerchantRegistrySession) IsMerchantVerified(_merchantId [32]byte) (bool, error) {
	return _MerchantRegistry.Contract.IsMerchantVerified(&_MerchantRegistry.CallOpts, _merchantId)
}

// IsMerchantVerified is a free data retrieval call binding the contract method 0x24e922aa.
//
// Solidity: function isMerchantVerified(bytes32 _merchantId) view returns(bool isVerified)
func (_MerchantRegistry *MerchantRegistryCallerSession) IsMerchantVerified(_merchantId [32]byte) (bool, error) {
	return _MerchantRegistry.Contract.IsMerchantVerified(&_MerchantRegistry.CallOpts, _merchantId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MerchantRegistry *MerchantRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MerchantRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MerchantRegistry *MerchantRegistrySession) Owner() (common.Address, error) {
	return _MerchantRegistry.Contract.Owner(&_MerchantRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MerchantRegistry *MerchantRegistryCallerSession) Owner() (common.Address, error) {
	return _MerchantRegistry.Contract.Owner(&_MerchantRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MerchantRegistry *MerchantRegistryCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MerchantRegistry.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MerchantRegistry *MerchantRegistrySession) Paused() (bool, error) {
	return _MerchantRegistry.Contract.Paused(&_MerchantRegistry.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_MerchantRegistry *MerchantRegistryCallerSession) Paused() (bool, error) {
	return _MerchantRegistry.Contract.Paused(&_MerchantRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MerchantRegistry *MerchantRegistryCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MerchantRegistry.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MerchantRegistry *MerchantRegistrySession) PendingOwner() (common.Address, error) {
	return _MerchantRegistry.Contract.PendingOwner(&_MerchantRegistry.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_MerchantRegistry *MerchantRegistryCallerSession) PendingOwner() (common.Address, error) {
	return _MerchantRegistry.Contract.PendingOwner(&_MerchantRegistry.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MerchantRegistry *MerchantRegistryTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantRegistry.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MerchantRegistry *MerchantRegistrySession) AcceptOwnership() (*types.Transaction, error) {
	return _MerchantRegistry.Contract.AcceptOwnership(&_MerchantRegistry.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_MerchantRegistry *MerchantRegistryTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _MerchantRegistry.Contract.AcceptOwnership(&_MerchantRegistry.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_MerchantRegistry *MerchantRegistryTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _MerchantRegistry.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_MerchantRegistry *MerchantRegistrySession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.Initialize(&_MerchantRegistry.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_MerchantRegistry *MerchantRegistryTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.Initialize(&_MerchantRegistry.TransactOpts, initialOwner)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MerchantRegistry *MerchantRegistryTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantRegistry.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MerchantRegistry *MerchantRegistrySession) Pause() (*types.Transaction, error) {
	return _MerchantRegistry.Contract.Pause(&_MerchantRegistry.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_MerchantRegistry *MerchantRegistryTransactorSession) Pause() (*types.Transaction, error) {
	return _MerchantRegistry.Contract.Pause(&_MerchantRegistry.TransactOpts)
}

// RegisterMerchant is a paid mutator transaction binding the contract method 0xa6c8a384.
//
// Solidity: function registerMerchant(address _payoutWallet, string _metadataUri) returns(bytes32 merchantId)
func (_MerchantRegistry *MerchantRegistryTransactor) RegisterMerchant(opts *bind.TransactOpts, _payoutWallet common.Address, _metadataUri string) (*types.Transaction, error) {
	return _MerchantRegistry.contract.Transact(opts, "registerMerchant", _payoutWallet, _metadataUri)
}

// RegisterMerchant is a paid mutator transaction binding the contract method 0xa6c8a384.
//
// Solidity: function registerMerchant(address _payoutWallet, string _metadataUri) returns(bytes32 merchantId)
func (_MerchantRegistry *MerchantRegistrySession) RegisterMerchant(_payoutWallet common.Address, _metadataUri string) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.RegisterMerchant(&_MerchantRegistry.TransactOpts, _payoutWallet, _metadataUri)
}

// RegisterMerchant is a paid mutator transaction binding the contract method 0xa6c8a384.
//
// Solidity: function registerMerchant(address _payoutWallet, string _metadataUri) returns(bytes32 merchantId)
func (_MerchantRegistry *MerchantRegistryTransactorSession) RegisterMerchant(_payoutWallet common.Address, _metadataUri string) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.RegisterMerchant(&_MerchantRegistry.TransactOpts, _payoutWallet, _metadataUri)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantRegistry *MerchantRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantRegistry *MerchantRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _MerchantRegistry.Contract.RenounceOwnership(&_MerchantRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MerchantRegistry *MerchantRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MerchantRegistry.Contract.RenounceOwnership(&_MerchantRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MerchantRegistry *MerchantRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MerchantRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MerchantRegistry *MerchantRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.TransferOwnership(&_MerchantRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MerchantRegistry *MerchantRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.TransferOwnership(&_MerchantRegistry.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MerchantRegistry *MerchantRegistryTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerchantRegistry.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MerchantRegistry *MerchantRegistrySession) Unpause() (*types.Transaction, error) {
	return _MerchantRegistry.Contract.Unpause(&_MerchantRegistry.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_MerchantRegistry *MerchantRegistryTransactorSession) Unpause() (*types.Transaction, error) {
	return _MerchantRegistry.Contract.Unpause(&_MerchantRegistry.TransactOpts)
}

// UpdateMerchant is a paid mutator transaction binding the contract method 0x51397ffe.
//
// Solidity: function updateMerchant(bytes32 _merchantId, address _payoutWallet, string _metadataUri) returns()
func (_MerchantRegistry *MerchantRegistryTransactor) UpdateMerchant(opts *bind.TransactOpts, _merchantId [32]byte, _payoutWallet common.Address, _metadataUri string) (*types.Transaction, error) {
	return _MerchantRegistry.contract.Transact(opts, "updateMerchant", _merchantId, _payoutWallet, _metadataUri)
}

// UpdateMerchant is a paid mutator transaction binding the contract method 0x51397ffe.
//
// Solidity: function updateMerchant(bytes32 _merchantId, address _payoutWallet, string _metadataUri) returns()
func (_MerchantRegistry *MerchantRegistrySession) UpdateMerchant(_merchantId [32]byte, _payoutWallet common.Address, _metadataUri string) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.UpdateMerchant(&_MerchantRegistry.TransactOpts, _merchantId, _payoutWallet, _metadataUri)
}

// UpdateMerchant is a paid mutator transaction binding the contract method 0x51397ffe.
//
// Solidity: function updateMerchant(bytes32 _merchantId, address _payoutWallet, string _metadataUri) returns()
func (_MerchantRegistry *MerchantRegistryTransactorSession) UpdateMerchant(_merchantId [32]byte, _payoutWallet common.Address, _metadataUri string) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.UpdateMerchant(&_MerchantRegistry.TransactOpts, _merchantId, _payoutWallet, _metadataUri)
}

// UpdateMerchantVerificationStatus is a paid mutator transaction binding the contract method 0x1d206508.
//
// Solidity: function updateMerchantVerificationStatus(bytes32 _merchantId, uint8 _newStatus) returns()
func (_MerchantRegistry *MerchantRegistryTransactor) UpdateMerchantVerificationStatus(opts *bind.TransactOpts, _merchantId [32]byte, _newStatus uint8) (*types.Transaction, error) {
	return _MerchantRegistry.contract.Transact(opts, "updateMerchantVerificationStatus", _merchantId, _newStatus)
}

// UpdateMerchantVerificationStatus is a paid mutator transaction binding the contract method 0x1d206508.
//
// Solidity: function updateMerchantVerificationStatus(bytes32 _merchantId, uint8 _newStatus) returns()
func (_MerchantRegistry *MerchantRegistrySession) UpdateMerchantVerificationStatus(_merchantId [32]byte, _newStatus uint8) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.UpdateMerchantVerificationStatus(&_MerchantRegistry.TransactOpts, _merchantId, _newStatus)
}

// UpdateMerchantVerificationStatus is a paid mutator transaction binding the contract method 0x1d206508.
//
// Solidity: function updateMerchantVerificationStatus(bytes32 _merchantId, uint8 _newStatus) returns()
func (_MerchantRegistry *MerchantRegistryTransactorSession) UpdateMerchantVerificationStatus(_merchantId [32]byte, _newStatus uint8) (*types.Transaction, error) {
	return _MerchantRegistry.Contract.UpdateMerchantVerificationStatus(&_MerchantRegistry.TransactOpts, _merchantId, _newStatus)
}

// MerchantRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the MerchantRegistry contract.
type MerchantRegistryInitializedIterator struct {
	Event *MerchantRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *MerchantRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantRegistryInitialized)
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
		it.Event = new(MerchantRegistryInitialized)
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
func (it *MerchantRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantRegistryInitialized represents a Initialized event raised by the MerchantRegistry contract.
type MerchantRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MerchantRegistry *MerchantRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*MerchantRegistryInitializedIterator, error) {

	logs, sub, err := _MerchantRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryInitializedIterator{contract: _MerchantRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MerchantRegistry *MerchantRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *MerchantRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _MerchantRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantRegistryInitialized)
				if err := _MerchantRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_MerchantRegistry *MerchantRegistryFilterer) ParseInitialized(log types.Log) (*MerchantRegistryInitialized, error) {
	event := new(MerchantRegistryInitialized)
	if err := _MerchantRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantRegistryMerchantRegisteredIterator is returned from FilterMerchantRegistered and is used to iterate over the raw logs and unpacked data for MerchantRegistered events raised by the MerchantRegistry contract.
type MerchantRegistryMerchantRegisteredIterator struct {
	Event *MerchantRegistryMerchantRegistered // Event containing the contract specifics and raw log

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
func (it *MerchantRegistryMerchantRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantRegistryMerchantRegistered)
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
		it.Event = new(MerchantRegistryMerchantRegistered)
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
func (it *MerchantRegistryMerchantRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantRegistryMerchantRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantRegistryMerchantRegistered represents a MerchantRegistered event raised by the MerchantRegistry contract.
type MerchantRegistryMerchantRegistered struct {
	MerchantId   [32]byte
	Owner        common.Address
	PayoutWallet common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMerchantRegistered is a free log retrieval operation binding the contract event 0x71ef82f4d8c1d0e7d595951a0ed57ed7232d28b1f13c9106d26ca7ae31001164.
//
// Solidity: event MerchantRegistered(bytes32 indexed merchantId, address owner, address payoutWallet)
func (_MerchantRegistry *MerchantRegistryFilterer) FilterMerchantRegistered(opts *bind.FilterOpts, merchantId [][32]byte) (*MerchantRegistryMerchantRegisteredIterator, error) {

	var merchantIdRule []interface{}
	for _, merchantIdItem := range merchantId {
		merchantIdRule = append(merchantIdRule, merchantIdItem)
	}

	logs, sub, err := _MerchantRegistry.contract.FilterLogs(opts, "MerchantRegistered", merchantIdRule)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryMerchantRegisteredIterator{contract: _MerchantRegistry.contract, event: "MerchantRegistered", logs: logs, sub: sub}, nil
}

// WatchMerchantRegistered is a free log subscription operation binding the contract event 0x71ef82f4d8c1d0e7d595951a0ed57ed7232d28b1f13c9106d26ca7ae31001164.
//
// Solidity: event MerchantRegistered(bytes32 indexed merchantId, address owner, address payoutWallet)
func (_MerchantRegistry *MerchantRegistryFilterer) WatchMerchantRegistered(opts *bind.WatchOpts, sink chan<- *MerchantRegistryMerchantRegistered, merchantId [][32]byte) (event.Subscription, error) {

	var merchantIdRule []interface{}
	for _, merchantIdItem := range merchantId {
		merchantIdRule = append(merchantIdRule, merchantIdItem)
	}

	logs, sub, err := _MerchantRegistry.contract.WatchLogs(opts, "MerchantRegistered", merchantIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantRegistryMerchantRegistered)
				if err := _MerchantRegistry.contract.UnpackLog(event, "MerchantRegistered", log); err != nil {
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

// ParseMerchantRegistered is a log parse operation binding the contract event 0x71ef82f4d8c1d0e7d595951a0ed57ed7232d28b1f13c9106d26ca7ae31001164.
//
// Solidity: event MerchantRegistered(bytes32 indexed merchantId, address owner, address payoutWallet)
func (_MerchantRegistry *MerchantRegistryFilterer) ParseMerchantRegistered(log types.Log) (*MerchantRegistryMerchantRegistered, error) {
	event := new(MerchantRegistryMerchantRegistered)
	if err := _MerchantRegistry.contract.UnpackLog(event, "MerchantRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantRegistryMerchantUpdatedIterator is returned from FilterMerchantUpdated and is used to iterate over the raw logs and unpacked data for MerchantUpdated events raised by the MerchantRegistry contract.
type MerchantRegistryMerchantUpdatedIterator struct {
	Event *MerchantRegistryMerchantUpdated // Event containing the contract specifics and raw log

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
func (it *MerchantRegistryMerchantUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantRegistryMerchantUpdated)
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
		it.Event = new(MerchantRegistryMerchantUpdated)
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
func (it *MerchantRegistryMerchantUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantRegistryMerchantUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantRegistryMerchantUpdated represents a MerchantUpdated event raised by the MerchantRegistry contract.
type MerchantRegistryMerchantUpdated struct {
	MerchantId [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMerchantUpdated is a free log retrieval operation binding the contract event 0xbb3b2eb559c2003bd8feff754f7372c4cdc22fb3481b3d14026c53c875eb3d6b.
//
// Solidity: event MerchantUpdated(bytes32 indexed merchantId)
func (_MerchantRegistry *MerchantRegistryFilterer) FilterMerchantUpdated(opts *bind.FilterOpts, merchantId [][32]byte) (*MerchantRegistryMerchantUpdatedIterator, error) {

	var merchantIdRule []interface{}
	for _, merchantIdItem := range merchantId {
		merchantIdRule = append(merchantIdRule, merchantIdItem)
	}

	logs, sub, err := _MerchantRegistry.contract.FilterLogs(opts, "MerchantUpdated", merchantIdRule)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryMerchantUpdatedIterator{contract: _MerchantRegistry.contract, event: "MerchantUpdated", logs: logs, sub: sub}, nil
}

// WatchMerchantUpdated is a free log subscription operation binding the contract event 0xbb3b2eb559c2003bd8feff754f7372c4cdc22fb3481b3d14026c53c875eb3d6b.
//
// Solidity: event MerchantUpdated(bytes32 indexed merchantId)
func (_MerchantRegistry *MerchantRegistryFilterer) WatchMerchantUpdated(opts *bind.WatchOpts, sink chan<- *MerchantRegistryMerchantUpdated, merchantId [][32]byte) (event.Subscription, error) {

	var merchantIdRule []interface{}
	for _, merchantIdItem := range merchantId {
		merchantIdRule = append(merchantIdRule, merchantIdItem)
	}

	logs, sub, err := _MerchantRegistry.contract.WatchLogs(opts, "MerchantUpdated", merchantIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantRegistryMerchantUpdated)
				if err := _MerchantRegistry.contract.UnpackLog(event, "MerchantUpdated", log); err != nil {
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

// ParseMerchantUpdated is a log parse operation binding the contract event 0xbb3b2eb559c2003bd8feff754f7372c4cdc22fb3481b3d14026c53c875eb3d6b.
//
// Solidity: event MerchantUpdated(bytes32 indexed merchantId)
func (_MerchantRegistry *MerchantRegistryFilterer) ParseMerchantUpdated(log types.Log) (*MerchantRegistryMerchantUpdated, error) {
	event := new(MerchantRegistryMerchantUpdated)
	if err := _MerchantRegistry.contract.UnpackLog(event, "MerchantUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantRegistryMerchantVerificationStatusUpdatedIterator is returned from FilterMerchantVerificationStatusUpdated and is used to iterate over the raw logs and unpacked data for MerchantVerificationStatusUpdated events raised by the MerchantRegistry contract.
type MerchantRegistryMerchantVerificationStatusUpdatedIterator struct {
	Event *MerchantRegistryMerchantVerificationStatusUpdated // Event containing the contract specifics and raw log

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
func (it *MerchantRegistryMerchantVerificationStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantRegistryMerchantVerificationStatusUpdated)
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
		it.Event = new(MerchantRegistryMerchantVerificationStatusUpdated)
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
func (it *MerchantRegistryMerchantVerificationStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantRegistryMerchantVerificationStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantRegistryMerchantVerificationStatusUpdated represents a MerchantVerificationStatusUpdated event raised by the MerchantRegistry contract.
type MerchantRegistryMerchantVerificationStatusUpdated struct {
	MerchantId [32]byte
	OldStatus  uint8
	NewStatus  uint8
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMerchantVerificationStatusUpdated is a free log retrieval operation binding the contract event 0x2b13d71d3626fce95a015b9f1758cad49dad02c848bf55afc4102cb0ddb0b3af.
//
// Solidity: event MerchantVerificationStatusUpdated(bytes32 indexed merchantId, uint8 oldStatus, uint8 newStatus)
func (_MerchantRegistry *MerchantRegistryFilterer) FilterMerchantVerificationStatusUpdated(opts *bind.FilterOpts, merchantId [][32]byte) (*MerchantRegistryMerchantVerificationStatusUpdatedIterator, error) {

	var merchantIdRule []interface{}
	for _, merchantIdItem := range merchantId {
		merchantIdRule = append(merchantIdRule, merchantIdItem)
	}

	logs, sub, err := _MerchantRegistry.contract.FilterLogs(opts, "MerchantVerificationStatusUpdated", merchantIdRule)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryMerchantVerificationStatusUpdatedIterator{contract: _MerchantRegistry.contract, event: "MerchantVerificationStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchMerchantVerificationStatusUpdated is a free log subscription operation binding the contract event 0x2b13d71d3626fce95a015b9f1758cad49dad02c848bf55afc4102cb0ddb0b3af.
//
// Solidity: event MerchantVerificationStatusUpdated(bytes32 indexed merchantId, uint8 oldStatus, uint8 newStatus)
func (_MerchantRegistry *MerchantRegistryFilterer) WatchMerchantVerificationStatusUpdated(opts *bind.WatchOpts, sink chan<- *MerchantRegistryMerchantVerificationStatusUpdated, merchantId [][32]byte) (event.Subscription, error) {

	var merchantIdRule []interface{}
	for _, merchantIdItem := range merchantId {
		merchantIdRule = append(merchantIdRule, merchantIdItem)
	}

	logs, sub, err := _MerchantRegistry.contract.WatchLogs(opts, "MerchantVerificationStatusUpdated", merchantIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantRegistryMerchantVerificationStatusUpdated)
				if err := _MerchantRegistry.contract.UnpackLog(event, "MerchantVerificationStatusUpdated", log); err != nil {
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

// ParseMerchantVerificationStatusUpdated is a log parse operation binding the contract event 0x2b13d71d3626fce95a015b9f1758cad49dad02c848bf55afc4102cb0ddb0b3af.
//
// Solidity: event MerchantVerificationStatusUpdated(bytes32 indexed merchantId, uint8 oldStatus, uint8 newStatus)
func (_MerchantRegistry *MerchantRegistryFilterer) ParseMerchantVerificationStatusUpdated(log types.Log) (*MerchantRegistryMerchantVerificationStatusUpdated, error) {
	event := new(MerchantRegistryMerchantVerificationStatusUpdated)
	if err := _MerchantRegistry.contract.UnpackLog(event, "MerchantVerificationStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantRegistryOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the MerchantRegistry contract.
type MerchantRegistryOwnershipTransferStartedIterator struct {
	Event *MerchantRegistryOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *MerchantRegistryOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantRegistryOwnershipTransferStarted)
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
		it.Event = new(MerchantRegistryOwnershipTransferStarted)
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
func (it *MerchantRegistryOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantRegistryOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantRegistryOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the MerchantRegistry contract.
type MerchantRegistryOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MerchantRegistry *MerchantRegistryFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MerchantRegistryOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerchantRegistry.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryOwnershipTransferStartedIterator{contract: _MerchantRegistry.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MerchantRegistry *MerchantRegistryFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *MerchantRegistryOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerchantRegistry.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantRegistryOwnershipTransferStarted)
				if err := _MerchantRegistry.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_MerchantRegistry *MerchantRegistryFilterer) ParseOwnershipTransferStarted(log types.Log) (*MerchantRegistryOwnershipTransferStarted, error) {
	event := new(MerchantRegistryOwnershipTransferStarted)
	if err := _MerchantRegistry.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MerchantRegistry contract.
type MerchantRegistryOwnershipTransferredIterator struct {
	Event *MerchantRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MerchantRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantRegistryOwnershipTransferred)
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
		it.Event = new(MerchantRegistryOwnershipTransferred)
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
func (it *MerchantRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the MerchantRegistry contract.
type MerchantRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MerchantRegistry *MerchantRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MerchantRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerchantRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryOwnershipTransferredIterator{contract: _MerchantRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MerchantRegistry *MerchantRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MerchantRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerchantRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantRegistryOwnershipTransferred)
				if err := _MerchantRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MerchantRegistry *MerchantRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*MerchantRegistryOwnershipTransferred, error) {
	event := new(MerchantRegistryOwnershipTransferred)
	if err := _MerchantRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantRegistryPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the MerchantRegistry contract.
type MerchantRegistryPausedIterator struct {
	Event *MerchantRegistryPaused // Event containing the contract specifics and raw log

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
func (it *MerchantRegistryPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantRegistryPaused)
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
		it.Event = new(MerchantRegistryPaused)
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
func (it *MerchantRegistryPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantRegistryPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantRegistryPaused represents a Paused event raised by the MerchantRegistry contract.
type MerchantRegistryPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MerchantRegistry *MerchantRegistryFilterer) FilterPaused(opts *bind.FilterOpts) (*MerchantRegistryPausedIterator, error) {

	logs, sub, err := _MerchantRegistry.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryPausedIterator{contract: _MerchantRegistry.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MerchantRegistry *MerchantRegistryFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *MerchantRegistryPaused) (event.Subscription, error) {

	logs, sub, err := _MerchantRegistry.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantRegistryPaused)
				if err := _MerchantRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_MerchantRegistry *MerchantRegistryFilterer) ParsePaused(log types.Log) (*MerchantRegistryPaused, error) {
	event := new(MerchantRegistryPaused)
	if err := _MerchantRegistry.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerchantRegistryUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the MerchantRegistry contract.
type MerchantRegistryUnpausedIterator struct {
	Event *MerchantRegistryUnpaused // Event containing the contract specifics and raw log

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
func (it *MerchantRegistryUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerchantRegistryUnpaused)
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
		it.Event = new(MerchantRegistryUnpaused)
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
func (it *MerchantRegistryUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerchantRegistryUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerchantRegistryUnpaused represents a Unpaused event raised by the MerchantRegistry contract.
type MerchantRegistryUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MerchantRegistry *MerchantRegistryFilterer) FilterUnpaused(opts *bind.FilterOpts) (*MerchantRegistryUnpausedIterator, error) {

	logs, sub, err := _MerchantRegistry.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &MerchantRegistryUnpausedIterator{contract: _MerchantRegistry.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MerchantRegistry *MerchantRegistryFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *MerchantRegistryUnpaused) (event.Subscription, error) {

	logs, sub, err := _MerchantRegistry.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerchantRegistryUnpaused)
				if err := _MerchantRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_MerchantRegistry *MerchantRegistryFilterer) ParseUnpaused(log types.Log) (*MerchantRegistryUnpaused, error) {
	event := new(MerchantRegistryUnpaused)
	if err := _MerchantRegistry.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
