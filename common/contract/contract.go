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

// DxlnTradeTradeArg is an auto generated low-level Go binding around an user-defined struct.
type DxlnTradeTradeArg struct {
	TakerIndex *big.Int
	MakerIndex *big.Int
	Trader     common.Address
	Data       []byte
}

// DxlnTypesBalance is an auto generated low-level Go binding around an user-defined struct.
type DxlnTypesBalance struct {
	MarginIsPositive   bool
	PositionIsPositive bool
	Margin             *big.Int
	Position           *big.Int
}

// DxlnTypesIndex is an auto generated low-level Go binding around an user-defined struct.
type DxlnTypesIndex struct {
	Timestamp  uint32
	IsPositive bool
	Value      *big.Int
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isPositive\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"balance\",\"type\":\"bytes32\"}],\"name\":\"LogAccountSettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"balance\",\"type\":\"bytes32\"}],\"name\":\"LogDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"settlementPrice\",\"type\":\"uint256\"}],\"name\":\"LogFinalSettlementEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"index\",\"type\":\"bytes32\"}],\"name\":\"LogIndex\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"funder\",\"type\":\"address\"}],\"name\":\"LogSetFunder\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"LogSetGlobalOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"LogSetLocalOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minCollateral\",\"type\":\"uint256\"}],\"name\":\"LogSetMinCollateral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"LogSetOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"marginAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"positionAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isBuy\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"makerBalance\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"takerBalance\",\"type\":\"bytes32\"}],\"name\":\"LogTrade\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"balance\",\"type\":\"bytes32\"}],\"name\":\"LogWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"balance\",\"type\":\"bytes32\"}],\"name\":\"LogWithdrawFinalSettlement\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"priceLowerBound\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"priceUpperBound\",\"type\":\"uint256\"}],\"name\":\"enableFinalSettlement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountBalance\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"marginIsPositive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"positionIsPositive\",\"type\":\"bool\"},{\"internalType\":\"uint120\",\"name\":\"margin\",\"type\":\"uint120\"},{\"internalType\":\"uint120\",\"name\":\"position\",\"type\":\"uint120\"}],\"internalType\":\"structDxlnTypes.Balance\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isPositive\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"internalType\":\"structDxlnTypes.Index\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFinalSettlementEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFunderContract\",\"outputs\":[{\"internalType\":\"contractI_DxlnFunder\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGlobalIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isPositive\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"value\",\"type\":\"uint128\"}],\"internalType\":\"structDxlnTypes.Index\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getIsGlobalOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"getIsLocalOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracleContract\",\"outputs\":[{\"internalType\":\"contractI_DxlnOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOraclePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"hasAccountPermissions\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"funder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minCollateral\",\"type\":\"uint256\"}],\"name\":\"initializeV1\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"funder\",\"type\":\"address\"}],\"name\":\"setFunder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setGlobalOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setLocalOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minCollateral\",\"type\":\"uint256\"}],\"name\":\"setMinCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"takerIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"makerIndex\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structDxlnTrade.TradeArg[]\",\"name\":\"trades\",\"type\":\"tuple[]\"}],\"name\":\"trade\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFinalSettlement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetAccountBalance is a free data retrieval call binding the contract method 0x93423e9c.
//
// Solidity: function getAccountBalance(address account) view returns((bool,bool,uint120,uint120))
func (_Contract *ContractCaller) GetAccountBalance(opts *bind.CallOpts, account common.Address) (DxlnTypesBalance, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAccountBalance", account)

	if err != nil {
		return *new(DxlnTypesBalance), err
	}

	out0 := *abi.ConvertType(out[0], new(DxlnTypesBalance)).(*DxlnTypesBalance)

	return out0, err

}

// GetAccountBalance is a free data retrieval call binding the contract method 0x93423e9c.
//
// Solidity: function getAccountBalance(address account) view returns((bool,bool,uint120,uint120))
func (_Contract *ContractSession) GetAccountBalance(account common.Address) (DxlnTypesBalance, error) {
	return _Contract.Contract.GetAccountBalance(&_Contract.CallOpts, account)
}

// GetAccountBalance is a free data retrieval call binding the contract method 0x93423e9c.
//
// Solidity: function getAccountBalance(address account) view returns((bool,bool,uint120,uint120))
func (_Contract *ContractCallerSession) GetAccountBalance(account common.Address) (DxlnTypesBalance, error) {
	return _Contract.Contract.GetAccountBalance(&_Contract.CallOpts, account)
}

// GetAccountIndex is a free data retrieval call binding the contract method 0x9ba63e9e.
//
// Solidity: function getAccountIndex(address account) view returns((uint32,bool,uint128))
func (_Contract *ContractCaller) GetAccountIndex(opts *bind.CallOpts, account common.Address) (DxlnTypesIndex, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAccountIndex", account)

	if err != nil {
		return *new(DxlnTypesIndex), err
	}

	out0 := *abi.ConvertType(out[0], new(DxlnTypesIndex)).(*DxlnTypesIndex)

	return out0, err

}

// GetAccountIndex is a free data retrieval call binding the contract method 0x9ba63e9e.
//
// Solidity: function getAccountIndex(address account) view returns((uint32,bool,uint128))
func (_Contract *ContractSession) GetAccountIndex(account common.Address) (DxlnTypesIndex, error) {
	return _Contract.Contract.GetAccountIndex(&_Contract.CallOpts, account)
}

// GetAccountIndex is a free data retrieval call binding the contract method 0x9ba63e9e.
//
// Solidity: function getAccountIndex(address account) view returns((uint32,bool,uint128))
func (_Contract *ContractCallerSession) GetAccountIndex(account common.Address) (DxlnTypesIndex, error) {
	return _Contract.Contract.GetAccountIndex(&_Contract.CallOpts, account)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Contract *ContractCaller) GetAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Contract *ContractSession) GetAdmin() (common.Address, error) {
	return _Contract.Contract.GetAdmin(&_Contract.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address)
func (_Contract *ContractCallerSession) GetAdmin() (common.Address, error) {
	return _Contract.Contract.GetAdmin(&_Contract.CallOpts)
}

// GetFinalSettlementEnabled is a free data retrieval call binding the contract method 0x7099366b.
//
// Solidity: function getFinalSettlementEnabled() view returns(bool)
func (_Contract *ContractCaller) GetFinalSettlementEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getFinalSettlementEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetFinalSettlementEnabled is a free data retrieval call binding the contract method 0x7099366b.
//
// Solidity: function getFinalSettlementEnabled() view returns(bool)
func (_Contract *ContractSession) GetFinalSettlementEnabled() (bool, error) {
	return _Contract.Contract.GetFinalSettlementEnabled(&_Contract.CallOpts)
}

// GetFinalSettlementEnabled is a free data retrieval call binding the contract method 0x7099366b.
//
// Solidity: function getFinalSettlementEnabled() view returns(bool)
func (_Contract *ContractCallerSession) GetFinalSettlementEnabled() (bool, error) {
	return _Contract.Contract.GetFinalSettlementEnabled(&_Contract.CallOpts)
}

// GetFunderContract is a free data retrieval call binding the contract method 0xdc4f3a0e.
//
// Solidity: function getFunderContract() view returns(address)
func (_Contract *ContractCaller) GetFunderContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getFunderContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFunderContract is a free data retrieval call binding the contract method 0xdc4f3a0e.
//
// Solidity: function getFunderContract() view returns(address)
func (_Contract *ContractSession) GetFunderContract() (common.Address, error) {
	return _Contract.Contract.GetFunderContract(&_Contract.CallOpts)
}

// GetFunderContract is a free data retrieval call binding the contract method 0xdc4f3a0e.
//
// Solidity: function getFunderContract() view returns(address)
func (_Contract *ContractCallerSession) GetFunderContract() (common.Address, error) {
	return _Contract.Contract.GetFunderContract(&_Contract.CallOpts)
}

// GetGlobalIndex is a free data retrieval call binding the contract method 0x80d63681.
//
// Solidity: function getGlobalIndex() view returns((uint32,bool,uint128))
func (_Contract *ContractCaller) GetGlobalIndex(opts *bind.CallOpts) (DxlnTypesIndex, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getGlobalIndex")

	if err != nil {
		return *new(DxlnTypesIndex), err
	}

	out0 := *abi.ConvertType(out[0], new(DxlnTypesIndex)).(*DxlnTypesIndex)

	return out0, err

}

// GetGlobalIndex is a free data retrieval call binding the contract method 0x80d63681.
//
// Solidity: function getGlobalIndex() view returns((uint32,bool,uint128))
func (_Contract *ContractSession) GetGlobalIndex() (DxlnTypesIndex, error) {
	return _Contract.Contract.GetGlobalIndex(&_Contract.CallOpts)
}

// GetGlobalIndex is a free data retrieval call binding the contract method 0x80d63681.
//
// Solidity: function getGlobalIndex() view returns((uint32,bool,uint128))
func (_Contract *ContractCallerSession) GetGlobalIndex() (DxlnTypesIndex, error) {
	return _Contract.Contract.GetGlobalIndex(&_Contract.CallOpts)
}

// GetIsGlobalOperator is a free data retrieval call binding the contract method 0x052f72d7.
//
// Solidity: function getIsGlobalOperator(address operator) view returns(bool)
func (_Contract *ContractCaller) GetIsGlobalOperator(opts *bind.CallOpts, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getIsGlobalOperator", operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsGlobalOperator is a free data retrieval call binding the contract method 0x052f72d7.
//
// Solidity: function getIsGlobalOperator(address operator) view returns(bool)
func (_Contract *ContractSession) GetIsGlobalOperator(operator common.Address) (bool, error) {
	return _Contract.Contract.GetIsGlobalOperator(&_Contract.CallOpts, operator)
}

// GetIsGlobalOperator is a free data retrieval call binding the contract method 0x052f72d7.
//
// Solidity: function getIsGlobalOperator(address operator) view returns(bool)
func (_Contract *ContractCallerSession) GetIsGlobalOperator(operator common.Address) (bool, error) {
	return _Contract.Contract.GetIsGlobalOperator(&_Contract.CallOpts, operator)
}

// GetIsLocalOperator is a free data retrieval call binding the contract method 0x3a031bf0.
//
// Solidity: function getIsLocalOperator(address account, address operator) view returns(bool)
func (_Contract *ContractCaller) GetIsLocalOperator(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getIsLocalOperator", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsLocalOperator is a free data retrieval call binding the contract method 0x3a031bf0.
//
// Solidity: function getIsLocalOperator(address account, address operator) view returns(bool)
func (_Contract *ContractSession) GetIsLocalOperator(account common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.GetIsLocalOperator(&_Contract.CallOpts, account, operator)
}

// GetIsLocalOperator is a free data retrieval call binding the contract method 0x3a031bf0.
//
// Solidity: function getIsLocalOperator(address account, address operator) view returns(bool)
func (_Contract *ContractCallerSession) GetIsLocalOperator(account common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.GetIsLocalOperator(&_Contract.CallOpts, account, operator)
}

// GetMinCollateral is a free data retrieval call binding the contract method 0xe830b690.
//
// Solidity: function getMinCollateral() view returns(uint256)
func (_Contract *ContractCaller) GetMinCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getMinCollateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinCollateral is a free data retrieval call binding the contract method 0xe830b690.
//
// Solidity: function getMinCollateral() view returns(uint256)
func (_Contract *ContractSession) GetMinCollateral() (*big.Int, error) {
	return _Contract.Contract.GetMinCollateral(&_Contract.CallOpts)
}

// GetMinCollateral is a free data retrieval call binding the contract method 0xe830b690.
//
// Solidity: function getMinCollateral() view returns(uint256)
func (_Contract *ContractCallerSession) GetMinCollateral() (*big.Int, error) {
	return _Contract.Contract.GetMinCollateral(&_Contract.CallOpts)
}

// GetOracleContract is a free data retrieval call binding the contract method 0xe3bbb565.
//
// Solidity: function getOracleContract() view returns(address)
func (_Contract *ContractCaller) GetOracleContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOracleContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOracleContract is a free data retrieval call binding the contract method 0xe3bbb565.
//
// Solidity: function getOracleContract() view returns(address)
func (_Contract *ContractSession) GetOracleContract() (common.Address, error) {
	return _Contract.Contract.GetOracleContract(&_Contract.CallOpts)
}

// GetOracleContract is a free data retrieval call binding the contract method 0xe3bbb565.
//
// Solidity: function getOracleContract() view returns(address)
func (_Contract *ContractCallerSession) GetOracleContract() (common.Address, error) {
	return _Contract.Contract.GetOracleContract(&_Contract.CallOpts)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x796da7af.
//
// Solidity: function getOraclePrice() view returns(uint256)
func (_Contract *ContractCaller) GetOraclePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getOraclePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOraclePrice is a free data retrieval call binding the contract method 0x796da7af.
//
// Solidity: function getOraclePrice() view returns(uint256)
func (_Contract *ContractSession) GetOraclePrice() (*big.Int, error) {
	return _Contract.Contract.GetOraclePrice(&_Contract.CallOpts)
}

// GetOraclePrice is a free data retrieval call binding the contract method 0x796da7af.
//
// Solidity: function getOraclePrice() view returns(uint256)
func (_Contract *ContractCallerSession) GetOraclePrice() (*big.Int, error) {
	return _Contract.Contract.GetOraclePrice(&_Contract.CallOpts)
}

// GetTokenContract is a free data retrieval call binding the contract method 0x28b7bede.
//
// Solidity: function getTokenContract() view returns(address)
func (_Contract *ContractCaller) GetTokenContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getTokenContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenContract is a free data retrieval call binding the contract method 0x28b7bede.
//
// Solidity: function getTokenContract() view returns(address)
func (_Contract *ContractSession) GetTokenContract() (common.Address, error) {
	return _Contract.Contract.GetTokenContract(&_Contract.CallOpts)
}

// GetTokenContract is a free data retrieval call binding the contract method 0x28b7bede.
//
// Solidity: function getTokenContract() view returns(address)
func (_Contract *ContractCallerSession) GetTokenContract() (common.Address, error) {
	return _Contract.Contract.GetTokenContract(&_Contract.CallOpts)
}

// HasAccountPermissions is a free data retrieval call binding the contract method 0x84ea2862.
//
// Solidity: function hasAccountPermissions(address account, address operator) view returns(bool)
func (_Contract *ContractCaller) HasAccountPermissions(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "hasAccountPermissions", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAccountPermissions is a free data retrieval call binding the contract method 0x84ea2862.
//
// Solidity: function hasAccountPermissions(address account, address operator) view returns(bool)
func (_Contract *ContractSession) HasAccountPermissions(account common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.HasAccountPermissions(&_Contract.CallOpts, account, operator)
}

// HasAccountPermissions is a free data retrieval call binding the contract method 0x84ea2862.
//
// Solidity: function hasAccountPermissions(address account, address operator) view returns(bool)
func (_Contract *ContractCallerSession) HasAccountPermissions(account common.Address, operator common.Address) (bool, error) {
	return _Contract.Contract.HasAccountPermissions(&_Contract.CallOpts, account, operator)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns()
func (_Contract *ContractTransactor) Deposit(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "deposit", account, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns()
func (_Contract *ContractSession) Deposit(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, account, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns()
func (_Contract *ContractTransactorSession) Deposit(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Deposit(&_Contract.TransactOpts, account, amount)
}

// EnableFinalSettlement is a paid mutator transaction binding the contract method 0xf40c3699.
//
// Solidity: function enableFinalSettlement(uint256 priceLowerBound, uint256 priceUpperBound) returns()
func (_Contract *ContractTransactor) EnableFinalSettlement(opts *bind.TransactOpts, priceLowerBound *big.Int, priceUpperBound *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "enableFinalSettlement", priceLowerBound, priceUpperBound)
}

// EnableFinalSettlement is a paid mutator transaction binding the contract method 0xf40c3699.
//
// Solidity: function enableFinalSettlement(uint256 priceLowerBound, uint256 priceUpperBound) returns()
func (_Contract *ContractSession) EnableFinalSettlement(priceLowerBound *big.Int, priceUpperBound *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.EnableFinalSettlement(&_Contract.TransactOpts, priceLowerBound, priceUpperBound)
}

// EnableFinalSettlement is a paid mutator transaction binding the contract method 0xf40c3699.
//
// Solidity: function enableFinalSettlement(uint256 priceLowerBound, uint256 priceUpperBound) returns()
func (_Contract *ContractTransactorSession) EnableFinalSettlement(priceLowerBound *big.Int, priceUpperBound *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.EnableFinalSettlement(&_Contract.TransactOpts, priceLowerBound, priceUpperBound)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0xa895155b.
//
// Solidity: function initializeV1(address token, address oracle, address funder, uint256 minCollateral) returns()
func (_Contract *ContractTransactor) InitializeV1(opts *bind.TransactOpts, token common.Address, oracle common.Address, funder common.Address, minCollateral *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "initializeV1", token, oracle, funder, minCollateral)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0xa895155b.
//
// Solidity: function initializeV1(address token, address oracle, address funder, uint256 minCollateral) returns()
func (_Contract *ContractSession) InitializeV1(token common.Address, oracle common.Address, funder common.Address, minCollateral *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.InitializeV1(&_Contract.TransactOpts, token, oracle, funder, minCollateral)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0xa895155b.
//
// Solidity: function initializeV1(address token, address oracle, address funder, uint256 minCollateral) returns()
func (_Contract *ContractTransactorSession) InitializeV1(token common.Address, oracle common.Address, funder common.Address, minCollateral *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.InitializeV1(&_Contract.TransactOpts, token, oracle, funder, minCollateral)
}

// SetFunder is a paid mutator transaction binding the contract method 0x0acc8cd1.
//
// Solidity: function setFunder(address funder) returns()
func (_Contract *ContractTransactor) SetFunder(opts *bind.TransactOpts, funder common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setFunder", funder)
}

// SetFunder is a paid mutator transaction binding the contract method 0x0acc8cd1.
//
// Solidity: function setFunder(address funder) returns()
func (_Contract *ContractSession) SetFunder(funder common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetFunder(&_Contract.TransactOpts, funder)
}

// SetFunder is a paid mutator transaction binding the contract method 0x0acc8cd1.
//
// Solidity: function setFunder(address funder) returns()
func (_Contract *ContractTransactorSession) SetFunder(funder common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetFunder(&_Contract.TransactOpts, funder)
}

// SetGlobalOperator is a paid mutator transaction binding the contract method 0x46d256c5.
//
// Solidity: function setGlobalOperator(address operator, bool approved) returns()
func (_Contract *ContractTransactor) SetGlobalOperator(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setGlobalOperator", operator, approved)
}

// SetGlobalOperator is a paid mutator transaction binding the contract method 0x46d256c5.
//
// Solidity: function setGlobalOperator(address operator, bool approved) returns()
func (_Contract *ContractSession) SetGlobalOperator(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetGlobalOperator(&_Contract.TransactOpts, operator, approved)
}

// SetGlobalOperator is a paid mutator transaction binding the contract method 0x46d256c5.
//
// Solidity: function setGlobalOperator(address operator, bool approved) returns()
func (_Contract *ContractTransactorSession) SetGlobalOperator(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetGlobalOperator(&_Contract.TransactOpts, operator, approved)
}

// SetLocalOperator is a paid mutator transaction binding the contract method 0xb4959e72.
//
// Solidity: function setLocalOperator(address operator, bool approved) returns()
func (_Contract *ContractTransactor) SetLocalOperator(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setLocalOperator", operator, approved)
}

// SetLocalOperator is a paid mutator transaction binding the contract method 0xb4959e72.
//
// Solidity: function setLocalOperator(address operator, bool approved) returns()
func (_Contract *ContractSession) SetLocalOperator(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetLocalOperator(&_Contract.TransactOpts, operator, approved)
}

// SetLocalOperator is a paid mutator transaction binding the contract method 0xb4959e72.
//
// Solidity: function setLocalOperator(address operator, bool approved) returns()
func (_Contract *ContractTransactorSession) SetLocalOperator(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetLocalOperator(&_Contract.TransactOpts, operator, approved)
}

// SetMinCollateral is a paid mutator transaction binding the contract method 0x846321a4.
//
// Solidity: function setMinCollateral(uint256 minCollateral) returns()
func (_Contract *ContractTransactor) SetMinCollateral(opts *bind.TransactOpts, minCollateral *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setMinCollateral", minCollateral)
}

// SetMinCollateral is a paid mutator transaction binding the contract method 0x846321a4.
//
// Solidity: function setMinCollateral(uint256 minCollateral) returns()
func (_Contract *ContractSession) SetMinCollateral(minCollateral *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMinCollateral(&_Contract.TransactOpts, minCollateral)
}

// SetMinCollateral is a paid mutator transaction binding the contract method 0x846321a4.
//
// Solidity: function setMinCollateral(uint256 minCollateral) returns()
func (_Contract *ContractTransactorSession) SetMinCollateral(minCollateral *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetMinCollateral(&_Contract.TransactOpts, minCollateral)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_Contract *ContractTransactor) SetOracle(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setOracle", oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_Contract *ContractSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOracle(&_Contract.TransactOpts, oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns()
func (_Contract *ContractTransactorSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _Contract.Contract.SetOracle(&_Contract.TransactOpts, oracle)
}

// Trade is a paid mutator transaction binding the contract method 0x68eec3f6.
//
// Solidity: function trade(address[] accounts, (uint256,uint256,address,bytes)[] trades) returns()
func (_Contract *ContractTransactor) Trade(opts *bind.TransactOpts, accounts []common.Address, trades []DxlnTradeTradeArg) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "trade", accounts, trades)
}

// Trade is a paid mutator transaction binding the contract method 0x68eec3f6.
//
// Solidity: function trade(address[] accounts, (uint256,uint256,address,bytes)[] trades) returns()
func (_Contract *ContractSession) Trade(accounts []common.Address, trades []DxlnTradeTradeArg) (*types.Transaction, error) {
	return _Contract.Contract.Trade(&_Contract.TransactOpts, accounts, trades)
}

// Trade is a paid mutator transaction binding the contract method 0x68eec3f6.
//
// Solidity: function trade(address[] accounts, (uint256,uint256,address,bytes)[] trades) returns()
func (_Contract *ContractTransactorSession) Trade(accounts []common.Address, trades []DxlnTradeTradeArg) (*types.Transaction, error) {
	return _Contract.Contract.Trade(&_Contract.TransactOpts, accounts, trades)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address account, address destination, uint256 amount) returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, account common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", account, destination, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address account, address destination, uint256 amount) returns()
func (_Contract *ContractSession) Withdraw(account common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, account, destination, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address account, address destination, uint256 amount) returns()
func (_Contract *ContractTransactorSession) Withdraw(account common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, account, destination, amount)
}

// WithdrawFinalSettlement is a paid mutator transaction binding the contract method 0x142c69b3.
//
// Solidity: function withdrawFinalSettlement() returns()
func (_Contract *ContractTransactor) WithdrawFinalSettlement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdrawFinalSettlement")
}

// WithdrawFinalSettlement is a paid mutator transaction binding the contract method 0x142c69b3.
//
// Solidity: function withdrawFinalSettlement() returns()
func (_Contract *ContractSession) WithdrawFinalSettlement() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawFinalSettlement(&_Contract.TransactOpts)
}

// WithdrawFinalSettlement is a paid mutator transaction binding the contract method 0x142c69b3.
//
// Solidity: function withdrawFinalSettlement() returns()
func (_Contract *ContractTransactorSession) WithdrawFinalSettlement() (*types.Transaction, error) {
	return _Contract.Contract.WithdrawFinalSettlement(&_Contract.TransactOpts)
}

// ContractLogAccountSettledIterator is returned from FilterLogAccountSettled and is used to iterate over the raw logs and unpacked data for LogAccountSettled events raised by the Contract contract.
type ContractLogAccountSettledIterator struct {
	Event *ContractLogAccountSettled // Event containing the contract specifics and raw log

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
func (it *ContractLogAccountSettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogAccountSettled)
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
		it.Event = new(ContractLogAccountSettled)
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
func (it *ContractLogAccountSettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogAccountSettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogAccountSettled represents a LogAccountSettled event raised by the Contract contract.
type ContractLogAccountSettled struct {
	Account    common.Address
	IsPositive bool
	Amount     *big.Int
	Balance    [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLogAccountSettled is a free log retrieval operation binding the contract event 0x022694ffbbd957d26de6b85c040be68ec582d13d40114b29130581793a1bf31e.
//
// Solidity: event LogAccountSettled(address indexed account, bool isPositive, uint256 amount, bytes32 balance)
func (_Contract *ContractFilterer) FilterLogAccountSettled(opts *bind.FilterOpts, account []common.Address) (*ContractLogAccountSettledIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogAccountSettled", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogAccountSettledIterator{contract: _Contract.contract, event: "LogAccountSettled", logs: logs, sub: sub}, nil
}

// WatchLogAccountSettled is a free log subscription operation binding the contract event 0x022694ffbbd957d26de6b85c040be68ec582d13d40114b29130581793a1bf31e.
//
// Solidity: event LogAccountSettled(address indexed account, bool isPositive, uint256 amount, bytes32 balance)
func (_Contract *ContractFilterer) WatchLogAccountSettled(opts *bind.WatchOpts, sink chan<- *ContractLogAccountSettled, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogAccountSettled", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogAccountSettled)
				if err := _Contract.contract.UnpackLog(event, "LogAccountSettled", log); err != nil {
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

// ParseLogAccountSettled is a log parse operation binding the contract event 0x022694ffbbd957d26de6b85c040be68ec582d13d40114b29130581793a1bf31e.
//
// Solidity: event LogAccountSettled(address indexed account, bool isPositive, uint256 amount, bytes32 balance)
func (_Contract *ContractFilterer) ParseLogAccountSettled(log types.Log) (*ContractLogAccountSettled, error) {
	event := new(ContractLogAccountSettled)
	if err := _Contract.contract.UnpackLog(event, "LogAccountSettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogDepositIterator is returned from FilterLogDeposit and is used to iterate over the raw logs and unpacked data for LogDeposit events raised by the Contract contract.
type ContractLogDepositIterator struct {
	Event *ContractLogDeposit // Event containing the contract specifics and raw log

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
func (it *ContractLogDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogDeposit)
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
		it.Event = new(ContractLogDeposit)
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
func (it *ContractLogDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogDeposit represents a LogDeposit event raised by the Contract contract.
type ContractLogDeposit struct {
	Account common.Address
	Amount  *big.Int
	Balance [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogDeposit is a free log retrieval operation binding the contract event 0x40a9cb3a9707d3a68091d8ef7ffd4158d01d0b2ad92b1e489abe8312dd543023.
//
// Solidity: event LogDeposit(address indexed account, uint256 amount, bytes32 balance)
func (_Contract *ContractFilterer) FilterLogDeposit(opts *bind.FilterOpts, account []common.Address) (*ContractLogDepositIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogDeposit", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogDepositIterator{contract: _Contract.contract, event: "LogDeposit", logs: logs, sub: sub}, nil
}

// WatchLogDeposit is a free log subscription operation binding the contract event 0x40a9cb3a9707d3a68091d8ef7ffd4158d01d0b2ad92b1e489abe8312dd543023.
//
// Solidity: event LogDeposit(address indexed account, uint256 amount, bytes32 balance)
func (_Contract *ContractFilterer) WatchLogDeposit(opts *bind.WatchOpts, sink chan<- *ContractLogDeposit, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogDeposit", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogDeposit)
				if err := _Contract.contract.UnpackLog(event, "LogDeposit", log); err != nil {
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

// ParseLogDeposit is a log parse operation binding the contract event 0x40a9cb3a9707d3a68091d8ef7ffd4158d01d0b2ad92b1e489abe8312dd543023.
//
// Solidity: event LogDeposit(address indexed account, uint256 amount, bytes32 balance)
func (_Contract *ContractFilterer) ParseLogDeposit(log types.Log) (*ContractLogDeposit, error) {
	event := new(ContractLogDeposit)
	if err := _Contract.contract.UnpackLog(event, "LogDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogFinalSettlementEnabledIterator is returned from FilterLogFinalSettlementEnabled and is used to iterate over the raw logs and unpacked data for LogFinalSettlementEnabled events raised by the Contract contract.
type ContractLogFinalSettlementEnabledIterator struct {
	Event *ContractLogFinalSettlementEnabled // Event containing the contract specifics and raw log

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
func (it *ContractLogFinalSettlementEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogFinalSettlementEnabled)
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
		it.Event = new(ContractLogFinalSettlementEnabled)
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
func (it *ContractLogFinalSettlementEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogFinalSettlementEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogFinalSettlementEnabled represents a LogFinalSettlementEnabled event raised by the Contract contract.
type ContractLogFinalSettlementEnabled struct {
	SettlementPrice *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogFinalSettlementEnabled is a free log retrieval operation binding the contract event 0x68e4c41627e835051be46337f1542645a60c7e6d6ea79efc5f20bdadae5f88d2.
//
// Solidity: event LogFinalSettlementEnabled(uint256 settlementPrice)
func (_Contract *ContractFilterer) FilterLogFinalSettlementEnabled(opts *bind.FilterOpts) (*ContractLogFinalSettlementEnabledIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogFinalSettlementEnabled")
	if err != nil {
		return nil, err
	}
	return &ContractLogFinalSettlementEnabledIterator{contract: _Contract.contract, event: "LogFinalSettlementEnabled", logs: logs, sub: sub}, nil
}

// WatchLogFinalSettlementEnabled is a free log subscription operation binding the contract event 0x68e4c41627e835051be46337f1542645a60c7e6d6ea79efc5f20bdadae5f88d2.
//
// Solidity: event LogFinalSettlementEnabled(uint256 settlementPrice)
func (_Contract *ContractFilterer) WatchLogFinalSettlementEnabled(opts *bind.WatchOpts, sink chan<- *ContractLogFinalSettlementEnabled) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogFinalSettlementEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogFinalSettlementEnabled)
				if err := _Contract.contract.UnpackLog(event, "LogFinalSettlementEnabled", log); err != nil {
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

// ParseLogFinalSettlementEnabled is a log parse operation binding the contract event 0x68e4c41627e835051be46337f1542645a60c7e6d6ea79efc5f20bdadae5f88d2.
//
// Solidity: event LogFinalSettlementEnabled(uint256 settlementPrice)
func (_Contract *ContractFilterer) ParseLogFinalSettlementEnabled(log types.Log) (*ContractLogFinalSettlementEnabled, error) {
	event := new(ContractLogFinalSettlementEnabled)
	if err := _Contract.contract.UnpackLog(event, "LogFinalSettlementEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogIndexIterator is returned from FilterLogIndex and is used to iterate over the raw logs and unpacked data for LogIndex events raised by the Contract contract.
type ContractLogIndexIterator struct {
	Event *ContractLogIndex // Event containing the contract specifics and raw log

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
func (it *ContractLogIndexIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogIndex)
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
		it.Event = new(ContractLogIndex)
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
func (it *ContractLogIndexIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogIndexIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogIndex represents a LogIndex event raised by the Contract contract.
type ContractLogIndex struct {
	Index [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogIndex is a free log retrieval operation binding the contract event 0x995e61c355733308eab39a59e1e1ac167274cdd1ad707fe4d13e127a01076428.
//
// Solidity: event LogIndex(bytes32 index)
func (_Contract *ContractFilterer) FilterLogIndex(opts *bind.FilterOpts) (*ContractLogIndexIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogIndex")
	if err != nil {
		return nil, err
	}
	return &ContractLogIndexIterator{contract: _Contract.contract, event: "LogIndex", logs: logs, sub: sub}, nil
}

// WatchLogIndex is a free log subscription operation binding the contract event 0x995e61c355733308eab39a59e1e1ac167274cdd1ad707fe4d13e127a01076428.
//
// Solidity: event LogIndex(bytes32 index)
func (_Contract *ContractFilterer) WatchLogIndex(opts *bind.WatchOpts, sink chan<- *ContractLogIndex) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogIndex")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogIndex)
				if err := _Contract.contract.UnpackLog(event, "LogIndex", log); err != nil {
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

// ParseLogIndex is a log parse operation binding the contract event 0x995e61c355733308eab39a59e1e1ac167274cdd1ad707fe4d13e127a01076428.
//
// Solidity: event LogIndex(bytes32 index)
func (_Contract *ContractFilterer) ParseLogIndex(log types.Log) (*ContractLogIndex, error) {
	event := new(ContractLogIndex)
	if err := _Contract.contract.UnpackLog(event, "LogIndex", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogSetFunderIterator is returned from FilterLogSetFunder and is used to iterate over the raw logs and unpacked data for LogSetFunder events raised by the Contract contract.
type ContractLogSetFunderIterator struct {
	Event *ContractLogSetFunder // Event containing the contract specifics and raw log

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
func (it *ContractLogSetFunderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogSetFunder)
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
		it.Event = new(ContractLogSetFunder)
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
func (it *ContractLogSetFunderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogSetFunderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogSetFunder represents a LogSetFunder event raised by the Contract contract.
type ContractLogSetFunder struct {
	Funder common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogSetFunder is a free log retrieval operation binding the contract event 0x433b5c8c9ff78f62114ee8804a916537fa42009ebac4965bfed953f771789e47.
//
// Solidity: event LogSetFunder(address funder)
func (_Contract *ContractFilterer) FilterLogSetFunder(opts *bind.FilterOpts) (*ContractLogSetFunderIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogSetFunder")
	if err != nil {
		return nil, err
	}
	return &ContractLogSetFunderIterator{contract: _Contract.contract, event: "LogSetFunder", logs: logs, sub: sub}, nil
}

// WatchLogSetFunder is a free log subscription operation binding the contract event 0x433b5c8c9ff78f62114ee8804a916537fa42009ebac4965bfed953f771789e47.
//
// Solidity: event LogSetFunder(address funder)
func (_Contract *ContractFilterer) WatchLogSetFunder(opts *bind.WatchOpts, sink chan<- *ContractLogSetFunder) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogSetFunder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogSetFunder)
				if err := _Contract.contract.UnpackLog(event, "LogSetFunder", log); err != nil {
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

// ParseLogSetFunder is a log parse operation binding the contract event 0x433b5c8c9ff78f62114ee8804a916537fa42009ebac4965bfed953f771789e47.
//
// Solidity: event LogSetFunder(address funder)
func (_Contract *ContractFilterer) ParseLogSetFunder(log types.Log) (*ContractLogSetFunder, error) {
	event := new(ContractLogSetFunder)
	if err := _Contract.contract.UnpackLog(event, "LogSetFunder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogSetGlobalOperatorIterator is returned from FilterLogSetGlobalOperator and is used to iterate over the raw logs and unpacked data for LogSetGlobalOperator events raised by the Contract contract.
type ContractLogSetGlobalOperatorIterator struct {
	Event *ContractLogSetGlobalOperator // Event containing the contract specifics and raw log

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
func (it *ContractLogSetGlobalOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogSetGlobalOperator)
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
		it.Event = new(ContractLogSetGlobalOperator)
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
func (it *ContractLogSetGlobalOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogSetGlobalOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogSetGlobalOperator represents a LogSetGlobalOperator event raised by the Contract contract.
type ContractLogSetGlobalOperator struct {
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogSetGlobalOperator is a free log retrieval operation binding the contract event 0xeaeee7699e70e6b31ac89ec999ef6936b97ac1e364f0e1fcf584772372caa0d3.
//
// Solidity: event LogSetGlobalOperator(address operator, bool approved)
func (_Contract *ContractFilterer) FilterLogSetGlobalOperator(opts *bind.FilterOpts) (*ContractLogSetGlobalOperatorIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogSetGlobalOperator")
	if err != nil {
		return nil, err
	}
	return &ContractLogSetGlobalOperatorIterator{contract: _Contract.contract, event: "LogSetGlobalOperator", logs: logs, sub: sub}, nil
}

// WatchLogSetGlobalOperator is a free log subscription operation binding the contract event 0xeaeee7699e70e6b31ac89ec999ef6936b97ac1e364f0e1fcf584772372caa0d3.
//
// Solidity: event LogSetGlobalOperator(address operator, bool approved)
func (_Contract *ContractFilterer) WatchLogSetGlobalOperator(opts *bind.WatchOpts, sink chan<- *ContractLogSetGlobalOperator) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogSetGlobalOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogSetGlobalOperator)
				if err := _Contract.contract.UnpackLog(event, "LogSetGlobalOperator", log); err != nil {
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

// ParseLogSetGlobalOperator is a log parse operation binding the contract event 0xeaeee7699e70e6b31ac89ec999ef6936b97ac1e364f0e1fcf584772372caa0d3.
//
// Solidity: event LogSetGlobalOperator(address operator, bool approved)
func (_Contract *ContractFilterer) ParseLogSetGlobalOperator(log types.Log) (*ContractLogSetGlobalOperator, error) {
	event := new(ContractLogSetGlobalOperator)
	if err := _Contract.contract.UnpackLog(event, "LogSetGlobalOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogSetLocalOperatorIterator is returned from FilterLogSetLocalOperator and is used to iterate over the raw logs and unpacked data for LogSetLocalOperator events raised by the Contract contract.
type ContractLogSetLocalOperatorIterator struct {
	Event *ContractLogSetLocalOperator // Event containing the contract specifics and raw log

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
func (it *ContractLogSetLocalOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogSetLocalOperator)
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
		it.Event = new(ContractLogSetLocalOperator)
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
func (it *ContractLogSetLocalOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogSetLocalOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogSetLocalOperator represents a LogSetLocalOperator event raised by the Contract contract.
type ContractLogSetLocalOperator struct {
	Sender   common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogSetLocalOperator is a free log retrieval operation binding the contract event 0xfe9fa8ad7dbd5e50cbcd1a903ea64717cb80b02e6b737e74f7e2f070b3e4d15f.
//
// Solidity: event LogSetLocalOperator(address indexed sender, address operator, bool approved)
func (_Contract *ContractFilterer) FilterLogSetLocalOperator(opts *bind.FilterOpts, sender []common.Address) (*ContractLogSetLocalOperatorIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogSetLocalOperator", senderRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogSetLocalOperatorIterator{contract: _Contract.contract, event: "LogSetLocalOperator", logs: logs, sub: sub}, nil
}

// WatchLogSetLocalOperator is a free log subscription operation binding the contract event 0xfe9fa8ad7dbd5e50cbcd1a903ea64717cb80b02e6b737e74f7e2f070b3e4d15f.
//
// Solidity: event LogSetLocalOperator(address indexed sender, address operator, bool approved)
func (_Contract *ContractFilterer) WatchLogSetLocalOperator(opts *bind.WatchOpts, sink chan<- *ContractLogSetLocalOperator, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogSetLocalOperator", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogSetLocalOperator)
				if err := _Contract.contract.UnpackLog(event, "LogSetLocalOperator", log); err != nil {
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

// ParseLogSetLocalOperator is a log parse operation binding the contract event 0xfe9fa8ad7dbd5e50cbcd1a903ea64717cb80b02e6b737e74f7e2f070b3e4d15f.
//
// Solidity: event LogSetLocalOperator(address indexed sender, address operator, bool approved)
func (_Contract *ContractFilterer) ParseLogSetLocalOperator(log types.Log) (*ContractLogSetLocalOperator, error) {
	event := new(ContractLogSetLocalOperator)
	if err := _Contract.contract.UnpackLog(event, "LogSetLocalOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogSetMinCollateralIterator is returned from FilterLogSetMinCollateral and is used to iterate over the raw logs and unpacked data for LogSetMinCollateral events raised by the Contract contract.
type ContractLogSetMinCollateralIterator struct {
	Event *ContractLogSetMinCollateral // Event containing the contract specifics and raw log

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
func (it *ContractLogSetMinCollateralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogSetMinCollateral)
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
		it.Event = new(ContractLogSetMinCollateral)
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
func (it *ContractLogSetMinCollateralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogSetMinCollateralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogSetMinCollateral represents a LogSetMinCollateral event raised by the Contract contract.
type ContractLogSetMinCollateral struct {
	MinCollateral *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLogSetMinCollateral is a free log retrieval operation binding the contract event 0x248b36ced4662a14c995e0872f00eb61be4e3dea3913226cdeb513d64728cdca.
//
// Solidity: event LogSetMinCollateral(uint256 minCollateral)
func (_Contract *ContractFilterer) FilterLogSetMinCollateral(opts *bind.FilterOpts) (*ContractLogSetMinCollateralIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogSetMinCollateral")
	if err != nil {
		return nil, err
	}
	return &ContractLogSetMinCollateralIterator{contract: _Contract.contract, event: "LogSetMinCollateral", logs: logs, sub: sub}, nil
}

// WatchLogSetMinCollateral is a free log subscription operation binding the contract event 0x248b36ced4662a14c995e0872f00eb61be4e3dea3913226cdeb513d64728cdca.
//
// Solidity: event LogSetMinCollateral(uint256 minCollateral)
func (_Contract *ContractFilterer) WatchLogSetMinCollateral(opts *bind.WatchOpts, sink chan<- *ContractLogSetMinCollateral) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogSetMinCollateral")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogSetMinCollateral)
				if err := _Contract.contract.UnpackLog(event, "LogSetMinCollateral", log); err != nil {
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

// ParseLogSetMinCollateral is a log parse operation binding the contract event 0x248b36ced4662a14c995e0872f00eb61be4e3dea3913226cdeb513d64728cdca.
//
// Solidity: event LogSetMinCollateral(uint256 minCollateral)
func (_Contract *ContractFilterer) ParseLogSetMinCollateral(log types.Log) (*ContractLogSetMinCollateral, error) {
	event := new(ContractLogSetMinCollateral)
	if err := _Contract.contract.UnpackLog(event, "LogSetMinCollateral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogSetOracleIterator is returned from FilterLogSetOracle and is used to iterate over the raw logs and unpacked data for LogSetOracle events raised by the Contract contract.
type ContractLogSetOracleIterator struct {
	Event *ContractLogSetOracle // Event containing the contract specifics and raw log

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
func (it *ContractLogSetOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogSetOracle)
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
		it.Event = new(ContractLogSetOracle)
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
func (it *ContractLogSetOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogSetOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogSetOracle represents a LogSetOracle event raised by the Contract contract.
type ContractLogSetOracle struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogSetOracle is a free log retrieval operation binding the contract event 0xad675642c3cba5442815383698d42cd28889533d9671a6d32cffea58ef0874da.
//
// Solidity: event LogSetOracle(address oracle)
func (_Contract *ContractFilterer) FilterLogSetOracle(opts *bind.FilterOpts) (*ContractLogSetOracleIterator, error) {

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogSetOracle")
	if err != nil {
		return nil, err
	}
	return &ContractLogSetOracleIterator{contract: _Contract.contract, event: "LogSetOracle", logs: logs, sub: sub}, nil
}

// WatchLogSetOracle is a free log subscription operation binding the contract event 0xad675642c3cba5442815383698d42cd28889533d9671a6d32cffea58ef0874da.
//
// Solidity: event LogSetOracle(address oracle)
func (_Contract *ContractFilterer) WatchLogSetOracle(opts *bind.WatchOpts, sink chan<- *ContractLogSetOracle) (event.Subscription, error) {

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogSetOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogSetOracle)
				if err := _Contract.contract.UnpackLog(event, "LogSetOracle", log); err != nil {
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

// ParseLogSetOracle is a log parse operation binding the contract event 0xad675642c3cba5442815383698d42cd28889533d9671a6d32cffea58ef0874da.
//
// Solidity: event LogSetOracle(address oracle)
func (_Contract *ContractFilterer) ParseLogSetOracle(log types.Log) (*ContractLogSetOracle, error) {
	event := new(ContractLogSetOracle)
	if err := _Contract.contract.UnpackLog(event, "LogSetOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogTradeIterator is returned from FilterLogTrade and is used to iterate over the raw logs and unpacked data for LogTrade events raised by the Contract contract.
type ContractLogTradeIterator struct {
	Event *ContractLogTrade // Event containing the contract specifics and raw log

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
func (it *ContractLogTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogTrade)
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
		it.Event = new(ContractLogTrade)
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
func (it *ContractLogTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogTrade represents a LogTrade event raised by the Contract contract.
type ContractLogTrade struct {
	Maker          common.Address
	Taker          common.Address
	Trader         common.Address
	MarginAmount   *big.Int
	PositionAmount *big.Int
	IsBuy          bool
	MakerBalance   [32]byte
	TakerBalance   [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLogTrade is a free log retrieval operation binding the contract event 0x5171a2ba3550a103fd09ca39b7dcfdf328a5acef18e290c7802d69c8ba73d8d9.
//
// Solidity: event LogTrade(address indexed maker, address indexed taker, address trader, uint256 marginAmount, uint256 positionAmount, bool isBuy, bytes32 makerBalance, bytes32 takerBalance)
func (_Contract *ContractFilterer) FilterLogTrade(opts *bind.FilterOpts, maker []common.Address, taker []common.Address) (*ContractLogTradeIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogTrade", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogTradeIterator{contract: _Contract.contract, event: "LogTrade", logs: logs, sub: sub}, nil
}

// WatchLogTrade is a free log subscription operation binding the contract event 0x5171a2ba3550a103fd09ca39b7dcfdf328a5acef18e290c7802d69c8ba73d8d9.
//
// Solidity: event LogTrade(address indexed maker, address indexed taker, address trader, uint256 marginAmount, uint256 positionAmount, bool isBuy, bytes32 makerBalance, bytes32 takerBalance)
func (_Contract *ContractFilterer) WatchLogTrade(opts *bind.WatchOpts, sink chan<- *ContractLogTrade, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogTrade", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogTrade)
				if err := _Contract.contract.UnpackLog(event, "LogTrade", log); err != nil {
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

// ParseLogTrade is a log parse operation binding the contract event 0x5171a2ba3550a103fd09ca39b7dcfdf328a5acef18e290c7802d69c8ba73d8d9.
//
// Solidity: event LogTrade(address indexed maker, address indexed taker, address trader, uint256 marginAmount, uint256 positionAmount, bool isBuy, bytes32 makerBalance, bytes32 takerBalance)
func (_Contract *ContractFilterer) ParseLogTrade(log types.Log) (*ContractLogTrade, error) {
	event := new(ContractLogTrade)
	if err := _Contract.contract.UnpackLog(event, "LogTrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogWithdrawIterator is returned from FilterLogWithdraw and is used to iterate over the raw logs and unpacked data for LogWithdraw events raised by the Contract contract.
type ContractLogWithdrawIterator struct {
	Event *ContractLogWithdraw // Event containing the contract specifics and raw log

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
func (it *ContractLogWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogWithdraw)
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
		it.Event = new(ContractLogWithdraw)
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
func (it *ContractLogWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogWithdraw represents a LogWithdraw event raised by the Contract contract.
type ContractLogWithdraw struct {
	Account     common.Address
	Destination common.Address
	Amount      *big.Int
	Balance     [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogWithdraw is a free log retrieval operation binding the contract event 0x74348e8cb927b5536fe550310d0cdf05914498fcb04ad61b99c29e3899b0bce9.
//
// Solidity: event LogWithdraw(address indexed account, address destination, uint256 amount, bytes32 balance)
func (_Contract *ContractFilterer) FilterLogWithdraw(opts *bind.FilterOpts, account []common.Address) (*ContractLogWithdrawIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogWithdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogWithdrawIterator{contract: _Contract.contract, event: "LogWithdraw", logs: logs, sub: sub}, nil
}

// WatchLogWithdraw is a free log subscription operation binding the contract event 0x74348e8cb927b5536fe550310d0cdf05914498fcb04ad61b99c29e3899b0bce9.
//
// Solidity: event LogWithdraw(address indexed account, address destination, uint256 amount, bytes32 balance)
func (_Contract *ContractFilterer) WatchLogWithdraw(opts *bind.WatchOpts, sink chan<- *ContractLogWithdraw, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogWithdraw", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogWithdraw)
				if err := _Contract.contract.UnpackLog(event, "LogWithdraw", log); err != nil {
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

// ParseLogWithdraw is a log parse operation binding the contract event 0x74348e8cb927b5536fe550310d0cdf05914498fcb04ad61b99c29e3899b0bce9.
//
// Solidity: event LogWithdraw(address indexed account, address destination, uint256 amount, bytes32 balance)
func (_Contract *ContractFilterer) ParseLogWithdraw(log types.Log) (*ContractLogWithdraw, error) {
	event := new(ContractLogWithdraw)
	if err := _Contract.contract.UnpackLog(event, "LogWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractLogWithdrawFinalSettlementIterator is returned from FilterLogWithdrawFinalSettlement and is used to iterate over the raw logs and unpacked data for LogWithdrawFinalSettlement events raised by the Contract contract.
type ContractLogWithdrawFinalSettlementIterator struct {
	Event *ContractLogWithdrawFinalSettlement // Event containing the contract specifics and raw log

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
func (it *ContractLogWithdrawFinalSettlementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractLogWithdrawFinalSettlement)
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
		it.Event = new(ContractLogWithdrawFinalSettlement)
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
func (it *ContractLogWithdrawFinalSettlementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractLogWithdrawFinalSettlementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractLogWithdrawFinalSettlement represents a LogWithdrawFinalSettlement event raised by the Contract contract.
type ContractLogWithdrawFinalSettlement struct {
	Account common.Address
	Amount  *big.Int
	Balance [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawFinalSettlement is a free log retrieval operation binding the contract event 0xc3b34c584e097adcd5d59ecaf4107928698a4f075c7753b5dbe28cd20d7ac1fd.
//
// Solidity: event LogWithdrawFinalSettlement(address indexed account, uint256 amount, bytes32 balance)
func (_Contract *ContractFilterer) FilterLogWithdrawFinalSettlement(opts *bind.FilterOpts, account []common.Address) (*ContractLogWithdrawFinalSettlementIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "LogWithdrawFinalSettlement", accountRule)
	if err != nil {
		return nil, err
	}
	return &ContractLogWithdrawFinalSettlementIterator{contract: _Contract.contract, event: "LogWithdrawFinalSettlement", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawFinalSettlement is a free log subscription operation binding the contract event 0xc3b34c584e097adcd5d59ecaf4107928698a4f075c7753b5dbe28cd20d7ac1fd.
//
// Solidity: event LogWithdrawFinalSettlement(address indexed account, uint256 amount, bytes32 balance)
func (_Contract *ContractFilterer) WatchLogWithdrawFinalSettlement(opts *bind.WatchOpts, sink chan<- *ContractLogWithdrawFinalSettlement, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "LogWithdrawFinalSettlement", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractLogWithdrawFinalSettlement)
				if err := _Contract.contract.UnpackLog(event, "LogWithdrawFinalSettlement", log); err != nil {
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

// ParseLogWithdrawFinalSettlement is a log parse operation binding the contract event 0xc3b34c584e097adcd5d59ecaf4107928698a4f075c7753b5dbe28cd20d7ac1fd.
//
// Solidity: event LogWithdrawFinalSettlement(address indexed account, uint256 amount, bytes32 balance)
func (_Contract *ContractFilterer) ParseLogWithdrawFinalSettlement(log types.Log) (*ContractLogWithdrawFinalSettlement, error) {
	event := new(ContractLogWithdrawFinalSettlement)
	if err := _Contract.contract.UnpackLog(event, "LogWithdrawFinalSettlement", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
