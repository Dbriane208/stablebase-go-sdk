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

// IPaymentProcessorOrder is an auto generated low-level Go binding around an user-defined struct.
type IPaymentProcessorOrder struct {
	Payer          common.Address
	Token          common.Address
	MerchantId     [32]byte
	MerchantPayout common.Address
	Amount         *big.Int
	Exists         bool
	CurrentBps     *big.Int
	CreatedAt      *big.Int
	MetadataUri    string
	Status         uint8
}

// PaymentProcessorMetaData contains all meta data concerning the PaymentProcessor contract.
var PaymentProcessorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOrder\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createOrder\",\"inputs\":[{\"name\":\"_merchantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_metadataUri\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultPlatformFeeBps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"emergencyWithdraw\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emergencyWithdrawalEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getContractTokenBalance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMerchantTokenBalance\",\"inputs\":[{\"name\":\"merchant\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrder\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIPaymentProcessor.Order\",\"components\":[{\"name\":\"payer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"merchantId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"merchantPayout\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"exists\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"currentBps\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"metadataUri\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumIPaymentProcessor.OrderStatus\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrderCreatedAt\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"createdAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrderRemainingTime\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"remainingTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOrderStatus\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumIPaymentProcessor.OrderStatus\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlatformTokenBalance\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPlatformWallet\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_platformWallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_defaultPlatformFeeBps\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_merchantRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_orderExpirationTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOrderExpired\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"expired\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isTokenSupported\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"merchantRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractMerchantRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"orderExpirationTime\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"payOrder\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"refundOrder\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setEmergencyWithdrawalEnabled\",\"inputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTokenSupport\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"settleOrder\",\"inputs\":[{\"name\":\"_orderId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateMerchantRegistry\",\"inputs\":[{\"name\":\"newRegistry\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOrderExpirationTime\",\"inputs\":[{\"name\":\"newExpirationTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateProtocolAddress\",\"inputs\":[{\"name\":\"what\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"value\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"EmergencyWithdrawalEnabledUpdated\",\"inputs\":[{\"name\":\"enabled\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EmergencyWithdrawalSuccess\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxSlippageBpsUpdated\",\"inputs\":[{\"name\":\"oldSlippage\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newSlippage\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MerchantRegistryUpdated\",\"inputs\":[{\"name\":\"oldRegistry\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newRegistry\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderCancelled\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"payer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderCreated\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"payer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"merchantId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"merchantPayout\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIPaymentProcessor.OrderStatus\"},{\"name\":\"metadataUri\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderExpirationTimeUpdated\",\"inputs\":[{\"name\":\"oldTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newTime\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderExpired\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"payer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderPaid\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"payer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIPaymentProcessor.OrderStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderRefunded\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"payer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OrderSettled\",\"inputs\":[{\"name\":\"orderId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"merchant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"payer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"netAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumIPaymentProcessor.OrderStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolAddressUpdated\",\"inputs\":[{\"name\":\"what\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"platformWallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenSupportUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"supported\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"PaymentProcessor__EmergencyDisabled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentProcessor__EmergencyWithdrawalFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentProcessor__InsufficientBalance\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentProcessor__InvalidAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentProcessor__InvalidMetadataUri\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentProcessor__InvalidStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentProcessor__InvalidToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentProcessor__OrderAlreadyExists\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentProcessor__OrderExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentProcessor__OrderNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentProcessor__ThrowZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentProcessor__TokenNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentProcessor__TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentProcessor__UnauthorizedAccess\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"PaymentProcessor__UnverifiedMerchant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"TokensManager__InvalidBps\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokensManager__InvalidProviderToMerchant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokensManager__InvalidStatus\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokensManager__ThrowZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokensManager__TokenNotSupported\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TokensManager__WalletAlreadySet\",\"inputs\":[]}]",
}

// PaymentProcessorABI is the input ABI used to generate the binding from.
// Deprecated: Use PaymentProcessorMetaData.ABI instead.
var PaymentProcessorABI = PaymentProcessorMetaData.ABI

// PaymentProcessor is an auto generated Go binding around an Ethereum contract.
type PaymentProcessor struct {
	PaymentProcessorCaller     // Read-only binding to the contract
	PaymentProcessorTransactor // Write-only binding to the contract
	PaymentProcessorFilterer   // Log filterer for contract events
}

// PaymentProcessorCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentProcessorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentProcessorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentProcessorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentProcessorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentProcessorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentProcessorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentProcessorSession struct {
	Contract     *PaymentProcessor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentProcessorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentProcessorCallerSession struct {
	Contract *PaymentProcessorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PaymentProcessorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentProcessorTransactorSession struct {
	Contract     *PaymentProcessorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PaymentProcessorRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentProcessorRaw struct {
	Contract *PaymentProcessor // Generic contract binding to access the raw methods on
}

// PaymentProcessorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentProcessorCallerRaw struct {
	Contract *PaymentProcessorCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentProcessorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentProcessorTransactorRaw struct {
	Contract *PaymentProcessorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPaymentProcessor creates a new instance of PaymentProcessor, bound to a specific deployed contract.
func NewPaymentProcessor(address common.Address, backend bind.ContractBackend) (*PaymentProcessor, error) {
	contract, err := bindPaymentProcessor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessor{PaymentProcessorCaller: PaymentProcessorCaller{contract: contract}, PaymentProcessorTransactor: PaymentProcessorTransactor{contract: contract}, PaymentProcessorFilterer: PaymentProcessorFilterer{contract: contract}}, nil
}

// NewPaymentProcessorCaller creates a new read-only instance of PaymentProcessor, bound to a specific deployed contract.
func NewPaymentProcessorCaller(address common.Address, caller bind.ContractCaller) (*PaymentProcessorCaller, error) {
	contract, err := bindPaymentProcessor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorCaller{contract: contract}, nil
}

// NewPaymentProcessorTransactor creates a new write-only instance of PaymentProcessor, bound to a specific deployed contract.
func NewPaymentProcessorTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentProcessorTransactor, error) {
	contract, err := bindPaymentProcessor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorTransactor{contract: contract}, nil
}

// NewPaymentProcessorFilterer creates a new log filterer instance of PaymentProcessor, bound to a specific deployed contract.
func NewPaymentProcessorFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentProcessorFilterer, error) {
	contract, err := bindPaymentProcessor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorFilterer{contract: contract}, nil
}

// bindPaymentProcessor binds a generic wrapper to an already deployed contract.
func bindPaymentProcessor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PaymentProcessorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentProcessor *PaymentProcessorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentProcessor.Contract.PaymentProcessorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentProcessor *PaymentProcessorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.PaymentProcessorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentProcessor *PaymentProcessorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.PaymentProcessorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PaymentProcessor *PaymentProcessorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PaymentProcessor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PaymentProcessor *PaymentProcessorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PaymentProcessor *PaymentProcessorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.contract.Transact(opts, method, params...)
}

// DefaultPlatformFeeBps is a free data retrieval call binding the contract method 0xed3b7e9f.
//
// Solidity: function defaultPlatformFeeBps() view returns(uint256)
func (_PaymentProcessor *PaymentProcessorCaller) DefaultPlatformFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "defaultPlatformFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DefaultPlatformFeeBps is a free data retrieval call binding the contract method 0xed3b7e9f.
//
// Solidity: function defaultPlatformFeeBps() view returns(uint256)
func (_PaymentProcessor *PaymentProcessorSession) DefaultPlatformFeeBps() (*big.Int, error) {
	return _PaymentProcessor.Contract.DefaultPlatformFeeBps(&_PaymentProcessor.CallOpts)
}

// DefaultPlatformFeeBps is a free data retrieval call binding the contract method 0xed3b7e9f.
//
// Solidity: function defaultPlatformFeeBps() view returns(uint256)
func (_PaymentProcessor *PaymentProcessorCallerSession) DefaultPlatformFeeBps() (*big.Int, error) {
	return _PaymentProcessor.Contract.DefaultPlatformFeeBps(&_PaymentProcessor.CallOpts)
}

// EmergencyWithdrawalEnabled is a free data retrieval call binding the contract method 0x7fe4e8ee.
//
// Solidity: function emergencyWithdrawalEnabled() view returns(bool)
func (_PaymentProcessor *PaymentProcessorCaller) EmergencyWithdrawalEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "emergencyWithdrawalEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// EmergencyWithdrawalEnabled is a free data retrieval call binding the contract method 0x7fe4e8ee.
//
// Solidity: function emergencyWithdrawalEnabled() view returns(bool)
func (_PaymentProcessor *PaymentProcessorSession) EmergencyWithdrawalEnabled() (bool, error) {
	return _PaymentProcessor.Contract.EmergencyWithdrawalEnabled(&_PaymentProcessor.CallOpts)
}

// EmergencyWithdrawalEnabled is a free data retrieval call binding the contract method 0x7fe4e8ee.
//
// Solidity: function emergencyWithdrawalEnabled() view returns(bool)
func (_PaymentProcessor *PaymentProcessorCallerSession) EmergencyWithdrawalEnabled() (bool, error) {
	return _PaymentProcessor.Contract.EmergencyWithdrawalEnabled(&_PaymentProcessor.CallOpts)
}

// GetContractTokenBalance is a free data retrieval call binding the contract method 0x14205e28.
//
// Solidity: function getContractTokenBalance(address token) view returns(uint256)
func (_PaymentProcessor *PaymentProcessorCaller) GetContractTokenBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "getContractTokenBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractTokenBalance is a free data retrieval call binding the contract method 0x14205e28.
//
// Solidity: function getContractTokenBalance(address token) view returns(uint256)
func (_PaymentProcessor *PaymentProcessorSession) GetContractTokenBalance(token common.Address) (*big.Int, error) {
	return _PaymentProcessor.Contract.GetContractTokenBalance(&_PaymentProcessor.CallOpts, token)
}

// GetContractTokenBalance is a free data retrieval call binding the contract method 0x14205e28.
//
// Solidity: function getContractTokenBalance(address token) view returns(uint256)
func (_PaymentProcessor *PaymentProcessorCallerSession) GetContractTokenBalance(token common.Address) (*big.Int, error) {
	return _PaymentProcessor.Contract.GetContractTokenBalance(&_PaymentProcessor.CallOpts, token)
}

// GetMerchantTokenBalance is a free data retrieval call binding the contract method 0xf3d94b8e.
//
// Solidity: function getMerchantTokenBalance(address merchant, address token) view returns(uint256)
func (_PaymentProcessor *PaymentProcessorCaller) GetMerchantTokenBalance(opts *bind.CallOpts, merchant common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "getMerchantTokenBalance", merchant, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMerchantTokenBalance is a free data retrieval call binding the contract method 0xf3d94b8e.
//
// Solidity: function getMerchantTokenBalance(address merchant, address token) view returns(uint256)
func (_PaymentProcessor *PaymentProcessorSession) GetMerchantTokenBalance(merchant common.Address, token common.Address) (*big.Int, error) {
	return _PaymentProcessor.Contract.GetMerchantTokenBalance(&_PaymentProcessor.CallOpts, merchant, token)
}

// GetMerchantTokenBalance is a free data retrieval call binding the contract method 0xf3d94b8e.
//
// Solidity: function getMerchantTokenBalance(address merchant, address token) view returns(uint256)
func (_PaymentProcessor *PaymentProcessorCallerSession) GetMerchantTokenBalance(merchant common.Address, token common.Address) (*big.Int, error) {
	return _PaymentProcessor.Contract.GetMerchantTokenBalance(&_PaymentProcessor.CallOpts, merchant, token)
}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 orderId) view returns((address,address,bytes32,address,uint256,bool,uint96,uint256,string,uint8))
func (_PaymentProcessor *PaymentProcessorCaller) GetOrder(opts *bind.CallOpts, orderId [32]byte) (IPaymentProcessorOrder, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "getOrder", orderId)

	if err != nil {
		return *new(IPaymentProcessorOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(IPaymentProcessorOrder)).(*IPaymentProcessorOrder)

	return out0, err

}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 orderId) view returns((address,address,bytes32,address,uint256,bool,uint96,uint256,string,uint8))
func (_PaymentProcessor *PaymentProcessorSession) GetOrder(orderId [32]byte) (IPaymentProcessorOrder, error) {
	return _PaymentProcessor.Contract.GetOrder(&_PaymentProcessor.CallOpts, orderId)
}

// GetOrder is a free data retrieval call binding the contract method 0x5778472a.
//
// Solidity: function getOrder(bytes32 orderId) view returns((address,address,bytes32,address,uint256,bool,uint96,uint256,string,uint8))
func (_PaymentProcessor *PaymentProcessorCallerSession) GetOrder(orderId [32]byte) (IPaymentProcessorOrder, error) {
	return _PaymentProcessor.Contract.GetOrder(&_PaymentProcessor.CallOpts, orderId)
}

// GetOrderCreatedAt is a free data retrieval call binding the contract method 0xbc476cda.
//
// Solidity: function getOrderCreatedAt(bytes32 orderId) view returns(uint256 createdAt)
func (_PaymentProcessor *PaymentProcessorCaller) GetOrderCreatedAt(opts *bind.CallOpts, orderId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "getOrderCreatedAt", orderId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderCreatedAt is a free data retrieval call binding the contract method 0xbc476cda.
//
// Solidity: function getOrderCreatedAt(bytes32 orderId) view returns(uint256 createdAt)
func (_PaymentProcessor *PaymentProcessorSession) GetOrderCreatedAt(orderId [32]byte) (*big.Int, error) {
	return _PaymentProcessor.Contract.GetOrderCreatedAt(&_PaymentProcessor.CallOpts, orderId)
}

// GetOrderCreatedAt is a free data retrieval call binding the contract method 0xbc476cda.
//
// Solidity: function getOrderCreatedAt(bytes32 orderId) view returns(uint256 createdAt)
func (_PaymentProcessor *PaymentProcessorCallerSession) GetOrderCreatedAt(orderId [32]byte) (*big.Int, error) {
	return _PaymentProcessor.Contract.GetOrderCreatedAt(&_PaymentProcessor.CallOpts, orderId)
}

// GetOrderRemainingTime is a free data retrieval call binding the contract method 0x01904028.
//
// Solidity: function getOrderRemainingTime(bytes32 orderId) view returns(uint256 remainingTime)
func (_PaymentProcessor *PaymentProcessorCaller) GetOrderRemainingTime(opts *bind.CallOpts, orderId [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "getOrderRemainingTime", orderId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOrderRemainingTime is a free data retrieval call binding the contract method 0x01904028.
//
// Solidity: function getOrderRemainingTime(bytes32 orderId) view returns(uint256 remainingTime)
func (_PaymentProcessor *PaymentProcessorSession) GetOrderRemainingTime(orderId [32]byte) (*big.Int, error) {
	return _PaymentProcessor.Contract.GetOrderRemainingTime(&_PaymentProcessor.CallOpts, orderId)
}

// GetOrderRemainingTime is a free data retrieval call binding the contract method 0x01904028.
//
// Solidity: function getOrderRemainingTime(bytes32 orderId) view returns(uint256 remainingTime)
func (_PaymentProcessor *PaymentProcessorCallerSession) GetOrderRemainingTime(orderId [32]byte) (*big.Int, error) {
	return _PaymentProcessor.Contract.GetOrderRemainingTime(&_PaymentProcessor.CallOpts, orderId)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderId) view returns(uint8)
func (_PaymentProcessor *PaymentProcessorCaller) GetOrderStatus(opts *bind.CallOpts, orderId [32]byte) (uint8, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "getOrderStatus", orderId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderId) view returns(uint8)
func (_PaymentProcessor *PaymentProcessorSession) GetOrderStatus(orderId [32]byte) (uint8, error) {
	return _PaymentProcessor.Contract.GetOrderStatus(&_PaymentProcessor.CallOpts, orderId)
}

// GetOrderStatus is a free data retrieval call binding the contract method 0x46423aa7.
//
// Solidity: function getOrderStatus(bytes32 orderId) view returns(uint8)
func (_PaymentProcessor *PaymentProcessorCallerSession) GetOrderStatus(orderId [32]byte) (uint8, error) {
	return _PaymentProcessor.Contract.GetOrderStatus(&_PaymentProcessor.CallOpts, orderId)
}

// GetPlatformTokenBalance is a free data retrieval call binding the contract method 0xe9386b0b.
//
// Solidity: function getPlatformTokenBalance(address token) view returns(uint256)
func (_PaymentProcessor *PaymentProcessorCaller) GetPlatformTokenBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "getPlatformTokenBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPlatformTokenBalance is a free data retrieval call binding the contract method 0xe9386b0b.
//
// Solidity: function getPlatformTokenBalance(address token) view returns(uint256)
func (_PaymentProcessor *PaymentProcessorSession) GetPlatformTokenBalance(token common.Address) (*big.Int, error) {
	return _PaymentProcessor.Contract.GetPlatformTokenBalance(&_PaymentProcessor.CallOpts, token)
}

// GetPlatformTokenBalance is a free data retrieval call binding the contract method 0xe9386b0b.
//
// Solidity: function getPlatformTokenBalance(address token) view returns(uint256)
func (_PaymentProcessor *PaymentProcessorCallerSession) GetPlatformTokenBalance(token common.Address) (*big.Int, error) {
	return _PaymentProcessor.Contract.GetPlatformTokenBalance(&_PaymentProcessor.CallOpts, token)
}

// GetPlatformWallet is a free data retrieval call binding the contract method 0xd76a09f4.
//
// Solidity: function getPlatformWallet() view returns(address)
func (_PaymentProcessor *PaymentProcessorCaller) GetPlatformWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "getPlatformWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPlatformWallet is a free data retrieval call binding the contract method 0xd76a09f4.
//
// Solidity: function getPlatformWallet() view returns(address)
func (_PaymentProcessor *PaymentProcessorSession) GetPlatformWallet() (common.Address, error) {
	return _PaymentProcessor.Contract.GetPlatformWallet(&_PaymentProcessor.CallOpts)
}

// GetPlatformWallet is a free data retrieval call binding the contract method 0xd76a09f4.
//
// Solidity: function getPlatformWallet() view returns(address)
func (_PaymentProcessor *PaymentProcessorCallerSession) GetPlatformWallet() (common.Address, error) {
	return _PaymentProcessor.Contract.GetPlatformWallet(&_PaymentProcessor.CallOpts)
}

// IsOrderExpired is a free data retrieval call binding the contract method 0xb8293da7.
//
// Solidity: function isOrderExpired(bytes32 orderId) view returns(bool expired)
func (_PaymentProcessor *PaymentProcessorCaller) IsOrderExpired(opts *bind.CallOpts, orderId [32]byte) (bool, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "isOrderExpired", orderId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrderExpired is a free data retrieval call binding the contract method 0xb8293da7.
//
// Solidity: function isOrderExpired(bytes32 orderId) view returns(bool expired)
func (_PaymentProcessor *PaymentProcessorSession) IsOrderExpired(orderId [32]byte) (bool, error) {
	return _PaymentProcessor.Contract.IsOrderExpired(&_PaymentProcessor.CallOpts, orderId)
}

// IsOrderExpired is a free data retrieval call binding the contract method 0xb8293da7.
//
// Solidity: function isOrderExpired(bytes32 orderId) view returns(bool expired)
func (_PaymentProcessor *PaymentProcessorCallerSession) IsOrderExpired(orderId [32]byte) (bool, error) {
	return _PaymentProcessor.Contract.IsOrderExpired(&_PaymentProcessor.CallOpts, orderId)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address _token) view returns(bool)
func (_PaymentProcessor *PaymentProcessorCaller) IsTokenSupported(opts *bind.CallOpts, _token common.Address) (bool, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "isTokenSupported", _token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address _token) view returns(bool)
func (_PaymentProcessor *PaymentProcessorSession) IsTokenSupported(_token common.Address) (bool, error) {
	return _PaymentProcessor.Contract.IsTokenSupported(&_PaymentProcessor.CallOpts, _token)
}

// IsTokenSupported is a free data retrieval call binding the contract method 0x75151b63.
//
// Solidity: function isTokenSupported(address _token) view returns(bool)
func (_PaymentProcessor *PaymentProcessorCallerSession) IsTokenSupported(_token common.Address) (bool, error) {
	return _PaymentProcessor.Contract.IsTokenSupported(&_PaymentProcessor.CallOpts, _token)
}

// MerchantRegistry is a free data retrieval call binding the contract method 0x6eca5042.
//
// Solidity: function merchantRegistry() view returns(address)
func (_PaymentProcessor *PaymentProcessorCaller) MerchantRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "merchantRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MerchantRegistry is a free data retrieval call binding the contract method 0x6eca5042.
//
// Solidity: function merchantRegistry() view returns(address)
func (_PaymentProcessor *PaymentProcessorSession) MerchantRegistry() (common.Address, error) {
	return _PaymentProcessor.Contract.MerchantRegistry(&_PaymentProcessor.CallOpts)
}

// MerchantRegistry is a free data retrieval call binding the contract method 0x6eca5042.
//
// Solidity: function merchantRegistry() view returns(address)
func (_PaymentProcessor *PaymentProcessorCallerSession) MerchantRegistry() (common.Address, error) {
	return _PaymentProcessor.Contract.MerchantRegistry(&_PaymentProcessor.CallOpts)
}

// OrderExpirationTime is a free data retrieval call binding the contract method 0xe6d05b93.
//
// Solidity: function orderExpirationTime() view returns(uint256)
func (_PaymentProcessor *PaymentProcessorCaller) OrderExpirationTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "orderExpirationTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderExpirationTime is a free data retrieval call binding the contract method 0xe6d05b93.
//
// Solidity: function orderExpirationTime() view returns(uint256)
func (_PaymentProcessor *PaymentProcessorSession) OrderExpirationTime() (*big.Int, error) {
	return _PaymentProcessor.Contract.OrderExpirationTime(&_PaymentProcessor.CallOpts)
}

// OrderExpirationTime is a free data retrieval call binding the contract method 0xe6d05b93.
//
// Solidity: function orderExpirationTime() view returns(uint256)
func (_PaymentProcessor *PaymentProcessorCallerSession) OrderExpirationTime() (*big.Int, error) {
	return _PaymentProcessor.Contract.OrderExpirationTime(&_PaymentProcessor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentProcessor *PaymentProcessorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentProcessor *PaymentProcessorSession) Owner() (common.Address, error) {
	return _PaymentProcessor.Contract.Owner(&_PaymentProcessor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PaymentProcessor *PaymentProcessorCallerSession) Owner() (common.Address, error) {
	return _PaymentProcessor.Contract.Owner(&_PaymentProcessor.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PaymentProcessor *PaymentProcessorCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PaymentProcessor *PaymentProcessorSession) Paused() (bool, error) {
	return _PaymentProcessor.Contract.Paused(&_PaymentProcessor.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_PaymentProcessor *PaymentProcessorCallerSession) Paused() (bool, error) {
	return _PaymentProcessor.Contract.Paused(&_PaymentProcessor.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PaymentProcessor *PaymentProcessorCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PaymentProcessor.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PaymentProcessor *PaymentProcessorSession) PendingOwner() (common.Address, error) {
	return _PaymentProcessor.Contract.PendingOwner(&_PaymentProcessor.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_PaymentProcessor *PaymentProcessorCallerSession) PendingOwner() (common.Address, error) {
	return _PaymentProcessor.Contract.PendingOwner(&_PaymentProcessor.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PaymentProcessor *PaymentProcessorTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PaymentProcessor *PaymentProcessorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PaymentProcessor.Contract.AcceptOwnership(&_PaymentProcessor.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_PaymentProcessor *PaymentProcessorTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _PaymentProcessor.Contract.AcceptOwnership(&_PaymentProcessor.TransactOpts)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 _orderId) returns(bool)
func (_PaymentProcessor *PaymentProcessorTransactor) CancelOrder(opts *bind.TransactOpts, _orderId [32]byte) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "cancelOrder", _orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 _orderId) returns(bool)
func (_PaymentProcessor *PaymentProcessorSession) CancelOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.CancelOrder(&_PaymentProcessor.TransactOpts, _orderId)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x7489ec23.
//
// Solidity: function cancelOrder(bytes32 _orderId) returns(bool)
func (_PaymentProcessor *PaymentProcessorTransactorSession) CancelOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.CancelOrder(&_PaymentProcessor.TransactOpts, _orderId)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x6beccaf8.
//
// Solidity: function createOrder(bytes32 _merchantId, address _token, uint256 _amount, string _metadataUri) returns(bytes32 _orderId)
func (_PaymentProcessor *PaymentProcessorTransactor) CreateOrder(opts *bind.TransactOpts, _merchantId [32]byte, _token common.Address, _amount *big.Int, _metadataUri string) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "createOrder", _merchantId, _token, _amount, _metadataUri)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x6beccaf8.
//
// Solidity: function createOrder(bytes32 _merchantId, address _token, uint256 _amount, string _metadataUri) returns(bytes32 _orderId)
func (_PaymentProcessor *PaymentProcessorSession) CreateOrder(_merchantId [32]byte, _token common.Address, _amount *big.Int, _metadataUri string) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.CreateOrder(&_PaymentProcessor.TransactOpts, _merchantId, _token, _amount, _metadataUri)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x6beccaf8.
//
// Solidity: function createOrder(bytes32 _merchantId, address _token, uint256 _amount, string _metadataUri) returns(bytes32 _orderId)
func (_PaymentProcessor *PaymentProcessorTransactorSession) CreateOrder(_merchantId [32]byte, _token common.Address, _amount *big.Int, _metadataUri string) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.CreateOrder(&_PaymentProcessor.TransactOpts, _merchantId, _token, _amount, _metadataUri)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xe63ea408.
//
// Solidity: function emergencyWithdraw(address token, address to, uint256 amount) returns()
func (_PaymentProcessor *PaymentProcessorTransactor) EmergencyWithdraw(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "emergencyWithdraw", token, to, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xe63ea408.
//
// Solidity: function emergencyWithdraw(address token, address to, uint256 amount) returns()
func (_PaymentProcessor *PaymentProcessorSession) EmergencyWithdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.EmergencyWithdraw(&_PaymentProcessor.TransactOpts, token, to, amount)
}

// EmergencyWithdraw is a paid mutator transaction binding the contract method 0xe63ea408.
//
// Solidity: function emergencyWithdraw(address token, address to, uint256 amount) returns()
func (_PaymentProcessor *PaymentProcessorTransactorSession) EmergencyWithdraw(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.EmergencyWithdraw(&_PaymentProcessor.TransactOpts, token, to, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xe251d7be.
//
// Solidity: function initialize(address _platformWallet, uint256 _defaultPlatformFeeBps, address _merchantRegistry, uint256 _orderExpirationTime, address initialOwner) returns()
func (_PaymentProcessor *PaymentProcessorTransactor) Initialize(opts *bind.TransactOpts, _platformWallet common.Address, _defaultPlatformFeeBps *big.Int, _merchantRegistry common.Address, _orderExpirationTime *big.Int, initialOwner common.Address) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "initialize", _platformWallet, _defaultPlatformFeeBps, _merchantRegistry, _orderExpirationTime, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xe251d7be.
//
// Solidity: function initialize(address _platformWallet, uint256 _defaultPlatformFeeBps, address _merchantRegistry, uint256 _orderExpirationTime, address initialOwner) returns()
func (_PaymentProcessor *PaymentProcessorSession) Initialize(_platformWallet common.Address, _defaultPlatformFeeBps *big.Int, _merchantRegistry common.Address, _orderExpirationTime *big.Int, initialOwner common.Address) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.Initialize(&_PaymentProcessor.TransactOpts, _platformWallet, _defaultPlatformFeeBps, _merchantRegistry, _orderExpirationTime, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xe251d7be.
//
// Solidity: function initialize(address _platformWallet, uint256 _defaultPlatformFeeBps, address _merchantRegistry, uint256 _orderExpirationTime, address initialOwner) returns()
func (_PaymentProcessor *PaymentProcessorTransactorSession) Initialize(_platformWallet common.Address, _defaultPlatformFeeBps *big.Int, _merchantRegistry common.Address, _orderExpirationTime *big.Int, initialOwner common.Address) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.Initialize(&_PaymentProcessor.TransactOpts, _platformWallet, _defaultPlatformFeeBps, _merchantRegistry, _orderExpirationTime, initialOwner)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PaymentProcessor *PaymentProcessorTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PaymentProcessor *PaymentProcessorSession) Pause() (*types.Transaction, error) {
	return _PaymentProcessor.Contract.Pause(&_PaymentProcessor.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PaymentProcessor *PaymentProcessorTransactorSession) Pause() (*types.Transaction, error) {
	return _PaymentProcessor.Contract.Pause(&_PaymentProcessor.TransactOpts)
}

// PayOrder is a paid mutator transaction binding the contract method 0x58504a36.
//
// Solidity: function payOrder(bytes32 _orderId) returns(bool)
func (_PaymentProcessor *PaymentProcessorTransactor) PayOrder(opts *bind.TransactOpts, _orderId [32]byte) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "payOrder", _orderId)
}

// PayOrder is a paid mutator transaction binding the contract method 0x58504a36.
//
// Solidity: function payOrder(bytes32 _orderId) returns(bool)
func (_PaymentProcessor *PaymentProcessorSession) PayOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.PayOrder(&_PaymentProcessor.TransactOpts, _orderId)
}

// PayOrder is a paid mutator transaction binding the contract method 0x58504a36.
//
// Solidity: function payOrder(bytes32 _orderId) returns(bool)
func (_PaymentProcessor *PaymentProcessorTransactorSession) PayOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.PayOrder(&_PaymentProcessor.TransactOpts, _orderId)
}

// RefundOrder is a paid mutator transaction binding the contract method 0x116a1d9a.
//
// Solidity: function refundOrder(bytes32 _orderId) returns(bool)
func (_PaymentProcessor *PaymentProcessorTransactor) RefundOrder(opts *bind.TransactOpts, _orderId [32]byte) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "refundOrder", _orderId)
}

// RefundOrder is a paid mutator transaction binding the contract method 0x116a1d9a.
//
// Solidity: function refundOrder(bytes32 _orderId) returns(bool)
func (_PaymentProcessor *PaymentProcessorSession) RefundOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.RefundOrder(&_PaymentProcessor.TransactOpts, _orderId)
}

// RefundOrder is a paid mutator transaction binding the contract method 0x116a1d9a.
//
// Solidity: function refundOrder(bytes32 _orderId) returns(bool)
func (_PaymentProcessor *PaymentProcessorTransactorSession) RefundOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.RefundOrder(&_PaymentProcessor.TransactOpts, _orderId)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentProcessor *PaymentProcessorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentProcessor *PaymentProcessorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PaymentProcessor.Contract.RenounceOwnership(&_PaymentProcessor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PaymentProcessor *PaymentProcessorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PaymentProcessor.Contract.RenounceOwnership(&_PaymentProcessor.TransactOpts)
}

// SetEmergencyWithdrawalEnabled is a paid mutator transaction binding the contract method 0x0a66bb45.
//
// Solidity: function setEmergencyWithdrawalEnabled(bool enabled) returns()
func (_PaymentProcessor *PaymentProcessorTransactor) SetEmergencyWithdrawalEnabled(opts *bind.TransactOpts, enabled bool) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "setEmergencyWithdrawalEnabled", enabled)
}

// SetEmergencyWithdrawalEnabled is a paid mutator transaction binding the contract method 0x0a66bb45.
//
// Solidity: function setEmergencyWithdrawalEnabled(bool enabled) returns()
func (_PaymentProcessor *PaymentProcessorSession) SetEmergencyWithdrawalEnabled(enabled bool) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.SetEmergencyWithdrawalEnabled(&_PaymentProcessor.TransactOpts, enabled)
}

// SetEmergencyWithdrawalEnabled is a paid mutator transaction binding the contract method 0x0a66bb45.
//
// Solidity: function setEmergencyWithdrawalEnabled(bool enabled) returns()
func (_PaymentProcessor *PaymentProcessorTransactorSession) SetEmergencyWithdrawalEnabled(enabled bool) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.SetEmergencyWithdrawalEnabled(&_PaymentProcessor.TransactOpts, enabled)
}

// SetTokenSupport is a paid mutator transaction binding the contract method 0xf89248a4.
//
// Solidity: function setTokenSupport(address token, uint256 status) returns()
func (_PaymentProcessor *PaymentProcessorTransactor) SetTokenSupport(opts *bind.TransactOpts, token common.Address, status *big.Int) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "setTokenSupport", token, status)
}

// SetTokenSupport is a paid mutator transaction binding the contract method 0xf89248a4.
//
// Solidity: function setTokenSupport(address token, uint256 status) returns()
func (_PaymentProcessor *PaymentProcessorSession) SetTokenSupport(token common.Address, status *big.Int) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.SetTokenSupport(&_PaymentProcessor.TransactOpts, token, status)
}

// SetTokenSupport is a paid mutator transaction binding the contract method 0xf89248a4.
//
// Solidity: function setTokenSupport(address token, uint256 status) returns()
func (_PaymentProcessor *PaymentProcessorTransactorSession) SetTokenSupport(token common.Address, status *big.Int) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.SetTokenSupport(&_PaymentProcessor.TransactOpts, token, status)
}

// SettleOrder is a paid mutator transaction binding the contract method 0x49085d8c.
//
// Solidity: function settleOrder(bytes32 _orderId) returns(bool)
func (_PaymentProcessor *PaymentProcessorTransactor) SettleOrder(opts *bind.TransactOpts, _orderId [32]byte) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "settleOrder", _orderId)
}

// SettleOrder is a paid mutator transaction binding the contract method 0x49085d8c.
//
// Solidity: function settleOrder(bytes32 _orderId) returns(bool)
func (_PaymentProcessor *PaymentProcessorSession) SettleOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.SettleOrder(&_PaymentProcessor.TransactOpts, _orderId)
}

// SettleOrder is a paid mutator transaction binding the contract method 0x49085d8c.
//
// Solidity: function settleOrder(bytes32 _orderId) returns(bool)
func (_PaymentProcessor *PaymentProcessorTransactorSession) SettleOrder(_orderId [32]byte) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.SettleOrder(&_PaymentProcessor.TransactOpts, _orderId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaymentProcessor *PaymentProcessorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaymentProcessor *PaymentProcessorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.TransferOwnership(&_PaymentProcessor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PaymentProcessor *PaymentProcessorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.TransferOwnership(&_PaymentProcessor.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PaymentProcessor *PaymentProcessorTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PaymentProcessor *PaymentProcessorSession) Unpause() (*types.Transaction, error) {
	return _PaymentProcessor.Contract.Unpause(&_PaymentProcessor.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PaymentProcessor *PaymentProcessorTransactorSession) Unpause() (*types.Transaction, error) {
	return _PaymentProcessor.Contract.Unpause(&_PaymentProcessor.TransactOpts)
}

// UpdateMerchantRegistry is a paid mutator transaction binding the contract method 0x85625ac3.
//
// Solidity: function updateMerchantRegistry(address newRegistry) returns()
func (_PaymentProcessor *PaymentProcessorTransactor) UpdateMerchantRegistry(opts *bind.TransactOpts, newRegistry common.Address) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "updateMerchantRegistry", newRegistry)
}

// UpdateMerchantRegistry is a paid mutator transaction binding the contract method 0x85625ac3.
//
// Solidity: function updateMerchantRegistry(address newRegistry) returns()
func (_PaymentProcessor *PaymentProcessorSession) UpdateMerchantRegistry(newRegistry common.Address) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.UpdateMerchantRegistry(&_PaymentProcessor.TransactOpts, newRegistry)
}

// UpdateMerchantRegistry is a paid mutator transaction binding the contract method 0x85625ac3.
//
// Solidity: function updateMerchantRegistry(address newRegistry) returns()
func (_PaymentProcessor *PaymentProcessorTransactorSession) UpdateMerchantRegistry(newRegistry common.Address) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.UpdateMerchantRegistry(&_PaymentProcessor.TransactOpts, newRegistry)
}

// UpdateOrderExpirationTime is a paid mutator transaction binding the contract method 0xb2721bd7.
//
// Solidity: function updateOrderExpirationTime(uint256 newExpirationTime) returns()
func (_PaymentProcessor *PaymentProcessorTransactor) UpdateOrderExpirationTime(opts *bind.TransactOpts, newExpirationTime *big.Int) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "updateOrderExpirationTime", newExpirationTime)
}

// UpdateOrderExpirationTime is a paid mutator transaction binding the contract method 0xb2721bd7.
//
// Solidity: function updateOrderExpirationTime(uint256 newExpirationTime) returns()
func (_PaymentProcessor *PaymentProcessorSession) UpdateOrderExpirationTime(newExpirationTime *big.Int) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.UpdateOrderExpirationTime(&_PaymentProcessor.TransactOpts, newExpirationTime)
}

// UpdateOrderExpirationTime is a paid mutator transaction binding the contract method 0xb2721bd7.
//
// Solidity: function updateOrderExpirationTime(uint256 newExpirationTime) returns()
func (_PaymentProcessor *PaymentProcessorTransactorSession) UpdateOrderExpirationTime(newExpirationTime *big.Int) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.UpdateOrderExpirationTime(&_PaymentProcessor.TransactOpts, newExpirationTime)
}

// UpdateProtocolAddress is a paid mutator transaction binding the contract method 0x40ebc677.
//
// Solidity: function updateProtocolAddress(bytes32 what, address value) returns()
func (_PaymentProcessor *PaymentProcessorTransactor) UpdateProtocolAddress(opts *bind.TransactOpts, what [32]byte, value common.Address) (*types.Transaction, error) {
	return _PaymentProcessor.contract.Transact(opts, "updateProtocolAddress", what, value)
}

// UpdateProtocolAddress is a paid mutator transaction binding the contract method 0x40ebc677.
//
// Solidity: function updateProtocolAddress(bytes32 what, address value) returns()
func (_PaymentProcessor *PaymentProcessorSession) UpdateProtocolAddress(what [32]byte, value common.Address) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.UpdateProtocolAddress(&_PaymentProcessor.TransactOpts, what, value)
}

// UpdateProtocolAddress is a paid mutator transaction binding the contract method 0x40ebc677.
//
// Solidity: function updateProtocolAddress(bytes32 what, address value) returns()
func (_PaymentProcessor *PaymentProcessorTransactorSession) UpdateProtocolAddress(what [32]byte, value common.Address) (*types.Transaction, error) {
	return _PaymentProcessor.Contract.UpdateProtocolAddress(&_PaymentProcessor.TransactOpts, what, value)
}

// PaymentProcessorEmergencyWithdrawalEnabledUpdatedIterator is returned from FilterEmergencyWithdrawalEnabledUpdated and is used to iterate over the raw logs and unpacked data for EmergencyWithdrawalEnabledUpdated events raised by the PaymentProcessor contract.
type PaymentProcessorEmergencyWithdrawalEnabledUpdatedIterator struct {
	Event *PaymentProcessorEmergencyWithdrawalEnabledUpdated // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorEmergencyWithdrawalEnabledUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorEmergencyWithdrawalEnabledUpdated)
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
		it.Event = new(PaymentProcessorEmergencyWithdrawalEnabledUpdated)
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
func (it *PaymentProcessorEmergencyWithdrawalEnabledUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorEmergencyWithdrawalEnabledUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorEmergencyWithdrawalEnabledUpdated represents a EmergencyWithdrawalEnabledUpdated event raised by the PaymentProcessor contract.
type PaymentProcessorEmergencyWithdrawalEnabledUpdated struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdrawalEnabledUpdated is a free log retrieval operation binding the contract event 0xd4f0990d03e7ad5243d0a8589d4a52df340d2bb32a2547814d93b80674675565.
//
// Solidity: event EmergencyWithdrawalEnabledUpdated(bool enabled)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterEmergencyWithdrawalEnabledUpdated(opts *bind.FilterOpts) (*PaymentProcessorEmergencyWithdrawalEnabledUpdatedIterator, error) {

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "EmergencyWithdrawalEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorEmergencyWithdrawalEnabledUpdatedIterator{contract: _PaymentProcessor.contract, event: "EmergencyWithdrawalEnabledUpdated", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdrawalEnabledUpdated is a free log subscription operation binding the contract event 0xd4f0990d03e7ad5243d0a8589d4a52df340d2bb32a2547814d93b80674675565.
//
// Solidity: event EmergencyWithdrawalEnabledUpdated(bool enabled)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchEmergencyWithdrawalEnabledUpdated(opts *bind.WatchOpts, sink chan<- *PaymentProcessorEmergencyWithdrawalEnabledUpdated) (event.Subscription, error) {

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "EmergencyWithdrawalEnabledUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorEmergencyWithdrawalEnabledUpdated)
				if err := _PaymentProcessor.contract.UnpackLog(event, "EmergencyWithdrawalEnabledUpdated", log); err != nil {
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

// ParseEmergencyWithdrawalEnabledUpdated is a log parse operation binding the contract event 0xd4f0990d03e7ad5243d0a8589d4a52df340d2bb32a2547814d93b80674675565.
//
// Solidity: event EmergencyWithdrawalEnabledUpdated(bool enabled)
func (_PaymentProcessor *PaymentProcessorFilterer) ParseEmergencyWithdrawalEnabledUpdated(log types.Log) (*PaymentProcessorEmergencyWithdrawalEnabledUpdated, error) {
	event := new(PaymentProcessorEmergencyWithdrawalEnabledUpdated)
	if err := _PaymentProcessor.contract.UnpackLog(event, "EmergencyWithdrawalEnabledUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorEmergencyWithdrawalSuccessIterator is returned from FilterEmergencyWithdrawalSuccess and is used to iterate over the raw logs and unpacked data for EmergencyWithdrawalSuccess events raised by the PaymentProcessor contract.
type PaymentProcessorEmergencyWithdrawalSuccessIterator struct {
	Event *PaymentProcessorEmergencyWithdrawalSuccess // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorEmergencyWithdrawalSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorEmergencyWithdrawalSuccess)
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
		it.Event = new(PaymentProcessorEmergencyWithdrawalSuccess)
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
func (it *PaymentProcessorEmergencyWithdrawalSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorEmergencyWithdrawalSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorEmergencyWithdrawalSuccess represents a EmergencyWithdrawalSuccess event raised by the PaymentProcessor contract.
type PaymentProcessorEmergencyWithdrawalSuccess struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEmergencyWithdrawalSuccess is a free log retrieval operation binding the contract event 0x5814d775ec31b0c5da55427172ebb946182e8686a551ad58c3727593119c4e3d.
//
// Solidity: event EmergencyWithdrawalSuccess(address indexed token, address indexed to, uint256 amount)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterEmergencyWithdrawalSuccess(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*PaymentProcessorEmergencyWithdrawalSuccessIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "EmergencyWithdrawalSuccess", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorEmergencyWithdrawalSuccessIterator{contract: _PaymentProcessor.contract, event: "EmergencyWithdrawalSuccess", logs: logs, sub: sub}, nil
}

// WatchEmergencyWithdrawalSuccess is a free log subscription operation binding the contract event 0x5814d775ec31b0c5da55427172ebb946182e8686a551ad58c3727593119c4e3d.
//
// Solidity: event EmergencyWithdrawalSuccess(address indexed token, address indexed to, uint256 amount)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchEmergencyWithdrawalSuccess(opts *bind.WatchOpts, sink chan<- *PaymentProcessorEmergencyWithdrawalSuccess, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "EmergencyWithdrawalSuccess", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorEmergencyWithdrawalSuccess)
				if err := _PaymentProcessor.contract.UnpackLog(event, "EmergencyWithdrawalSuccess", log); err != nil {
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

// ParseEmergencyWithdrawalSuccess is a log parse operation binding the contract event 0x5814d775ec31b0c5da55427172ebb946182e8686a551ad58c3727593119c4e3d.
//
// Solidity: event EmergencyWithdrawalSuccess(address indexed token, address indexed to, uint256 amount)
func (_PaymentProcessor *PaymentProcessorFilterer) ParseEmergencyWithdrawalSuccess(log types.Log) (*PaymentProcessorEmergencyWithdrawalSuccess, error) {
	event := new(PaymentProcessorEmergencyWithdrawalSuccess)
	if err := _PaymentProcessor.contract.UnpackLog(event, "EmergencyWithdrawalSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PaymentProcessor contract.
type PaymentProcessorInitializedIterator struct {
	Event *PaymentProcessorInitialized // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorInitialized)
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
		it.Event = new(PaymentProcessorInitialized)
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
func (it *PaymentProcessorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorInitialized represents a Initialized event raised by the PaymentProcessor contract.
type PaymentProcessorInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterInitialized(opts *bind.FilterOpts) (*PaymentProcessorInitializedIterator, error) {

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorInitializedIterator{contract: _PaymentProcessor.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PaymentProcessorInitialized) (event.Subscription, error) {

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorInitialized)
				if err := _PaymentProcessor.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PaymentProcessor *PaymentProcessorFilterer) ParseInitialized(log types.Log) (*PaymentProcessorInitialized, error) {
	event := new(PaymentProcessorInitialized)
	if err := _PaymentProcessor.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorMaxSlippageBpsUpdatedIterator is returned from FilterMaxSlippageBpsUpdated and is used to iterate over the raw logs and unpacked data for MaxSlippageBpsUpdated events raised by the PaymentProcessor contract.
type PaymentProcessorMaxSlippageBpsUpdatedIterator struct {
	Event *PaymentProcessorMaxSlippageBpsUpdated // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorMaxSlippageBpsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorMaxSlippageBpsUpdated)
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
		it.Event = new(PaymentProcessorMaxSlippageBpsUpdated)
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
func (it *PaymentProcessorMaxSlippageBpsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorMaxSlippageBpsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorMaxSlippageBpsUpdated represents a MaxSlippageBpsUpdated event raised by the PaymentProcessor contract.
type PaymentProcessorMaxSlippageBpsUpdated struct {
	OldSlippage *big.Int
	NewSlippage *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMaxSlippageBpsUpdated is a free log retrieval operation binding the contract event 0x828d2234d18d3fc572cfb754a3ec1d22b18061e6ba3d85faebd716bff3324dbf.
//
// Solidity: event MaxSlippageBpsUpdated(uint256 oldSlippage, uint256 newSlippage)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterMaxSlippageBpsUpdated(opts *bind.FilterOpts) (*PaymentProcessorMaxSlippageBpsUpdatedIterator, error) {

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "MaxSlippageBpsUpdated")
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorMaxSlippageBpsUpdatedIterator{contract: _PaymentProcessor.contract, event: "MaxSlippageBpsUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSlippageBpsUpdated is a free log subscription operation binding the contract event 0x828d2234d18d3fc572cfb754a3ec1d22b18061e6ba3d85faebd716bff3324dbf.
//
// Solidity: event MaxSlippageBpsUpdated(uint256 oldSlippage, uint256 newSlippage)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchMaxSlippageBpsUpdated(opts *bind.WatchOpts, sink chan<- *PaymentProcessorMaxSlippageBpsUpdated) (event.Subscription, error) {

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "MaxSlippageBpsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorMaxSlippageBpsUpdated)
				if err := _PaymentProcessor.contract.UnpackLog(event, "MaxSlippageBpsUpdated", log); err != nil {
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

// ParseMaxSlippageBpsUpdated is a log parse operation binding the contract event 0x828d2234d18d3fc572cfb754a3ec1d22b18061e6ba3d85faebd716bff3324dbf.
//
// Solidity: event MaxSlippageBpsUpdated(uint256 oldSlippage, uint256 newSlippage)
func (_PaymentProcessor *PaymentProcessorFilterer) ParseMaxSlippageBpsUpdated(log types.Log) (*PaymentProcessorMaxSlippageBpsUpdated, error) {
	event := new(PaymentProcessorMaxSlippageBpsUpdated)
	if err := _PaymentProcessor.contract.UnpackLog(event, "MaxSlippageBpsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorMerchantRegistryUpdatedIterator is returned from FilterMerchantRegistryUpdated and is used to iterate over the raw logs and unpacked data for MerchantRegistryUpdated events raised by the PaymentProcessor contract.
type PaymentProcessorMerchantRegistryUpdatedIterator struct {
	Event *PaymentProcessorMerchantRegistryUpdated // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorMerchantRegistryUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorMerchantRegistryUpdated)
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
		it.Event = new(PaymentProcessorMerchantRegistryUpdated)
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
func (it *PaymentProcessorMerchantRegistryUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorMerchantRegistryUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorMerchantRegistryUpdated represents a MerchantRegistryUpdated event raised by the PaymentProcessor contract.
type PaymentProcessorMerchantRegistryUpdated struct {
	OldRegistry common.Address
	NewRegistry common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMerchantRegistryUpdated is a free log retrieval operation binding the contract event 0x3c858684ec5e6e6fbc8550e3d904e0be2b41dcce856abe3f37c8f50c8865bb2a.
//
// Solidity: event MerchantRegistryUpdated(address indexed oldRegistry, address indexed newRegistry)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterMerchantRegistryUpdated(opts *bind.FilterOpts, oldRegistry []common.Address, newRegistry []common.Address) (*PaymentProcessorMerchantRegistryUpdatedIterator, error) {

	var oldRegistryRule []interface{}
	for _, oldRegistryItem := range oldRegistry {
		oldRegistryRule = append(oldRegistryRule, oldRegistryItem)
	}
	var newRegistryRule []interface{}
	for _, newRegistryItem := range newRegistry {
		newRegistryRule = append(newRegistryRule, newRegistryItem)
	}

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "MerchantRegistryUpdated", oldRegistryRule, newRegistryRule)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorMerchantRegistryUpdatedIterator{contract: _PaymentProcessor.contract, event: "MerchantRegistryUpdated", logs: logs, sub: sub}, nil
}

// WatchMerchantRegistryUpdated is a free log subscription operation binding the contract event 0x3c858684ec5e6e6fbc8550e3d904e0be2b41dcce856abe3f37c8f50c8865bb2a.
//
// Solidity: event MerchantRegistryUpdated(address indexed oldRegistry, address indexed newRegistry)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchMerchantRegistryUpdated(opts *bind.WatchOpts, sink chan<- *PaymentProcessorMerchantRegistryUpdated, oldRegistry []common.Address, newRegistry []common.Address) (event.Subscription, error) {

	var oldRegistryRule []interface{}
	for _, oldRegistryItem := range oldRegistry {
		oldRegistryRule = append(oldRegistryRule, oldRegistryItem)
	}
	var newRegistryRule []interface{}
	for _, newRegistryItem := range newRegistry {
		newRegistryRule = append(newRegistryRule, newRegistryItem)
	}

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "MerchantRegistryUpdated", oldRegistryRule, newRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorMerchantRegistryUpdated)
				if err := _PaymentProcessor.contract.UnpackLog(event, "MerchantRegistryUpdated", log); err != nil {
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

// ParseMerchantRegistryUpdated is a log parse operation binding the contract event 0x3c858684ec5e6e6fbc8550e3d904e0be2b41dcce856abe3f37c8f50c8865bb2a.
//
// Solidity: event MerchantRegistryUpdated(address indexed oldRegistry, address indexed newRegistry)
func (_PaymentProcessor *PaymentProcessorFilterer) ParseMerchantRegistryUpdated(log types.Log) (*PaymentProcessorMerchantRegistryUpdated, error) {
	event := new(PaymentProcessorMerchantRegistryUpdated)
	if err := _PaymentProcessor.contract.UnpackLog(event, "MerchantRegistryUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the PaymentProcessor contract.
type PaymentProcessorOrderCancelledIterator struct {
	Event *PaymentProcessorOrderCancelled // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorOrderCancelled)
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
		it.Event = new(PaymentProcessorOrderCancelled)
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
func (it *PaymentProcessorOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorOrderCancelled represents a OrderCancelled event raised by the PaymentProcessor contract.
type PaymentProcessorOrderCancelled struct {
	OrderId [32]byte
	Payer   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0xb2705df32ac67fc3101f496cd7036bf59074a603544d97d73650b6f09744986a.
//
// Solidity: event OrderCancelled(bytes32 indexed orderId, address indexed payer, uint256 amount)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterOrderCancelled(opts *bind.FilterOpts, orderId [][32]byte, payer []common.Address) (*PaymentProcessorOrderCancelledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "OrderCancelled", orderIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorOrderCancelledIterator{contract: _PaymentProcessor.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0xb2705df32ac67fc3101f496cd7036bf59074a603544d97d73650b6f09744986a.
//
// Solidity: event OrderCancelled(bytes32 indexed orderId, address indexed payer, uint256 amount)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *PaymentProcessorOrderCancelled, orderId [][32]byte, payer []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "OrderCancelled", orderIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorOrderCancelled)
				if err := _PaymentProcessor.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0xb2705df32ac67fc3101f496cd7036bf59074a603544d97d73650b6f09744986a.
//
// Solidity: event OrderCancelled(bytes32 indexed orderId, address indexed payer, uint256 amount)
func (_PaymentProcessor *PaymentProcessorFilterer) ParseOrderCancelled(log types.Log) (*PaymentProcessorOrderCancelled, error) {
	event := new(PaymentProcessorOrderCancelled)
	if err := _PaymentProcessor.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the PaymentProcessor contract.
type PaymentProcessorOrderCreatedIterator struct {
	Event *PaymentProcessorOrderCreated // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorOrderCreated)
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
		it.Event = new(PaymentProcessorOrderCreated)
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
func (it *PaymentProcessorOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorOrderCreated represents a OrderCreated event raised by the PaymentProcessor contract.
type PaymentProcessorOrderCreated struct {
	OrderId        [32]byte
	Payer          common.Address
	MerchantId     [32]byte
	MerchantPayout common.Address
	Token          common.Address
	Amount         *big.Int
	Status         uint8
	MetadataUri    string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x1b70765dd926c1ddb3b1ef39fe1af33547337283e270c9bc7c9c1caf8376377a.
//
// Solidity: event OrderCreated(bytes32 indexed orderId, address indexed payer, bytes32 indexed merchantId, address merchantPayout, address token, uint256 amount, uint8 status, string metadataUri)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterOrderCreated(opts *bind.FilterOpts, orderId [][32]byte, payer []common.Address, merchantId [][32]byte) (*PaymentProcessorOrderCreatedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var merchantIdRule []interface{}
	for _, merchantIdItem := range merchantId {
		merchantIdRule = append(merchantIdRule, merchantIdItem)
	}

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "OrderCreated", orderIdRule, payerRule, merchantIdRule)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorOrderCreatedIterator{contract: _PaymentProcessor.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x1b70765dd926c1ddb3b1ef39fe1af33547337283e270c9bc7c9c1caf8376377a.
//
// Solidity: event OrderCreated(bytes32 indexed orderId, address indexed payer, bytes32 indexed merchantId, address merchantPayout, address token, uint256 amount, uint8 status, string metadataUri)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *PaymentProcessorOrderCreated, orderId [][32]byte, payer []common.Address, merchantId [][32]byte) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}
	var merchantIdRule []interface{}
	for _, merchantIdItem := range merchantId {
		merchantIdRule = append(merchantIdRule, merchantIdItem)
	}

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "OrderCreated", orderIdRule, payerRule, merchantIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorOrderCreated)
				if err := _PaymentProcessor.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0x1b70765dd926c1ddb3b1ef39fe1af33547337283e270c9bc7c9c1caf8376377a.
//
// Solidity: event OrderCreated(bytes32 indexed orderId, address indexed payer, bytes32 indexed merchantId, address merchantPayout, address token, uint256 amount, uint8 status, string metadataUri)
func (_PaymentProcessor *PaymentProcessorFilterer) ParseOrderCreated(log types.Log) (*PaymentProcessorOrderCreated, error) {
	event := new(PaymentProcessorOrderCreated)
	if err := _PaymentProcessor.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorOrderExpirationTimeUpdatedIterator is returned from FilterOrderExpirationTimeUpdated and is used to iterate over the raw logs and unpacked data for OrderExpirationTimeUpdated events raised by the PaymentProcessor contract.
type PaymentProcessorOrderExpirationTimeUpdatedIterator struct {
	Event *PaymentProcessorOrderExpirationTimeUpdated // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorOrderExpirationTimeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorOrderExpirationTimeUpdated)
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
		it.Event = new(PaymentProcessorOrderExpirationTimeUpdated)
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
func (it *PaymentProcessorOrderExpirationTimeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorOrderExpirationTimeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorOrderExpirationTimeUpdated represents a OrderExpirationTimeUpdated event raised by the PaymentProcessor contract.
type PaymentProcessorOrderExpirationTimeUpdated struct {
	OldTime *big.Int
	NewTime *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderExpirationTimeUpdated is a free log retrieval operation binding the contract event 0xa26f70a3d60aec5b16495eaa74529905fa4cc5f2aa8022f9991d07eb5885da08.
//
// Solidity: event OrderExpirationTimeUpdated(uint256 oldTime, uint256 newTime)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterOrderExpirationTimeUpdated(opts *bind.FilterOpts) (*PaymentProcessorOrderExpirationTimeUpdatedIterator, error) {

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "OrderExpirationTimeUpdated")
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorOrderExpirationTimeUpdatedIterator{contract: _PaymentProcessor.contract, event: "OrderExpirationTimeUpdated", logs: logs, sub: sub}, nil
}

// WatchOrderExpirationTimeUpdated is a free log subscription operation binding the contract event 0xa26f70a3d60aec5b16495eaa74529905fa4cc5f2aa8022f9991d07eb5885da08.
//
// Solidity: event OrderExpirationTimeUpdated(uint256 oldTime, uint256 newTime)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchOrderExpirationTimeUpdated(opts *bind.WatchOpts, sink chan<- *PaymentProcessorOrderExpirationTimeUpdated) (event.Subscription, error) {

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "OrderExpirationTimeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorOrderExpirationTimeUpdated)
				if err := _PaymentProcessor.contract.UnpackLog(event, "OrderExpirationTimeUpdated", log); err != nil {
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

// ParseOrderExpirationTimeUpdated is a log parse operation binding the contract event 0xa26f70a3d60aec5b16495eaa74529905fa4cc5f2aa8022f9991d07eb5885da08.
//
// Solidity: event OrderExpirationTimeUpdated(uint256 oldTime, uint256 newTime)
func (_PaymentProcessor *PaymentProcessorFilterer) ParseOrderExpirationTimeUpdated(log types.Log) (*PaymentProcessorOrderExpirationTimeUpdated, error) {
	event := new(PaymentProcessorOrderExpirationTimeUpdated)
	if err := _PaymentProcessor.contract.UnpackLog(event, "OrderExpirationTimeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorOrderExpiredIterator is returned from FilterOrderExpired and is used to iterate over the raw logs and unpacked data for OrderExpired events raised by the PaymentProcessor contract.
type PaymentProcessorOrderExpiredIterator struct {
	Event *PaymentProcessorOrderExpired // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorOrderExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorOrderExpired)
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
		it.Event = new(PaymentProcessorOrderExpired)
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
func (it *PaymentProcessorOrderExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorOrderExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorOrderExpired represents a OrderExpired event raised by the PaymentProcessor contract.
type PaymentProcessorOrderExpired struct {
	OrderId [32]byte
	Payer   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderExpired is a free log retrieval operation binding the contract event 0xfb3d3c72741d97a417c79513cde55af3d64dc8844cc5b59f470935843eceff4c.
//
// Solidity: event OrderExpired(bytes32 indexed orderId, address indexed payer)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterOrderExpired(opts *bind.FilterOpts, orderId [][32]byte, payer []common.Address) (*PaymentProcessorOrderExpiredIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "OrderExpired", orderIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorOrderExpiredIterator{contract: _PaymentProcessor.contract, event: "OrderExpired", logs: logs, sub: sub}, nil
}

// WatchOrderExpired is a free log subscription operation binding the contract event 0xfb3d3c72741d97a417c79513cde55af3d64dc8844cc5b59f470935843eceff4c.
//
// Solidity: event OrderExpired(bytes32 indexed orderId, address indexed payer)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchOrderExpired(opts *bind.WatchOpts, sink chan<- *PaymentProcessorOrderExpired, orderId [][32]byte, payer []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "OrderExpired", orderIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorOrderExpired)
				if err := _PaymentProcessor.contract.UnpackLog(event, "OrderExpired", log); err != nil {
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

// ParseOrderExpired is a log parse operation binding the contract event 0xfb3d3c72741d97a417c79513cde55af3d64dc8844cc5b59f470935843eceff4c.
//
// Solidity: event OrderExpired(bytes32 indexed orderId, address indexed payer)
func (_PaymentProcessor *PaymentProcessorFilterer) ParseOrderExpired(log types.Log) (*PaymentProcessorOrderExpired, error) {
	event := new(PaymentProcessorOrderExpired)
	if err := _PaymentProcessor.contract.UnpackLog(event, "OrderExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorOrderPaidIterator is returned from FilterOrderPaid and is used to iterate over the raw logs and unpacked data for OrderPaid events raised by the PaymentProcessor contract.
type PaymentProcessorOrderPaidIterator struct {
	Event *PaymentProcessorOrderPaid // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorOrderPaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorOrderPaid)
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
		it.Event = new(PaymentProcessorOrderPaid)
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
func (it *PaymentProcessorOrderPaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorOrderPaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorOrderPaid represents a OrderPaid event raised by the PaymentProcessor contract.
type PaymentProcessorOrderPaid struct {
	OrderId [32]byte
	Payer   common.Address
	Amount  *big.Int
	Status  uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderPaid is a free log retrieval operation binding the contract event 0x80465c2698455bc6862471ab072190da60f525bebba942641614691b53a8239d.
//
// Solidity: event OrderPaid(bytes32 indexed orderId, address indexed payer, uint256 amount, uint8 status)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterOrderPaid(opts *bind.FilterOpts, orderId [][32]byte, payer []common.Address) (*PaymentProcessorOrderPaidIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "OrderPaid", orderIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorOrderPaidIterator{contract: _PaymentProcessor.contract, event: "OrderPaid", logs: logs, sub: sub}, nil
}

// WatchOrderPaid is a free log subscription operation binding the contract event 0x80465c2698455bc6862471ab072190da60f525bebba942641614691b53a8239d.
//
// Solidity: event OrderPaid(bytes32 indexed orderId, address indexed payer, uint256 amount, uint8 status)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchOrderPaid(opts *bind.WatchOpts, sink chan<- *PaymentProcessorOrderPaid, orderId [][32]byte, payer []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "OrderPaid", orderIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorOrderPaid)
				if err := _PaymentProcessor.contract.UnpackLog(event, "OrderPaid", log); err != nil {
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

// ParseOrderPaid is a log parse operation binding the contract event 0x80465c2698455bc6862471ab072190da60f525bebba942641614691b53a8239d.
//
// Solidity: event OrderPaid(bytes32 indexed orderId, address indexed payer, uint256 amount, uint8 status)
func (_PaymentProcessor *PaymentProcessorFilterer) ParseOrderPaid(log types.Log) (*PaymentProcessorOrderPaid, error) {
	event := new(PaymentProcessorOrderPaid)
	if err := _PaymentProcessor.contract.UnpackLog(event, "OrderPaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorOrderRefundedIterator is returned from FilterOrderRefunded and is used to iterate over the raw logs and unpacked data for OrderRefunded events raised by the PaymentProcessor contract.
type PaymentProcessorOrderRefundedIterator struct {
	Event *PaymentProcessorOrderRefunded // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorOrderRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorOrderRefunded)
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
		it.Event = new(PaymentProcessorOrderRefunded)
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
func (it *PaymentProcessorOrderRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorOrderRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorOrderRefunded represents a OrderRefunded event raised by the PaymentProcessor contract.
type PaymentProcessorOrderRefunded struct {
	OrderId [32]byte
	Payer   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOrderRefunded is a free log retrieval operation binding the contract event 0x5573a774401f48f00a6daf92475b23c21e67cb0de431cf3d63cd0558188d155c.
//
// Solidity: event OrderRefunded(bytes32 indexed orderId, address indexed payer, uint256 amount)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterOrderRefunded(opts *bind.FilterOpts, orderId [][32]byte, payer []common.Address) (*PaymentProcessorOrderRefundedIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "OrderRefunded", orderIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorOrderRefundedIterator{contract: _PaymentProcessor.contract, event: "OrderRefunded", logs: logs, sub: sub}, nil
}

// WatchOrderRefunded is a free log subscription operation binding the contract event 0x5573a774401f48f00a6daf92475b23c21e67cb0de431cf3d63cd0558188d155c.
//
// Solidity: event OrderRefunded(bytes32 indexed orderId, address indexed payer, uint256 amount)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchOrderRefunded(opts *bind.WatchOpts, sink chan<- *PaymentProcessorOrderRefunded, orderId [][32]byte, payer []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var payerRule []interface{}
	for _, payerItem := range payer {
		payerRule = append(payerRule, payerItem)
	}

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "OrderRefunded", orderIdRule, payerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorOrderRefunded)
				if err := _PaymentProcessor.contract.UnpackLog(event, "OrderRefunded", log); err != nil {
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

// ParseOrderRefunded is a log parse operation binding the contract event 0x5573a774401f48f00a6daf92475b23c21e67cb0de431cf3d63cd0558188d155c.
//
// Solidity: event OrderRefunded(bytes32 indexed orderId, address indexed payer, uint256 amount)
func (_PaymentProcessor *PaymentProcessorFilterer) ParseOrderRefunded(log types.Log) (*PaymentProcessorOrderRefunded, error) {
	event := new(PaymentProcessorOrderRefunded)
	if err := _PaymentProcessor.contract.UnpackLog(event, "OrderRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorOrderSettledIterator is returned from FilterOrderSettled and is used to iterate over the raw logs and unpacked data for OrderSettled events raised by the PaymentProcessor contract.
type PaymentProcessorOrderSettledIterator struct {
	Event *PaymentProcessorOrderSettled // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorOrderSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorOrderSettled)
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
		it.Event = new(PaymentProcessorOrderSettled)
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
func (it *PaymentProcessorOrderSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorOrderSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorOrderSettled represents a OrderSettled event raised by the PaymentProcessor contract.
type PaymentProcessorOrderSettled struct {
	OrderId   [32]byte
	Merchant  common.Address
	Payer     common.Address
	NetAmount *big.Int
	Fee       *big.Int
	Status    uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOrderSettled is a free log retrieval operation binding the contract event 0x38d2bd53295e3b0b06d39a1d4d4e09ed446cdca372069f2403ea3c1409077eb6.
//
// Solidity: event OrderSettled(bytes32 indexed orderId, address indexed merchant, address payer, uint256 netAmount, uint256 fee, uint8 status)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterOrderSettled(opts *bind.FilterOpts, orderId [][32]byte, merchant []common.Address) (*PaymentProcessorOrderSettledIterator, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var merchantRule []interface{}
	for _, merchantItem := range merchant {
		merchantRule = append(merchantRule, merchantItem)
	}

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "OrderSettled", orderIdRule, merchantRule)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorOrderSettledIterator{contract: _PaymentProcessor.contract, event: "OrderSettled", logs: logs, sub: sub}, nil
}

// WatchOrderSettled is a free log subscription operation binding the contract event 0x38d2bd53295e3b0b06d39a1d4d4e09ed446cdca372069f2403ea3c1409077eb6.
//
// Solidity: event OrderSettled(bytes32 indexed orderId, address indexed merchant, address payer, uint256 netAmount, uint256 fee, uint8 status)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchOrderSettled(opts *bind.WatchOpts, sink chan<- *PaymentProcessorOrderSettled, orderId [][32]byte, merchant []common.Address) (event.Subscription, error) {

	var orderIdRule []interface{}
	for _, orderIdItem := range orderId {
		orderIdRule = append(orderIdRule, orderIdItem)
	}
	var merchantRule []interface{}
	for _, merchantItem := range merchant {
		merchantRule = append(merchantRule, merchantItem)
	}

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "OrderSettled", orderIdRule, merchantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorOrderSettled)
				if err := _PaymentProcessor.contract.UnpackLog(event, "OrderSettled", log); err != nil {
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

// ParseOrderSettled is a log parse operation binding the contract event 0x38d2bd53295e3b0b06d39a1d4d4e09ed446cdca372069f2403ea3c1409077eb6.
//
// Solidity: event OrderSettled(bytes32 indexed orderId, address indexed merchant, address payer, uint256 netAmount, uint256 fee, uint8 status)
func (_PaymentProcessor *PaymentProcessorFilterer) ParseOrderSettled(log types.Log) (*PaymentProcessorOrderSettled, error) {
	event := new(PaymentProcessorOrderSettled)
	if err := _PaymentProcessor.contract.UnpackLog(event, "OrderSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the PaymentProcessor contract.
type PaymentProcessorOwnershipTransferStartedIterator struct {
	Event *PaymentProcessorOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorOwnershipTransferStarted)
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
		it.Event = new(PaymentProcessorOwnershipTransferStarted)
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
func (it *PaymentProcessorOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the PaymentProcessor contract.
type PaymentProcessorOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaymentProcessorOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorOwnershipTransferStartedIterator{contract: _PaymentProcessor.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *PaymentProcessorOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorOwnershipTransferStarted)
				if err := _PaymentProcessor.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_PaymentProcessor *PaymentProcessorFilterer) ParseOwnershipTransferStarted(log types.Log) (*PaymentProcessorOwnershipTransferStarted, error) {
	event := new(PaymentProcessorOwnershipTransferStarted)
	if err := _PaymentProcessor.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PaymentProcessor contract.
type PaymentProcessorOwnershipTransferredIterator struct {
	Event *PaymentProcessorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorOwnershipTransferred)
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
		it.Event = new(PaymentProcessorOwnershipTransferred)
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
func (it *PaymentProcessorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorOwnershipTransferred represents a OwnershipTransferred event raised by the PaymentProcessor contract.
type PaymentProcessorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PaymentProcessorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorOwnershipTransferredIterator{contract: _PaymentProcessor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PaymentProcessorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorOwnershipTransferred)
				if err := _PaymentProcessor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PaymentProcessor *PaymentProcessorFilterer) ParseOwnershipTransferred(log types.Log) (*PaymentProcessorOwnershipTransferred, error) {
	event := new(PaymentProcessorOwnershipTransferred)
	if err := _PaymentProcessor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the PaymentProcessor contract.
type PaymentProcessorPausedIterator struct {
	Event *PaymentProcessorPaused // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorPaused)
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
		it.Event = new(PaymentProcessorPaused)
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
func (it *PaymentProcessorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorPaused represents a Paused event raised by the PaymentProcessor contract.
type PaymentProcessorPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterPaused(opts *bind.FilterOpts) (*PaymentProcessorPausedIterator, error) {

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorPausedIterator{contract: _PaymentProcessor.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *PaymentProcessorPaused) (event.Subscription, error) {

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorPaused)
				if err := _PaymentProcessor.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_PaymentProcessor *PaymentProcessorFilterer) ParsePaused(log types.Log) (*PaymentProcessorPaused, error) {
	event := new(PaymentProcessorPaused)
	if err := _PaymentProcessor.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorProtocolAddressUpdatedIterator is returned from FilterProtocolAddressUpdated and is used to iterate over the raw logs and unpacked data for ProtocolAddressUpdated events raised by the PaymentProcessor contract.
type PaymentProcessorProtocolAddressUpdatedIterator struct {
	Event *PaymentProcessorProtocolAddressUpdated // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorProtocolAddressUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorProtocolAddressUpdated)
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
		it.Event = new(PaymentProcessorProtocolAddressUpdated)
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
func (it *PaymentProcessorProtocolAddressUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorProtocolAddressUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorProtocolAddressUpdated represents a ProtocolAddressUpdated event raised by the PaymentProcessor contract.
type PaymentProcessorProtocolAddressUpdated struct {
	What           [32]byte
	PlatformWallet common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProtocolAddressUpdated is a free log retrieval operation binding the contract event 0xbbc5b96e57cfecb3dbeeadf92e87f15e58e64fcd75cbe256dcc5d9ef2e51e8a4.
//
// Solidity: event ProtocolAddressUpdated(bytes32 indexed what, address indexed platformWallet)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterProtocolAddressUpdated(opts *bind.FilterOpts, what [][32]byte, platformWallet []common.Address) (*PaymentProcessorProtocolAddressUpdatedIterator, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}
	var platformWalletRule []interface{}
	for _, platformWalletItem := range platformWallet {
		platformWalletRule = append(platformWalletRule, platformWalletItem)
	}

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "ProtocolAddressUpdated", whatRule, platformWalletRule)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorProtocolAddressUpdatedIterator{contract: _PaymentProcessor.contract, event: "ProtocolAddressUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolAddressUpdated is a free log subscription operation binding the contract event 0xbbc5b96e57cfecb3dbeeadf92e87f15e58e64fcd75cbe256dcc5d9ef2e51e8a4.
//
// Solidity: event ProtocolAddressUpdated(bytes32 indexed what, address indexed platformWallet)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchProtocolAddressUpdated(opts *bind.WatchOpts, sink chan<- *PaymentProcessorProtocolAddressUpdated, what [][32]byte, platformWallet []common.Address) (event.Subscription, error) {

	var whatRule []interface{}
	for _, whatItem := range what {
		whatRule = append(whatRule, whatItem)
	}
	var platformWalletRule []interface{}
	for _, platformWalletItem := range platformWallet {
		platformWalletRule = append(platformWalletRule, platformWalletItem)
	}

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "ProtocolAddressUpdated", whatRule, platformWalletRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorProtocolAddressUpdated)
				if err := _PaymentProcessor.contract.UnpackLog(event, "ProtocolAddressUpdated", log); err != nil {
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

// ParseProtocolAddressUpdated is a log parse operation binding the contract event 0xbbc5b96e57cfecb3dbeeadf92e87f15e58e64fcd75cbe256dcc5d9ef2e51e8a4.
//
// Solidity: event ProtocolAddressUpdated(bytes32 indexed what, address indexed platformWallet)
func (_PaymentProcessor *PaymentProcessorFilterer) ParseProtocolAddressUpdated(log types.Log) (*PaymentProcessorProtocolAddressUpdated, error) {
	event := new(PaymentProcessorProtocolAddressUpdated)
	if err := _PaymentProcessor.contract.UnpackLog(event, "ProtocolAddressUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorTokenSupportUpdatedIterator is returned from FilterTokenSupportUpdated and is used to iterate over the raw logs and unpacked data for TokenSupportUpdated events raised by the PaymentProcessor contract.
type PaymentProcessorTokenSupportUpdatedIterator struct {
	Event *PaymentProcessorTokenSupportUpdated // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorTokenSupportUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorTokenSupportUpdated)
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
		it.Event = new(PaymentProcessorTokenSupportUpdated)
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
func (it *PaymentProcessorTokenSupportUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorTokenSupportUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorTokenSupportUpdated represents a TokenSupportUpdated event raised by the PaymentProcessor contract.
type PaymentProcessorTokenSupportUpdated struct {
	Token     [32]byte
	Supported bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenSupportUpdated is a free log retrieval operation binding the contract event 0x475cd53258eb63abbbc7ae999308fd56beb3f4d5c002442aaa733a5010295618.
//
// Solidity: event TokenSupportUpdated(bytes32 indexed token, bool supported)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterTokenSupportUpdated(opts *bind.FilterOpts, token [][32]byte) (*PaymentProcessorTokenSupportUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "TokenSupportUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorTokenSupportUpdatedIterator{contract: _PaymentProcessor.contract, event: "TokenSupportUpdated", logs: logs, sub: sub}, nil
}

// WatchTokenSupportUpdated is a free log subscription operation binding the contract event 0x475cd53258eb63abbbc7ae999308fd56beb3f4d5c002442aaa733a5010295618.
//
// Solidity: event TokenSupportUpdated(bytes32 indexed token, bool supported)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchTokenSupportUpdated(opts *bind.WatchOpts, sink chan<- *PaymentProcessorTokenSupportUpdated, token [][32]byte) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "TokenSupportUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorTokenSupportUpdated)
				if err := _PaymentProcessor.contract.UnpackLog(event, "TokenSupportUpdated", log); err != nil {
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

// ParseTokenSupportUpdated is a log parse operation binding the contract event 0x475cd53258eb63abbbc7ae999308fd56beb3f4d5c002442aaa733a5010295618.
//
// Solidity: event TokenSupportUpdated(bytes32 indexed token, bool supported)
func (_PaymentProcessor *PaymentProcessorFilterer) ParseTokenSupportUpdated(log types.Log) (*PaymentProcessorTokenSupportUpdated, error) {
	event := new(PaymentProcessorTokenSupportUpdated)
	if err := _PaymentProcessor.contract.UnpackLog(event, "TokenSupportUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PaymentProcessorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the PaymentProcessor contract.
type PaymentProcessorUnpausedIterator struct {
	Event *PaymentProcessorUnpaused // Event containing the contract specifics and raw log

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
func (it *PaymentProcessorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentProcessorUnpaused)
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
		it.Event = new(PaymentProcessorUnpaused)
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
func (it *PaymentProcessorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentProcessorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentProcessorUnpaused represents a Unpaused event raised by the PaymentProcessor contract.
type PaymentProcessorUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PaymentProcessor *PaymentProcessorFilterer) FilterUnpaused(opts *bind.FilterOpts) (*PaymentProcessorUnpausedIterator, error) {

	logs, sub, err := _PaymentProcessor.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &PaymentProcessorUnpausedIterator{contract: _PaymentProcessor.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_PaymentProcessor *PaymentProcessorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *PaymentProcessorUnpaused) (event.Subscription, error) {

	logs, sub, err := _PaymentProcessor.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentProcessorUnpaused)
				if err := _PaymentProcessor.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_PaymentProcessor *PaymentProcessorFilterer) ParseUnpaused(log types.Log) (*PaymentProcessorUnpaused, error) {
	event := new(PaymentProcessorUnpaused)
	if err := _PaymentProcessor.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
