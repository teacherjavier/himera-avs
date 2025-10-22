// core/chainio/avs_reader.go
package chainio

import (
	sdkmath "cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/imua-xyz/imua-avs-sdk/logging"

	himera_avs "github.com/imua-xyz/imua-avs/contracts/bindings/avs"
	"github.com/imua-xyz/imua-avs/core/chainio/eth"
)

// AvsReader defines the interface for reading data from the AVS contract.
type AvsReader interface {
	// GetOptInOperators returns the list of operators who have opted into the AVS.
	GetOptInOperators(opts *bind.CallOpts, avsAddress string) ([]gethcommon.Address, error)
	// GetRegisteredPubkey returns the registered BLS public key of an operator.
	GetRegisteredPubkey(opts *bind.CallOpts, avsAddress string, operator gethcommon.Address) ([]byte, error)
	// GtAVSUSDValue returns the total USD value of the AVS.
	GtAVSUSDValue(opts *bind.CallOpts, avsAddress string) (sdkmath.LegacyDec, error)
	// GetOperatorOptedUSDValue returns the USD value of an operator's stake in the AVS.
	GetOperatorOptedUSDValue(opts *bind.CallOpts, avsAddress string, operatorAddr gethcommon.Address) (sdkmath.LegacyDec, error)
	// GetAVSEpochIdentifier returns the epoch identifier for the AVS.
	GetAVSEpochIdentifier(opts *bind.CallOpts, avsAddress string) (string, error)
	// GetTaskInfo returns information about a specific task.
	GetTaskInfo(opts *bind.CallOpts, taskAddress string, taskID uint64) (himera_avs.TaskInfo, error)
	// IsOperator returns true if the given address is a registered operator.
	IsOperator(opts *bind.CallOpts, operator gethcommon.Address) (bool, error)
	// GetCurrentEpoch returns the current epoch number.
	GetCurrentEpoch(opts *bind.CallOpts, epochIdentifier string) (int64, error)
	// GetChallengeInfo returns information about a challenge for a specific task.
	GetChallengeInfo(opts *bind.CallOpts, taskAddress string, taskID uint64) (gethcommon.Address, error)
	// GetOperatorTaskResponse returns the task response of a specific operator for a given task.
	GetOperatorTaskResponse(opts *bind.CallOpts, taskAddress string, operatorAddress gethcommon.Address, taskID uint64) (himera_avs.TaskResultInfo, error)
	// GetOperatorTaskResponseList returns a list of all operator responses for a given task.
	GetOperatorTaskResponseList(opts *bind.CallOpts, taskAddress string, taskID uint64) ([]himera_avs.OperatorResInfo, error)
}

// ChainReader implements the AvsReader interface for reading data from the blockchain.
type ChainReader struct {
	logger     logging.Logger
	avsManager *himera_avs.ContractHimeraAvs
	ethClient  eth.EthClient
}

// ensure ChainReader implements AvsReader
var _ AvsReader = (*ChainReader)(nil)

// NewChainReader creates a new ChainReader instance.
func NewChainReader(
	avsManager *himera_avs.ContractHimeraAvs,
	logger logging.Logger,
	ethClient eth.EthClient,
) *ChainReader {
	return &ChainReader{
		avsManager: avsManager,
		logger:     logger,
		ethClient:  ethClient,
	}
}

// BuildChainReader builds a new ChainReader instance with the necessary contract bindings.
func BuildChainReader(
	avsAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
) (*ChainReader, error) {
	contractBindings, err := NewContractBindings(avsAddr, ethClient, logger)
	if err != nil {
		return nil, err
	}
	return NewChainReader(contractBindings.AVSManager, logger, ethClient), nil
}

// GetOptInOperators returns the list of operators who have opted into the AVS.
func (r *ChainReader) GetOptInOperators(opts *bind.CallOpts, avsAddress string) ([]gethcommon.Address, error) {
	return r.avsManager.GetOptInOperators(opts, gethcommon.HexToAddress(avsAddress))
}

// GetRegisteredPubkey returns the registered BLS public key of an operator.
func (r *ChainReader) GetRegisteredPubkey(opts *bind.CallOpts, avsAddress string, operator gethcommon.Address) ([]byte, error) {
	return r.avsManager.GetRegisteredPubkey(opts, gethcommon.HexToAddress(avsAddress), operator)
}

// GtAVSUSDValue returns the total USD value of the AVS.
func (r *ChainReader) GtAVSUSDValue(opts *bind.CallOpts, avsAddress string) (sdkmath.LegacyDec, error) {
	amount, err := r.avsManager.GetAVSUSDValue(opts, gethcommon.HexToAddress(avsAddress))
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}
	return sdkmath.LegacyNewDecFromBigInt(amount), nil
}

// GetOperatorOptedUSDValue returns the USD value of an operator's stake in the AVS.
func (r *ChainReader) GetOperatorOptedUSDValue(opts *bind.CallOpts, avsAddress string, operatorAddr gethcommon.Address) (sdkmath.LegacyDec, error) {
	amount, err := r.avsManager.GetOperatorOptedUSDValue(opts, gethcommon.HexToAddress(avsAddress), operatorAddr)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}
	return sdkmath.LegacyNewDecFromBigInt(amount), nil
}

// GetAVSEpochIdentifier returns the epoch identifier for the AVS.
func (r *ChainReader) GetAVSEpochIdentifier(opts *bind.CallOpts, avsAddress string) (string, error) {
	return r.avsManager.GetAVSEpochIdentifier(opts, gethcommon.HexToAddress(avsAddress))
}

// GetTaskInfo returns information about a specific task.
func (r *ChainReader) GetTaskInfo(opts *bind.CallOpts, avsAddress string, taskID uint64) (himera_avs.TaskInfo, error) {
	return r.avsManager.GetTaskInfo(opts, gethcommon.HexToAddress(avsAddress), taskID)
}

// IsOperator returns true if the given address is a registered operator.
func (r *ChainReader) IsOperator(opts *bind.CallOpts, operator gethcommon.Address) (bool, error) {
	return r.avsManager.IsOperator(opts, operator)
}

// GetCurrentEpoch returns the current epoch number.
func (r *ChainReader) GetCurrentEpoch(opts *bind.CallOpts, epochIdentifier string) (int64, error) {
	return r.avsManager.GetCurrentEpoch(opts, epochIdentifier)
}

// GetChallengeInfo returns information about a challenge for a specific task.
func (r *ChainReader) GetChallengeInfo(opts *bind.CallOpts, taskAddress string, taskID uint64) (gethcommon.Address, error) {
	return r.avsManager.GetChallengeInfo(opts, gethcommon.HexToAddress(taskAddress), taskID)
}

// GetOperatorTaskResponse returns the task response of a specific operator for a given task.
func (r *ChainReader) GetOperatorTaskResponse(opts *bind.CallOpts, taskAddress string, operatorAddress gethcommon.Address, taskID uint64) (himera_avs.TaskResultInfo, error) {
	return r.avsManager.GetOperatorTaskResponse(opts, gethcommon.HexToAddress(taskAddress), operatorAddress, taskID)
}

// GetOperatorTaskResponseList returns a list of all operator responses for a given task.
func (r *ChainReader) GetOperatorTaskResponseList(opts *bind.CallOpts, taskAddress string, taskID uint64) ([]himera_avs.OperatorResInfo, error) {
	return r.avsManager.GetOperatorTaskResponseList(opts, gethcommon.HexToAddress(taskAddress), taskID)
}
