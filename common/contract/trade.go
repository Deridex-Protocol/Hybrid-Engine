// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// DxlnOrdersFill is an auto generated low-level Go binding around an user-defined struct.
type DxlnOrdersFill struct {
	Amount        *big.Int
	Price         *big.Int
	Fee           *big.Int
	IsNegativeFee bool
}

// DxlnOrdersOrder is an auto generated low-level Go binding around an user-defined struct.
type DxlnOrdersOrder struct {
	Flags        [32]byte
	Amount       *big.Int
	LimitPrice   *big.Int
	TriggerPrice *big.Int
	LimitFee     *big.Int
	Maker        common.Address
	Taker        common.Address
	Expiration   *big.Int
}

// DxlnOrdersTradeData is an auto generated low-level Go binding around an user-defined struct.
type DxlnOrdersTradeData struct {
	Order     DxlnOrdersOrder
	Fill      DxlnOrdersFill
	Signature TypedSignatureSignature
}

// DxlnTypesTradeResult is an auto generated low-level Go binding around an user-defined struct.
type DxlnTypesTradeResult struct {
	MarginAmount   *big.Int
	PositionAmount *big.Int
	IsBuy          bool
	TraderFlags    [32]byte
}

// TypedSignatureSignature is an auto generated low-level Go binding around an user-defined struct.
type TypedSignatureSignature struct {
	R     [32]byte
	S     [32]byte
	VType [2]byte
}

// TradeMetaData contains all meta data concerning the Trade contract.
var TradeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"flags\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"triggerPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"internalType\":\"structDxlnOrders.Order\",\"name\":\"order\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isNegativeFee\",\"type\":\"bool\"}],\"internalType\":\"structDxlnOrders.Fill\",\"name\":\"fill\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes2\",\"name\":\"vType\",\"type\":\"bytes2\"}],\"internalType\":\"structTypedSignature.Signature\",\"name\":\"signature\",\"type\":\"tuple\"}],\"internalType\":\"structDxlnOrders.TradeData\",\"name\":\"data\",\"type\":\"tuple\"}],\"name\":\"tradeExample\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"marginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"traderFlags\",\"type\":\"bytes32\"}],\"internalType\":\"structDxlnTypes.TradeResult\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TradeABI is the input ABI used to generate the binding from.
// Deprecated: Use TradeMetaData.ABI instead.
var TradeABI = TradeMetaData.ABI

// Trade is an auto generated Go binding around an Ethereum contract.
type Trade struct {
	TradeCaller     // Read-only binding to the contract
	TradeTransactor // Write-only binding to the contract
	TradeFilterer   // Log filterer for contract events
}

// TradeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradeSession struct {
	Contract     *Trade            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradeCallerSession struct {
	Contract *TradeCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TradeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradeTransactorSession struct {
	Contract     *TradeTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradeRaw struct {
	Contract *Trade // Generic contract binding to access the raw methods on
}

// TradeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradeCallerRaw struct {
	Contract *TradeCaller // Generic read-only contract binding to access the raw methods on
}

// TradeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradeTransactorRaw struct {
	Contract *TradeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrade creates a new instance of Trade, bound to a specific deployed contract.
func NewTrade(address common.Address, backend bind.ContractBackend) (*Trade, error) {
	contract, err := bindTrade(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trade{TradeCaller: TradeCaller{contract: contract}, TradeTransactor: TradeTransactor{contract: contract}, TradeFilterer: TradeFilterer{contract: contract}}, nil
}

// NewTradeCaller creates a new read-only instance of Trade, bound to a specific deployed contract.
func NewTradeCaller(address common.Address, caller bind.ContractCaller) (*TradeCaller, error) {
	contract, err := bindTrade(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradeCaller{contract: contract}, nil
}

// NewTradeTransactor creates a new write-only instance of Trade, bound to a specific deployed contract.
func NewTradeTransactor(address common.Address, transactor bind.ContractTransactor) (*TradeTransactor, error) {
	contract, err := bindTrade(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradeTransactor{contract: contract}, nil
}

// NewTradeFilterer creates a new log filterer instance of Trade, bound to a specific deployed contract.
func NewTradeFilterer(address common.Address, filterer bind.ContractFilterer) (*TradeFilterer, error) {
	contract, err := bindTrade(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradeFilterer{contract: contract}, nil
}

// bindTrade binds a generic wrapper to an already deployed contract.
func bindTrade(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TradeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trade *TradeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trade.Contract.TradeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trade *TradeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.Contract.TradeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trade *TradeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trade.Contract.TradeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trade *TradeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Trade.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trade *TradeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trade.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trade *TradeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trade.Contract.contract.Transact(opts, method, params...)
}

// TradeExample is a paid mutator transaction binding the contract method 0x637e8d9b.
//
// Solidity: function tradeExample(((bytes32,uint256,uint256,uint256,uint256,address,address,uint256),(uint256,uint256,uint256,bool),(bytes32,bytes32,bytes2)) data) returns((uint256,uint256,bool,bytes32))
func (_Trade *TradeTransactor) TradeExample(opts *bind.TransactOpts, data DxlnOrdersTradeData) (*types.Transaction, error) {
	return _Trade.contract.Transact(opts, "tradeExample", data)
}

// TradeExample is a paid mutator transaction binding the contract method 0x637e8d9b.
//
// Solidity: function tradeExample(((bytes32,uint256,uint256,uint256,uint256,address,address,uint256),(uint256,uint256,uint256,bool),(bytes32,bytes32,bytes2)) data) returns((uint256,uint256,bool,bytes32))
func (_Trade *TradeSession) TradeExample(data DxlnOrdersTradeData) (*types.Transaction, error) {
	return _Trade.Contract.TradeExample(&_Trade.TransactOpts, data)
}

// TradeExample is a paid mutator transaction binding the contract method 0x637e8d9b.
//
// Solidity: function tradeExample(((bytes32,uint256,uint256,uint256,uint256,address,address,uint256),(uint256,uint256,uint256,bool),(bytes32,bytes32,bytes2)) data) returns((uint256,uint256,bool,bytes32))
func (_Trade *TradeTransactorSession) TradeExample(data DxlnOrdersTradeData) (*types.Transaction, error) {
	return _Trade.Contract.TradeExample(&_Trade.TransactOpts, data)
}
