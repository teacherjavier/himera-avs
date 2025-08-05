// core/chainio/avs_writer.go
package chainio

import (
	"context"
	"errors"
	
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	gethcommon "github.com/ethereum/go-ethereum/common"
	gethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/imua-xyz/imua-avs-sdk/client/txmgr"
	"github.com/imua-xyz/imua-avs-sdk/logging"
	
	himera_avs "github.com/imua-xyz/imua-avs/contracts/bindings/avs"
	"github.com/imua-xyz/imua-avs/types"
)

type AvsWriter interface {
	RegisterAVSToChain(ctx context.Context, params types.AVSParams) (*gethtypes.Receipt, error)
	RegisterBLSPublicKey(ctx context.Context, pubKey []byte, pubKeyRegistrationSignature []byte) (*gethtypes.Receipt, error)
	CreateHimeraTask(ctx context.Context, taskDefId uint8, taskInput []byte) (*gethtypes.Receipt, error)
	OperatorSubmitTask(ctx context.Context, taskID uint64, taskResponse []byte, blsSignature []byte, phase uint8) (*gethtypes.Receipt, error)
	Challenge(ctx context.Context, taskID uint64, actualThreshold uint8, isExpected bool, eligibleRewardOperators []gethcommon.Address, eligibleSlashOperators []gethcommon.Address) (*gethtypes.Receipt, error)
	RegisterOperatorToAVS(ctx context.Context) (*gethtypes.Receipt, error)
	GetSigner() bind.SignerFn
}

type ChainWriter struct {
	avsManager  *himera_avs.ContractHimeraAvs
	chainReader AvsReader
	ethClient   eth.EthClient
	logger      logging.Logger
	txMgr       txmgr.TxManager
}

var _ AvsWriter = (*ChainWriter)(nil)

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

func BuildChainWriter(
	avsAddr gethcommon.Address,
	ethClient eth.EthClient,
	logger logging.Logger,
	txMgr txmgr.TxManager,
) (*ChainWriter, error) {
	contractBindings, err := NewContractBindings(avsAddr, ethClient, logger)
	if err != nil { return nil, err }
	chainReader, err := BuildChainReader(avsAddr, ethClient, logger)
	if err != nil { return nil, err }
	return NewChainWriter(contractBindings.AVSManager, chainReader, ethClient, logger, txMgr), nil
}

func (w *ChainWriter) GetSigner() bind.SignerFn {
	return w.txMgr.GetSigner()
}

func (w *ChainWriter) RegisterAVSToChain(ctx context.Context, params types.AVSParams) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil { return nil, err }
	tx, err := w.avsManager.RegisterAVS(noSendTxOpts, params)
	if err != nil { return nil, err }
	return w.txMgr.Send(ctx, tx)
}

func (w *ChainWriter) RegisterBLSPublicKey(ctx context.Context, pubKey []byte, pubKeyRegistrationSignature []byte) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil { return nil, err }
	tx, err := w.avsManager.RegisterBLSPublicKey(noSendTxOpts, pubKey, pubKeyRegistrationSignature)
	if err != nil { return nil, err }
	return w.txMgr.Send(ctx, tx)
}

func (w *ChainWriter) CreateHimeraTask(ctx context.Context, taskDefId uint8, taskInput []byte) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil { return nil, err }
	tx, err := w.avsManager.CreateHimeraTask(noSendTxOpts, taskDefId, taskInput)
	if err != nil { return nil, err }
	return w.txMgr.Send(ctx, tx)
}

func (w *ChainWriter) OperatorSubmitTask(ctx context.Context, taskID uint64, taskResponse []byte, blsSignature []byte, phase uint8) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil { return nil, err }
	tx, err := w.avsManager.OperatorSubmitTask(noSendTxOpts, taskID, taskResponse, blsSignature, phase)
	if err != nil { return nil, err }
	return w.txMgr.Send(ctx, tx)
}

func (w *ChainWriter) Challenge(ctx context.Context, taskID uint64, actualThreshold uint8, isExpected bool, eligibleRewardOperators []gethcommon.Address, eligibleSlashOperators []gethcommon.Address) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil { return nil, err }
	tx, err := w.avsManager.Challenge(noSendTxOpts, taskID, actualThreshold, isExpected, eligibleRewardOperators, eligibleSlashOperators)
	if err != nil { return nil, err }
	return w.txMgr.Send(ctx, tx)
}

func (w *ChainWriter) RegisterOperatorToAVS(ctx context.Context) (*gethtypes.Receipt, error) {
	noSendTxOpts, err := w.txMgr.GetNoSendTxOpts()
	if err != nil { return nil, err }
	tx, err := w.avsManager.RegisterOperatorToAVS(noSendTxOpts)
	if err != nil { return nil, err }
	return w.txMgr.Send(ctx, tx)
}
