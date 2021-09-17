// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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

// DistrictContractMetaData contains all meta data concerning the DistrictContract contract.
var DistrictContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"district_id\",\"type\":\"uint256\"}],\"name\":\"DistrictName\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"}],\"name\":\"MetaTransactionExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"x\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"z\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"plotId\",\"type\":\"uint256\"}],\"name\":\"PlotCreation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"origin_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"target_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"plotId\",\"type\":\"uint256\"}],\"name\":\"PlotTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ERC712_VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128[]\",\"name\":\"_xs\",\"type\":\"int128[]\"},{\"internalType\":\"int128[]\",\"name\":\"_zs\",\"type\":\"int128[]\"},{\"internalType\":\"uint256\",\"name\":\"_districtId\",\"type\":\"uint256\"}],\"name\":\"adminClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128[]\",\"name\":\"_xs\",\"type\":\"int128[]\"},{\"internalType\":\"int128[]\",\"name\":\"_zs\",\"type\":\"int128[]\"},{\"internalType\":\"uint256\",\"name\":\"_districtId\",\"type\":\"uint256\"},{\"internalType\":\"bytes24\",\"name\":\"_nickname\",\"type\":\"bytes24\"}],\"name\":\"claimDistrictLands\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"districtNameOf\",\"outputs\":[{\"internalType\":\"bytes24\",\"name\":\"\",\"type\":\"bytes24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"districtPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"functionSignature\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"sigR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"sigS\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"sigV\",\"type\":\"uint8\"}],\"name\":\"executeMetaTransaction\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeperator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOperator\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes24\",\"name\":\"\",\"type\":\"bytes24\"}],\"name\":\"nameDistrictOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plotDistrictOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"},{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"name\":\"plotIdOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plotPriceDistances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plotPrices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plot_x\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"plot_z\",\"outputs\":[{\"internalType\":\"int128\",\"name\":\"\",\"type\":\"int128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_claimable\",\"type\":\"bool\"}],\"name\":\"setClaimable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"district_id\",\"type\":\"uint256\"},{\"internalType\":\"bytes24\",\"name\":\"districtName\",\"type\":\"bytes24\"}],\"name\":\"setDistrictName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_districtPrice\",\"type\":\"uint256\"}],\"name\":\"setDistrictPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_prices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_distances\",\"type\":\"uint256[]\"}],\"name\":\"setPlotPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_worldSize\",\"type\":\"uint128\"}],\"name\":\"setWorldSize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalPlots\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"origin_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"target_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"plot_ids\",\"type\":\"uint256[]\"}],\"name\":\"transferPlot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"worldSize\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DistrictContractABI is the input ABI used to generate the binding from.
// Deprecated: Use DistrictContractMetaData.ABI instead.
var DistrictContractABI = DistrictContractMetaData.ABI

// DistrictContract is an auto generated Go binding around an Ethereum contract.
type DistrictContract struct {
	DistrictContractCaller     // Read-only binding to the contract
	DistrictContractTransactor // Write-only binding to the contract
	DistrictContractFilterer   // Log filterer for contract events
}

// DistrictContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type DistrictContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistrictContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DistrictContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistrictContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistrictContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistrictContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistrictContractSession struct {
	Contract     *DistrictContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DistrictContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistrictContractCallerSession struct {
	Contract *DistrictContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// DistrictContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistrictContractTransactorSession struct {
	Contract     *DistrictContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// DistrictContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type DistrictContractRaw struct {
	Contract *DistrictContract // Generic contract binding to access the raw methods on
}

// DistrictContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistrictContractCallerRaw struct {
	Contract *DistrictContractCaller // Generic read-only contract binding to access the raw methods on
}

// DistrictContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistrictContractTransactorRaw struct {
	Contract *DistrictContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDistrictContract creates a new instance of DistrictContract, bound to a specific deployed contract.
func NewDistrictContract(address common.Address, backend bind.ContractBackend) (*DistrictContract, error) {
	contract, err := bindDistrictContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DistrictContract{DistrictContractCaller: DistrictContractCaller{contract: contract}, DistrictContractTransactor: DistrictContractTransactor{contract: contract}, DistrictContractFilterer: DistrictContractFilterer{contract: contract}}, nil
}

// NewDistrictContractCaller creates a new read-only instance of DistrictContract, bound to a specific deployed contract.
func NewDistrictContractCaller(address common.Address, caller bind.ContractCaller) (*DistrictContractCaller, error) {
	contract, err := bindDistrictContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistrictContractCaller{contract: contract}, nil
}

// NewDistrictContractTransactor creates a new write-only instance of DistrictContract, bound to a specific deployed contract.
func NewDistrictContractTransactor(address common.Address, transactor bind.ContractTransactor) (*DistrictContractTransactor, error) {
	contract, err := bindDistrictContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistrictContractTransactor{contract: contract}, nil
}

// NewDistrictContractFilterer creates a new log filterer instance of DistrictContract, bound to a specific deployed contract.
func NewDistrictContractFilterer(address common.Address, filterer bind.ContractFilterer) (*DistrictContractFilterer, error) {
	contract, err := bindDistrictContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistrictContractFilterer{contract: contract}, nil
}

// bindDistrictContract binds a generic wrapper to an already deployed contract.
func bindDistrictContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DistrictContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DistrictContract *DistrictContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DistrictContract.Contract.DistrictContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DistrictContract *DistrictContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistrictContract.Contract.DistrictContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DistrictContract *DistrictContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DistrictContract.Contract.DistrictContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DistrictContract *DistrictContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DistrictContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DistrictContract *DistrictContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistrictContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DistrictContract *DistrictContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DistrictContract.Contract.contract.Transact(opts, method, params...)
}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_DistrictContract *DistrictContractCaller) ERC712VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "ERC712_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_DistrictContract *DistrictContractSession) ERC712VERSION() (string, error) {
	return _DistrictContract.Contract.ERC712VERSION(&_DistrictContract.CallOpts)
}

// ERC712VERSION is a free data retrieval call binding the contract method 0x0f7e5970.
//
// Solidity: function ERC712_VERSION() view returns(string)
func (_DistrictContract *DistrictContractCallerSession) ERC712VERSION() (string, error) {
	return _DistrictContract.Contract.ERC712VERSION(&_DistrictContract.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DistrictContract *DistrictContractCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DistrictContract *DistrictContractSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DistrictContract.Contract.BalanceOf(&_DistrictContract.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_DistrictContract *DistrictContractCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _DistrictContract.Contract.BalanceOf(&_DistrictContract.CallOpts, owner)
}

// Claimable is a free data retrieval call binding the contract method 0xaf38d757.
//
// Solidity: function claimable() view returns(bool)
func (_DistrictContract *DistrictContractCaller) Claimable(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "claimable")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0xaf38d757.
//
// Solidity: function claimable() view returns(bool)
func (_DistrictContract *DistrictContractSession) Claimable() (bool, error) {
	return _DistrictContract.Contract.Claimable(&_DistrictContract.CallOpts)
}

// Claimable is a free data retrieval call binding the contract method 0xaf38d757.
//
// Solidity: function claimable() view returns(bool)
func (_DistrictContract *DistrictContractCallerSession) Claimable() (bool, error) {
	return _DistrictContract.Contract.Claimable(&_DistrictContract.CallOpts)
}

// DistrictNameOf is a free data retrieval call binding the contract method 0x779fe02c.
//
// Solidity: function districtNameOf(uint256 ) view returns(bytes24)
func (_DistrictContract *DistrictContractCaller) DistrictNameOf(opts *bind.CallOpts, arg0 *big.Int) ([24]byte, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "districtNameOf", arg0)

	if err != nil {
		return *new([24]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([24]byte)).(*[24]byte)

	return out0, err

}

// DistrictNameOf is a free data retrieval call binding the contract method 0x779fe02c.
//
// Solidity: function districtNameOf(uint256 ) view returns(bytes24)
func (_DistrictContract *DistrictContractSession) DistrictNameOf(arg0 *big.Int) ([24]byte, error) {
	return _DistrictContract.Contract.DistrictNameOf(&_DistrictContract.CallOpts, arg0)
}

// DistrictNameOf is a free data retrieval call binding the contract method 0x779fe02c.
//
// Solidity: function districtNameOf(uint256 ) view returns(bytes24)
func (_DistrictContract *DistrictContractCallerSession) DistrictNameOf(arg0 *big.Int) ([24]byte, error) {
	return _DistrictContract.Contract.DistrictNameOf(&_DistrictContract.CallOpts, arg0)
}

// DistrictPrice is a free data retrieval call binding the contract method 0xfd58c368.
//
// Solidity: function districtPrice() view returns(uint256)
func (_DistrictContract *DistrictContractCaller) DistrictPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "districtPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DistrictPrice is a free data retrieval call binding the contract method 0xfd58c368.
//
// Solidity: function districtPrice() view returns(uint256)
func (_DistrictContract *DistrictContractSession) DistrictPrice() (*big.Int, error) {
	return _DistrictContract.Contract.DistrictPrice(&_DistrictContract.CallOpts)
}

// DistrictPrice is a free data retrieval call binding the contract method 0xfd58c368.
//
// Solidity: function districtPrice() view returns(uint256)
func (_DistrictContract *DistrictContractCallerSession) DistrictPrice() (*big.Int, error) {
	return _DistrictContract.Contract.DistrictPrice(&_DistrictContract.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DistrictContract *DistrictContractCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DistrictContract *DistrictContractSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DistrictContract.Contract.GetApproved(&_DistrictContract.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_DistrictContract *DistrictContractCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _DistrictContract.Contract.GetApproved(&_DistrictContract.CallOpts, tokenId)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_DistrictContract *DistrictContractCaller) GetChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "getChainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_DistrictContract *DistrictContractSession) GetChainId() (*big.Int, error) {
	return _DistrictContract.Contract.GetChainId(&_DistrictContract.CallOpts)
}

// GetChainId is a free data retrieval call binding the contract method 0x3408e470.
//
// Solidity: function getChainId() view returns(uint256)
func (_DistrictContract *DistrictContractCallerSession) GetChainId() (*big.Int, error) {
	return _DistrictContract.Contract.GetChainId(&_DistrictContract.CallOpts)
}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_DistrictContract *DistrictContractCaller) GetDomainSeperator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "getDomainSeperator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_DistrictContract *DistrictContractSession) GetDomainSeperator() ([32]byte, error) {
	return _DistrictContract.Contract.GetDomainSeperator(&_DistrictContract.CallOpts)
}

// GetDomainSeperator is a free data retrieval call binding the contract method 0x20379ee5.
//
// Solidity: function getDomainSeperator() view returns(bytes32)
func (_DistrictContract *DistrictContractCallerSession) GetDomainSeperator() ([32]byte, error) {
	return _DistrictContract.Contract.GetDomainSeperator(&_DistrictContract.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_DistrictContract *DistrictContractCaller) GetNonce(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "getNonce", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_DistrictContract *DistrictContractSession) GetNonce(user common.Address) (*big.Int, error) {
	return _DistrictContract.Contract.GetNonce(&_DistrictContract.CallOpts, user)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address user) view returns(uint256 nonce)
func (_DistrictContract *DistrictContractCallerSession) GetNonce(user common.Address) (*big.Int, error) {
	return _DistrictContract.Contract.GetNonce(&_DistrictContract.CallOpts, user)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_DistrictContract *DistrictContractCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "isApprovedForAll", _owner, _operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_DistrictContract *DistrictContractSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _DistrictContract.Contract.IsApprovedForAll(&_DistrictContract.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address _owner, address _operator) view returns(bool isOperator)
func (_DistrictContract *DistrictContractCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _DistrictContract.Contract.IsApprovedForAll(&_DistrictContract.CallOpts, _owner, _operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DistrictContract *DistrictContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DistrictContract *DistrictContractSession) Name() (string, error) {
	return _DistrictContract.Contract.Name(&_DistrictContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_DistrictContract *DistrictContractCallerSession) Name() (string, error) {
	return _DistrictContract.Contract.Name(&_DistrictContract.CallOpts)
}

// NameDistrictOf is a free data retrieval call binding the contract method 0xe98e860c.
//
// Solidity: function nameDistrictOf(bytes24 ) view returns(uint256)
func (_DistrictContract *DistrictContractCaller) NameDistrictOf(opts *bind.CallOpts, arg0 [24]byte) (*big.Int, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "nameDistrictOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NameDistrictOf is a free data retrieval call binding the contract method 0xe98e860c.
//
// Solidity: function nameDistrictOf(bytes24 ) view returns(uint256)
func (_DistrictContract *DistrictContractSession) NameDistrictOf(arg0 [24]byte) (*big.Int, error) {
	return _DistrictContract.Contract.NameDistrictOf(&_DistrictContract.CallOpts, arg0)
}

// NameDistrictOf is a free data retrieval call binding the contract method 0xe98e860c.
//
// Solidity: function nameDistrictOf(bytes24 ) view returns(uint256)
func (_DistrictContract *DistrictContractCallerSession) NameDistrictOf(arg0 [24]byte) (*big.Int, error) {
	return _DistrictContract.Contract.NameDistrictOf(&_DistrictContract.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DistrictContract *DistrictContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DistrictContract *DistrictContractSession) Owner() (common.Address, error) {
	return _DistrictContract.Contract.Owner(&_DistrictContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DistrictContract *DistrictContractCallerSession) Owner() (common.Address, error) {
	return _DistrictContract.Contract.Owner(&_DistrictContract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DistrictContract *DistrictContractCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DistrictContract *DistrictContractSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DistrictContract.Contract.OwnerOf(&_DistrictContract.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_DistrictContract *DistrictContractCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _DistrictContract.Contract.OwnerOf(&_DistrictContract.CallOpts, tokenId)
}

// PlotDistrictOf is a free data retrieval call binding the contract method 0x2d276f7a.
//
// Solidity: function plotDistrictOf(uint256 ) view returns(uint256)
func (_DistrictContract *DistrictContractCaller) PlotDistrictOf(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "plotDistrictOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlotDistrictOf is a free data retrieval call binding the contract method 0x2d276f7a.
//
// Solidity: function plotDistrictOf(uint256 ) view returns(uint256)
func (_DistrictContract *DistrictContractSession) PlotDistrictOf(arg0 *big.Int) (*big.Int, error) {
	return _DistrictContract.Contract.PlotDistrictOf(&_DistrictContract.CallOpts, arg0)
}

// PlotDistrictOf is a free data retrieval call binding the contract method 0x2d276f7a.
//
// Solidity: function plotDistrictOf(uint256 ) view returns(uint256)
func (_DistrictContract *DistrictContractCallerSession) PlotDistrictOf(arg0 *big.Int) (*big.Int, error) {
	return _DistrictContract.Contract.PlotDistrictOf(&_DistrictContract.CallOpts, arg0)
}

// PlotIdOf is a free data retrieval call binding the contract method 0x316859aa.
//
// Solidity: function plotIdOf(int128 , int128 ) view returns(uint256)
func (_DistrictContract *DistrictContractCaller) PlotIdOf(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "plotIdOf", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlotIdOf is a free data retrieval call binding the contract method 0x316859aa.
//
// Solidity: function plotIdOf(int128 , int128 ) view returns(uint256)
func (_DistrictContract *DistrictContractSession) PlotIdOf(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _DistrictContract.Contract.PlotIdOf(&_DistrictContract.CallOpts, arg0, arg1)
}

// PlotIdOf is a free data retrieval call binding the contract method 0x316859aa.
//
// Solidity: function plotIdOf(int128 , int128 ) view returns(uint256)
func (_DistrictContract *DistrictContractCallerSession) PlotIdOf(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _DistrictContract.Contract.PlotIdOf(&_DistrictContract.CallOpts, arg0, arg1)
}

// PlotPriceDistances is a free data retrieval call binding the contract method 0x1b81c552.
//
// Solidity: function plotPriceDistances(uint256 ) view returns(uint256)
func (_DistrictContract *DistrictContractCaller) PlotPriceDistances(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "plotPriceDistances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlotPriceDistances is a free data retrieval call binding the contract method 0x1b81c552.
//
// Solidity: function plotPriceDistances(uint256 ) view returns(uint256)
func (_DistrictContract *DistrictContractSession) PlotPriceDistances(arg0 *big.Int) (*big.Int, error) {
	return _DistrictContract.Contract.PlotPriceDistances(&_DistrictContract.CallOpts, arg0)
}

// PlotPriceDistances is a free data retrieval call binding the contract method 0x1b81c552.
//
// Solidity: function plotPriceDistances(uint256 ) view returns(uint256)
func (_DistrictContract *DistrictContractCallerSession) PlotPriceDistances(arg0 *big.Int) (*big.Int, error) {
	return _DistrictContract.Contract.PlotPriceDistances(&_DistrictContract.CallOpts, arg0)
}

// PlotPrices is a free data retrieval call binding the contract method 0xa65e6c0d.
//
// Solidity: function plotPrices(uint256 ) view returns(uint256)
func (_DistrictContract *DistrictContractCaller) PlotPrices(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "plotPrices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlotPrices is a free data retrieval call binding the contract method 0xa65e6c0d.
//
// Solidity: function plotPrices(uint256 ) view returns(uint256)
func (_DistrictContract *DistrictContractSession) PlotPrices(arg0 *big.Int) (*big.Int, error) {
	return _DistrictContract.Contract.PlotPrices(&_DistrictContract.CallOpts, arg0)
}

// PlotPrices is a free data retrieval call binding the contract method 0xa65e6c0d.
//
// Solidity: function plotPrices(uint256 ) view returns(uint256)
func (_DistrictContract *DistrictContractCallerSession) PlotPrices(arg0 *big.Int) (*big.Int, error) {
	return _DistrictContract.Contract.PlotPrices(&_DistrictContract.CallOpts, arg0)
}

// PlotX is a free data retrieval call binding the contract method 0xb1e5ec99.
//
// Solidity: function plot_x(uint256 ) view returns(int128)
func (_DistrictContract *DistrictContractCaller) PlotX(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "plot_x", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlotX is a free data retrieval call binding the contract method 0xb1e5ec99.
//
// Solidity: function plot_x(uint256 ) view returns(int128)
func (_DistrictContract *DistrictContractSession) PlotX(arg0 *big.Int) (*big.Int, error) {
	return _DistrictContract.Contract.PlotX(&_DistrictContract.CallOpts, arg0)
}

// PlotX is a free data retrieval call binding the contract method 0xb1e5ec99.
//
// Solidity: function plot_x(uint256 ) view returns(int128)
func (_DistrictContract *DistrictContractCallerSession) PlotX(arg0 *big.Int) (*big.Int, error) {
	return _DistrictContract.Contract.PlotX(&_DistrictContract.CallOpts, arg0)
}

// PlotZ is a free data retrieval call binding the contract method 0xa96c4485.
//
// Solidity: function plot_z(uint256 ) view returns(int128)
func (_DistrictContract *DistrictContractCaller) PlotZ(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "plot_z", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PlotZ is a free data retrieval call binding the contract method 0xa96c4485.
//
// Solidity: function plot_z(uint256 ) view returns(int128)
func (_DistrictContract *DistrictContractSession) PlotZ(arg0 *big.Int) (*big.Int, error) {
	return _DistrictContract.Contract.PlotZ(&_DistrictContract.CallOpts, arg0)
}

// PlotZ is a free data retrieval call binding the contract method 0xa96c4485.
//
// Solidity: function plot_z(uint256 ) view returns(int128)
func (_DistrictContract *DistrictContractCallerSession) PlotZ(arg0 *big.Int) (*big.Int, error) {
	return _DistrictContract.Contract.PlotZ(&_DistrictContract.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DistrictContract *DistrictContractCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DistrictContract *DistrictContractSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DistrictContract.Contract.SupportsInterface(&_DistrictContract.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_DistrictContract *DistrictContractCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _DistrictContract.Contract.SupportsInterface(&_DistrictContract.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DistrictContract *DistrictContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DistrictContract *DistrictContractSession) Symbol() (string, error) {
	return _DistrictContract.Contract.Symbol(&_DistrictContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_DistrictContract *DistrictContractCallerSession) Symbol() (string, error) {
	return _DistrictContract.Contract.Symbol(&_DistrictContract.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DistrictContract *DistrictContractCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DistrictContract *DistrictContractSession) TokenURI(tokenId *big.Int) (string, error) {
	return _DistrictContract.Contract.TokenURI(&_DistrictContract.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_DistrictContract *DistrictContractCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _DistrictContract.Contract.TokenURI(&_DistrictContract.CallOpts, tokenId)
}

// TotalPlots is a free data retrieval call binding the contract method 0x52158f35.
//
// Solidity: function totalPlots() view returns(uint256)
func (_DistrictContract *DistrictContractCaller) TotalPlots(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "totalPlots")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalPlots is a free data retrieval call binding the contract method 0x52158f35.
//
// Solidity: function totalPlots() view returns(uint256)
func (_DistrictContract *DistrictContractSession) TotalPlots() (*big.Int, error) {
	return _DistrictContract.Contract.TotalPlots(&_DistrictContract.CallOpts)
}

// TotalPlots is a free data retrieval call binding the contract method 0x52158f35.
//
// Solidity: function totalPlots() view returns(uint256)
func (_DistrictContract *DistrictContractCallerSession) TotalPlots() (*big.Int, error) {
	return _DistrictContract.Contract.TotalPlots(&_DistrictContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DistrictContract *DistrictContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DistrictContract *DistrictContractSession) TotalSupply() (*big.Int, error) {
	return _DistrictContract.Contract.TotalSupply(&_DistrictContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_DistrictContract *DistrictContractCallerSession) TotalSupply() (*big.Int, error) {
	return _DistrictContract.Contract.TotalSupply(&_DistrictContract.CallOpts)
}

// WorldSize is a free data retrieval call binding the contract method 0x5677afbc.
//
// Solidity: function worldSize() view returns(uint128)
func (_DistrictContract *DistrictContractCaller) WorldSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DistrictContract.contract.Call(opts, &out, "worldSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WorldSize is a free data retrieval call binding the contract method 0x5677afbc.
//
// Solidity: function worldSize() view returns(uint128)
func (_DistrictContract *DistrictContractSession) WorldSize() (*big.Int, error) {
	return _DistrictContract.Contract.WorldSize(&_DistrictContract.CallOpts)
}

// WorldSize is a free data retrieval call binding the contract method 0x5677afbc.
//
// Solidity: function worldSize() view returns(uint128)
func (_DistrictContract *DistrictContractCallerSession) WorldSize() (*big.Int, error) {
	return _DistrictContract.Contract.WorldSize(&_DistrictContract.CallOpts)
}

// AdminClaim is a paid mutator transaction binding the contract method 0xb729daa5.
//
// Solidity: function adminClaim(int128[] _xs, int128[] _zs, uint256 _districtId) returns()
func (_DistrictContract *DistrictContractTransactor) AdminClaim(opts *bind.TransactOpts, _xs []*big.Int, _zs []*big.Int, _districtId *big.Int) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "adminClaim", _xs, _zs, _districtId)
}

// AdminClaim is a paid mutator transaction binding the contract method 0xb729daa5.
//
// Solidity: function adminClaim(int128[] _xs, int128[] _zs, uint256 _districtId) returns()
func (_DistrictContract *DistrictContractSession) AdminClaim(_xs []*big.Int, _zs []*big.Int, _districtId *big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.AdminClaim(&_DistrictContract.TransactOpts, _xs, _zs, _districtId)
}

// AdminClaim is a paid mutator transaction binding the contract method 0xb729daa5.
//
// Solidity: function adminClaim(int128[] _xs, int128[] _zs, uint256 _districtId) returns()
func (_DistrictContract *DistrictContractTransactorSession) AdminClaim(_xs []*big.Int, _zs []*big.Int, _districtId *big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.AdminClaim(&_DistrictContract.TransactOpts, _xs, _zs, _districtId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DistrictContract *DistrictContractTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DistrictContract *DistrictContractSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.Approve(&_DistrictContract.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_DistrictContract *DistrictContractTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.Approve(&_DistrictContract.TransactOpts, to, tokenId)
}

// ClaimDistrictLands is a paid mutator transaction binding the contract method 0x4bf1e0e2.
//
// Solidity: function claimDistrictLands(int128[] _xs, int128[] _zs, uint256 _districtId, bytes24 _nickname) payable returns()
func (_DistrictContract *DistrictContractTransactor) ClaimDistrictLands(opts *bind.TransactOpts, _xs []*big.Int, _zs []*big.Int, _districtId *big.Int, _nickname [24]byte) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "claimDistrictLands", _xs, _zs, _districtId, _nickname)
}

// ClaimDistrictLands is a paid mutator transaction binding the contract method 0x4bf1e0e2.
//
// Solidity: function claimDistrictLands(int128[] _xs, int128[] _zs, uint256 _districtId, bytes24 _nickname) payable returns()
func (_DistrictContract *DistrictContractSession) ClaimDistrictLands(_xs []*big.Int, _zs []*big.Int, _districtId *big.Int, _nickname [24]byte) (*types.Transaction, error) {
	return _DistrictContract.Contract.ClaimDistrictLands(&_DistrictContract.TransactOpts, _xs, _zs, _districtId, _nickname)
}

// ClaimDistrictLands is a paid mutator transaction binding the contract method 0x4bf1e0e2.
//
// Solidity: function claimDistrictLands(int128[] _xs, int128[] _zs, uint256 _districtId, bytes24 _nickname) payable returns()
func (_DistrictContract *DistrictContractTransactorSession) ClaimDistrictLands(_xs []*big.Int, _zs []*big.Int, _districtId *big.Int, _nickname [24]byte) (*types.Transaction, error) {
	return _DistrictContract.Contract.ClaimDistrictLands(&_DistrictContract.TransactOpts, _xs, _zs, _districtId, _nickname)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_DistrictContract *DistrictContractTransactor) ExecuteMetaTransaction(opts *bind.TransactOpts, userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "executeMetaTransaction", userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_DistrictContract *DistrictContractSession) ExecuteMetaTransaction(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _DistrictContract.Contract.ExecuteMetaTransaction(&_DistrictContract.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// ExecuteMetaTransaction is a paid mutator transaction binding the contract method 0x0c53c51c.
//
// Solidity: function executeMetaTransaction(address userAddress, bytes functionSignature, bytes32 sigR, bytes32 sigS, uint8 sigV) payable returns(bytes)
func (_DistrictContract *DistrictContractTransactorSession) ExecuteMetaTransaction(userAddress common.Address, functionSignature []byte, sigR [32]byte, sigS [32]byte, sigV uint8) (*types.Transaction, error) {
	return _DistrictContract.Contract.ExecuteMetaTransaction(&_DistrictContract.TransactOpts, userAddress, functionSignature, sigR, sigS, sigV)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_DistrictContract *DistrictContractTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "initialize", _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_DistrictContract *DistrictContractSession) Initialize(_name string, _symbol string) (*types.Transaction, error) {
	return _DistrictContract.Contract.Initialize(&_DistrictContract.TransactOpts, _name, _symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string _name, string _symbol) returns()
func (_DistrictContract *DistrictContractTransactorSession) Initialize(_name string, _symbol string) (*types.Transaction, error) {
	return _DistrictContract.Contract.Initialize(&_DistrictContract.TransactOpts, _name, _symbol)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DistrictContract *DistrictContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DistrictContract *DistrictContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _DistrictContract.Contract.RenounceOwnership(&_DistrictContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DistrictContract *DistrictContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DistrictContract.Contract.RenounceOwnership(&_DistrictContract.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DistrictContract *DistrictContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DistrictContract *DistrictContractSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.SafeTransferFrom(&_DistrictContract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_DistrictContract *DistrictContractTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.SafeTransferFrom(&_DistrictContract.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_DistrictContract *DistrictContractTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_DistrictContract *DistrictContractSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _DistrictContract.Contract.SafeTransferFrom0(&_DistrictContract.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_DistrictContract *DistrictContractTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _DistrictContract.Contract.SafeTransferFrom0(&_DistrictContract.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DistrictContract *DistrictContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DistrictContract *DistrictContractSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DistrictContract.Contract.SetApprovalForAll(&_DistrictContract.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_DistrictContract *DistrictContractTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _DistrictContract.Contract.SetApprovalForAll(&_DistrictContract.TransactOpts, operator, approved)
}

// SetClaimable is a paid mutator transaction binding the contract method 0x378c93ad.
//
// Solidity: function setClaimable(bool _claimable) returns()
func (_DistrictContract *DistrictContractTransactor) SetClaimable(opts *bind.TransactOpts, _claimable bool) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "setClaimable", _claimable)
}

// SetClaimable is a paid mutator transaction binding the contract method 0x378c93ad.
//
// Solidity: function setClaimable(bool _claimable) returns()
func (_DistrictContract *DistrictContractSession) SetClaimable(_claimable bool) (*types.Transaction, error) {
	return _DistrictContract.Contract.SetClaimable(&_DistrictContract.TransactOpts, _claimable)
}

// SetClaimable is a paid mutator transaction binding the contract method 0x378c93ad.
//
// Solidity: function setClaimable(bool _claimable) returns()
func (_DistrictContract *DistrictContractTransactorSession) SetClaimable(_claimable bool) (*types.Transaction, error) {
	return _DistrictContract.Contract.SetClaimable(&_DistrictContract.TransactOpts, _claimable)
}

// SetDistrictName is a paid mutator transaction binding the contract method 0x960d82ae.
//
// Solidity: function setDistrictName(uint256 district_id, bytes24 districtName) returns()
func (_DistrictContract *DistrictContractTransactor) SetDistrictName(opts *bind.TransactOpts, district_id *big.Int, districtName [24]byte) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "setDistrictName", district_id, districtName)
}

// SetDistrictName is a paid mutator transaction binding the contract method 0x960d82ae.
//
// Solidity: function setDistrictName(uint256 district_id, bytes24 districtName) returns()
func (_DistrictContract *DistrictContractSession) SetDistrictName(district_id *big.Int, districtName [24]byte) (*types.Transaction, error) {
	return _DistrictContract.Contract.SetDistrictName(&_DistrictContract.TransactOpts, district_id, districtName)
}

// SetDistrictName is a paid mutator transaction binding the contract method 0x960d82ae.
//
// Solidity: function setDistrictName(uint256 district_id, bytes24 districtName) returns()
func (_DistrictContract *DistrictContractTransactorSession) SetDistrictName(district_id *big.Int, districtName [24]byte) (*types.Transaction, error) {
	return _DistrictContract.Contract.SetDistrictName(&_DistrictContract.TransactOpts, district_id, districtName)
}

// SetDistrictPrice is a paid mutator transaction binding the contract method 0x8c32f1f1.
//
// Solidity: function setDistrictPrice(uint256 _districtPrice) returns()
func (_DistrictContract *DistrictContractTransactor) SetDistrictPrice(opts *bind.TransactOpts, _districtPrice *big.Int) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "setDistrictPrice", _districtPrice)
}

// SetDistrictPrice is a paid mutator transaction binding the contract method 0x8c32f1f1.
//
// Solidity: function setDistrictPrice(uint256 _districtPrice) returns()
func (_DistrictContract *DistrictContractSession) SetDistrictPrice(_districtPrice *big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.SetDistrictPrice(&_DistrictContract.TransactOpts, _districtPrice)
}

// SetDistrictPrice is a paid mutator transaction binding the contract method 0x8c32f1f1.
//
// Solidity: function setDistrictPrice(uint256 _districtPrice) returns()
func (_DistrictContract *DistrictContractTransactorSession) SetDistrictPrice(_districtPrice *big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.SetDistrictPrice(&_DistrictContract.TransactOpts, _districtPrice)
}

// SetPlotPrices is a paid mutator transaction binding the contract method 0xf6552d2f.
//
// Solidity: function setPlotPrices(uint256[] _prices, uint256[] _distances) returns()
func (_DistrictContract *DistrictContractTransactor) SetPlotPrices(opts *bind.TransactOpts, _prices []*big.Int, _distances []*big.Int) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "setPlotPrices", _prices, _distances)
}

// SetPlotPrices is a paid mutator transaction binding the contract method 0xf6552d2f.
//
// Solidity: function setPlotPrices(uint256[] _prices, uint256[] _distances) returns()
func (_DistrictContract *DistrictContractSession) SetPlotPrices(_prices []*big.Int, _distances []*big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.SetPlotPrices(&_DistrictContract.TransactOpts, _prices, _distances)
}

// SetPlotPrices is a paid mutator transaction binding the contract method 0xf6552d2f.
//
// Solidity: function setPlotPrices(uint256[] _prices, uint256[] _distances) returns()
func (_DistrictContract *DistrictContractTransactorSession) SetPlotPrices(_prices []*big.Int, _distances []*big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.SetPlotPrices(&_DistrictContract.TransactOpts, _prices, _distances)
}

// SetWorldSize is a paid mutator transaction binding the contract method 0xd7b5b567.
//
// Solidity: function setWorldSize(uint128 _worldSize) returns()
func (_DistrictContract *DistrictContractTransactor) SetWorldSize(opts *bind.TransactOpts, _worldSize *big.Int) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "setWorldSize", _worldSize)
}

// SetWorldSize is a paid mutator transaction binding the contract method 0xd7b5b567.
//
// Solidity: function setWorldSize(uint128 _worldSize) returns()
func (_DistrictContract *DistrictContractSession) SetWorldSize(_worldSize *big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.SetWorldSize(&_DistrictContract.TransactOpts, _worldSize)
}

// SetWorldSize is a paid mutator transaction binding the contract method 0xd7b5b567.
//
// Solidity: function setWorldSize(uint128 _worldSize) returns()
func (_DistrictContract *DistrictContractTransactorSession) SetWorldSize(_worldSize *big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.SetWorldSize(&_DistrictContract.TransactOpts, _worldSize)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DistrictContract *DistrictContractTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DistrictContract *DistrictContractSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.TransferFrom(&_DistrictContract.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_DistrictContract *DistrictContractTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.TransferFrom(&_DistrictContract.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DistrictContract *DistrictContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DistrictContract *DistrictContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DistrictContract.Contract.TransferOwnership(&_DistrictContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DistrictContract *DistrictContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DistrictContract.Contract.TransferOwnership(&_DistrictContract.TransactOpts, newOwner)
}

// TransferPlot is a paid mutator transaction binding the contract method 0x69ae1456.
//
// Solidity: function transferPlot(uint256 origin_id, uint256 target_id, uint256[] plot_ids) returns()
func (_DistrictContract *DistrictContractTransactor) TransferPlot(opts *bind.TransactOpts, origin_id *big.Int, target_id *big.Int, plot_ids []*big.Int) (*types.Transaction, error) {
	return _DistrictContract.contract.Transact(opts, "transferPlot", origin_id, target_id, plot_ids)
}

// TransferPlot is a paid mutator transaction binding the contract method 0x69ae1456.
//
// Solidity: function transferPlot(uint256 origin_id, uint256 target_id, uint256[] plot_ids) returns()
func (_DistrictContract *DistrictContractSession) TransferPlot(origin_id *big.Int, target_id *big.Int, plot_ids []*big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.TransferPlot(&_DistrictContract.TransactOpts, origin_id, target_id, plot_ids)
}

// TransferPlot is a paid mutator transaction binding the contract method 0x69ae1456.
//
// Solidity: function transferPlot(uint256 origin_id, uint256 target_id, uint256[] plot_ids) returns()
func (_DistrictContract *DistrictContractTransactorSession) TransferPlot(origin_id *big.Int, target_id *big.Int, plot_ids []*big.Int) (*types.Transaction, error) {
	return _DistrictContract.Contract.TransferPlot(&_DistrictContract.TransactOpts, origin_id, target_id, plot_ids)
}

// DistrictContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the DistrictContract contract.
type DistrictContractApprovalIterator struct {
	Event *DistrictContractApproval // Event containing the contract specifics and raw log

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
func (it *DistrictContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistrictContractApproval)
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
		it.Event = new(DistrictContractApproval)
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
func (it *DistrictContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistrictContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistrictContractApproval represents a Approval event raised by the DistrictContract contract.
type DistrictContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DistrictContract *DistrictContractFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*DistrictContractApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DistrictContract.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DistrictContractApprovalIterator{contract: _DistrictContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DistrictContract *DistrictContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DistrictContractApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DistrictContract.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistrictContractApproval)
				if err := _DistrictContract.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_DistrictContract *DistrictContractFilterer) ParseApproval(log types.Log) (*DistrictContractApproval, error) {
	event := new(DistrictContractApproval)
	if err := _DistrictContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistrictContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the DistrictContract contract.
type DistrictContractApprovalForAllIterator struct {
	Event *DistrictContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *DistrictContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistrictContractApprovalForAll)
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
		it.Event = new(DistrictContractApprovalForAll)
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
func (it *DistrictContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistrictContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistrictContractApprovalForAll represents a ApprovalForAll event raised by the DistrictContract contract.
type DistrictContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DistrictContract *DistrictContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*DistrictContractApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DistrictContract.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &DistrictContractApprovalForAllIterator{contract: _DistrictContract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DistrictContract *DistrictContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *DistrictContractApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _DistrictContract.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistrictContractApprovalForAll)
				if err := _DistrictContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_DistrictContract *DistrictContractFilterer) ParseApprovalForAll(log types.Log) (*DistrictContractApprovalForAll, error) {
	event := new(DistrictContractApprovalForAll)
	if err := _DistrictContract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistrictContractDistrictNameIterator is returned from FilterDistrictName and is used to iterate over the raw logs and unpacked data for DistrictName events raised by the DistrictContract contract.
type DistrictContractDistrictNameIterator struct {
	Event *DistrictContractDistrictName // Event containing the contract specifics and raw log

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
func (it *DistrictContractDistrictNameIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistrictContractDistrictName)
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
		it.Event = new(DistrictContractDistrictName)
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
func (it *DistrictContractDistrictNameIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistrictContractDistrictNameIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistrictContractDistrictName represents a DistrictName event raised by the DistrictContract contract.
type DistrictContractDistrictName struct {
	DistrictId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDistrictName is a free log retrieval operation binding the contract event 0x257242c53ca9802d7b970b3ae41707c259e0b36ccce43bc618cfe649768642b3.
//
// Solidity: event DistrictName(uint256 district_id)
func (_DistrictContract *DistrictContractFilterer) FilterDistrictName(opts *bind.FilterOpts) (*DistrictContractDistrictNameIterator, error) {

	logs, sub, err := _DistrictContract.contract.FilterLogs(opts, "DistrictName")
	if err != nil {
		return nil, err
	}
	return &DistrictContractDistrictNameIterator{contract: _DistrictContract.contract, event: "DistrictName", logs: logs, sub: sub}, nil
}

// WatchDistrictName is a free log subscription operation binding the contract event 0x257242c53ca9802d7b970b3ae41707c259e0b36ccce43bc618cfe649768642b3.
//
// Solidity: event DistrictName(uint256 district_id)
func (_DistrictContract *DistrictContractFilterer) WatchDistrictName(opts *bind.WatchOpts, sink chan<- *DistrictContractDistrictName) (event.Subscription, error) {

	logs, sub, err := _DistrictContract.contract.WatchLogs(opts, "DistrictName")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistrictContractDistrictName)
				if err := _DistrictContract.contract.UnpackLog(event, "DistrictName", log); err != nil {
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

// ParseDistrictName is a log parse operation binding the contract event 0x257242c53ca9802d7b970b3ae41707c259e0b36ccce43bc618cfe649768642b3.
//
// Solidity: event DistrictName(uint256 district_id)
func (_DistrictContract *DistrictContractFilterer) ParseDistrictName(log types.Log) (*DistrictContractDistrictName, error) {
	event := new(DistrictContractDistrictName)
	if err := _DistrictContract.contract.UnpackLog(event, "DistrictName", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistrictContractMetaTransactionExecutedIterator is returned from FilterMetaTransactionExecuted and is used to iterate over the raw logs and unpacked data for MetaTransactionExecuted events raised by the DistrictContract contract.
type DistrictContractMetaTransactionExecutedIterator struct {
	Event *DistrictContractMetaTransactionExecuted // Event containing the contract specifics and raw log

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
func (it *DistrictContractMetaTransactionExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistrictContractMetaTransactionExecuted)
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
		it.Event = new(DistrictContractMetaTransactionExecuted)
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
func (it *DistrictContractMetaTransactionExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistrictContractMetaTransactionExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistrictContractMetaTransactionExecuted represents a MetaTransactionExecuted event raised by the DistrictContract contract.
type DistrictContractMetaTransactionExecuted struct {
	UserAddress       common.Address
	RelayerAddress    common.Address
	FunctionSignature []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterMetaTransactionExecuted is a free log retrieval operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_DistrictContract *DistrictContractFilterer) FilterMetaTransactionExecuted(opts *bind.FilterOpts) (*DistrictContractMetaTransactionExecutedIterator, error) {

	logs, sub, err := _DistrictContract.contract.FilterLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return &DistrictContractMetaTransactionExecutedIterator{contract: _DistrictContract.contract, event: "MetaTransactionExecuted", logs: logs, sub: sub}, nil
}

// WatchMetaTransactionExecuted is a free log subscription operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_DistrictContract *DistrictContractFilterer) WatchMetaTransactionExecuted(opts *bind.WatchOpts, sink chan<- *DistrictContractMetaTransactionExecuted) (event.Subscription, error) {

	logs, sub, err := _DistrictContract.contract.WatchLogs(opts, "MetaTransactionExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistrictContractMetaTransactionExecuted)
				if err := _DistrictContract.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
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

// ParseMetaTransactionExecuted is a log parse operation binding the contract event 0x5845892132946850460bff5a0083f71031bc5bf9aadcd40f1de79423eac9b10b.
//
// Solidity: event MetaTransactionExecuted(address userAddress, address relayerAddress, bytes functionSignature)
func (_DistrictContract *DistrictContractFilterer) ParseMetaTransactionExecuted(log types.Log) (*DistrictContractMetaTransactionExecuted, error) {
	event := new(DistrictContractMetaTransactionExecuted)
	if err := _DistrictContract.contract.UnpackLog(event, "MetaTransactionExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistrictContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DistrictContract contract.
type DistrictContractOwnershipTransferredIterator struct {
	Event *DistrictContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DistrictContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistrictContractOwnershipTransferred)
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
		it.Event = new(DistrictContractOwnershipTransferred)
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
func (it *DistrictContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistrictContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistrictContractOwnershipTransferred represents a OwnershipTransferred event raised by the DistrictContract contract.
type DistrictContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DistrictContract *DistrictContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DistrictContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DistrictContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DistrictContractOwnershipTransferredIterator{contract: _DistrictContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DistrictContract *DistrictContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DistrictContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DistrictContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistrictContractOwnershipTransferred)
				if err := _DistrictContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DistrictContract *DistrictContractFilterer) ParseOwnershipTransferred(log types.Log) (*DistrictContractOwnershipTransferred, error) {
	event := new(DistrictContractOwnershipTransferred)
	if err := _DistrictContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistrictContractPlotCreationIterator is returned from FilterPlotCreation and is used to iterate over the raw logs and unpacked data for PlotCreation events raised by the DistrictContract contract.
type DistrictContractPlotCreationIterator struct {
	Event *DistrictContractPlotCreation // Event containing the contract specifics and raw log

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
func (it *DistrictContractPlotCreationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistrictContractPlotCreation)
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
		it.Event = new(DistrictContractPlotCreation)
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
func (it *DistrictContractPlotCreationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistrictContractPlotCreationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistrictContractPlotCreation represents a PlotCreation event raised by the DistrictContract contract.
type DistrictContractPlotCreation struct {
	X      *big.Int
	Z      *big.Int
	PlotId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPlotCreation is a free log retrieval operation binding the contract event 0x70111e952c6021484b2dbe6ba7539974d0faaf4722b98cab13eba5249333067e.
//
// Solidity: event PlotCreation(int128 x, int128 z, uint256 plotId)
func (_DistrictContract *DistrictContractFilterer) FilterPlotCreation(opts *bind.FilterOpts) (*DistrictContractPlotCreationIterator, error) {

	logs, sub, err := _DistrictContract.contract.FilterLogs(opts, "PlotCreation")
	if err != nil {
		return nil, err
	}
	return &DistrictContractPlotCreationIterator{contract: _DistrictContract.contract, event: "PlotCreation", logs: logs, sub: sub}, nil
}

// WatchPlotCreation is a free log subscription operation binding the contract event 0x70111e952c6021484b2dbe6ba7539974d0faaf4722b98cab13eba5249333067e.
//
// Solidity: event PlotCreation(int128 x, int128 z, uint256 plotId)
func (_DistrictContract *DistrictContractFilterer) WatchPlotCreation(opts *bind.WatchOpts, sink chan<- *DistrictContractPlotCreation) (event.Subscription, error) {

	logs, sub, err := _DistrictContract.contract.WatchLogs(opts, "PlotCreation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistrictContractPlotCreation)
				if err := _DistrictContract.contract.UnpackLog(event, "PlotCreation", log); err != nil {
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

// ParsePlotCreation is a log parse operation binding the contract event 0x70111e952c6021484b2dbe6ba7539974d0faaf4722b98cab13eba5249333067e.
//
// Solidity: event PlotCreation(int128 x, int128 z, uint256 plotId)
func (_DistrictContract *DistrictContractFilterer) ParsePlotCreation(log types.Log) (*DistrictContractPlotCreation, error) {
	event := new(DistrictContractPlotCreation)
	if err := _DistrictContract.contract.UnpackLog(event, "PlotCreation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistrictContractPlotTransferIterator is returned from FilterPlotTransfer and is used to iterate over the raw logs and unpacked data for PlotTransfer events raised by the DistrictContract contract.
type DistrictContractPlotTransferIterator struct {
	Event *DistrictContractPlotTransfer // Event containing the contract specifics and raw log

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
func (it *DistrictContractPlotTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistrictContractPlotTransfer)
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
		it.Event = new(DistrictContractPlotTransfer)
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
func (it *DistrictContractPlotTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistrictContractPlotTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistrictContractPlotTransfer represents a PlotTransfer event raised by the DistrictContract contract.
type DistrictContractPlotTransfer struct {
	OriginId *big.Int
	TargetId *big.Int
	PlotId   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPlotTransfer is a free log retrieval operation binding the contract event 0x5112ef0e7d99d6ff5fcfc318db540adef82456873e39b67493ce8cf18f2af76c.
//
// Solidity: event PlotTransfer(uint256 origin_id, uint256 target_id, uint256 plotId)
func (_DistrictContract *DistrictContractFilterer) FilterPlotTransfer(opts *bind.FilterOpts) (*DistrictContractPlotTransferIterator, error) {

	logs, sub, err := _DistrictContract.contract.FilterLogs(opts, "PlotTransfer")
	if err != nil {
		return nil, err
	}
	return &DistrictContractPlotTransferIterator{contract: _DistrictContract.contract, event: "PlotTransfer", logs: logs, sub: sub}, nil
}

// WatchPlotTransfer is a free log subscription operation binding the contract event 0x5112ef0e7d99d6ff5fcfc318db540adef82456873e39b67493ce8cf18f2af76c.
//
// Solidity: event PlotTransfer(uint256 origin_id, uint256 target_id, uint256 plotId)
func (_DistrictContract *DistrictContractFilterer) WatchPlotTransfer(opts *bind.WatchOpts, sink chan<- *DistrictContractPlotTransfer) (event.Subscription, error) {

	logs, sub, err := _DistrictContract.contract.WatchLogs(opts, "PlotTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistrictContractPlotTransfer)
				if err := _DistrictContract.contract.UnpackLog(event, "PlotTransfer", log); err != nil {
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

// ParsePlotTransfer is a log parse operation binding the contract event 0x5112ef0e7d99d6ff5fcfc318db540adef82456873e39b67493ce8cf18f2af76c.
//
// Solidity: event PlotTransfer(uint256 origin_id, uint256 target_id, uint256 plotId)
func (_DistrictContract *DistrictContractFilterer) ParsePlotTransfer(log types.Log) (*DistrictContractPlotTransfer, error) {
	event := new(DistrictContractPlotTransfer)
	if err := _DistrictContract.contract.UnpackLog(event, "PlotTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistrictContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the DistrictContract contract.
type DistrictContractTransferIterator struct {
	Event *DistrictContractTransfer // Event containing the contract specifics and raw log

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
func (it *DistrictContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistrictContractTransfer)
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
		it.Event = new(DistrictContractTransfer)
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
func (it *DistrictContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistrictContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistrictContractTransfer represents a Transfer event raised by the DistrictContract contract.
type DistrictContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DistrictContract *DistrictContractFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*DistrictContractTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DistrictContract.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &DistrictContractTransferIterator{contract: _DistrictContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DistrictContract *DistrictContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DistrictContractTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _DistrictContract.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistrictContractTransfer)
				if err := _DistrictContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_DistrictContract *DistrictContractFilterer) ParseTransfer(log types.Log) (*DistrictContractTransfer, error) {
	event := new(DistrictContractTransfer)
	if err := _DistrictContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
