// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package himera_avs

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

// AVSParams is an auto generated low-level Go binding around an user-defined struct.
type AVSParams struct {
	Sender              common.Address
	AvsName             string
	MinStakeAmount      uint64
	TaskAddress         common.Address
	SlashAddress        common.Address
	RewardAddress       common.Address
	AvsOwnerAddresses   []common.Address
	WhitelistAddresses  []common.Address
	AssetIDs            []string
	AvsUnbondingPeriod  uint64
	MinSelfDelegation   uint64
	EpochIdentifier     string
	MiniOptInOperators  uint64
	MinTotalStakeAmount uint64
	AvsRewardProportion uint64
	AvsSlashProportion  uint64
}

// HimeraTaskLibraryTaskDefinition is an auto generated low-level Go binding around an user-defined struct.
type HimeraTaskLibraryTaskDefinition struct {
	Id          uint8
	Name        string
	TaskType    uint8
	Description string
}

// OperatorActivePower is an auto generated low-level Go binding around an user-defined struct.
type OperatorActivePower struct {
	Operator common.Address
	Power    *big.Int
}

// TaskInfo is an auto generated low-level Go binding around an user-defined struct.
type TaskInfo struct {
	TaskContractAddress     common.Address
	Name                    string
	Hash                    []byte
	TaskID                  uint64
	TaskResponsePeriod      uint64
	TaskStatisticalPeriod   uint64
	TaskChallengePeriod     uint64
	ThresholdPercentage     uint8
	StartingEpoch           uint64
	ActualThreshold         string
	OptInOperators          []common.Address
	SignedOperators         []common.Address
	NoSignedOperators       []common.Address
	ErrSignedOperators      []common.Address
	TaskTotalPower          string
	OperatorActivePower     []OperatorActivePower
	IsExpected              bool
	EligibleRewardOperators []common.Address
	EligibleSlashOperators  []common.Address
}

// ContractHimeraAvsMetaData contains all meta data concerning the ContractHimeraAvs contract.
var ContractHimeraAvsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"challenge\",\"inputs\":[{\"name\":\"taskID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actualThreshold\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"isExpected\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"eligibleRewardOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"eligibleSlashOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"challengerAddress\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createHimeraTask\",\"inputs\":[{\"name\":\"himeraTaskDefId\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"taskInput\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"imuaTaskId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deregisterOperatorFromAVS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOptInOperators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskDefinition\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structHimeraTaskLibrary.TaskDefinition\",\"components\":[{\"name\":\"id\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"taskType\",\"type\":\"uint8\",\"internalType\":\"enumHimeraTaskLibrary.HimeraTaskType\"},{\"name\":\"description\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTaskInfo\",\"inputs\":[{\"name\":\"taskID\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTaskInfo\",\"components\":[{\"name\":\"taskContractAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"hash\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"taskID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"taskResponsePeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"taskStatisticalPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"taskChallengePeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"thresholdPercentage\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"startingEpoch\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"actualThreshold\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"optInOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"signedOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"noSignedOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"errSignedOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"taskTotalPower\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"operatorActivePower\",\"type\":\"tuple[]\",\"internalType\":\"structOperatorActivePower[]\",\"components\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"power\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"isExpected\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"eligibleRewardOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"eligibleSlashOperators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"operatorSubmitTask\",\"inputs\":[{\"name\":\"taskID\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"taskResponse\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"blsSignature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"phase\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerAVS\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structAVSParams\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"minStakeAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"taskAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slashAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsOwnerAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"whitelistAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"assetIDs\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minSelfDelegation\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochIdentifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"miniOptInOperators\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minTotalStakeAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"avsRewardProportion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"avsSlashProportion\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerBLSPublicKey\",\"inputs\":[{\"name\":\"pubKey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"pubKeyRegSig\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"registerOperatorToAVS\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setChallenger\",\"inputs\":[{\"name\":\"_challenger\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardManager\",\"inputs\":[{\"name\":\"_rewardManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSlasher\",\"inputs\":[{\"name\":\"_slasher\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setupTaskDefinitions\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateAVS\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple\",\"internalType\":\"structAVSParams\",\"components\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"minStakeAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"taskAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slashAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rewardAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"avsOwnerAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"whitelistAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"assetIDs\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"avsUnbondingPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minSelfDelegation\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"epochIdentifier\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"miniOptInOperators\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"minTotalStakeAmount\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"avsRewardProportion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"avsSlashProportion\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"BLSPublicKeyRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"avsAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubKey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChallengeSubmitted\",\"inputs\":[{\"name\":\"taskID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"challenger\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"isExpected\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"HimeraTaskCreated\",\"inputs\":[{\"name\":\"imuaTaskId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"definitionHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"himeraTaskDefId\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"taskInput\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorOptedIn\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorOptedOut\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RewardManagerUpdated\",\"inputs\":[{\"name\":\"newRewardManager\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SlasherUpdated\",\"inputs\":[{\"name\":\"newSlasher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskDefinitionCreated\",\"inputs\":[{\"name\":\"taskDefinitionId\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"taskType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumHimeraTaskLibrary.HimeraTaskType\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaskSubmitted\",\"inputs\":[{\"name\":\"taskID\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"phase\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// ContractHimeraAvsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractHimeraAvsMetaData.ABI instead.
var ContractHimeraAvsABI = ContractHimeraAvsMetaData.ABI

// ContractHimeraAvs is an auto generated Go binding around an Ethereum contract.
type ContractHimeraAvs struct {
	ContractHimeraAvsCaller     // Read-only binding to the contract
	ContractHimeraAvsTransactor // Write-only binding to the contract
	ContractHimeraAvsFilterer   // Log filterer for contract events
}

// ContractHimeraAvsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractHimeraAvsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractHimeraAvsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractHimeraAvsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractHimeraAvsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractHimeraAvsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractHimeraAvsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractHimeraAvsSession struct {
	Contract     *ContractHimeraAvs // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContractHimeraAvsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractHimeraAvsCallerSession struct {
	Contract *ContractHimeraAvsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ContractHimeraAvsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractHimeraAvsTransactorSession struct {
	Contract     *ContractHimeraAvsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ContractHimeraAvsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractHimeraAvsRaw struct {
	Contract *ContractHimeraAvs // Generic contract binding to access the raw methods on
}

// ContractHimeraAvsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractHimeraAvsCallerRaw struct {
	Contract *ContractHimeraAvsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractHimeraAvsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractHimeraAvsTransactorRaw struct {
	Contract *ContractHimeraAvsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractHimeraAvs creates a new instance of ContractHimeraAvs, bound to a specific deployed contract.
func NewContractHimeraAvs(address common.Address, backend bind.ContractBackend) (*ContractHimeraAvs, error) {
	contract, err := bindContractHimeraAvs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvs{ContractHimeraAvsCaller: ContractHimeraAvsCaller{contract: contract}, ContractHimeraAvsTransactor: ContractHimeraAvsTransactor{contract: contract}, ContractHimeraAvsFilterer: ContractHimeraAvsFilterer{contract: contract}}, nil
}

// NewContractHimeraAvsCaller creates a new read-only instance of ContractHimeraAvs, bound to a specific deployed contract.
func NewContractHimeraAvsCaller(address common.Address, caller bind.ContractCaller) (*ContractHimeraAvsCaller, error) {
	contract, err := bindContractHimeraAvs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsCaller{contract: contract}, nil
}

// NewContractHimeraAvsTransactor creates a new write-only instance of ContractHimeraAvs, bound to a specific deployed contract.
func NewContractHimeraAvsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractHimeraAvsTransactor, error) {
	contract, err := bindContractHimeraAvs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsTransactor{contract: contract}, nil
}

// NewContractHimeraAvsFilterer creates a new log filterer instance of ContractHimeraAvs, bound to a specific deployed contract.
func NewContractHimeraAvsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractHimeraAvsFilterer, error) {
	contract, err := bindContractHimeraAvs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsFilterer{contract: contract}, nil
}

// bindContractHimeraAvs binds a generic wrapper to an already deployed contract.
func bindContractHimeraAvs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractHimeraAvsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractHimeraAvs *ContractHimeraAvsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractHimeraAvs.Contract.ContractHimeraAvsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractHimeraAvs *ContractHimeraAvsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.ContractHimeraAvsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractHimeraAvs *ContractHimeraAvsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.ContractHimeraAvsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractHimeraAvs *ContractHimeraAvsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractHimeraAvs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractHimeraAvs *ContractHimeraAvsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractHimeraAvs *ContractHimeraAvsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.contract.Transact(opts, method, params...)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractHimeraAvs *ContractHimeraAvsCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ContractHimeraAvs.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractHimeraAvs *ContractHimeraAvsSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractHimeraAvs.Contract.UPGRADEINTERFACEVERSION(&_ContractHimeraAvs.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_ContractHimeraAvs *ContractHimeraAvsCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _ContractHimeraAvs.Contract.UPGRADEINTERFACEVERSION(&_ContractHimeraAvs.CallOpts)
}

// ChallengerAddress is a free data retrieval call binding the contract method 0x82f8d845.
//
// Solidity: function challengerAddress() view returns(address)
func (_ContractHimeraAvs *ContractHimeraAvsCaller) ChallengerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractHimeraAvs.contract.Call(opts, &out, "challengerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChallengerAddress is a free data retrieval call binding the contract method 0x82f8d845.
//
// Solidity: function challengerAddress() view returns(address)
func (_ContractHimeraAvs *ContractHimeraAvsSession) ChallengerAddress() (common.Address, error) {
	return _ContractHimeraAvs.Contract.ChallengerAddress(&_ContractHimeraAvs.CallOpts)
}

// ChallengerAddress is a free data retrieval call binding the contract method 0x82f8d845.
//
// Solidity: function challengerAddress() view returns(address)
func (_ContractHimeraAvs *ContractHimeraAvsCallerSession) ChallengerAddress() (common.Address, error) {
	return _ContractHimeraAvs.Contract.ChallengerAddress(&_ContractHimeraAvs.CallOpts)
}

// GetOptInOperators is a free data retrieval call binding the contract method 0xa528113f.
//
// Solidity: function getOptInOperators() view returns(address[])
func (_ContractHimeraAvs *ContractHimeraAvsCaller) GetOptInOperators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ContractHimeraAvs.contract.Call(opts, &out, "getOptInOperators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOptInOperators is a free data retrieval call binding the contract method 0xa528113f.
//
// Solidity: function getOptInOperators() view returns(address[])
func (_ContractHimeraAvs *ContractHimeraAvsSession) GetOptInOperators() ([]common.Address, error) {
	return _ContractHimeraAvs.Contract.GetOptInOperators(&_ContractHimeraAvs.CallOpts)
}

// GetOptInOperators is a free data retrieval call binding the contract method 0xa528113f.
//
// Solidity: function getOptInOperators() view returns(address[])
func (_ContractHimeraAvs *ContractHimeraAvsCallerSession) GetOptInOperators() ([]common.Address, error) {
	return _ContractHimeraAvs.Contract.GetOptInOperators(&_ContractHimeraAvs.CallOpts)
}

// GetTaskDefinition is a free data retrieval call binding the contract method 0xbd0d5737.
//
// Solidity: function getTaskDefinition(uint8 id) view returns((uint8,string,uint8,string))
func (_ContractHimeraAvs *ContractHimeraAvsCaller) GetTaskDefinition(opts *bind.CallOpts, id uint8) (HimeraTaskLibraryTaskDefinition, error) {
	var out []interface{}
	err := _ContractHimeraAvs.contract.Call(opts, &out, "getTaskDefinition", id)

	if err != nil {
		return *new(HimeraTaskLibraryTaskDefinition), err
	}

	out0 := *abi.ConvertType(out[0], new(HimeraTaskLibraryTaskDefinition)).(*HimeraTaskLibraryTaskDefinition)

	return out0, err

}

// GetTaskDefinition is a free data retrieval call binding the contract method 0xbd0d5737.
//
// Solidity: function getTaskDefinition(uint8 id) view returns((uint8,string,uint8,string))
func (_ContractHimeraAvs *ContractHimeraAvsSession) GetTaskDefinition(id uint8) (HimeraTaskLibraryTaskDefinition, error) {
	return _ContractHimeraAvs.Contract.GetTaskDefinition(&_ContractHimeraAvs.CallOpts, id)
}

// GetTaskDefinition is a free data retrieval call binding the contract method 0xbd0d5737.
//
// Solidity: function getTaskDefinition(uint8 id) view returns((uint8,string,uint8,string))
func (_ContractHimeraAvs *ContractHimeraAvsCallerSession) GetTaskDefinition(id uint8) (HimeraTaskLibraryTaskDefinition, error) {
	return _ContractHimeraAvs.Contract.GetTaskDefinition(&_ContractHimeraAvs.CallOpts, id)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0xe73e8a71.
//
// Solidity: function getTaskInfo(uint64 taskID) view returns((address,string,bytes,uint64,uint64,uint64,uint64,uint8,uint64,string,address[],address[],address[],address[],string,(address,uint256)[],bool,address[],address[]))
func (_ContractHimeraAvs *ContractHimeraAvsCaller) GetTaskInfo(opts *bind.CallOpts, taskID uint64) (TaskInfo, error) {
	var out []interface{}
	err := _ContractHimeraAvs.contract.Call(opts, &out, "getTaskInfo", taskID)

	if err != nil {
		return *new(TaskInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(TaskInfo)).(*TaskInfo)

	return out0, err

}

// GetTaskInfo is a free data retrieval call binding the contract method 0xe73e8a71.
//
// Solidity: function getTaskInfo(uint64 taskID) view returns((address,string,bytes,uint64,uint64,uint64,uint64,uint8,uint64,string,address[],address[],address[],address[],string,(address,uint256)[],bool,address[],address[]))
func (_ContractHimeraAvs *ContractHimeraAvsSession) GetTaskInfo(taskID uint64) (TaskInfo, error) {
	return _ContractHimeraAvs.Contract.GetTaskInfo(&_ContractHimeraAvs.CallOpts, taskID)
}

// GetTaskInfo is a free data retrieval call binding the contract method 0xe73e8a71.
//
// Solidity: function getTaskInfo(uint64 taskID) view returns((address,string,bytes,uint64,uint64,uint64,uint64,uint8,uint64,string,address[],address[],address[],address[],string,(address,uint256)[],bool,address[],address[]))
func (_ContractHimeraAvs *ContractHimeraAvsCallerSession) GetTaskInfo(taskID uint64) (TaskInfo, error) {
	return _ContractHimeraAvs.Contract.GetTaskInfo(&_ContractHimeraAvs.CallOpts, taskID)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractHimeraAvs *ContractHimeraAvsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractHimeraAvs.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractHimeraAvs *ContractHimeraAvsSession) Owner() (common.Address, error) {
	return _ContractHimeraAvs.Contract.Owner(&_ContractHimeraAvs.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ContractHimeraAvs *ContractHimeraAvsCallerSession) Owner() (common.Address, error) {
	return _ContractHimeraAvs.Contract.Owner(&_ContractHimeraAvs.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractHimeraAvs *ContractHimeraAvsCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ContractHimeraAvs.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractHimeraAvs *ContractHimeraAvsSession) ProxiableUUID() ([32]byte, error) {
	return _ContractHimeraAvs.Contract.ProxiableUUID(&_ContractHimeraAvs.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ContractHimeraAvs *ContractHimeraAvsCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ContractHimeraAvs.Contract.ProxiableUUID(&_ContractHimeraAvs.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_ContractHimeraAvs *ContractHimeraAvsCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractHimeraAvs.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_ContractHimeraAvs *ContractHimeraAvsSession) RewardManager() (common.Address, error) {
	return _ContractHimeraAvs.Contract.RewardManager(&_ContractHimeraAvs.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_ContractHimeraAvs *ContractHimeraAvsCallerSession) RewardManager() (common.Address, error) {
	return _ContractHimeraAvs.Contract.RewardManager(&_ContractHimeraAvs.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractHimeraAvs *ContractHimeraAvsCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ContractHimeraAvs.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractHimeraAvs *ContractHimeraAvsSession) Slasher() (common.Address, error) {
	return _ContractHimeraAvs.Contract.Slasher(&_ContractHimeraAvs.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_ContractHimeraAvs *ContractHimeraAvsCallerSession) Slasher() (common.Address, error) {
	return _ContractHimeraAvs.Contract.Slasher(&_ContractHimeraAvs.CallOpts)
}

// Challenge is a paid mutator transaction binding the contract method 0xe20529ce.
//
// Solidity: function challenge(uint64 taskID, uint8 actualThreshold, bool isExpected, address[] eligibleRewardOperators, address[] eligibleSlashOperators) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) Challenge(opts *bind.TransactOpts, taskID uint64, actualThreshold uint8, isExpected bool, eligibleRewardOperators []common.Address, eligibleSlashOperators []common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "challenge", taskID, actualThreshold, isExpected, eligibleRewardOperators, eligibleSlashOperators)
}

// Challenge is a paid mutator transaction binding the contract method 0xe20529ce.
//
// Solidity: function challenge(uint64 taskID, uint8 actualThreshold, bool isExpected, address[] eligibleRewardOperators, address[] eligibleSlashOperators) returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) Challenge(taskID uint64, actualThreshold uint8, isExpected bool, eligibleRewardOperators []common.Address, eligibleSlashOperators []common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.Challenge(&_ContractHimeraAvs.TransactOpts, taskID, actualThreshold, isExpected, eligibleRewardOperators, eligibleSlashOperators)
}

// Challenge is a paid mutator transaction binding the contract method 0xe20529ce.
//
// Solidity: function challenge(uint64 taskID, uint8 actualThreshold, bool isExpected, address[] eligibleRewardOperators, address[] eligibleSlashOperators) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) Challenge(taskID uint64, actualThreshold uint8, isExpected bool, eligibleRewardOperators []common.Address, eligibleSlashOperators []common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.Challenge(&_ContractHimeraAvs.TransactOpts, taskID, actualThreshold, isExpected, eligibleRewardOperators, eligibleSlashOperators)
}

// CreateHimeraTask is a paid mutator transaction binding the contract method 0x29c0c3e1.
//
// Solidity: function createHimeraTask(uint8 himeraTaskDefId, bytes taskInput) returns(uint64 imuaTaskId)
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) CreateHimeraTask(opts *bind.TransactOpts, himeraTaskDefId uint8, taskInput []byte) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "createHimeraTask", himeraTaskDefId, taskInput)
}

// CreateHimeraTask is a paid mutator transaction binding the contract method 0x29c0c3e1.
//
// Solidity: function createHimeraTask(uint8 himeraTaskDefId, bytes taskInput) returns(uint64 imuaTaskId)
func (_ContractHimeraAvs *ContractHimeraAvsSession) CreateHimeraTask(himeraTaskDefId uint8, taskInput []byte) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.CreateHimeraTask(&_ContractHimeraAvs.TransactOpts, himeraTaskDefId, taskInput)
}

// CreateHimeraTask is a paid mutator transaction binding the contract method 0x29c0c3e1.
//
// Solidity: function createHimeraTask(uint8 himeraTaskDefId, bytes taskInput) returns(uint64 imuaTaskId)
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) CreateHimeraTask(himeraTaskDefId uint8, taskInput []byte) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.CreateHimeraTask(&_ContractHimeraAvs.TransactOpts, himeraTaskDefId, taskInput)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) DeregisterOperatorFromAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "deregisterOperatorFromAVS")
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) DeregisterOperatorFromAVS() (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.DeregisterOperatorFromAVS(&_ContractHimeraAvs.TransactOpts)
}

// DeregisterOperatorFromAVS is a paid mutator transaction binding the contract method 0xde16bf46.
//
// Solidity: function deregisterOperatorFromAVS() returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) DeregisterOperatorFromAVS() (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.DeregisterOperatorFromAVS(&_ContractHimeraAvs.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) Initialize(opts *bind.TransactOpts, initialOwner common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "initialize", initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.Initialize(&_ContractHimeraAvs.TransactOpts, initialOwner)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address initialOwner) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) Initialize(initialOwner common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.Initialize(&_ContractHimeraAvs.TransactOpts, initialOwner)
}

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0x6d392e69.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, uint8 phase) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) OperatorSubmitTask(opts *bind.TransactOpts, taskID uint64, taskResponse []byte, blsSignature []byte, phase uint8) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "operatorSubmitTask", taskID, taskResponse, blsSignature, phase)
}

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0x6d392e69.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, uint8 phase) returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) OperatorSubmitTask(taskID uint64, taskResponse []byte, blsSignature []byte, phase uint8) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.OperatorSubmitTask(&_ContractHimeraAvs.TransactOpts, taskID, taskResponse, blsSignature, phase)
}

// OperatorSubmitTask is a paid mutator transaction binding the contract method 0x6d392e69.
//
// Solidity: function operatorSubmitTask(uint64 taskID, bytes taskResponse, bytes blsSignature, uint8 phase) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) OperatorSubmitTask(taskID uint64, taskResponse []byte, blsSignature []byte, phase uint8) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.OperatorSubmitTask(&_ContractHimeraAvs.TransactOpts, taskID, taskResponse, blsSignature, phase)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0x0b70f322.
//
// Solidity: function registerAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) RegisterAVS(opts *bind.TransactOpts, params AVSParams) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "registerAVS", params)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0x0b70f322.
//
// Solidity: function registerAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) RegisterAVS(params AVSParams) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.RegisterAVS(&_ContractHimeraAvs.TransactOpts, params)
}

// RegisterAVS is a paid mutator transaction binding the contract method 0x0b70f322.
//
// Solidity: function registerAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) RegisterAVS(params AVSParams) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.RegisterAVS(&_ContractHimeraAvs.TransactOpts, params)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xf75816e9.
//
// Solidity: function registerBLSPublicKey(bytes pubKey, bytes pubKeyRegSig) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) RegisterBLSPublicKey(opts *bind.TransactOpts, pubKey []byte, pubKeyRegSig []byte) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "registerBLSPublicKey", pubKey, pubKeyRegSig)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xf75816e9.
//
// Solidity: function registerBLSPublicKey(bytes pubKey, bytes pubKeyRegSig) returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) RegisterBLSPublicKey(pubKey []byte, pubKeyRegSig []byte) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.RegisterBLSPublicKey(&_ContractHimeraAvs.TransactOpts, pubKey, pubKeyRegSig)
}

// RegisterBLSPublicKey is a paid mutator transaction binding the contract method 0xf75816e9.
//
// Solidity: function registerBLSPublicKey(bytes pubKey, bytes pubKeyRegSig) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) RegisterBLSPublicKey(pubKey []byte, pubKeyRegSig []byte) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.RegisterBLSPublicKey(&_ContractHimeraAvs.TransactOpts, pubKey, pubKeyRegSig)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) RegisterOperatorToAVS(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "registerOperatorToAVS")
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) RegisterOperatorToAVS() (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.RegisterOperatorToAVS(&_ContractHimeraAvs.TransactOpts)
}

// RegisterOperatorToAVS is a paid mutator transaction binding the contract method 0xc208dd99.
//
// Solidity: function registerOperatorToAVS() returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) RegisterOperatorToAVS() (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.RegisterOperatorToAVS(&_ContractHimeraAvs.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.RenounceOwnership(&_ContractHimeraAvs.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.RenounceOwnership(&_ContractHimeraAvs.TransactOpts)
}

// SetChallenger is a paid mutator transaction binding the contract method 0xb6b46535.
//
// Solidity: function setChallenger(address _challenger) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) SetChallenger(opts *bind.TransactOpts, _challenger common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "setChallenger", _challenger)
}

// SetChallenger is a paid mutator transaction binding the contract method 0xb6b46535.
//
// Solidity: function setChallenger(address _challenger) returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) SetChallenger(_challenger common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.SetChallenger(&_ContractHimeraAvs.TransactOpts, _challenger)
}

// SetChallenger is a paid mutator transaction binding the contract method 0xb6b46535.
//
// Solidity: function setChallenger(address _challenger) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) SetChallenger(_challenger common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.SetChallenger(&_ContractHimeraAvs.TransactOpts, _challenger)
}

// SetRewardManager is a paid mutator transaction binding the contract method 0x153ee554.
//
// Solidity: function setRewardManager(address _rewardManager) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) SetRewardManager(opts *bind.TransactOpts, _rewardManager common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "setRewardManager", _rewardManager)
}

// SetRewardManager is a paid mutator transaction binding the contract method 0x153ee554.
//
// Solidity: function setRewardManager(address _rewardManager) returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) SetRewardManager(_rewardManager common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.SetRewardManager(&_ContractHimeraAvs.TransactOpts, _rewardManager)
}

// SetRewardManager is a paid mutator transaction binding the contract method 0x153ee554.
//
// Solidity: function setRewardManager(address _rewardManager) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) SetRewardManager(_rewardManager common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.SetRewardManager(&_ContractHimeraAvs.TransactOpts, _rewardManager)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address _slasher) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) SetSlasher(opts *bind.TransactOpts, _slasher common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "setSlasher", _slasher)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address _slasher) returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) SetSlasher(_slasher common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.SetSlasher(&_ContractHimeraAvs.TransactOpts, _slasher)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address _slasher) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) SetSlasher(_slasher common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.SetSlasher(&_ContractHimeraAvs.TransactOpts, _slasher)
}

// SetupTaskDefinitions is a paid mutator transaction binding the contract method 0x2b1561ef.
//
// Solidity: function setupTaskDefinitions() returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) SetupTaskDefinitions(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "setupTaskDefinitions")
}

// SetupTaskDefinitions is a paid mutator transaction binding the contract method 0x2b1561ef.
//
// Solidity: function setupTaskDefinitions() returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) SetupTaskDefinitions() (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.SetupTaskDefinitions(&_ContractHimeraAvs.TransactOpts)
}

// SetupTaskDefinitions is a paid mutator transaction binding the contract method 0x2b1561ef.
//
// Solidity: function setupTaskDefinitions() returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) SetupTaskDefinitions() (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.SetupTaskDefinitions(&_ContractHimeraAvs.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.TransferOwnership(&_ContractHimeraAvs.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.TransferOwnership(&_ContractHimeraAvs.TransactOpts, newOwner)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x3a72b900.
//
// Solidity: function updateAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) UpdateAVS(opts *bind.TransactOpts, params AVSParams) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "updateAVS", params)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x3a72b900.
//
// Solidity: function updateAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) UpdateAVS(params AVSParams) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.UpdateAVS(&_ContractHimeraAvs.TransactOpts, params)
}

// UpdateAVS is a paid mutator transaction binding the contract method 0x3a72b900.
//
// Solidity: function updateAVS((address,string,uint64,address,address,address,address[],address[],string[],uint64,uint64,string,uint64,uint64,uint64,uint64) params) returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) UpdateAVS(params AVSParams) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.UpdateAVS(&_ContractHimeraAvs.TransactOpts, params)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractHimeraAvs.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractHimeraAvs *ContractHimeraAvsSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.UpgradeToAndCall(&_ContractHimeraAvs.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ContractHimeraAvs *ContractHimeraAvsTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ContractHimeraAvs.Contract.UpgradeToAndCall(&_ContractHimeraAvs.TransactOpts, newImplementation, data)
}

// ContractHimeraAvsBLSPublicKeyRegisteredIterator is returned from FilterBLSPublicKeyRegistered and is used to iterate over the raw logs and unpacked data for BLSPublicKeyRegistered events raised by the ContractHimeraAvs contract.
type ContractHimeraAvsBLSPublicKeyRegisteredIterator struct {
	Event *ContractHimeraAvsBLSPublicKeyRegistered // Event containing the contract specifics and raw log

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
func (it *ContractHimeraAvsBLSPublicKeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractHimeraAvsBLSPublicKeyRegistered)
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
		it.Event = new(ContractHimeraAvsBLSPublicKeyRegistered)
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
func (it *ContractHimeraAvsBLSPublicKeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractHimeraAvsBLSPublicKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractHimeraAvsBLSPublicKeyRegistered represents a BLSPublicKeyRegistered event raised by the ContractHimeraAvs contract.
type ContractHimeraAvsBLSPublicKeyRegistered struct {
	Operator   common.Address
	AvsAddress common.Address
	PubKey     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBLSPublicKeyRegistered is a free log retrieval operation binding the contract event 0xd7a56e64a4cd1aeb35e575da573ffdbd29cbafdf2ef88c1772197d7c72be405f.
//
// Solidity: event BLSPublicKeyRegistered(address indexed operator, address indexed avsAddress, bytes pubKey)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) FilterBLSPublicKeyRegistered(opts *bind.FilterOpts, operator []common.Address, avsAddress []common.Address) (*ContractHimeraAvsBLSPublicKeyRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsAddressRule []interface{}
	for _, avsAddressItem := range avsAddress {
		avsAddressRule = append(avsAddressRule, avsAddressItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.FilterLogs(opts, "BLSPublicKeyRegistered", operatorRule, avsAddressRule)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsBLSPublicKeyRegisteredIterator{contract: _ContractHimeraAvs.contract, event: "BLSPublicKeyRegistered", logs: logs, sub: sub}, nil
}

// WatchBLSPublicKeyRegistered is a free log subscription operation binding the contract event 0xd7a56e64a4cd1aeb35e575da573ffdbd29cbafdf2ef88c1772197d7c72be405f.
//
// Solidity: event BLSPublicKeyRegistered(address indexed operator, address indexed avsAddress, bytes pubKey)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) WatchBLSPublicKeyRegistered(opts *bind.WatchOpts, sink chan<- *ContractHimeraAvsBLSPublicKeyRegistered, operator []common.Address, avsAddress []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var avsAddressRule []interface{}
	for _, avsAddressItem := range avsAddress {
		avsAddressRule = append(avsAddressRule, avsAddressItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.WatchLogs(opts, "BLSPublicKeyRegistered", operatorRule, avsAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractHimeraAvsBLSPublicKeyRegistered)
				if err := _ContractHimeraAvs.contract.UnpackLog(event, "BLSPublicKeyRegistered", log); err != nil {
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

// ParseBLSPublicKeyRegistered is a log parse operation binding the contract event 0xd7a56e64a4cd1aeb35e575da573ffdbd29cbafdf2ef88c1772197d7c72be405f.
//
// Solidity: event BLSPublicKeyRegistered(address indexed operator, address indexed avsAddress, bytes pubKey)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) ParseBLSPublicKeyRegistered(log types.Log) (*ContractHimeraAvsBLSPublicKeyRegistered, error) {
	event := new(ContractHimeraAvsBLSPublicKeyRegistered)
	if err := _ContractHimeraAvs.contract.UnpackLog(event, "BLSPublicKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractHimeraAvsChallengeSubmittedIterator is returned from FilterChallengeSubmitted and is used to iterate over the raw logs and unpacked data for ChallengeSubmitted events raised by the ContractHimeraAvs contract.
type ContractHimeraAvsChallengeSubmittedIterator struct {
	Event *ContractHimeraAvsChallengeSubmitted // Event containing the contract specifics and raw log

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
func (it *ContractHimeraAvsChallengeSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractHimeraAvsChallengeSubmitted)
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
		it.Event = new(ContractHimeraAvsChallengeSubmitted)
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
func (it *ContractHimeraAvsChallengeSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractHimeraAvsChallengeSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractHimeraAvsChallengeSubmitted represents a ChallengeSubmitted event raised by the ContractHimeraAvs contract.
type ContractHimeraAvsChallengeSubmitted struct {
	TaskID     uint64
	Challenger common.Address
	IsExpected bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChallengeSubmitted is a free log retrieval operation binding the contract event 0x9de5978cdd4b224702e27e43ab9f28631c0ab8ffe4b20e4cd8301d7a6da8ae72.
//
// Solidity: event ChallengeSubmitted(uint64 indexed taskID, address indexed challenger, bool isExpected)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) FilterChallengeSubmitted(opts *bind.FilterOpts, taskID []uint64, challenger []common.Address) (*ContractHimeraAvsChallengeSubmittedIterator, error) {

	var taskIDRule []interface{}
	for _, taskIDItem := range taskID {
		taskIDRule = append(taskIDRule, taskIDItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.FilterLogs(opts, "ChallengeSubmitted", taskIDRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsChallengeSubmittedIterator{contract: _ContractHimeraAvs.contract, event: "ChallengeSubmitted", logs: logs, sub: sub}, nil
}

// WatchChallengeSubmitted is a free log subscription operation binding the contract event 0x9de5978cdd4b224702e27e43ab9f28631c0ab8ffe4b20e4cd8301d7a6da8ae72.
//
// Solidity: event ChallengeSubmitted(uint64 indexed taskID, address indexed challenger, bool isExpected)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) WatchChallengeSubmitted(opts *bind.WatchOpts, sink chan<- *ContractHimeraAvsChallengeSubmitted, taskID []uint64, challenger []common.Address) (event.Subscription, error) {

	var taskIDRule []interface{}
	for _, taskIDItem := range taskID {
		taskIDRule = append(taskIDRule, taskIDItem)
	}
	var challengerRule []interface{}
	for _, challengerItem := range challenger {
		challengerRule = append(challengerRule, challengerItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.WatchLogs(opts, "ChallengeSubmitted", taskIDRule, challengerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractHimeraAvsChallengeSubmitted)
				if err := _ContractHimeraAvs.contract.UnpackLog(event, "ChallengeSubmitted", log); err != nil {
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

// ParseChallengeSubmitted is a log parse operation binding the contract event 0x9de5978cdd4b224702e27e43ab9f28631c0ab8ffe4b20e4cd8301d7a6da8ae72.
//
// Solidity: event ChallengeSubmitted(uint64 indexed taskID, address indexed challenger, bool isExpected)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) ParseChallengeSubmitted(log types.Log) (*ContractHimeraAvsChallengeSubmitted, error) {
	event := new(ContractHimeraAvsChallengeSubmitted)
	if err := _ContractHimeraAvs.contract.UnpackLog(event, "ChallengeSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractHimeraAvsHimeraTaskCreatedIterator is returned from FilterHimeraTaskCreated and is used to iterate over the raw logs and unpacked data for HimeraTaskCreated events raised by the ContractHimeraAvs contract.
type ContractHimeraAvsHimeraTaskCreatedIterator struct {
	Event *ContractHimeraAvsHimeraTaskCreated // Event containing the contract specifics and raw log

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
func (it *ContractHimeraAvsHimeraTaskCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractHimeraAvsHimeraTaskCreated)
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
		it.Event = new(ContractHimeraAvsHimeraTaskCreated)
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
func (it *ContractHimeraAvsHimeraTaskCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractHimeraAvsHimeraTaskCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractHimeraAvsHimeraTaskCreated represents a HimeraTaskCreated event raised by the ContractHimeraAvs contract.
type ContractHimeraAvsHimeraTaskCreated struct {
	ImuaTaskId      uint64
	DefinitionHash  [32]byte
	HimeraTaskDefId uint8
	TaskInput       []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterHimeraTaskCreated is a free log retrieval operation binding the contract event 0x1aed6be81feb12ccef30c1cd152e1d55d0f204797de8d6240ffc01dffc56d286.
//
// Solidity: event HimeraTaskCreated(uint64 indexed imuaTaskId, bytes32 definitionHash, uint8 himeraTaskDefId, bytes taskInput)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) FilterHimeraTaskCreated(opts *bind.FilterOpts, imuaTaskId []uint64) (*ContractHimeraAvsHimeraTaskCreatedIterator, error) {

	var imuaTaskIdRule []interface{}
	for _, imuaTaskIdItem := range imuaTaskId {
		imuaTaskIdRule = append(imuaTaskIdRule, imuaTaskIdItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.FilterLogs(opts, "HimeraTaskCreated", imuaTaskIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsHimeraTaskCreatedIterator{contract: _ContractHimeraAvs.contract, event: "HimeraTaskCreated", logs: logs, sub: sub}, nil
}

// WatchHimeraTaskCreated is a free log subscription operation binding the contract event 0x1aed6be81feb12ccef30c1cd152e1d55d0f204797de8d6240ffc01dffc56d286.
//
// Solidity: event HimeraTaskCreated(uint64 indexed imuaTaskId, bytes32 definitionHash, uint8 himeraTaskDefId, bytes taskInput)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) WatchHimeraTaskCreated(opts *bind.WatchOpts, sink chan<- *ContractHimeraAvsHimeraTaskCreated, imuaTaskId []uint64) (event.Subscription, error) {

	var imuaTaskIdRule []interface{}
	for _, imuaTaskIdItem := range imuaTaskId {
		imuaTaskIdRule = append(imuaTaskIdRule, imuaTaskIdItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.WatchLogs(opts, "HimeraTaskCreated", imuaTaskIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractHimeraAvsHimeraTaskCreated)
				if err := _ContractHimeraAvs.contract.UnpackLog(event, "HimeraTaskCreated", log); err != nil {
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

// ParseHimeraTaskCreated is a log parse operation binding the contract event 0x1aed6be81feb12ccef30c1cd152e1d55d0f204797de8d6240ffc01dffc56d286.
//
// Solidity: event HimeraTaskCreated(uint64 indexed imuaTaskId, bytes32 definitionHash, uint8 himeraTaskDefId, bytes taskInput)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) ParseHimeraTaskCreated(log types.Log) (*ContractHimeraAvsHimeraTaskCreated, error) {
	event := new(ContractHimeraAvsHimeraTaskCreated)
	if err := _ContractHimeraAvs.contract.UnpackLog(event, "HimeraTaskCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractHimeraAvsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContractHimeraAvs contract.
type ContractHimeraAvsInitializedIterator struct {
	Event *ContractHimeraAvsInitialized // Event containing the contract specifics and raw log

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
func (it *ContractHimeraAvsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractHimeraAvsInitialized)
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
		it.Event = new(ContractHimeraAvsInitialized)
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
func (it *ContractHimeraAvsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractHimeraAvsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractHimeraAvsInitialized represents a Initialized event raised by the ContractHimeraAvs contract.
type ContractHimeraAvsInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContractHimeraAvsInitializedIterator, error) {

	logs, sub, err := _ContractHimeraAvs.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsInitializedIterator{contract: _ContractHimeraAvs.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContractHimeraAvsInitialized) (event.Subscription, error) {

	logs, sub, err := _ContractHimeraAvs.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractHimeraAvsInitialized)
				if err := _ContractHimeraAvs.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) ParseInitialized(log types.Log) (*ContractHimeraAvsInitialized, error) {
	event := new(ContractHimeraAvsInitialized)
	if err := _ContractHimeraAvs.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractHimeraAvsOperatorOptedInIterator is returned from FilterOperatorOptedIn and is used to iterate over the raw logs and unpacked data for OperatorOptedIn events raised by the ContractHimeraAvs contract.
type ContractHimeraAvsOperatorOptedInIterator struct {
	Event *ContractHimeraAvsOperatorOptedIn // Event containing the contract specifics and raw log

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
func (it *ContractHimeraAvsOperatorOptedInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractHimeraAvsOperatorOptedIn)
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
		it.Event = new(ContractHimeraAvsOperatorOptedIn)
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
func (it *ContractHimeraAvsOperatorOptedInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractHimeraAvsOperatorOptedInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractHimeraAvsOperatorOptedIn represents a OperatorOptedIn event raised by the ContractHimeraAvs contract.
type ContractHimeraAvsOperatorOptedIn struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorOptedIn is a free log retrieval operation binding the contract event 0x3eb9749f7e08a89f04537594ff4cf7885d8f0b56138c8cfa999b44771fb1a148.
//
// Solidity: event OperatorOptedIn(address indexed operator)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) FilterOperatorOptedIn(opts *bind.FilterOpts, operator []common.Address) (*ContractHimeraAvsOperatorOptedInIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.FilterLogs(opts, "OperatorOptedIn", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsOperatorOptedInIterator{contract: _ContractHimeraAvs.contract, event: "OperatorOptedIn", logs: logs, sub: sub}, nil
}

// WatchOperatorOptedIn is a free log subscription operation binding the contract event 0x3eb9749f7e08a89f04537594ff4cf7885d8f0b56138c8cfa999b44771fb1a148.
//
// Solidity: event OperatorOptedIn(address indexed operator)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) WatchOperatorOptedIn(opts *bind.WatchOpts, sink chan<- *ContractHimeraAvsOperatorOptedIn, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.WatchLogs(opts, "OperatorOptedIn", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractHimeraAvsOperatorOptedIn)
				if err := _ContractHimeraAvs.contract.UnpackLog(event, "OperatorOptedIn", log); err != nil {
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

// ParseOperatorOptedIn is a log parse operation binding the contract event 0x3eb9749f7e08a89f04537594ff4cf7885d8f0b56138c8cfa999b44771fb1a148.
//
// Solidity: event OperatorOptedIn(address indexed operator)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) ParseOperatorOptedIn(log types.Log) (*ContractHimeraAvsOperatorOptedIn, error) {
	event := new(ContractHimeraAvsOperatorOptedIn)
	if err := _ContractHimeraAvs.contract.UnpackLog(event, "OperatorOptedIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractHimeraAvsOperatorOptedOutIterator is returned from FilterOperatorOptedOut and is used to iterate over the raw logs and unpacked data for OperatorOptedOut events raised by the ContractHimeraAvs contract.
type ContractHimeraAvsOperatorOptedOutIterator struct {
	Event *ContractHimeraAvsOperatorOptedOut // Event containing the contract specifics and raw log

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
func (it *ContractHimeraAvsOperatorOptedOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractHimeraAvsOperatorOptedOut)
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
		it.Event = new(ContractHimeraAvsOperatorOptedOut)
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
func (it *ContractHimeraAvsOperatorOptedOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractHimeraAvsOperatorOptedOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractHimeraAvsOperatorOptedOut represents a OperatorOptedOut event raised by the ContractHimeraAvs contract.
type ContractHimeraAvsOperatorOptedOut struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorOptedOut is a free log retrieval operation binding the contract event 0x70245217442e9b16c38a95d39c7516c388c244266dd7c91b7bd8dbe285adf1a2.
//
// Solidity: event OperatorOptedOut(address indexed operator)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) FilterOperatorOptedOut(opts *bind.FilterOpts, operator []common.Address) (*ContractHimeraAvsOperatorOptedOutIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.FilterLogs(opts, "OperatorOptedOut", operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsOperatorOptedOutIterator{contract: _ContractHimeraAvs.contract, event: "OperatorOptedOut", logs: logs, sub: sub}, nil
}

// WatchOperatorOptedOut is a free log subscription operation binding the contract event 0x70245217442e9b16c38a95d39c7516c388c244266dd7c91b7bd8dbe285adf1a2.
//
// Solidity: event OperatorOptedOut(address indexed operator)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) WatchOperatorOptedOut(opts *bind.WatchOpts, sink chan<- *ContractHimeraAvsOperatorOptedOut, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.WatchLogs(opts, "OperatorOptedOut", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractHimeraAvsOperatorOptedOut)
				if err := _ContractHimeraAvs.contract.UnpackLog(event, "OperatorOptedOut", log); err != nil {
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

// ParseOperatorOptedOut is a log parse operation binding the contract event 0x70245217442e9b16c38a95d39c7516c388c244266dd7c91b7bd8dbe285adf1a2.
//
// Solidity: event OperatorOptedOut(address indexed operator)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) ParseOperatorOptedOut(log types.Log) (*ContractHimeraAvsOperatorOptedOut, error) {
	event := new(ContractHimeraAvsOperatorOptedOut)
	if err := _ContractHimeraAvs.contract.UnpackLog(event, "OperatorOptedOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractHimeraAvsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ContractHimeraAvs contract.
type ContractHimeraAvsOwnershipTransferredIterator struct {
	Event *ContractHimeraAvsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ContractHimeraAvsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractHimeraAvsOwnershipTransferred)
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
		it.Event = new(ContractHimeraAvsOwnershipTransferred)
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
func (it *ContractHimeraAvsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractHimeraAvsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractHimeraAvsOwnershipTransferred represents a OwnershipTransferred event raised by the ContractHimeraAvs contract.
type ContractHimeraAvsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ContractHimeraAvsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsOwnershipTransferredIterator{contract: _ContractHimeraAvs.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ContractHimeraAvsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractHimeraAvsOwnershipTransferred)
				if err := _ContractHimeraAvs.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) ParseOwnershipTransferred(log types.Log) (*ContractHimeraAvsOwnershipTransferred, error) {
	event := new(ContractHimeraAvsOwnershipTransferred)
	if err := _ContractHimeraAvs.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractHimeraAvsRewardManagerUpdatedIterator is returned from FilterRewardManagerUpdated and is used to iterate over the raw logs and unpacked data for RewardManagerUpdated events raised by the ContractHimeraAvs contract.
type ContractHimeraAvsRewardManagerUpdatedIterator struct {
	Event *ContractHimeraAvsRewardManagerUpdated // Event containing the contract specifics and raw log

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
func (it *ContractHimeraAvsRewardManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractHimeraAvsRewardManagerUpdated)
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
		it.Event = new(ContractHimeraAvsRewardManagerUpdated)
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
func (it *ContractHimeraAvsRewardManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractHimeraAvsRewardManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractHimeraAvsRewardManagerUpdated represents a RewardManagerUpdated event raised by the ContractHimeraAvs contract.
type ContractHimeraAvsRewardManagerUpdated struct {
	NewRewardManager common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRewardManagerUpdated is a free log retrieval operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) FilterRewardManagerUpdated(opts *bind.FilterOpts, newRewardManager []common.Address) (*ContractHimeraAvsRewardManagerUpdatedIterator, error) {

	var newRewardManagerRule []interface{}
	for _, newRewardManagerItem := range newRewardManager {
		newRewardManagerRule = append(newRewardManagerRule, newRewardManagerItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.FilterLogs(opts, "RewardManagerUpdated", newRewardManagerRule)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsRewardManagerUpdatedIterator{contract: _ContractHimeraAvs.contract, event: "RewardManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchRewardManagerUpdated is a free log subscription operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) WatchRewardManagerUpdated(opts *bind.WatchOpts, sink chan<- *ContractHimeraAvsRewardManagerUpdated, newRewardManager []common.Address) (event.Subscription, error) {

	var newRewardManagerRule []interface{}
	for _, newRewardManagerItem := range newRewardManager {
		newRewardManagerRule = append(newRewardManagerRule, newRewardManagerItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.WatchLogs(opts, "RewardManagerUpdated", newRewardManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractHimeraAvsRewardManagerUpdated)
				if err := _ContractHimeraAvs.contract.UnpackLog(event, "RewardManagerUpdated", log); err != nil {
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

// ParseRewardManagerUpdated is a log parse operation binding the contract event 0x3d94d9e8342a65edb95eef4f65059294d45e5192603632d8dddb2344e7078053.
//
// Solidity: event RewardManagerUpdated(address indexed newRewardManager)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) ParseRewardManagerUpdated(log types.Log) (*ContractHimeraAvsRewardManagerUpdated, error) {
	event := new(ContractHimeraAvsRewardManagerUpdated)
	if err := _ContractHimeraAvs.contract.UnpackLog(event, "RewardManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractHimeraAvsSlasherUpdatedIterator is returned from FilterSlasherUpdated and is used to iterate over the raw logs and unpacked data for SlasherUpdated events raised by the ContractHimeraAvs contract.
type ContractHimeraAvsSlasherUpdatedIterator struct {
	Event *ContractHimeraAvsSlasherUpdated // Event containing the contract specifics and raw log

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
func (it *ContractHimeraAvsSlasherUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractHimeraAvsSlasherUpdated)
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
		it.Event = new(ContractHimeraAvsSlasherUpdated)
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
func (it *ContractHimeraAvsSlasherUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractHimeraAvsSlasherUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractHimeraAvsSlasherUpdated represents a SlasherUpdated event raised by the ContractHimeraAvs contract.
type ContractHimeraAvsSlasherUpdated struct {
	NewSlasher common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSlasherUpdated is a free log retrieval operation binding the contract event 0x0adf62081dae4c128a0af3a933748637b1d874a033588518f810559e6bdb23ff.
//
// Solidity: event SlasherUpdated(address indexed newSlasher)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) FilterSlasherUpdated(opts *bind.FilterOpts, newSlasher []common.Address) (*ContractHimeraAvsSlasherUpdatedIterator, error) {

	var newSlasherRule []interface{}
	for _, newSlasherItem := range newSlasher {
		newSlasherRule = append(newSlasherRule, newSlasherItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.FilterLogs(opts, "SlasherUpdated", newSlasherRule)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsSlasherUpdatedIterator{contract: _ContractHimeraAvs.contract, event: "SlasherUpdated", logs: logs, sub: sub}, nil
}

// WatchSlasherUpdated is a free log subscription operation binding the contract event 0x0adf62081dae4c128a0af3a933748637b1d874a033588518f810559e6bdb23ff.
//
// Solidity: event SlasherUpdated(address indexed newSlasher)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) WatchSlasherUpdated(opts *bind.WatchOpts, sink chan<- *ContractHimeraAvsSlasherUpdated, newSlasher []common.Address) (event.Subscription, error) {

	var newSlasherRule []interface{}
	for _, newSlasherItem := range newSlasher {
		newSlasherRule = append(newSlasherRule, newSlasherItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.WatchLogs(opts, "SlasherUpdated", newSlasherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractHimeraAvsSlasherUpdated)
				if err := _ContractHimeraAvs.contract.UnpackLog(event, "SlasherUpdated", log); err != nil {
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

// ParseSlasherUpdated is a log parse operation binding the contract event 0x0adf62081dae4c128a0af3a933748637b1d874a033588518f810559e6bdb23ff.
//
// Solidity: event SlasherUpdated(address indexed newSlasher)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) ParseSlasherUpdated(log types.Log) (*ContractHimeraAvsSlasherUpdated, error) {
	event := new(ContractHimeraAvsSlasherUpdated)
	if err := _ContractHimeraAvs.contract.UnpackLog(event, "SlasherUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractHimeraAvsTaskDefinitionCreatedIterator is returned from FilterTaskDefinitionCreated and is used to iterate over the raw logs and unpacked data for TaskDefinitionCreated events raised by the ContractHimeraAvs contract.
type ContractHimeraAvsTaskDefinitionCreatedIterator struct {
	Event *ContractHimeraAvsTaskDefinitionCreated // Event containing the contract specifics and raw log

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
func (it *ContractHimeraAvsTaskDefinitionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractHimeraAvsTaskDefinitionCreated)
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
		it.Event = new(ContractHimeraAvsTaskDefinitionCreated)
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
func (it *ContractHimeraAvsTaskDefinitionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractHimeraAvsTaskDefinitionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractHimeraAvsTaskDefinitionCreated represents a TaskDefinitionCreated event raised by the ContractHimeraAvs contract.
type ContractHimeraAvsTaskDefinitionCreated struct {
	TaskDefinitionId uint8
	Name             string
	TaskType         uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTaskDefinitionCreated is a free log retrieval operation binding the contract event 0x1b6bb8ffcd76909ae520aa61bd004fe8cf42b17b448db5d7c04e940b6834ccb5.
//
// Solidity: event TaskDefinitionCreated(uint8 indexed taskDefinitionId, string name, uint8 taskType)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) FilterTaskDefinitionCreated(opts *bind.FilterOpts, taskDefinitionId []uint8) (*ContractHimeraAvsTaskDefinitionCreatedIterator, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.FilterLogs(opts, "TaskDefinitionCreated", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsTaskDefinitionCreatedIterator{contract: _ContractHimeraAvs.contract, event: "TaskDefinitionCreated", logs: logs, sub: sub}, nil
}

// WatchTaskDefinitionCreated is a free log subscription operation binding the contract event 0x1b6bb8ffcd76909ae520aa61bd004fe8cf42b17b448db5d7c04e940b6834ccb5.
//
// Solidity: event TaskDefinitionCreated(uint8 indexed taskDefinitionId, string name, uint8 taskType)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) WatchTaskDefinitionCreated(opts *bind.WatchOpts, sink chan<- *ContractHimeraAvsTaskDefinitionCreated, taskDefinitionId []uint8) (event.Subscription, error) {

	var taskDefinitionIdRule []interface{}
	for _, taskDefinitionIdItem := range taskDefinitionId {
		taskDefinitionIdRule = append(taskDefinitionIdRule, taskDefinitionIdItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.WatchLogs(opts, "TaskDefinitionCreated", taskDefinitionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractHimeraAvsTaskDefinitionCreated)
				if err := _ContractHimeraAvs.contract.UnpackLog(event, "TaskDefinitionCreated", log); err != nil {
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

// ParseTaskDefinitionCreated is a log parse operation binding the contract event 0x1b6bb8ffcd76909ae520aa61bd004fe8cf42b17b448db5d7c04e940b6834ccb5.
//
// Solidity: event TaskDefinitionCreated(uint8 indexed taskDefinitionId, string name, uint8 taskType)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) ParseTaskDefinitionCreated(log types.Log) (*ContractHimeraAvsTaskDefinitionCreated, error) {
	event := new(ContractHimeraAvsTaskDefinitionCreated)
	if err := _ContractHimeraAvs.contract.UnpackLog(event, "TaskDefinitionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractHimeraAvsTaskSubmittedIterator is returned from FilterTaskSubmitted and is used to iterate over the raw logs and unpacked data for TaskSubmitted events raised by the ContractHimeraAvs contract.
type ContractHimeraAvsTaskSubmittedIterator struct {
	Event *ContractHimeraAvsTaskSubmitted // Event containing the contract specifics and raw log

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
func (it *ContractHimeraAvsTaskSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractHimeraAvsTaskSubmitted)
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
		it.Event = new(ContractHimeraAvsTaskSubmitted)
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
func (it *ContractHimeraAvsTaskSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractHimeraAvsTaskSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractHimeraAvsTaskSubmitted represents a TaskSubmitted event raised by the ContractHimeraAvs contract.
type ContractHimeraAvsTaskSubmitted struct {
	TaskID   uint64
	Operator common.Address
	Phase    uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTaskSubmitted is a free log retrieval operation binding the contract event 0x852ef925dfb4a67eaf327f40c6c555596a5c1b2099e05cad9362356ec8916ae3.
//
// Solidity: event TaskSubmitted(uint64 indexed taskID, address indexed operator, uint8 phase)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) FilterTaskSubmitted(opts *bind.FilterOpts, taskID []uint64, operator []common.Address) (*ContractHimeraAvsTaskSubmittedIterator, error) {

	var taskIDRule []interface{}
	for _, taskIDItem := range taskID {
		taskIDRule = append(taskIDRule, taskIDItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.FilterLogs(opts, "TaskSubmitted", taskIDRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsTaskSubmittedIterator{contract: _ContractHimeraAvs.contract, event: "TaskSubmitted", logs: logs, sub: sub}, nil
}

// WatchTaskSubmitted is a free log subscription operation binding the contract event 0x852ef925dfb4a67eaf327f40c6c555596a5c1b2099e05cad9362356ec8916ae3.
//
// Solidity: event TaskSubmitted(uint64 indexed taskID, address indexed operator, uint8 phase)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) WatchTaskSubmitted(opts *bind.WatchOpts, sink chan<- *ContractHimeraAvsTaskSubmitted, taskID []uint64, operator []common.Address) (event.Subscription, error) {

	var taskIDRule []interface{}
	for _, taskIDItem := range taskID {
		taskIDRule = append(taskIDRule, taskIDItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.WatchLogs(opts, "TaskSubmitted", taskIDRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractHimeraAvsTaskSubmitted)
				if err := _ContractHimeraAvs.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
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

// ParseTaskSubmitted is a log parse operation binding the contract event 0x852ef925dfb4a67eaf327f40c6c555596a5c1b2099e05cad9362356ec8916ae3.
//
// Solidity: event TaskSubmitted(uint64 indexed taskID, address indexed operator, uint8 phase)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) ParseTaskSubmitted(log types.Log) (*ContractHimeraAvsTaskSubmitted, error) {
	event := new(ContractHimeraAvsTaskSubmitted)
	if err := _ContractHimeraAvs.contract.UnpackLog(event, "TaskSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractHimeraAvsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ContractHimeraAvs contract.
type ContractHimeraAvsUpgradedIterator struct {
	Event *ContractHimeraAvsUpgraded // Event containing the contract specifics and raw log

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
func (it *ContractHimeraAvsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractHimeraAvsUpgraded)
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
		it.Event = new(ContractHimeraAvsUpgraded)
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
func (it *ContractHimeraAvsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractHimeraAvsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractHimeraAvsUpgraded represents a Upgraded event raised by the ContractHimeraAvs contract.
type ContractHimeraAvsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ContractHimeraAvsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ContractHimeraAvsUpgradedIterator{contract: _ContractHimeraAvs.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ContractHimeraAvsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ContractHimeraAvs.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractHimeraAvsUpgraded)
				if err := _ContractHimeraAvs.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ContractHimeraAvs *ContractHimeraAvsFilterer) ParseUpgraded(log types.Log) (*ContractHimeraAvsUpgraded, error) {
	event := new(ContractHimeraAvsUpgraded)
	if err := _ContractHimeraAvs.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
