// core/chainio/avs_writer.go
package chainio

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/imua-xyz/imua-avs-sdk/client/txmgr"
	"github.com/imua-xyz/imua-avs-sdk/logging"

	himera_avs "github.com/imua-xyz/imua-avs/contracts/bindings/avs"
	"github.com/imua-xyz/imua-avs/core/chainio/eth"
)

// AvsWriter defines the interface for writing data to the AVS contract.
type AvsWriter interface {
	// RegisterAVSToChain registers the AVS with the chain.
	RegisterAVSToChain(ctx context.Context, params himera_avs.AVSParams) (*gethtypes.Receipt, error)
	// RegisterBLSPublicKey registers the operator's BLS public key.
	RegisterBLSPublicKey(ctx context.Context, avsAddr gethcommon.Address, pubKey []byte, pubKeyRegistrationSignature []byte) (*gethtypes.Receipt, error)
	// CreateNewTask creates a new task in the AVS.
	CreateNewTask(ctx context.Context, name string, numberToBeSquared uint64, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint8, taskStatisticalPeriod uint64) (*gethtypes.Receipt, error)
	// OperatorSubmitTask submits a task response from an operator.
	OperatorSubmitTask(ctx context.Context, taskID uint64, taskResponse []byte, blsSignature []byte, taskContractAddress gethcommon.Address, phase uint8) (*gethtypes.Receipt, error)
	// RaiseAndResolveChallenge raises a challenge for a task.
	RaiseAndResolveChallenge(ctx context.Context, req himera_avs.AvsServiceContractChallengeReq) (*gethtypes.Receipt, error)
	// RegisterOperatorToAVS registers the operator with the AVS.
	RegisterOperatorToAVS(ctx context.Context) (*gethtypes.Receipt, error)
	// GetSigner returns the signer function for the transaction manager.
	GetSigner() bind.SignerFn
}

// ChainWriter implements the AvsWriter interface for writing data to the blockchain.
type ChainWriter struct {
	avsManager  *himera_avs.ContractHimeraAvs
	chainReader AvsReader
	ethClient   eth.EthClient
	logger      logging.Logger
	txMgr       txmgr.TxManager
}

// ensure ChainWriter implements AvsWriter
var _ AvsWriter = (*ChainWriter)(nil)

// NewChainWriter creates a new ChainWriter instance.
func NewChainWriter(
	avsManager *himera_avs.ContractHimeraAvs,
	chainReader AvsReader,
	ethClient eth.EthClient,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) *ChainWriter {
	return &ChainWriter{
		avsManager:  avsManager,
		chainReader: chainReader,
		logger:      logger,
		ethClient:   ethClient,
		txMgr:       txMgr,
	}
}

// BuildChainWriter builds a new ChainWriter instance with the necessary contract bindings.
func BuildChainWriter(
	avsAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) (*ChainWriter, error) {
	contractBindings, err := NewContractBindings(avsAddr, ethClient, logger)
	if err != nil {
		return nil, err
	}
	chainReader, err := BuildChainReader(avsAddr, ethClient, logger)
	if err != nil {
		return nil, err
	}
	return NewChainWriter(contractBindings.AVSManager, chainReader, ethClient, logger, txMgr), nil
}

// GetSigner returns the signer function for the transaction manager.
func (w *ChainWriter) GetSigner() bind.SignerFn {
	return w.txMgr.GetSigner()
}

// RegisterAVSToChain registers the AVS with the chain.
func (w *ChainWriter) RegisterAVSToChain(ctx context.Context, params himera_avs.AVSParams) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.RegisterAVS(noSendTxOpts, params)
	if err != nil {
		return nil, err
	}
	return w.txMgr.Send(ctx, tx)
}

// RegisterBLSPublicKey registers the operator's BLS public key.
func (w *ChainWriter) RegisterBLSPublicKey(ctx context.Context, avsAddr gethcommon.Address, pubKey []byte, pubKeyRegistrationSignature []byte) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.RegisterBLSPublicKey(noSendTxOpts, avsAddr, pubKey, pubKeyRegistrationSignature)
	if err != nil {
		return nil, err
	}
	return w.txMgr.Send(ctx, tx)
}

// CreateNewTask creates a new task in the AVS.
func (w *ChainWriter) CreateNewTask(ctx context.Context, name string, numberToBeSquared uint64, taskResponsePeriod uint64, taskChallengePeriod uint64, thresholdPercentage uint8, taskStatisticalPeriod uint64) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.CreateNewTask(noSendTxOpts, name, numberToBeSquared, taskResponsePeriod, taskChallengePeriod, thresholdPercentage, taskStatisticalPeriod)
	if err != nil {
		return nil, err
	}
	return w.txMgr.Send(ctx, tx)
}

// OperatorSubmitTask submits a task response from an operator.
func (w *ChainWriter) OperatorSubmitTask(ctx context.Context, taskID uint64, taskResponse []byte, blsSignature []byte, taskContractAddress gethcommon.Address, phase uint8) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.OperatorSubmitTask(noSendTxOpts, taskID, taskResponse, blsSignature, taskContractAddress, phase)
	if err != nil {
		return nil, err
	}
	return w.txMgr.Send(ctx, tx)
}

// RaiseAndResolveChallenge raises a challenge for a task.
func (w *ChainWriter) RaiseAndResolveChallenge(ctx context.Context, req himera_avs.AvsServiceContractChallengeReq) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.RaiseAndResolveChallenge(noSendTxOpts, req)
	if err != nil {
		return nil, err
	}
	return w.txMgr.Send(ctx, tx)
}

// RegisterOperatorToAVS registers the operator with the AVS.
func (w *ChainWriter) RegisterOperatorToAVS(ctx context.Context) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil {
		return nil, err
	}
	tx, err := w.avsManager.RegisterOperatorToAVS(noSendTxOpts)
	if err != nil {
		return nil, err
	}
	return w.txMgr.Send(ctx, tx)
}
