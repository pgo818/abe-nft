// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mainnft

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

// MainnftMetaData contains all meta data concerning the Mainnft contract.
var MainnftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ChildNFTCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"childNFTContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"createChildNFTWithURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipfsURI\",\"type\":\"string\"}],\"name\":\"extractIPFSHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mintTo\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"ipfsURI\",\"type\":\"string\"}],\"name\":\"mintWithHttpGateway\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_childNFTContract\",\"type\":\"address\"}],\"name\":\"setChildNFTContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newTokenURI\",\"type\":\"string\"}],\"name\":\"setSpecificTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MainnftABI is the input ABI used to generate the binding from.
// Deprecated: Use MainnftMetaData.ABI instead.
var MainnftABI = MainnftMetaData.ABI

// Mainnft is an auto generated Go binding around an Ethereum contract.
type Mainnft struct {
	MainnftCaller     // Read-only binding to the contract
	MainnftTransactor // Write-only binding to the contract
	MainnftFilterer   // Log filterer for contract events
}

// MainnftCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainnftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainnftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainnftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainnftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainnftSession struct {
	Contract     *Mainnft          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainnftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainnftCallerSession struct {
	Contract *MainnftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MainnftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainnftTransactorSession struct {
	Contract     *MainnftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MainnftRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainnftRaw struct {
	Contract *Mainnft // Generic contract binding to access the raw methods on
}

// MainnftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainnftCallerRaw struct {
	Contract *MainnftCaller // Generic read-only contract binding to access the raw methods on
}

// MainnftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainnftTransactorRaw struct {
	Contract *MainnftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMainnft creates a new instance of Mainnft, bound to a specific deployed contract.
func NewMainnft(address common.Address, backend bind.ContractBackend) (*Mainnft, error) {
	contract, err := bindMainnft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mainnft{MainnftCaller: MainnftCaller{contract: contract}, MainnftTransactor: MainnftTransactor{contract: contract}, MainnftFilterer: MainnftFilterer{contract: contract}}, nil
}

// NewMainnftCaller creates a new read-only instance of Mainnft, bound to a specific deployed contract.
func NewMainnftCaller(address common.Address, caller bind.ContractCaller) (*MainnftCaller, error) {
	contract, err := bindMainnft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainnftCaller{contract: contract}, nil
}

// NewMainnftTransactor creates a new write-only instance of Mainnft, bound to a specific deployed contract.
func NewMainnftTransactor(address common.Address, transactor bind.ContractTransactor) (*MainnftTransactor, error) {
	contract, err := bindMainnft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainnftTransactor{contract: contract}, nil
}

// NewMainnftFilterer creates a new log filterer instance of Mainnft, bound to a specific deployed contract.
func NewMainnftFilterer(address common.Address, filterer bind.ContractFilterer) (*MainnftFilterer, error) {
	contract, err := bindMainnft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainnftFilterer{contract: contract}, nil
}

// bindMainnft binds a generic wrapper to an already deployed contract.
func bindMainnft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainnftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mainnft *MainnftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mainnft.Contract.MainnftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mainnft *MainnftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mainnft.Contract.MainnftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mainnft *MainnftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mainnft.Contract.MainnftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mainnft *MainnftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mainnft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mainnft *MainnftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mainnft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mainnft *MainnftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mainnft.Contract.contract.Transact(opts, method, params...)
}

// MAXAMOUNT is a free data retrieval call binding the contract method 0xd40dc870.
//
// Solidity: function MAX_AMOUNT() view returns(uint256)
func (_Mainnft *MainnftCaller) MAXAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "MAX_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXAMOUNT is a free data retrieval call binding the contract method 0xd40dc870.
//
// Solidity: function MAX_AMOUNT() view returns(uint256)
func (_Mainnft *MainnftSession) MAXAMOUNT() (*big.Int, error) {
	return _Mainnft.Contract.MAXAMOUNT(&_Mainnft.CallOpts)
}

// MAXAMOUNT is a free data retrieval call binding the contract method 0xd40dc870.
//
// Solidity: function MAX_AMOUNT() view returns(uint256)
func (_Mainnft *MainnftCallerSession) MAXAMOUNT() (*big.Int, error) {
	return _Mainnft.Contract.MAXAMOUNT(&_Mainnft.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mainnft *MainnftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mainnft *MainnftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Mainnft.Contract.BalanceOf(&_Mainnft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mainnft *MainnftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Mainnft.Contract.BalanceOf(&_Mainnft.CallOpts, owner)
}

// ChildNFTContract is a free data retrieval call binding the contract method 0x4ed3c7d5.
//
// Solidity: function childNFTContract() view returns(address)
func (_Mainnft *MainnftCaller) ChildNFTContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "childNFTContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChildNFTContract is a free data retrieval call binding the contract method 0x4ed3c7d5.
//
// Solidity: function childNFTContract() view returns(address)
func (_Mainnft *MainnftSession) ChildNFTContract() (common.Address, error) {
	return _Mainnft.Contract.ChildNFTContract(&_Mainnft.CallOpts)
}

// ChildNFTContract is a free data retrieval call binding the contract method 0x4ed3c7d5.
//
// Solidity: function childNFTContract() view returns(address)
func (_Mainnft *MainnftCallerSession) ChildNFTContract() (common.Address, error) {
	return _Mainnft.Contract.ChildNFTContract(&_Mainnft.CallOpts)
}

// ExtractIPFSHash is a free data retrieval call binding the contract method 0x2c1d8933.
//
// Solidity: function extractIPFSHash(string ipfsURI) pure returns(string)
func (_Mainnft *MainnftCaller) ExtractIPFSHash(opts *bind.CallOpts, ipfsURI string) (string, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "extractIPFSHash", ipfsURI)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ExtractIPFSHash is a free data retrieval call binding the contract method 0x2c1d8933.
//
// Solidity: function extractIPFSHash(string ipfsURI) pure returns(string)
func (_Mainnft *MainnftSession) ExtractIPFSHash(ipfsURI string) (string, error) {
	return _Mainnft.Contract.ExtractIPFSHash(&_Mainnft.CallOpts, ipfsURI)
}

// ExtractIPFSHash is a free data retrieval call binding the contract method 0x2c1d8933.
//
// Solidity: function extractIPFSHash(string ipfsURI) pure returns(string)
func (_Mainnft *MainnftCallerSession) ExtractIPFSHash(ipfsURI string) (string, error) {
	return _Mainnft.Contract.ExtractIPFSHash(&_Mainnft.CallOpts, ipfsURI)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mainnft *MainnftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mainnft *MainnftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Mainnft.Contract.GetApproved(&_Mainnft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mainnft *MainnftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Mainnft.Contract.GetApproved(&_Mainnft.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mainnft *MainnftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mainnft *MainnftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Mainnft.Contract.IsApprovedForAll(&_Mainnft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mainnft *MainnftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Mainnft.Contract.IsApprovedForAll(&_Mainnft.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mainnft *MainnftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mainnft *MainnftSession) Name() (string, error) {
	return _Mainnft.Contract.Name(&_Mainnft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mainnft *MainnftCallerSession) Name() (string, error) {
	return _Mainnft.Contract.Name(&_Mainnft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mainnft *MainnftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mainnft *MainnftSession) Owner() (common.Address, error) {
	return _Mainnft.Contract.Owner(&_Mainnft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mainnft *MainnftCallerSession) Owner() (common.Address, error) {
	return _Mainnft.Contract.Owner(&_Mainnft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mainnft *MainnftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mainnft *MainnftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Mainnft.Contract.OwnerOf(&_Mainnft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mainnft *MainnftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Mainnft.Contract.OwnerOf(&_Mainnft.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mainnft *MainnftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mainnft *MainnftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Mainnft.Contract.SupportsInterface(&_Mainnft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mainnft *MainnftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Mainnft.Contract.SupportsInterface(&_Mainnft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mainnft *MainnftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mainnft *MainnftSession) Symbol() (string, error) {
	return _Mainnft.Contract.Symbol(&_Mainnft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mainnft *MainnftCallerSession) Symbol() (string, error) {
	return _Mainnft.Contract.Symbol(&_Mainnft.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Mainnft *MainnftCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Mainnft *MainnftSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Mainnft.Contract.TokenByIndex(&_Mainnft.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Mainnft *MainnftCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Mainnft.Contract.TokenByIndex(&_Mainnft.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Mainnft *MainnftCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Mainnft *MainnftSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Mainnft.Contract.TokenOfOwnerByIndex(&_Mainnft.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Mainnft *MainnftCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Mainnft.Contract.TokenOfOwnerByIndex(&_Mainnft.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mainnft *MainnftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mainnft *MainnftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Mainnft.Contract.TokenURI(&_Mainnft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mainnft *MainnftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Mainnft.Contract.TokenURI(&_Mainnft.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mainnft *MainnftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mainnft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mainnft *MainnftSession) TotalSupply() (*big.Int, error) {
	return _Mainnft.Contract.TotalSupply(&_Mainnft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Mainnft *MainnftCallerSession) TotalSupply() (*big.Int, error) {
	return _Mainnft.Contract.TotalSupply(&_Mainnft.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mainnft *MainnftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mainnft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mainnft *MainnftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mainnft.Contract.Approve(&_Mainnft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mainnft *MainnftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mainnft.Contract.Approve(&_Mainnft.TransactOpts, to, tokenId)
}

// CreateChildNFTWithURI is a paid mutator transaction binding the contract method 0x29360da0.
//
// Solidity: function createChildNFTWithURI(address to, uint256 tokenId, address receiver, string uri) returns()
func (_Mainnft *MainnftTransactor) CreateChildNFTWithURI(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, receiver common.Address, uri string) (*types.Transaction, error) {
	return _Mainnft.contract.Transact(opts, "createChildNFTWithURI", to, tokenId, receiver, uri)
}

// CreateChildNFTWithURI is a paid mutator transaction binding the contract method 0x29360da0.
//
// Solidity: function createChildNFTWithURI(address to, uint256 tokenId, address receiver, string uri) returns()
func (_Mainnft *MainnftSession) CreateChildNFTWithURI(to common.Address, tokenId *big.Int, receiver common.Address, uri string) (*types.Transaction, error) {
	return _Mainnft.Contract.CreateChildNFTWithURI(&_Mainnft.TransactOpts, to, tokenId, receiver, uri)
}

// CreateChildNFTWithURI is a paid mutator transaction binding the contract method 0x29360da0.
//
// Solidity: function createChildNFTWithURI(address to, uint256 tokenId, address receiver, string uri) returns()
func (_Mainnft *MainnftTransactorSession) CreateChildNFTWithURI(to common.Address, tokenId *big.Int, receiver common.Address, uri string) (*types.Transaction, error) {
	return _Mainnft.Contract.CreateChildNFTWithURI(&_Mainnft.TransactOpts, to, tokenId, receiver, uri)
}

// MintTo is a paid mutator transaction binding the contract method 0x0075a317.
//
// Solidity: function mintTo(address to, string uri) payable returns()
func (_Mainnft *MainnftTransactor) MintTo(opts *bind.TransactOpts, to common.Address, uri string) (*types.Transaction, error) {
	return _Mainnft.contract.Transact(opts, "mintTo", to, uri)
}

// MintTo is a paid mutator transaction binding the contract method 0x0075a317.
//
// Solidity: function mintTo(address to, string uri) payable returns()
func (_Mainnft *MainnftSession) MintTo(to common.Address, uri string) (*types.Transaction, error) {
	return _Mainnft.Contract.MintTo(&_Mainnft.TransactOpts, to, uri)
}

// MintTo is a paid mutator transaction binding the contract method 0x0075a317.
//
// Solidity: function mintTo(address to, string uri) payable returns()
func (_Mainnft *MainnftTransactorSession) MintTo(to common.Address, uri string) (*types.Transaction, error) {
	return _Mainnft.Contract.MintTo(&_Mainnft.TransactOpts, to, uri)
}

// MintWithHttpGateway is a paid mutator transaction binding the contract method 0xf042bccd.
//
// Solidity: function mintWithHttpGateway(address to, string ipfsURI) payable returns()
func (_Mainnft *MainnftTransactor) MintWithHttpGateway(opts *bind.TransactOpts, to common.Address, ipfsURI string) (*types.Transaction, error) {
	return _Mainnft.contract.Transact(opts, "mintWithHttpGateway", to, ipfsURI)
}

// MintWithHttpGateway is a paid mutator transaction binding the contract method 0xf042bccd.
//
// Solidity: function mintWithHttpGateway(address to, string ipfsURI) payable returns()
func (_Mainnft *MainnftSession) MintWithHttpGateway(to common.Address, ipfsURI string) (*types.Transaction, error) {
	return _Mainnft.Contract.MintWithHttpGateway(&_Mainnft.TransactOpts, to, ipfsURI)
}

// MintWithHttpGateway is a paid mutator transaction binding the contract method 0xf042bccd.
//
// Solidity: function mintWithHttpGateway(address to, string ipfsURI) payable returns()
func (_Mainnft *MainnftTransactorSession) MintWithHttpGateway(to common.Address, ipfsURI string) (*types.Transaction, error) {
	return _Mainnft.Contract.MintWithHttpGateway(&_Mainnft.TransactOpts, to, ipfsURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mainnft *MainnftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mainnft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mainnft *MainnftSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mainnft.Contract.RenounceOwnership(&_Mainnft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mainnft *MainnftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mainnft.Contract.RenounceOwnership(&_Mainnft.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mainnft *MainnftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mainnft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mainnft *MainnftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mainnft.Contract.SafeTransferFrom(&_Mainnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mainnft *MainnftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mainnft.Contract.SafeTransferFrom(&_Mainnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Mainnft *MainnftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Mainnft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Mainnft *MainnftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Mainnft.Contract.SafeTransferFrom0(&_Mainnft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Mainnft *MainnftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Mainnft.Contract.SafeTransferFrom0(&_Mainnft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mainnft *MainnftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mainnft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mainnft *MainnftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mainnft.Contract.SetApprovalForAll(&_Mainnft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mainnft *MainnftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mainnft.Contract.SetApprovalForAll(&_Mainnft.TransactOpts, operator, approved)
}

// SetChildNFTContract is a paid mutator transaction binding the contract method 0x8265ea47.
//
// Solidity: function setChildNFTContract(address _childNFTContract) returns()
func (_Mainnft *MainnftTransactor) SetChildNFTContract(opts *bind.TransactOpts, _childNFTContract common.Address) (*types.Transaction, error) {
	return _Mainnft.contract.Transact(opts, "setChildNFTContract", _childNFTContract)
}

// SetChildNFTContract is a paid mutator transaction binding the contract method 0x8265ea47.
//
// Solidity: function setChildNFTContract(address _childNFTContract) returns()
func (_Mainnft *MainnftSession) SetChildNFTContract(_childNFTContract common.Address) (*types.Transaction, error) {
	return _Mainnft.Contract.SetChildNFTContract(&_Mainnft.TransactOpts, _childNFTContract)
}

// SetChildNFTContract is a paid mutator transaction binding the contract method 0x8265ea47.
//
// Solidity: function setChildNFTContract(address _childNFTContract) returns()
func (_Mainnft *MainnftTransactorSession) SetChildNFTContract(_childNFTContract common.Address) (*types.Transaction, error) {
	return _Mainnft.Contract.SetChildNFTContract(&_Mainnft.TransactOpts, _childNFTContract)
}

// SetSpecificTokenURI is a paid mutator transaction binding the contract method 0x537c1e42.
//
// Solidity: function setSpecificTokenURI(address to, uint256 tokenId, string newTokenURI) returns()
func (_Mainnft *MainnftTransactor) SetSpecificTokenURI(opts *bind.TransactOpts, to common.Address, tokenId *big.Int, newTokenURI string) (*types.Transaction, error) {
	return _Mainnft.contract.Transact(opts, "setSpecificTokenURI", to, tokenId, newTokenURI)
}

// SetSpecificTokenURI is a paid mutator transaction binding the contract method 0x537c1e42.
//
// Solidity: function setSpecificTokenURI(address to, uint256 tokenId, string newTokenURI) returns()
func (_Mainnft *MainnftSession) SetSpecificTokenURI(to common.Address, tokenId *big.Int, newTokenURI string) (*types.Transaction, error) {
	return _Mainnft.Contract.SetSpecificTokenURI(&_Mainnft.TransactOpts, to, tokenId, newTokenURI)
}

// SetSpecificTokenURI is a paid mutator transaction binding the contract method 0x537c1e42.
//
// Solidity: function setSpecificTokenURI(address to, uint256 tokenId, string newTokenURI) returns()
func (_Mainnft *MainnftTransactorSession) SetSpecificTokenURI(to common.Address, tokenId *big.Int, newTokenURI string) (*types.Transaction, error) {
	return _Mainnft.Contract.SetSpecificTokenURI(&_Mainnft.TransactOpts, to, tokenId, newTokenURI)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mainnft *MainnftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mainnft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mainnft *MainnftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mainnft.Contract.TransferFrom(&_Mainnft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mainnft *MainnftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mainnft.Contract.TransferFrom(&_Mainnft.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mainnft *MainnftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mainnft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mainnft *MainnftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mainnft.Contract.TransferOwnership(&_Mainnft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mainnft *MainnftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mainnft.Contract.TransferOwnership(&_Mainnft.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Mainnft *MainnftTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mainnft.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Mainnft *MainnftSession) Withdraw() (*types.Transaction, error) {
	return _Mainnft.Contract.Withdraw(&_Mainnft.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Mainnft *MainnftTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Mainnft.Contract.Withdraw(&_Mainnft.TransactOpts)
}

// MainnftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Mainnft contract.
type MainnftApprovalIterator struct {
	Event *MainnftApproval // Event containing the contract specifics and raw log

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
func (it *MainnftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnftApproval)
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
		it.Event = new(MainnftApproval)
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
func (it *MainnftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnftApproval represents a Approval event raised by the Mainnft contract.
type MainnftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Mainnft *MainnftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MainnftApprovalIterator, error) {

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

	logs, sub, err := _Mainnft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MainnftApprovalIterator{contract: _Mainnft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Mainnft *MainnftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MainnftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Mainnft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnftApproval)
				if err := _Mainnft.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Mainnft *MainnftFilterer) ParseApproval(log types.Log) (*MainnftApproval, error) {
	event := new(MainnftApproval)
	if err := _Mainnft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainnftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Mainnft contract.
type MainnftApprovalForAllIterator struct {
	Event *MainnftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MainnftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnftApprovalForAll)
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
		it.Event = new(MainnftApprovalForAll)
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
func (it *MainnftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnftApprovalForAll represents a ApprovalForAll event raised by the Mainnft contract.
type MainnftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Mainnft *MainnftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MainnftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Mainnft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MainnftApprovalForAllIterator{contract: _Mainnft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Mainnft *MainnftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MainnftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Mainnft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnftApprovalForAll)
				if err := _Mainnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Mainnft *MainnftFilterer) ParseApprovalForAll(log types.Log) (*MainnftApprovalForAll, error) {
	event := new(MainnftApprovalForAll)
	if err := _Mainnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainnftBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Mainnft contract.
type MainnftBatchMetadataUpdateIterator struct {
	Event *MainnftBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *MainnftBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnftBatchMetadataUpdate)
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
		it.Event = new(MainnftBatchMetadataUpdate)
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
func (it *MainnftBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnftBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnftBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Mainnft contract.
type MainnftBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Mainnft *MainnftFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*MainnftBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Mainnft.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &MainnftBatchMetadataUpdateIterator{contract: _Mainnft.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Mainnft *MainnftFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *MainnftBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Mainnft.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnftBatchMetadataUpdate)
				if err := _Mainnft.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Mainnft *MainnftFilterer) ParseBatchMetadataUpdate(log types.Log) (*MainnftBatchMetadataUpdate, error) {
	event := new(MainnftBatchMetadataUpdate)
	if err := _Mainnft.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainnftChildNFTCreatedIterator is returned from FilterChildNFTCreated and is used to iterate over the raw logs and unpacked data for ChildNFTCreated events raised by the Mainnft contract.
type MainnftChildNFTCreatedIterator struct {
	Event *MainnftChildNFTCreated // Event containing the contract specifics and raw log

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
func (it *MainnftChildNFTCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnftChildNFTCreated)
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
		it.Event = new(MainnftChildNFTCreated)
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
func (it *MainnftChildNFTCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnftChildNFTCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnftChildNFTCreated represents a ChildNFTCreated event raised by the Mainnft contract.
type MainnftChildNFTCreated struct {
	TokenId  *big.Int
	Receiver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChildNFTCreated is a free log retrieval operation binding the contract event 0x2cbffa88a077b6f4874d4225e9cf4de15ee94033ef2e7d4d0b1bf128101a3e2b.
//
// Solidity: event ChildNFTCreated(uint256 indexed tokenId, address indexed receiver)
func (_Mainnft *MainnftFilterer) FilterChildNFTCreated(opts *bind.FilterOpts, tokenId []*big.Int, receiver []common.Address) (*MainnftChildNFTCreatedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Mainnft.contract.FilterLogs(opts, "ChildNFTCreated", tokenIdRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &MainnftChildNFTCreatedIterator{contract: _Mainnft.contract, event: "ChildNFTCreated", logs: logs, sub: sub}, nil
}

// WatchChildNFTCreated is a free log subscription operation binding the contract event 0x2cbffa88a077b6f4874d4225e9cf4de15ee94033ef2e7d4d0b1bf128101a3e2b.
//
// Solidity: event ChildNFTCreated(uint256 indexed tokenId, address indexed receiver)
func (_Mainnft *MainnftFilterer) WatchChildNFTCreated(opts *bind.WatchOpts, sink chan<- *MainnftChildNFTCreated, tokenId []*big.Int, receiver []common.Address) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Mainnft.contract.WatchLogs(opts, "ChildNFTCreated", tokenIdRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnftChildNFTCreated)
				if err := _Mainnft.contract.UnpackLog(event, "ChildNFTCreated", log); err != nil {
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

// ParseChildNFTCreated is a log parse operation binding the contract event 0x2cbffa88a077b6f4874d4225e9cf4de15ee94033ef2e7d4d0b1bf128101a3e2b.
//
// Solidity: event ChildNFTCreated(uint256 indexed tokenId, address indexed receiver)
func (_Mainnft *MainnftFilterer) ParseChildNFTCreated(log types.Log) (*MainnftChildNFTCreated, error) {
	event := new(MainnftChildNFTCreated)
	if err := _Mainnft.contract.UnpackLog(event, "ChildNFTCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainnftMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Mainnft contract.
type MainnftMetadataUpdateIterator struct {
	Event *MainnftMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *MainnftMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnftMetadataUpdate)
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
		it.Event = new(MainnftMetadataUpdate)
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
func (it *MainnftMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnftMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnftMetadataUpdate represents a MetadataUpdate event raised by the Mainnft contract.
type MainnftMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Mainnft *MainnftFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*MainnftMetadataUpdateIterator, error) {

	logs, sub, err := _Mainnft.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &MainnftMetadataUpdateIterator{contract: _Mainnft.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Mainnft *MainnftFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *MainnftMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Mainnft.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnftMetadataUpdate)
				if err := _Mainnft.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Mainnft *MainnftFilterer) ParseMetadataUpdate(log types.Log) (*MainnftMetadataUpdate, error) {
	event := new(MainnftMetadataUpdate)
	if err := _Mainnft.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainnftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mainnft contract.
type MainnftOwnershipTransferredIterator struct {
	Event *MainnftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MainnftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnftOwnershipTransferred)
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
		it.Event = new(MainnftOwnershipTransferred)
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
func (it *MainnftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnftOwnershipTransferred represents a OwnershipTransferred event raised by the Mainnft contract.
type MainnftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mainnft *MainnftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MainnftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mainnft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MainnftOwnershipTransferredIterator{contract: _Mainnft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mainnft *MainnftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MainnftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mainnft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnftOwnershipTransferred)
				if err := _Mainnft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Mainnft *MainnftFilterer) ParseOwnershipTransferred(log types.Log) (*MainnftOwnershipTransferred, error) {
	event := new(MainnftOwnershipTransferred)
	if err := _Mainnft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainnftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mainnft contract.
type MainnftTransferIterator struct {
	Event *MainnftTransfer // Event containing the contract specifics and raw log

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
func (it *MainnftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainnftTransfer)
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
		it.Event = new(MainnftTransfer)
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
func (it *MainnftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainnftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainnftTransfer represents a Transfer event raised by the Mainnft contract.
type MainnftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Mainnft *MainnftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MainnftTransferIterator, error) {

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

	logs, sub, err := _Mainnft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MainnftTransferIterator{contract: _Mainnft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Mainnft *MainnftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MainnftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Mainnft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainnftTransfer)
				if err := _Mainnft.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Mainnft *MainnftFilterer) ParseTransfer(log types.Log) (*MainnftTransfer, error) {
	event := new(MainnftTransfer)
	if err := _Mainnft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
