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

// La interfaz se actualiza para usar los tipos de IAVSManager que genera abigen para HIMERA
type AvsReader interface {
	GetOptInOperators(opts *bind.CallOpts) ([]gethcommon.Address, error)
	GetRegisteredPubkey(opts *bind.CallOpts, operator gethcommon.Address) ([]byte, error)
	GtAVSUSDValue(opts *bind.CallOpts) (sdkmath.LegacyDec, error)
	GetOperatorOptedUSDValue(opts *bind.CallOpts, operatorAddr gethcommon.Address) (sdkmath.LegacyDec, error)
	GetAVSEpochIdentifier(opts *bind.CallOpts) (string, error)
	GetTaskInfo(opts *bind.CallOpts, taskID uint64) (himera_avs.TaskInfo, error)
	IsOperator(opts *bind.CallOpts, operator gethcommon.Address) (bool, error)
	GetCurrentEpoch(opts *bind.CallOpts, epochIdentifier string) (int64, error)
	GetChallengeInfo(opts *bind.CallOpts, taskID uint64) (gethcommon.Address, error)
	GetOperatorTaskResponse(opts *bind.CallOpts, operatorAddress gethcommon.Address, taskID uint64) (himera_avs.IAVSManagerTaskResultInfo, error)
	GetOperatorTaskResponseList(opts *bind.CallOpts, taskID uint64) ([]himera_avs.OperatorResInfo, error)
}

type ChainReader struct {
	logger     logging.Logger
	avsManager *himera_avs.ContractHimeraAvs
	ethClient  eth.EthClient
}

var _ AvsReader = (*ChainReader)(nil)

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

// El resto de las funciones ahora usan r.avsManager que es del tipo correcto
func (r *ChainReader) GetOptInOperators(opts *bind.CallOpts) ([]gethcommon.Address, error) {
	return r.avsManager.GetOptInOperators(opts)
}
func (r *ChainReader) GetRegisteredPubkey(opts *bind.CallOpts, operator gethcommon.Address) ([]byte, error) {
	return r.avsManager.GetRegisteredPubkey(opts, operator)
}
func (r *ChainReader) GtAVSUSDValue(opts *bind.CallOpts) (sdkmath.LegacyDec, error) {
	amount, err := r.avsManager.GetAVSUSDValue(opts)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}
	return sdkmath.LegacyNewDecFromBigInt(amount), nil
}
func (r *ChainReader) GetOperatorOptedUSDValue(opts *bind.CallOpts, operatorAddr gethcommon.Address) (sdkmath.LegacyDec, error) {
	amount, err := r.avsManager.GetOperatorOptedUSDValue(opts, operatorAddr)
	if err != nil {
		return sdkmath.LegacyDec{}, err
	}
	return sdkmath.LegacyNewDecFromBigInt(amount), nil
}
func (r *ChainReader) GetAVSEpochIdentifier(opts *bind.CallOpts) (string, error) {
	return r.avsManager.GetAVSEpochIdentifier(opts)
}
func (r *ChainReader) GetTaskInfo(opts *bind.CallOpts, taskID uint64) (himera_avs.IAVSManagerTaskInfo, error) {
	return r.avsManager.GetTaskInfo(opts, taskID)
}
func (r *ChainReader) IsOperator(opts *bind.CallOpts, operator gethcommon.Address) (bool, error) {
	return r.avsManager.IsOperator(opts, operator)
}
func (r *ChainReader) GetCurrentEpoch(opts *bind.CallOpts, epochIdentifier string) (int64, error) {
	return r.avsManager.GetCurrentEpoch(opts, epochIdentifier)
}
func (r *ChainReader) GetChallengeInfo(opts *bind.CallOpts, taskID uint64) (gethcommon.Address, error) {
	return r.avsManager.GetChallengeInfo(opts, taskID)
}
func (r *ChainReader) GetOperatorTaskResponse(opts *bind.CallOpts, operatorAddress gethcommon.Address, taskID uint64) (himera_avs.IAVSManagerTaskResultInfo, error) {
	return r.avsManager.GetOperatorTaskResponse(opts, operatorAddress, taskID)
}
func (r *ChainReader) GetOperatorTaskResponseList(opts *bind.CallOpts, taskID uint64) ([]himera_avs.IAVSManagerOperatorResInfo, error) {
	return r.avsManager.GetOperatorTaskResponseList(opts, taskID)
}
