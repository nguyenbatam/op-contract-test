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

// ControllerDataDoc is an auto generated low-level Go binding around an user-defined struct.
type ControllerDataDoc struct {
	Id          string
	HashContent string
}

// ControllerUploadSession is an auto generated low-level Go binding around an user-defined struct.
type ControllerUploadSession struct {
	Id        *big.Int
	User      common.Address
	Proof     [][32]byte
	Confirmed bool
}

// ControllerMetaData contains all meta data concerning the Controller contract.
var ControllerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pcspAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"riskScore\",\"type\":\"uint256\"}],\"name\":\"Confirm\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"}],\"name\":\"UploadData\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"contentHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"riskScore\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"name\":\"confirm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"geneNFT\",\"outputs\":[{\"internalType\":\"contractGeneNFT\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"}],\"name\":\"getDoc\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"hashContent\",\"type\":\"string\"}],\"internalType\":\"structController.DataDoc\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sessionId\",\"type\":\"uint256\"}],\"name\":\"getSession\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"confirmed\",\"type\":\"bool\"}],\"internalType\":\"structController.UploadSession\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"date\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"_merkleRoot\",\"type\":\"bytes32\"}],\"name\":\"initializeMerkleRoot\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"merkleRootDates\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pcspToken\",\"outputs\":[{\"internalType\":\"contractPostCovidStrokePrevention\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"docId\",\"type\":\"string\"}],\"name\":\"uploadData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"leaf\",\"type\":\"bytes32\"}],\"name\":\"verifyProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
}

// ControllerABI is the input ABI used to generate the binding from.
// Deprecated: Use ControllerMetaData.ABI instead.
var ControllerABI = ControllerMetaData.ABI

// Controller is an auto generated Go binding around an Ethereum contract.
type Controller struct {
	ControllerCaller     // Read-only binding to the contract
	ControllerTransactor // Write-only binding to the contract
	ControllerFilterer   // Log filterer for contract events
}

// ControllerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ControllerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ControllerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ControllerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ControllerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ControllerSession struct {
	Contract     *Controller       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ControllerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ControllerCallerSession struct {
	Contract *ControllerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ControllerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ControllerTransactorSession struct {
	Contract     *ControllerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ControllerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ControllerRaw struct {
	Contract *Controller // Generic contract binding to access the raw methods on
}

// ControllerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ControllerCallerRaw struct {
	Contract *ControllerCaller // Generic read-only contract binding to access the raw methods on
}

// ControllerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ControllerTransactorRaw struct {
	Contract *ControllerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewController creates a new instance of Controller, bound to a specific deployed contract.
func NewController(address common.Address, backend bind.ContractBackend) (*Controller, error) {
	contract, err := bindController(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Controller{ControllerCaller: ControllerCaller{contract: contract}, ControllerTransactor: ControllerTransactor{contract: contract}, ControllerFilterer: ControllerFilterer{contract: contract}}, nil
}

// NewControllerCaller creates a new read-only instance of Controller, bound to a specific deployed contract.
func NewControllerCaller(address common.Address, caller bind.ContractCaller) (*ControllerCaller, error) {
	contract, err := bindController(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerCaller{contract: contract}, nil
}

// NewControllerTransactor creates a new write-only instance of Controller, bound to a specific deployed contract.
func NewControllerTransactor(address common.Address, transactor bind.ContractTransactor) (*ControllerTransactor, error) {
	contract, err := bindController(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ControllerTransactor{contract: contract}, nil
}

// NewControllerFilterer creates a new log filterer instance of Controller, bound to a specific deployed contract.
func NewControllerFilterer(address common.Address, filterer bind.ContractFilterer) (*ControllerFilterer, error) {
	contract, err := bindController(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ControllerFilterer{contract: contract}, nil
}

// bindController binds a generic wrapper to an already deployed contract.
func bindController(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ControllerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.ControllerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.ControllerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Controller *ControllerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Controller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Controller *ControllerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Controller *ControllerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Controller.Contract.contract.Transact(opts, method, params...)
}

// GeneNFT is a free data retrieval call binding the contract method 0x5231f627.
//
// Solidity: function geneNFT() view returns(address)
func (_Controller *ControllerCaller) GeneNFT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "geneNFT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GeneNFT is a free data retrieval call binding the contract method 0x5231f627.
//
// Solidity: function geneNFT() view returns(address)
func (_Controller *ControllerSession) GeneNFT() (common.Address, error) {
	return _Controller.Contract.GeneNFT(&_Controller.CallOpts)
}

// GeneNFT is a free data retrieval call binding the contract method 0x5231f627.
//
// Solidity: function geneNFT() view returns(address)
func (_Controller *ControllerCallerSession) GeneNFT() (common.Address, error) {
	return _Controller.Contract.GeneNFT(&_Controller.CallOpts)
}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string docId) view returns((string,string))
func (_Controller *ControllerCaller) GetDoc(opts *bind.CallOpts, docId string) (ControllerDataDoc, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getDoc", docId)

	if err != nil {
		return *new(ControllerDataDoc), err
	}

	out0 := *abi.ConvertType(out[0], new(ControllerDataDoc)).(*ControllerDataDoc)

	return out0, err

}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string docId) view returns((string,string))
func (_Controller *ControllerSession) GetDoc(docId string) (ControllerDataDoc, error) {
	return _Controller.Contract.GetDoc(&_Controller.CallOpts, docId)
}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string docId) view returns((string,string))
func (_Controller *ControllerCallerSession) GetDoc(docId string) (ControllerDataDoc, error) {
	return _Controller.Contract.GetDoc(&_Controller.CallOpts, docId)
}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 sessionId) view returns((uint256,address,bytes32[],bool))
func (_Controller *ControllerCaller) GetSession(opts *bind.CallOpts, sessionId *big.Int) (ControllerUploadSession, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "getSession", sessionId)

	if err != nil {
		return *new(ControllerUploadSession), err
	}

	out0 := *abi.ConvertType(out[0], new(ControllerUploadSession)).(*ControllerUploadSession)

	return out0, err

}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 sessionId) view returns((uint256,address,bytes32[],bool))
func (_Controller *ControllerSession) GetSession(sessionId *big.Int) (ControllerUploadSession, error) {
	return _Controller.Contract.GetSession(&_Controller.CallOpts, sessionId)
}

// GetSession is a free data retrieval call binding the contract method 0x402ff0db.
//
// Solidity: function getSession(uint256 sessionId) view returns((uint256,address,bytes32[],bool))
func (_Controller *ControllerCallerSession) GetSession(sessionId *big.Int) (ControllerUploadSession, error) {
	return _Controller.Contract.GetSession(&_Controller.CallOpts, sessionId)
}

// MerkleRootDates is a free data retrieval call binding the contract method 0x92ac7f5f.
//
// Solidity: function merkleRootDates(string ) view returns(bytes32)
func (_Controller *ControllerCaller) MerkleRootDates(opts *bind.CallOpts, arg0 string) ([32]byte, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "merkleRootDates", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MerkleRootDates is a free data retrieval call binding the contract method 0x92ac7f5f.
//
// Solidity: function merkleRootDates(string ) view returns(bytes32)
func (_Controller *ControllerSession) MerkleRootDates(arg0 string) ([32]byte, error) {
	return _Controller.Contract.MerkleRootDates(&_Controller.CallOpts, arg0)
}

// MerkleRootDates is a free data retrieval call binding the contract method 0x92ac7f5f.
//
// Solidity: function merkleRootDates(string ) view returns(bytes32)
func (_Controller *ControllerCallerSession) MerkleRootDates(arg0 string) ([32]byte, error) {
	return _Controller.Contract.MerkleRootDates(&_Controller.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Controller *ControllerCallerSession) Owner() (common.Address, error) {
	return _Controller.Contract.Owner(&_Controller.CallOpts)
}

// PcspToken is a free data retrieval call binding the contract method 0xdab3761e.
//
// Solidity: function pcspToken() view returns(address)
func (_Controller *ControllerCaller) PcspToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "pcspToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PcspToken is a free data retrieval call binding the contract method 0xdab3761e.
//
// Solidity: function pcspToken() view returns(address)
func (_Controller *ControllerSession) PcspToken() (common.Address, error) {
	return _Controller.Contract.PcspToken(&_Controller.CallOpts)
}

// PcspToken is a free data retrieval call binding the contract method 0xdab3761e.
//
// Solidity: function pcspToken() view returns(address)
func (_Controller *ControllerCallerSession) PcspToken() (common.Address, error) {
	return _Controller.Contract.PcspToken(&_Controller.CallOpts)
}

// VerifyProof is a free data retrieval call binding the contract method 0x4bc935d7.
//
// Solidity: function verifyProof(bytes32[] proof, bytes32 merkleRoot, bytes32 leaf) pure returns(bool)
func (_Controller *ControllerCaller) VerifyProof(opts *bind.CallOpts, proof [][32]byte, merkleRoot [32]byte, leaf [32]byte) (bool, error) {
	var out []interface{}
	err := _Controller.contract.Call(opts, &out, "verifyProof", proof, merkleRoot, leaf)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyProof is a free data retrieval call binding the contract method 0x4bc935d7.
//
// Solidity: function verifyProof(bytes32[] proof, bytes32 merkleRoot, bytes32 leaf) pure returns(bool)
func (_Controller *ControllerSession) VerifyProof(proof [][32]byte, merkleRoot [32]byte, leaf [32]byte) (bool, error) {
	return _Controller.Contract.VerifyProof(&_Controller.CallOpts, proof, merkleRoot, leaf)
}

// VerifyProof is a free data retrieval call binding the contract method 0x4bc935d7.
//
// Solidity: function verifyProof(bytes32[] proof, bytes32 merkleRoot, bytes32 leaf) pure returns(bool)
func (_Controller *ControllerCallerSession) VerifyProof(proof [][32]byte, merkleRoot [32]byte, leaf [32]byte) (bool, error) {
	return _Controller.Contract.VerifyProof(&_Controller.CallOpts, proof, merkleRoot, leaf)
}

// Confirm is a paid mutator transaction binding the contract method 0x4e0c0e26.
//
// Solidity: function confirm(string docId, string contentHash, uint256 sessionId, uint256 riskScore, string date, bytes32[] proof, bytes32 leaf) returns()
func (_Controller *ControllerTransactor) Confirm(opts *bind.TransactOpts, docId string, contentHash string, sessionId *big.Int, riskScore *big.Int, date string, proof [][32]byte, leaf [32]byte) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "confirm", docId, contentHash, sessionId, riskScore, date, proof, leaf)
}

// Confirm is a paid mutator transaction binding the contract method 0x4e0c0e26.
//
// Solidity: function confirm(string docId, string contentHash, uint256 sessionId, uint256 riskScore, string date, bytes32[] proof, bytes32 leaf) returns()
func (_Controller *ControllerSession) Confirm(docId string, contentHash string, sessionId *big.Int, riskScore *big.Int, date string, proof [][32]byte, leaf [32]byte) (*types.Transaction, error) {
	return _Controller.Contract.Confirm(&_Controller.TransactOpts, docId, contentHash, sessionId, riskScore, date, proof, leaf)
}

// Confirm is a paid mutator transaction binding the contract method 0x4e0c0e26.
//
// Solidity: function confirm(string docId, string contentHash, uint256 sessionId, uint256 riskScore, string date, bytes32[] proof, bytes32 leaf) returns()
func (_Controller *ControllerTransactorSession) Confirm(docId string, contentHash string, sessionId *big.Int, riskScore *big.Int, date string, proof [][32]byte, leaf [32]byte) (*types.Transaction, error) {
	return _Controller.Contract.Confirm(&_Controller.TransactOpts, docId, contentHash, sessionId, riskScore, date, proof, leaf)
}

// InitializeMerkleRoot is a paid mutator transaction binding the contract method 0xa7951bb9.
//
// Solidity: function initializeMerkleRoot(string date, bytes32 _merkleRoot) returns()
func (_Controller *ControllerTransactor) InitializeMerkleRoot(opts *bind.TransactOpts, date string, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "initializeMerkleRoot", date, _merkleRoot)
}

// InitializeMerkleRoot is a paid mutator transaction binding the contract method 0xa7951bb9.
//
// Solidity: function initializeMerkleRoot(string date, bytes32 _merkleRoot) returns()
func (_Controller *ControllerSession) InitializeMerkleRoot(date string, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _Controller.Contract.InitializeMerkleRoot(&_Controller.TransactOpts, date, _merkleRoot)
}

// InitializeMerkleRoot is a paid mutator transaction binding the contract method 0xa7951bb9.
//
// Solidity: function initializeMerkleRoot(string date, bytes32 _merkleRoot) returns()
func (_Controller *ControllerTransactorSession) InitializeMerkleRoot(date string, _merkleRoot [32]byte) (*types.Transaction, error) {
	return _Controller.Contract.InitializeMerkleRoot(&_Controller.TransactOpts, date, _merkleRoot)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Controller.Contract.RenounceOwnership(&_Controller.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Controller *ControllerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Controller.Contract.RenounceOwnership(&_Controller.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Controller *ControllerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Controller *ControllerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Controller.Contract.TransferOwnership(&_Controller.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Controller *ControllerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Controller.Contract.TransferOwnership(&_Controller.TransactOpts, newOwner)
}

// UploadData is a paid mutator transaction binding the contract method 0x50969f44.
//
// Solidity: function uploadData(string docId) returns(uint256)
func (_Controller *ControllerTransactor) UploadData(opts *bind.TransactOpts, docId string) (*types.Transaction, error) {
	return _Controller.contract.Transact(opts, "uploadData", docId)
}

// UploadData is a paid mutator transaction binding the contract method 0x50969f44.
//
// Solidity: function uploadData(string docId) returns(uint256)
func (_Controller *ControllerSession) UploadData(docId string) (*types.Transaction, error) {
	return _Controller.Contract.UploadData(&_Controller.TransactOpts, docId)
}

// UploadData is a paid mutator transaction binding the contract method 0x50969f44.
//
// Solidity: function uploadData(string docId) returns(uint256)
func (_Controller *ControllerTransactorSession) UploadData(docId string) (*types.Transaction, error) {
	return _Controller.Contract.UploadData(&_Controller.TransactOpts, docId)
}

// ControllerConfirmIterator is returned from FilterConfirm and is used to iterate over the raw logs and unpacked data for Confirm events raised by the Controller contract.
type ControllerConfirmIterator struct {
	Event *ControllerConfirm // Event containing the contract specifics and raw log

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
func (it *ControllerConfirmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerConfirm)
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
		it.Event = new(ControllerConfirm)
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
func (it *ControllerConfirmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerConfirmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerConfirm represents a Confirm event raised by the Controller contract.
type ControllerConfirm struct {
	DocId     string
	SessionId *big.Int
	RiskScore *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterConfirm is a free log retrieval operation binding the contract event 0xdaa0a0d8915f73504d446b0a1bc73a101a113cc357d88e31b16459e9ce50fc0f.
//
// Solidity: event Confirm(string docId, uint256 sessionId, uint256 riskScore)
func (_Controller *ControllerFilterer) FilterConfirm(opts *bind.FilterOpts) (*ControllerConfirmIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "Confirm")
	if err != nil {
		return nil, err
	}
	return &ControllerConfirmIterator{contract: _Controller.contract, event: "Confirm", logs: logs, sub: sub}, nil
}

// WatchConfirm is a free log subscription operation binding the contract event 0xdaa0a0d8915f73504d446b0a1bc73a101a113cc357d88e31b16459e9ce50fc0f.
//
// Solidity: event Confirm(string docId, uint256 sessionId, uint256 riskScore)
func (_Controller *ControllerFilterer) WatchConfirm(opts *bind.WatchOpts, sink chan<- *ControllerConfirm) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "Confirm")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerConfirm)
				if err := _Controller.contract.UnpackLog(event, "Confirm", log); err != nil {
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

// ParseConfirm is a log parse operation binding the contract event 0xdaa0a0d8915f73504d446b0a1bc73a101a113cc357d88e31b16459e9ce50fc0f.
//
// Solidity: event Confirm(string docId, uint256 sessionId, uint256 riskScore)
func (_Controller *ControllerFilterer) ParseConfirm(log types.Log) (*ControllerConfirm, error) {
	event := new(ControllerConfirm)
	if err := _Controller.contract.UnpackLog(event, "Confirm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Controller contract.
type ControllerOwnershipTransferredIterator struct {
	Event *ControllerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ControllerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerOwnershipTransferred)
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
		it.Event = new(ControllerOwnershipTransferred)
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
func (it *ControllerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerOwnershipTransferred represents a OwnershipTransferred event raised by the Controller contract.
type ControllerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Controller *ControllerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ControllerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Controller.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ControllerOwnershipTransferredIterator{contract: _Controller.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Controller *ControllerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ControllerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Controller.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerOwnershipTransferred)
				if err := _Controller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Controller *ControllerFilterer) ParseOwnershipTransferred(log types.Log) (*ControllerOwnershipTransferred, error) {
	event := new(ControllerOwnershipTransferred)
	if err := _Controller.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ControllerUploadDataIterator is returned from FilterUploadData and is used to iterate over the raw logs and unpacked data for UploadData events raised by the Controller contract.
type ControllerUploadDataIterator struct {
	Event *ControllerUploadData // Event containing the contract specifics and raw log

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
func (it *ControllerUploadDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ControllerUploadData)
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
		it.Event = new(ControllerUploadData)
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
func (it *ControllerUploadDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ControllerUploadDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ControllerUploadData represents a UploadData event raised by the Controller contract.
type ControllerUploadData struct {
	DocId     string
	SessionId *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUploadData is a free log retrieval operation binding the contract event 0x698b35ede3baa51dbaa3b9a040c287690e40d0101d312c80eb364c7b17c458bc.
//
// Solidity: event UploadData(string docId, uint256 sessionId)
func (_Controller *ControllerFilterer) FilterUploadData(opts *bind.FilterOpts) (*ControllerUploadDataIterator, error) {

	logs, sub, err := _Controller.contract.FilterLogs(opts, "UploadData")
	if err != nil {
		return nil, err
	}
	return &ControllerUploadDataIterator{contract: _Controller.contract, event: "UploadData", logs: logs, sub: sub}, nil
}

// WatchUploadData is a free log subscription operation binding the contract event 0x698b35ede3baa51dbaa3b9a040c287690e40d0101d312c80eb364c7b17c458bc.
//
// Solidity: event UploadData(string docId, uint256 sessionId)
func (_Controller *ControllerFilterer) WatchUploadData(opts *bind.WatchOpts, sink chan<- *ControllerUploadData) (event.Subscription, error) {

	logs, sub, err := _Controller.contract.WatchLogs(opts, "UploadData")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ControllerUploadData)
				if err := _Controller.contract.UnpackLog(event, "UploadData", log); err != nil {
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

// ParseUploadData is a log parse operation binding the contract event 0x698b35ede3baa51dbaa3b9a040c287690e40d0101d312c80eb364c7b17c458bc.
//
// Solidity: event UploadData(string docId, uint256 sessionId)
func (_Controller *ControllerFilterer) ParseUploadData(log types.Log) (*ControllerUploadData, error) {
	event := new(ControllerUploadData)
	if err := _Controller.contract.UnpackLog(event, "UploadData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
