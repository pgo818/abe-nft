// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package childnft

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

// ChildnftMetaData contains all meta data concerning the Childnft contract.
var ChildnftMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ERC721EnumerableForbiddenBatchMint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"ERC721OutOfBoundsIndex\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"parentTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"childTokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ChildTokenMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ipfsURI\",\"type\":\"string\"}],\"name\":\"extractIPFSHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"childTokenId\",\"type\":\"uint256\"}],\"name\":\"getChildCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"childTokenId\",\"type\":\"uint256\"}],\"name\":\"getParentTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mainNFTContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"parentTokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"}],\"name\":\"mintChildNFTWithURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newURI\",\"type\":\"string\"}],\"name\":\"setChildTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mainNFTContract\",\"type\":\"address\"}],\"name\":\"setMainNFTContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newTokenURI\",\"type\":\"string\"}],\"name\":\"setSpecificTokenURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ChildnftABI is the input ABI used to generate the binding from.
// Deprecated: Use ChildnftMetaData.ABI instead.
var ChildnftABI = ChildnftMetaData.ABI

// Childnft is an auto generated Go binding around an Ethereum contract.
type Childnft struct {
	ChildnftCaller     // Read-only binding to the contract
	ChildnftTransactor // Write-only binding to the contract
	ChildnftFilterer   // Log filterer for contract events
}

// ChildnftCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChildnftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildnftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChildnftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildnftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChildnftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChildnftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChildnftSession struct {
	Contract     *Childnft         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChildnftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChildnftCallerSession struct {
	Contract *ChildnftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ChildnftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChildnftTransactorSession struct {
	Contract     *ChildnftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ChildnftRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChildnftRaw struct {
	Contract *Childnft // Generic contract binding to access the raw methods on
}

// ChildnftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChildnftCallerRaw struct {
	Contract *ChildnftCaller // Generic read-only contract binding to access the raw methods on
}

// ChildnftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChildnftTransactorRaw struct {
	Contract *ChildnftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChildnft creates a new instance of Childnft, bound to a specific deployed contract.
func NewChildnft(address common.Address, backend bind.ContractBackend) (*Childnft, error) {
	contract, err := bindChildnft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Childnft{ChildnftCaller: ChildnftCaller{contract: contract}, ChildnftTransactor: ChildnftTransactor{contract: contract}, ChildnftFilterer: ChildnftFilterer{contract: contract}}, nil
}

// NewChildnftCaller creates a new read-only instance of Childnft, bound to a specific deployed contract.
func NewChildnftCaller(address common.Address, caller bind.ContractCaller) (*ChildnftCaller, error) {
	contract, err := bindChildnft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChildnftCaller{contract: contract}, nil
}

// NewChildnftTransactor creates a new write-only instance of Childnft, bound to a specific deployed contract.
func NewChildnftTransactor(address common.Address, transactor bind.ContractTransactor) (*ChildnftTransactor, error) {
	contract, err := bindChildnft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChildnftTransactor{contract: contract}, nil
}

// NewChildnftFilterer creates a new log filterer instance of Childnft, bound to a specific deployed contract.
func NewChildnftFilterer(address common.Address, filterer bind.ContractFilterer) (*ChildnftFilterer, error) {
	contract, err := bindChildnft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChildnftFilterer{contract: contract}, nil
}

// bindChildnft binds a generic wrapper to an already deployed contract.
func bindChildnft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChildnftMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Childnft *ChildnftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Childnft.Contract.ChildnftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Childnft *ChildnftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Childnft.Contract.ChildnftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Childnft *ChildnftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Childnft.Contract.ChildnftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Childnft *ChildnftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Childnft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Childnft *ChildnftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Childnft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Childnft *ChildnftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Childnft.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Childnft *ChildnftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Childnft *ChildnftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Childnft.Contract.BalanceOf(&_Childnft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Childnft *ChildnftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Childnft.Contract.BalanceOf(&_Childnft.CallOpts, owner)
}

// ExtractIPFSHash is a free data retrieval call binding the contract method 0x2c1d8933.
//
// Solidity: function extractIPFSHash(string ipfsURI) pure returns(string)
func (_Childnft *ChildnftCaller) ExtractIPFSHash(opts *bind.CallOpts, ipfsURI string) (string, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "extractIPFSHash", ipfsURI)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ExtractIPFSHash is a free data retrieval call binding the contract method 0x2c1d8933.
//
// Solidity: function extractIPFSHash(string ipfsURI) pure returns(string)
func (_Childnft *ChildnftSession) ExtractIPFSHash(ipfsURI string) (string, error) {
	return _Childnft.Contract.ExtractIPFSHash(&_Childnft.CallOpts, ipfsURI)
}

// ExtractIPFSHash is a free data retrieval call binding the contract method 0x2c1d8933.
//
// Solidity: function extractIPFSHash(string ipfsURI) pure returns(string)
func (_Childnft *ChildnftCallerSession) ExtractIPFSHash(ipfsURI string) (string, error) {
	return _Childnft.Contract.ExtractIPFSHash(&_Childnft.CallOpts, ipfsURI)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Childnft *ChildnftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Childnft *ChildnftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Childnft.Contract.GetApproved(&_Childnft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Childnft *ChildnftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Childnft.Contract.GetApproved(&_Childnft.CallOpts, tokenId)
}

// GetChildCreator is a free data retrieval call binding the contract method 0xe59b355b.
//
// Solidity: function getChildCreator(uint256 childTokenId) view returns(address)
func (_Childnft *ChildnftCaller) GetChildCreator(opts *bind.CallOpts, childTokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "getChildCreator", childTokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetChildCreator is a free data retrieval call binding the contract method 0xe59b355b.
//
// Solidity: function getChildCreator(uint256 childTokenId) view returns(address)
func (_Childnft *ChildnftSession) GetChildCreator(childTokenId *big.Int) (common.Address, error) {
	return _Childnft.Contract.GetChildCreator(&_Childnft.CallOpts, childTokenId)
}

// GetChildCreator is a free data retrieval call binding the contract method 0xe59b355b.
//
// Solidity: function getChildCreator(uint256 childTokenId) view returns(address)
func (_Childnft *ChildnftCallerSession) GetChildCreator(childTokenId *big.Int) (common.Address, error) {
	return _Childnft.Contract.GetChildCreator(&_Childnft.CallOpts, childTokenId)
}

// GetParentTokenId is a free data retrieval call binding the contract method 0xcfa651f3.
//
// Solidity: function getParentTokenId(uint256 childTokenId) view returns(uint256)
func (_Childnft *ChildnftCaller) GetParentTokenId(opts *bind.CallOpts, childTokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "getParentTokenId", childTokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetParentTokenId is a free data retrieval call binding the contract method 0xcfa651f3.
//
// Solidity: function getParentTokenId(uint256 childTokenId) view returns(uint256)
func (_Childnft *ChildnftSession) GetParentTokenId(childTokenId *big.Int) (*big.Int, error) {
	return _Childnft.Contract.GetParentTokenId(&_Childnft.CallOpts, childTokenId)
}

// GetParentTokenId is a free data retrieval call binding the contract method 0xcfa651f3.
//
// Solidity: function getParentTokenId(uint256 childTokenId) view returns(uint256)
func (_Childnft *ChildnftCallerSession) GetParentTokenId(childTokenId *big.Int) (*big.Int, error) {
	return _Childnft.Contract.GetParentTokenId(&_Childnft.CallOpts, childTokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Childnft *ChildnftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Childnft *ChildnftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Childnft.Contract.IsApprovedForAll(&_Childnft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Childnft *ChildnftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Childnft.Contract.IsApprovedForAll(&_Childnft.CallOpts, owner, operator)
}

// MainNFTContract is a free data retrieval call binding the contract method 0x5ad317b7.
//
// Solidity: function mainNFTContract() view returns(address)
func (_Childnft *ChildnftCaller) MainNFTContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "mainNFTContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MainNFTContract is a free data retrieval call binding the contract method 0x5ad317b7.
//
// Solidity: function mainNFTContract() view returns(address)
func (_Childnft *ChildnftSession) MainNFTContract() (common.Address, error) {
	return _Childnft.Contract.MainNFTContract(&_Childnft.CallOpts)
}

// MainNFTContract is a free data retrieval call binding the contract method 0x5ad317b7.
//
// Solidity: function mainNFTContract() view returns(address)
func (_Childnft *ChildnftCallerSession) MainNFTContract() (common.Address, error) {
	return _Childnft.Contract.MainNFTContract(&_Childnft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Childnft *ChildnftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Childnft *ChildnftSession) Name() (string, error) {
	return _Childnft.Contract.Name(&_Childnft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Childnft *ChildnftCallerSession) Name() (string, error) {
	return _Childnft.Contract.Name(&_Childnft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Childnft *ChildnftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Childnft *ChildnftSession) Owner() (common.Address, error) {
	return _Childnft.Contract.Owner(&_Childnft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Childnft *ChildnftCallerSession) Owner() (common.Address, error) {
	return _Childnft.Contract.Owner(&_Childnft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Childnft *ChildnftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Childnft *ChildnftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Childnft.Contract.OwnerOf(&_Childnft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Childnft *ChildnftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Childnft.Contract.OwnerOf(&_Childnft.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Childnft *ChildnftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Childnft *ChildnftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Childnft.Contract.SupportsInterface(&_Childnft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Childnft *ChildnftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Childnft.Contract.SupportsInterface(&_Childnft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Childnft *ChildnftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Childnft *ChildnftSession) Symbol() (string, error) {
	return _Childnft.Contract.Symbol(&_Childnft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Childnft *ChildnftCallerSession) Symbol() (string, error) {
	return _Childnft.Contract.Symbol(&_Childnft.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Childnft *ChildnftCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Childnft *ChildnftSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Childnft.Contract.TokenByIndex(&_Childnft.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Childnft *ChildnftCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Childnft.Contract.TokenByIndex(&_Childnft.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Childnft *ChildnftCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Childnft *ChildnftSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Childnft.Contract.TokenOfOwnerByIndex(&_Childnft.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Childnft *ChildnftCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Childnft.Contract.TokenOfOwnerByIndex(&_Childnft.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Childnft *ChildnftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Childnft *ChildnftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Childnft.Contract.TokenURI(&_Childnft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Childnft *ChildnftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Childnft.Contract.TokenURI(&_Childnft.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Childnft *ChildnftCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Childnft.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Childnft *ChildnftSession) TotalSupply() (*big.Int, error) {
	return _Childnft.Contract.TotalSupply(&_Childnft.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Childnft *ChildnftCallerSession) TotalSupply() (*big.Int, error) {
	return _Childnft.Contract.TotalSupply(&_Childnft.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Childnft *ChildnftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Childnft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Childnft *ChildnftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Childnft.Contract.Approve(&_Childnft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Childnft *ChildnftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Childnft.Contract.Approve(&_Childnft.TransactOpts, to, tokenId)
}

// MintChildNFTWithURI is a paid mutator transaction binding the contract method 0xda682f3d.
//
// Solidity: function mintChildNFTWithURI(uint256 parentTokenId, address receiver, address creator, string uri) returns()
func (_Childnft *ChildnftTransactor) MintChildNFTWithURI(opts *bind.TransactOpts, parentTokenId *big.Int, receiver common.Address, creator common.Address, uri string) (*types.Transaction, error) {
	return _Childnft.contract.Transact(opts, "mintChildNFTWithURI", parentTokenId, receiver, creator, uri)
}

// MintChildNFTWithURI is a paid mutator transaction binding the contract method 0xda682f3d.
//
// Solidity: function mintChildNFTWithURI(uint256 parentTokenId, address receiver, address creator, string uri) returns()
func (_Childnft *ChildnftSession) MintChildNFTWithURI(parentTokenId *big.Int, receiver common.Address, creator common.Address, uri string) (*types.Transaction, error) {
	return _Childnft.Contract.MintChildNFTWithURI(&_Childnft.TransactOpts, parentTokenId, receiver, creator, uri)
}

// MintChildNFTWithURI is a paid mutator transaction binding the contract method 0xda682f3d.
//
// Solidity: function mintChildNFTWithURI(uint256 parentTokenId, address receiver, address creator, string uri) returns()
func (_Childnft *ChildnftTransactorSession) MintChildNFTWithURI(parentTokenId *big.Int, receiver common.Address, creator common.Address, uri string) (*types.Transaction, error) {
	return _Childnft.Contract.MintChildNFTWithURI(&_Childnft.TransactOpts, parentTokenId, receiver, creator, uri)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Childnft *ChildnftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Childnft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Childnft *ChildnftSession) RenounceOwnership() (*types.Transaction, error) {
	return _Childnft.Contract.RenounceOwnership(&_Childnft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Childnft *ChildnftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Childnft.Contract.RenounceOwnership(&_Childnft.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Childnft *ChildnftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Childnft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Childnft *ChildnftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Childnft.Contract.SafeTransferFrom(&_Childnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Childnft *ChildnftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Childnft.Contract.SafeTransferFrom(&_Childnft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Childnft *ChildnftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Childnft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Childnft *ChildnftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Childnft.Contract.SafeTransferFrom0(&_Childnft.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Childnft *ChildnftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Childnft.Contract.SafeTransferFrom0(&_Childnft.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Childnft *ChildnftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Childnft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Childnft *ChildnftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Childnft.Contract.SetApprovalForAll(&_Childnft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Childnft *ChildnftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Childnft.Contract.SetApprovalForAll(&_Childnft.TransactOpts, operator, approved)
}

// SetChildTokenURI is a paid mutator transaction binding the contract method 0xd738870e.
//
// Solidity: function setChildTokenURI(string newURI) returns()
func (_Childnft *ChildnftTransactor) SetChildTokenURI(opts *bind.TransactOpts, newURI string) (*types.Transaction, error) {
	return _Childnft.contract.Transact(opts, "setChildTokenURI", newURI)
}

// SetChildTokenURI is a paid mutator transaction binding the contract method 0xd738870e.
//
// Solidity: function setChildTokenURI(string newURI) returns()
func (_Childnft *ChildnftSession) SetChildTokenURI(newURI string) (*types.Transaction, error) {
	return _Childnft.Contract.SetChildTokenURI(&_Childnft.TransactOpts, newURI)
}

// SetChildTokenURI is a paid mutator transaction binding the contract method 0xd738870e.
//
// Solidity: function setChildTokenURI(string newURI) returns()
func (_Childnft *ChildnftTransactorSession) SetChildTokenURI(newURI string) (*types.Transaction, error) {
	return _Childnft.Contract.SetChildTokenURI(&_Childnft.TransactOpts, newURI)
}

// SetMainNFTContract is a paid mutator transaction binding the contract method 0x1a83fa4c.
//
// Solidity: function setMainNFTContract(address _mainNFTContract) returns()
func (_Childnft *ChildnftTransactor) SetMainNFTContract(opts *bind.TransactOpts, _mainNFTContract common.Address) (*types.Transaction, error) {
	return _Childnft.contract.Transact(opts, "setMainNFTContract", _mainNFTContract)
}

// SetMainNFTContract is a paid mutator transaction binding the contract method 0x1a83fa4c.
//
// Solidity: function setMainNFTContract(address _mainNFTContract) returns()
func (_Childnft *ChildnftSession) SetMainNFTContract(_mainNFTContract common.Address) (*types.Transaction, error) {
	return _Childnft.Contract.SetMainNFTContract(&_Childnft.TransactOpts, _mainNFTContract)
}

// SetMainNFTContract is a paid mutator transaction binding the contract method 0x1a83fa4c.
//
// Solidity: function setMainNFTContract(address _mainNFTContract) returns()
func (_Childnft *ChildnftTransactorSession) SetMainNFTContract(_mainNFTContract common.Address) (*types.Transaction, error) {
	return _Childnft.Contract.SetMainNFTContract(&_Childnft.TransactOpts, _mainNFTContract)
}

// SetSpecificTokenURI is a paid mutator transaction binding the contract method 0xf655f04c.
//
// Solidity: function setSpecificTokenURI(uint256 tokenId, string newTokenURI) returns()
func (_Childnft *ChildnftTransactor) SetSpecificTokenURI(opts *bind.TransactOpts, tokenId *big.Int, newTokenURI string) (*types.Transaction, error) {
	return _Childnft.contract.Transact(opts, "setSpecificTokenURI", tokenId, newTokenURI)
}

// SetSpecificTokenURI is a paid mutator transaction binding the contract method 0xf655f04c.
//
// Solidity: function setSpecificTokenURI(uint256 tokenId, string newTokenURI) returns()
func (_Childnft *ChildnftSession) SetSpecificTokenURI(tokenId *big.Int, newTokenURI string) (*types.Transaction, error) {
	return _Childnft.Contract.SetSpecificTokenURI(&_Childnft.TransactOpts, tokenId, newTokenURI)
}

// SetSpecificTokenURI is a paid mutator transaction binding the contract method 0xf655f04c.
//
// Solidity: function setSpecificTokenURI(uint256 tokenId, string newTokenURI) returns()
func (_Childnft *ChildnftTransactorSession) SetSpecificTokenURI(tokenId *big.Int, newTokenURI string) (*types.Transaction, error) {
	return _Childnft.Contract.SetSpecificTokenURI(&_Childnft.TransactOpts, tokenId, newTokenURI)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Childnft *ChildnftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Childnft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Childnft *ChildnftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Childnft.Contract.TransferFrom(&_Childnft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Childnft *ChildnftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Childnft.Contract.TransferFrom(&_Childnft.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Childnft *ChildnftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Childnft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Childnft *ChildnftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Childnft.Contract.TransferOwnership(&_Childnft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Childnft *ChildnftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Childnft.Contract.TransferOwnership(&_Childnft.TransactOpts, newOwner)
}

// ChildnftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Childnft contract.
type ChildnftApprovalIterator struct {
	Event *ChildnftApproval // Event containing the contract specifics and raw log

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
func (it *ChildnftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildnftApproval)
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
		it.Event = new(ChildnftApproval)
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
func (it *ChildnftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildnftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildnftApproval represents a Approval event raised by the Childnft contract.
type ChildnftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Childnft *ChildnftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ChildnftApprovalIterator, error) {

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

	logs, sub, err := _Childnft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ChildnftApprovalIterator{contract: _Childnft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Childnft *ChildnftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ChildnftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Childnft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildnftApproval)
				if err := _Childnft.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Childnft *ChildnftFilterer) ParseApproval(log types.Log) (*ChildnftApproval, error) {
	event := new(ChildnftApproval)
	if err := _Childnft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildnftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Childnft contract.
type ChildnftApprovalForAllIterator struct {
	Event *ChildnftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ChildnftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildnftApprovalForAll)
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
		it.Event = new(ChildnftApprovalForAll)
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
func (it *ChildnftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildnftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildnftApprovalForAll represents a ApprovalForAll event raised by the Childnft contract.
type ChildnftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Childnft *ChildnftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ChildnftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Childnft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ChildnftApprovalForAllIterator{contract: _Childnft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Childnft *ChildnftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ChildnftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Childnft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildnftApprovalForAll)
				if err := _Childnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_Childnft *ChildnftFilterer) ParseApprovalForAll(log types.Log) (*ChildnftApprovalForAll, error) {
	event := new(ChildnftApprovalForAll)
	if err := _Childnft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildnftBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Childnft contract.
type ChildnftBatchMetadataUpdateIterator struct {
	Event *ChildnftBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ChildnftBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildnftBatchMetadataUpdate)
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
		it.Event = new(ChildnftBatchMetadataUpdate)
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
func (it *ChildnftBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildnftBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildnftBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Childnft contract.
type ChildnftBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Childnft *ChildnftFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*ChildnftBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Childnft.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ChildnftBatchMetadataUpdateIterator{contract: _Childnft.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Childnft *ChildnftFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ChildnftBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Childnft.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildnftBatchMetadataUpdate)
				if err := _Childnft.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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
func (_Childnft *ChildnftFilterer) ParseBatchMetadataUpdate(log types.Log) (*ChildnftBatchMetadataUpdate, error) {
	event := new(ChildnftBatchMetadataUpdate)
	if err := _Childnft.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildnftChildTokenMintedIterator is returned from FilterChildTokenMinted and is used to iterate over the raw logs and unpacked data for ChildTokenMinted events raised by the Childnft contract.
type ChildnftChildTokenMintedIterator struct {
	Event *ChildnftChildTokenMinted // Event containing the contract specifics and raw log

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
func (it *ChildnftChildTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildnftChildTokenMinted)
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
		it.Event = new(ChildnftChildTokenMinted)
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
func (it *ChildnftChildTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildnftChildTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildnftChildTokenMinted represents a ChildTokenMinted event raised by the Childnft contract.
type ChildnftChildTokenMinted struct {
	ParentTokenId *big.Int
	ChildTokenId  *big.Int
	Receiver      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterChildTokenMinted is a free log retrieval operation binding the contract event 0x622d1d28c81c1622c1194f8d132fbf1fc580cb75e5ca86dda9b7987c60fbe764.
//
// Solidity: event ChildTokenMinted(uint256 indexed parentTokenId, uint256 indexed childTokenId, address indexed receiver)
func (_Childnft *ChildnftFilterer) FilterChildTokenMinted(opts *bind.FilterOpts, parentTokenId []*big.Int, childTokenId []*big.Int, receiver []common.Address) (*ChildnftChildTokenMintedIterator, error) {

	var parentTokenIdRule []interface{}
	for _, parentTokenIdItem := range parentTokenId {
		parentTokenIdRule = append(parentTokenIdRule, parentTokenIdItem)
	}
	var childTokenIdRule []interface{}
	for _, childTokenIdItem := range childTokenId {
		childTokenIdRule = append(childTokenIdRule, childTokenIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Childnft.contract.FilterLogs(opts, "ChildTokenMinted", parentTokenIdRule, childTokenIdRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ChildnftChildTokenMintedIterator{contract: _Childnft.contract, event: "ChildTokenMinted", logs: logs, sub: sub}, nil
}

// WatchChildTokenMinted is a free log subscription operation binding the contract event 0x622d1d28c81c1622c1194f8d132fbf1fc580cb75e5ca86dda9b7987c60fbe764.
//
// Solidity: event ChildTokenMinted(uint256 indexed parentTokenId, uint256 indexed childTokenId, address indexed receiver)
func (_Childnft *ChildnftFilterer) WatchChildTokenMinted(opts *bind.WatchOpts, sink chan<- *ChildnftChildTokenMinted, parentTokenId []*big.Int, childTokenId []*big.Int, receiver []common.Address) (event.Subscription, error) {

	var parentTokenIdRule []interface{}
	for _, parentTokenIdItem := range parentTokenId {
		parentTokenIdRule = append(parentTokenIdRule, parentTokenIdItem)
	}
	var childTokenIdRule []interface{}
	for _, childTokenIdItem := range childTokenId {
		childTokenIdRule = append(childTokenIdRule, childTokenIdItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Childnft.contract.WatchLogs(opts, "ChildTokenMinted", parentTokenIdRule, childTokenIdRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildnftChildTokenMinted)
				if err := _Childnft.contract.UnpackLog(event, "ChildTokenMinted", log); err != nil {
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

// ParseChildTokenMinted is a log parse operation binding the contract event 0x622d1d28c81c1622c1194f8d132fbf1fc580cb75e5ca86dda9b7987c60fbe764.
//
// Solidity: event ChildTokenMinted(uint256 indexed parentTokenId, uint256 indexed childTokenId, address indexed receiver)
func (_Childnft *ChildnftFilterer) ParseChildTokenMinted(log types.Log) (*ChildnftChildTokenMinted, error) {
	event := new(ChildnftChildTokenMinted)
	if err := _Childnft.contract.UnpackLog(event, "ChildTokenMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildnftMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Childnft contract.
type ChildnftMetadataUpdateIterator struct {
	Event *ChildnftMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *ChildnftMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildnftMetadataUpdate)
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
		it.Event = new(ChildnftMetadataUpdate)
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
func (it *ChildnftMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildnftMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildnftMetadataUpdate represents a MetadataUpdate event raised by the Childnft contract.
type ChildnftMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Childnft *ChildnftFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*ChildnftMetadataUpdateIterator, error) {

	logs, sub, err := _Childnft.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &ChildnftMetadataUpdateIterator{contract: _Childnft.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Childnft *ChildnftFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *ChildnftMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Childnft.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildnftMetadataUpdate)
				if err := _Childnft.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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
func (_Childnft *ChildnftFilterer) ParseMetadataUpdate(log types.Log) (*ChildnftMetadataUpdate, error) {
	event := new(ChildnftMetadataUpdate)
	if err := _Childnft.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildnftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Childnft contract.
type ChildnftOwnershipTransferredIterator struct {
	Event *ChildnftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ChildnftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildnftOwnershipTransferred)
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
		it.Event = new(ChildnftOwnershipTransferred)
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
func (it *ChildnftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildnftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildnftOwnershipTransferred represents a OwnershipTransferred event raised by the Childnft contract.
type ChildnftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Childnft *ChildnftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ChildnftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Childnft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ChildnftOwnershipTransferredIterator{contract: _Childnft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Childnft *ChildnftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ChildnftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Childnft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildnftOwnershipTransferred)
				if err := _Childnft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Childnft *ChildnftFilterer) ParseOwnershipTransferred(log types.Log) (*ChildnftOwnershipTransferred, error) {
	event := new(ChildnftOwnershipTransferred)
	if err := _Childnft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChildnftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Childnft contract.
type ChildnftTransferIterator struct {
	Event *ChildnftTransfer // Event containing the contract specifics and raw log

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
func (it *ChildnftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChildnftTransfer)
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
		it.Event = new(ChildnftTransfer)
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
func (it *ChildnftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChildnftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChildnftTransfer represents a Transfer event raised by the Childnft contract.
type ChildnftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Childnft *ChildnftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ChildnftTransferIterator, error) {

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

	logs, sub, err := _Childnft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ChildnftTransferIterator{contract: _Childnft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Childnft *ChildnftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ChildnftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _Childnft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChildnftTransfer)
				if err := _Childnft.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Childnft *ChildnftFilterer) ParseTransfer(log types.Log) (*ChildnftTransfer, error) {
	event := new(ChildnftTransfer)
	if err := _Childnft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
