package operator

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/imua-xyz/imua-avs-sdk/client/txmgr"
	"github.com/imua-xyz/imua-avs-sdk/crypto/bls"
	sdklogging "github.com/imua-xyz/imua-avs-sdk/logging"
	"github.com/imua-xyz/imua-avs-sdk/signer"
	
	// ¡IMPORTANTE! Asegúrate de que estas rutas sean correctas para tu proyecto
	// himera_avs "github.com/himera-avs/src/bindings/bindings_himera_avs"
	// "github.com/himera-avs/utils"

	himera_avs "github.com/imua-xyz/imua-avs/contracts/bindings/avs"
	"github.com/imua-xyz/imua-avs-sdk/utils"
	"github.com/imua-xyz/imua-avs/core"
	chain "github.com/imua-xyz/imua-avs/core/chainio"
	"github.com/imua-xyz/imua-avs/core/chainio/eth"
	"github.com/imua-xyz/imua-avs/types"
	blscommon "github.com/prysmaticlabs/prysm/v5/crypto/bls/common"
)

const retryDelay = 3 * time.Second

type Operator struct {
	config             types.NodeConfig
	logger             sdklogging.Logger
	ethClient          eth.EthClient
	avsWriter          chain.AvsWriter
	avsReader          chain.ChainReader
	ethWsClient        *ethclient.Client
	blsKeypair         blscommon.SecretKey
	operatorAddr       common.Address
	avsAddr            common.Address
	contractABI        abi.ABI
	newTaskCreatedChan chan *himera_avs.ContractHimeraAvsHimeraTaskCreated
}

func NewOperatorFromConfig(c types.NodeConfig) (*Operator, error) {
	logger, err := sdklogging.NewZapLogger(sdklogging.Development)
	if err != nil { return nil, fmt.Errorf("cannot create logger: %w", err) }

	ethRpcClient, err := eth.NewClient(c.EthRpcUrl)
	if err != nil { return nil, fmt.Errorf("cannot create http eth client: %w", err) }

	ethWsClient, err := ethclient.Dial(c.EthWsUrl)
	if err != nil { return nil, fmt.Errorf("cannot create ws eth client: %w", err) }

	blsKeyPassword, _ := os.LookupEnv("OPERATOR_BLS_KEY_PASSWORD")
	blsKeyPair, err := bls.ReadPrivateKeyFromFile(c.BlsPrivateKeyStorePath, blsKeyPassword)
	if err != nil { return nil, fmt.Errorf("cannot read BLS private key: %w", err) }

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil { return nil, fmt.Errorf("cannot get chainId: %w", err) }

	ecdsaKeyPassword, _ := os.LookupEnv("OPERATOR_ECDSA_KEY_PASSWORD")
	signer, operatorSender, err := signer.SignerFromConfig(signer.Config{
		KeystorePath: c.OperatorEcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil { return nil, fmt.Errorf("cannot create signer: %w", err) }

	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, logger, signer, operatorSender)
	avsAddr := common.HexToAddress(c.AVSAddress)

	avsReader, err := chain.BuildChainReader(avsAddr, ethRpcClient, logger)
	if err != nil { return nil, fmt.Errorf("cannot build avsReader: %w", err) }

	avsWriter, err := chain.BuildChainWriter(avsAddr, ethRpcClient, logger, txMgr)
	if err != nil { return nil, fmt.Errorf("cannot build avsWriter: %w", err) }

	contractABI, err := himera_avs.HimeraAvsMetaData.GetAbi()
	if err != nil { return nil, fmt.Errorf("cannot get HimeraAvs ABI: %w", err) }

	operator := &Operator{
		config:             c,
		logger:             logger,
		ethClient:          ethRpcClient,
		avsWriter:          avsWriter,
		avsReader:          *avsReader,
		ethWsClient:        ethWsClient,
		blsKeypair:         blsKeyPair,
		operatorAddr:       operatorSender,
		avsAddr:            avsAddr,
		contractABI:        *contractABI,
		newTaskCreatedChan: make(chan *himera_avs.ContractHimeraAvsHimeraTaskCreated),
	}

	logger.Info("HIMERA Operator Initialized", "operatorAddress", operatorSender.String())
	return operator, nil
}

func (o *Operator) Start(ctx context.Context) error {
	o.logger.Info("Starting HIMERA Operator Service...")

	go o.taskProcessor(ctx)

	query := ethereum.FilterQuery{Addresses: []common.Address{o.avsAddr}}
	logs := make(chan types.Log)
	sub, err := o.ethWsClient.SubscribeFilterLogs(ctx, query, logs)
	if err != nil { return fmt.Errorf("failed to subscribe to logs: %w", err) }
	defer sub.Unsubscribe()

	o.logger.Info("Listening for HimeraTaskCreated events...")

	for {
		select {
		case <-ctx.Done():
			o.logger.Info("Operator shutting down.")
			return nil
		case err := <-sub.Err():
			o.logger.Error("Subscription error", "err", err)
			return err
		case vLog := <-logs:
			if vLog.Topics[0] == o.contractABI.Events["HimeraTaskCreated"].ID {
				var event himera_avs.ContractHimeraAvsHimeraTaskCreated
				if err := o.contractABI.UnpackIntoInterface(&event, "HimeraTaskCreated", vLog.Data); err != nil {
					o.logger.Error("Failed to unpack event", "err", err)
					continue
				}
				o.newTaskCreatedChan <- &event
			}
		}
	}
}

func (o *Operator) taskProcessor(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case taskEvent := <-o.newTaskCreatedChan:
			taskResponse := o.ProcessNewTaskCreatedLog(taskEvent)
			
			sig, resBytes, err := o.SignTaskResponse(taskResponse)
			if err != nil {
				o.logger.Error("Failed to sign task response", "err", err, "TaskID", taskResponse.TaskID)
				continue
			}
			
			taskInfo, err := o.avsReader.GetTaskInfo(&bind.CallOpts{}, o.avsAddr.String(), taskResponse.TaskID)
			if err != nil {
				o.logger.Error("Failed to get task info", "err", err, "TaskID", taskResponse.TaskID)
				continue
			}

			go o.SendSignedTaskResponseToChain(context.Background(), taskResponse.TaskID, resBytes, sig, taskInfo)
		}
	}
}

func (o *Operator) ProcessNewTaskCreatedLog(e *himera_avs.ContractHimeraAvsHimeraTaskCreated) *core.TaskResponse {
	o.logger.Info("New HIMERA Task Received", "TaskID", e.ImuaTaskId.Uint64(), "TaskDefId", e.HimeraTaskDefId)
	var taskResult bool
	var err error

	switch e.HimeraTaskDefId {
	case 1: // VERIFY_EVM_AUTHORIZATION
		taskResult, err = o.handleVerifyEvmAuth(e.TaskInput)
	case 2: // AUDIT_STRATEGY_EXECUTION
		taskResult, err = o.handleAuditStrategy(e.TaskInput)
	case 3: // AUDIT_ACCOUNT_BALANCE
		taskResult, err = o.handleAuditBalance(e.TaskInput)
	default:
		err = fmt.Errorf("unknown task definition id: %d", e.HimeraTaskDefId)
	}

	if err != nil {
		o.logger.Error("Error processing task", "TaskID", e.ImuaTaskId.Uint64(), "err", err)
		taskResult = false
	}
	o.logger.Info("Task Processed", "TaskID", e.ImuaTaskId.Uint64(), "Result", taskResult)
	return &core.TaskResponse{TaskID: e.ImuaTaskId.Uint64(), TaskResult: taskResult}
}

func (o *Operator) SignTaskResponse(taskResponse *core.TaskResponse) ([]byte, []byte, error) {
	boolType, _ := abi.NewType("bool", "", nil)
    arguments := abi.Arguments{{Type: boolType}}
    responseBytes, err := arguments.Pack(taskResponse.TaskResult.(bool))
    if err != nil { return nil, nil, fmt.Errorf("failed to pack task result: %w", err) }
	
	responseHash := crypto.Keccak256Hash(responseBytes)
	sig := o.blsKeypair.Sign(responseHash[:])
	return sig.Marshal(), responseBytes, nil
}

func (o *Operator) SendSignedTaskResponseToChain(ctx context.Context, taskId uint64, taskResponse []byte, blsSignature []byte, taskInfo types.TaskInfo) (string, error) {
	startingEpoch := taskInfo.StartingEpoch
	taskResponsePeriod := taskInfo.TaskResponsePeriod
	taskStatisticalPeriod := taskInfo.TaskStatisticalPeriod
	phaseOneSubmitted := false
	phaseTwoSubmitted := false

	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		default:
			epochIdentifier, err := o.avsReader.GetAVSEpochIdentifier(&bind.CallOpts{}, o.avsAddr.String())
			if err != nil { return "", fmt.Errorf("failed to get AVS info: %w", err) }
			num, err := o.avsReader.GetCurrentEpoch(&bind.CallOpts{}, epochIdentifier)
			if err != nil { return "", fmt.Errorf("failed to get current epoch: %w", err) }
			
			currentEpoch := uint64(num)
			if currentEpoch > startingEpoch+taskResponsePeriod+taskStatisticalPeriod {
				o.logger.Info("Task period has passed, exiting.", "TaskID", taskId)
				return "Task period has passed.", nil
			}

			if currentEpoch <= startingEpoch+taskResponsePeriod {
				if !phaseOneSubmitted {
					o.logger.Info("Executing Phase One Submission (Commit)", "TaskID", taskId)
					_, err := o.avsWriter.OperatorSubmitTask(ctx, taskId, nil, blsSignature, o.avsAddr, 1)
					if err != nil { o.logger.Error("Failed Phase One Submission", "err", err); time.Sleep(retryDelay); continue }
					phaseOneSubmitted = true
					o.logger.Info("Phase One submitted successfully", "TaskID", taskId)
				}
			} else {
				if !phaseTwoSubmitted {
					o.logger.Info("Executing Phase Two Submission (Reveal)", "TaskID", taskId)
					_, err := o.avsWriter.OperatorSubmitTask(ctx, taskId, taskResponse, blsSignature, o.avsAddr, 2)
					if err != nil { o.logger.Error("Failed Phase Two Submission", "err", err); time.Sleep(retryDelay); continue }
					phaseTwoSubmitted = true
					o.logger.Info("Phase Two submitted successfully", "TaskID", taskId)
				}
			}
			if phaseOneSubmitted && phaseTwoSubmitted {
				o.logger.Info("Both phases completed successfully", "TaskID", taskId)
				return "Both phases completed.", nil
			}
			time.Sleep(retryDelay)
		}
	}
}

// --- Task Handlers (Lógica de Negocio de HIMERA) ---

func (o *Operator) handleVerifyEvmAuth(taskInput []byte) (bool, error) {
	addressType, _ := abi.NewType("address", "", nil)
	bytes32Type, _ := abi.NewType("bytes32", "", nil)
	bytesType, _ := abi.NewType("bytes", "", nil)
	arguments := abi.Arguments{{Type: addressType}, {Type: bytes32Type}, {Type: bytesType}}
	unpacked, err := arguments.Unpack(taskInput)
	if err != nil { return false, fmt.Errorf("failed to unpack taskInput: %w", err) }
	userAddress := unpacked[0].(common.Address)
	messageHash := unpacked[1].([32]byte)
	signature := unpacked[2].([]byte)
	o.logger.Info("Verifying signature...", "userAddress", userAddress.Hex(), "messageHash", hexutil.Encode(messageHash[:]))
	if len(signature) == 65 && (signature[64] == 27 || signature[64] == 28) {
		signature[64] -= 27
	}
	pubKey, err := crypto.Ecrecover(messageHash[:], signature)
	if err != nil { return false, fmt.Errorf("failed to recover public key: %w", err) }
	recoveredAddr := crypto.PubkeyToAddress(*crypto.UnmarshalPubkey(pubKey))
	isValid := userAddress == recoveredAddr
	o.logger.Info("Verification complete", "isValid", isValid, "recoveredAddress", recoveredAddr.Hex())
	return isValid, nil
}

func (o *Operator) handleAuditStrategy(taskInput []byte) (bool, error) {
	addrType, _ := abi.NewType("address", "", nil)
    bytes32Type, _ := abi.NewType("bytes32", "", nil)
    uint256Type, _ := abi.NewType("uint256", "", nil)
    arguments := abi.Arguments{{Type: addrType}, {Type: bytes32Type}, {Type: bytes32Type}, {Type: uint256Type}}
    unpacked, err := arguments.Unpack(taskInput)
	if err != nil { return false, err }
    userEvmAA := unpacked[0].(common.Address)
    sourceTxHash := unpacked[1].([32]byte)
    // destTxHash := unpacked[2].([32]byte)
    expectedAmount := unpacked[3].(*big.Int)
    o.logger.Info("Auditing strategy...", "user", userEvmAA.Hex(), "sourceTx", hexutil.Encode(sourceTxHash[:]))
    sourceRpc := os.Getenv("BASE_SEPOLIA_RPC_URL")
    bridgeContractAddress := common.HexToAddress("0x...ADDRESS_DEL_PUENTE") // Placeholder
    isValid, err := utils.AuditEvmTransaction(sourceRpc, common.BytesToHash(sourceTxHash[:]), userEvmAA, bridgeContractAddress, expectedAmount)
    if err != nil { return false, err }
    return isValid, nil
}

func (o *Operator) handleAuditBalance(taskInput []byte) (bool, error) {
    addrType, _ := abi.NewType("address", "", nil)
    balType, _ := abi.NewType("uint256", "", nil)
    arguments := abi.Arguments{
        {Type: addrType}, {Type: balType}, {Type: addrType}, {Type: balType}, {Type: balType}, {Type: balType},
    }
    unpacked, err := arguments.Unpack(taskInput)
	if err != nil { return false, err }
    userEthAddr := unpacked[0].(common.Address)
    expectedEthBal := unpacked[1].(*big.Int)
    userBaseAddr := unpacked[2].(common.Address)
    expectedBaseBal := unpacked[3].(*big.Int)
    userStarknetAddrBig := unpacked[4].(*big.Int)
    expectedStarknetBal := unpacked[5].(*big.Int)
    userStarknetAddr := fmt.Sprintf("0x%s", userStarknetAddrBig.Text(16))
    o.logger.Info("Auditing multi-chain balance...", "userEth", userEthAddr.Hex(), "userStarknet", userStarknetAddr)

    ethRpc := os.Getenv("ETH_SEPOLIA_RPC_URL")
    baseRpc := os.Getenv("BASE_SEPOLIA_RPC_URL")
    starknetRpc := os.Getenv("STARKNET_SEPOLIA_RPC_URL")
    usdcEthAddr := common.HexToAddress("0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7a08")
    usdcBaseAddr := common.HexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e")

    actualEthBal, err1 := utils.GetEvmBalance(ethRpc, usdcEthAddr, userEthAddr)
    actualBaseBal, err2 := utils.GetEvmBalance(baseRpc, usdcBaseAddr, userBaseAddr)
    actualStarknetBal, err3 := utils.GetStarknetStrkBalance(starknetRpc, userStarknetAddr)
    if err1 != nil || err2 != nil || err3 != nil {
        return false, fmt.Errorf("failed fetching balances: %v, %v, %v", err1, err2, err3)
    }

    ethMatch := actualEthBal.Cmp(expectedEthBal) == 0
    baseMatch := actualBaseBal.Cmp(expectedBaseBal) == 0
    starknetMatch := actualStarknetBal.Cmp(expectedStarknetBal) == 0
    o.logger.Info("Balance Audit Result", "ethMatch", ethMatch, "baseMatch", baseMatch, "starknetMatch", starknetMatch)
    return ethMatch && baseMatch && starknetMatch, nil
}
