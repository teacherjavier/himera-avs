package challenge

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
	sdklogging "github.com/imua-xyz/imua-avs-sdk/logging"
	"github.com/imua-xyz/imua-avs-sdk/signer"
	
	// ¡IMPORTANTE! Asegúrate de que estas rutas de importación sean correctas
	// himera_avs "github.com/himera-avs/src/bindings/bindings_himera_avs"
	// "github.com/himera-avs/utils"

	himera_avs "github.com/imua-xyz/imua-avs/contracts/bindings/avs"
	"github.com/imua-xyz/imua-avs-sdk/utils"

	chain "github.com/imua-xyz/imua-avs/core/chainio"
	"github.com/imua-xyz/imua-avs/core/chainio/eth"
	"github.com/imua-xyz/imua-avs/types"
)

type Challenger struct {
	config         types.NodeConfig
	logger         sdklogging.Logger
	ethClient      eth.EthClient
	ethWsClient    *ethclient.Client
	avsWriter      chain.AvsWriter
	avsReader      chain.ChainReader
	avsAddr        common.Address
	contractABI    abi.ABI
	taskDefHashMap map[common.Hash]uint8
}

func NewChallengerFromConfig(c types.NodeConfig) (*Challenger, error) {
	logger, err := sdklogging.NewZapLogger(sdklogging.Development)
	if err != nil { return nil, fmt.Errorf("cannot create logger: %w", err) }

	ethRpcClient, err := eth.NewClient(c.EthRpcUrl)
	if err != nil { return nil, fmt.Errorf("cannot create http eth client: %w", err) }
	
	ethWsClient, err := ethclient.Dial(c.EthWsUrl)
	if err != nil { return nil, fmt.Errorf("cannot create ws eth client: %w", err) }

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil { return nil, fmt.Errorf("cannot get chainId: %w", err) }

	ecdsaKeyPassword, _ := os.LookupEnv("AVS_ECDSA_KEY_PASSWORD")
	signer, challengeSender, err := signer.SignerFromConfig(signer.Config{
		KeystorePath: c.AVSEcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil { return nil, fmt.Errorf("cannot create signer: %w", err) }

	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, logger, signer, challengeSender)
	avsAddr := common.HexToAddress(c.AVSAddress)

	avsReader, err := chain.BuildChainReader(avsAddr, ethRpcClient, logger)
	if err != nil { return nil, err }

	avsWriter, err := chain.BuildChainWriter(avsAddr, ethRpcClient, logger, txMgr)
	if err != nil { return nil, err }

	contractABI, err := himera_avs.HimeraAvsMetaData.GetAbi()
	if err != nil { return nil, fmt.Errorf("cannot get HimeraAvs ABI: %w", err) }
	
	challenger := &Challenger{
		config:         c,
		logger:         logger,
		ethClient:      ethRpcClient,
		ethWsClient:    ethWsClient,
		avsWriter:      avsWriter,
		avsReader:      *avsReader,
		avsAddr:        avsAddr,
		contractABI:    *contractABI,
		taskDefHashMap: make(map[common.Hash]uint8),
	}

	if err := challenger.populateTaskDefMap(); err != nil {
		return nil, fmt.Errorf("failed to populate task definition map: %w", err)
	}
	
	logger.Info("HIMERA Challenger Initialized", "challengerAddress", challengeSender.String())
	return challenger, nil
}

func (c *Challenger) populateTaskDefMap() error {
	himeraContract, err := himera_avs.NewContractHimeraAvs(c.avsAddr, c.ethClient)
	if err != nil { return err }

	for i := uint8(1); i < 10; i++ {
		taskDef, err := himeraContract.GetTaskDefinition(&bind.CallOpts{}, i)
		if err != nil { break }
		
        taskDefStructForEncoding := himera_avs.HimeraTaskLibraryTaskDefinition{
            Id:          taskDef.Id,
            Name:        taskDef.Name,
            TaskType:    taskDef.TaskType,
            Description: taskDef.Description,
        }
        
        defBytes, err := c.contractABI.Methods["getTaskDefinition"].Inputs.Pack(taskDefStructForEncoding)
        if err != nil { continue }
        defHash := crypto.Keccak256Hash(defBytes)

		c.taskDefHashMap[defHash] = taskDef.Id
		c.logger.Info("Cached task definition", "id", taskDef.Id, "name", taskDef.Name, "hash", defHash.Hex())
	}
	return nil
}

func (c *Challenger) Start(ctx context.Context) error {
	c.logger.Info("Starting HIMERA Challenger Service...")
	query := ethereum.FilterQuery{Addresses: []common.Address{c.avsAddr}}
	logs := make(chan types.Log)
	sub, err := c.ethWsClient.SubscribeFilterLogs(ctx, query, logs)
	if err != nil { return err }
	defer sub.Unsubscribe()
	c.logger.Info("Listening for HimeraTaskCreated events to schedule challenges...")

	for {
		select {
		case <-ctx.Done():
			return nil
		case err := <-sub.Err():
			c.logger.Error("Subscription error", "err", err)
			return err
		case vLog := <-logs:
			if vLog.Topics[0] == c.contractABI.Events["HimeraTaskCreated"].ID {
				var event himera_avs.ContractHimeraAvsHimeraTaskCreated
				if err := c.contractABI.UnpackIntoInterface(&event, "HimeraTaskCreated", vLog.Data); err != nil {
					c.logger.Error("Failed to unpack event", "err", err)
					continue
				}
				go c.handleNewTaskForChallenge(ctx, event)
			}
		}
	}
}

func (c *Challenger) handleNewTaskForChallenge(ctx context.Context, taskEvent himera_avs.ContractHimeraAvsHimeraTaskCreated) {
    taskID := taskEvent.ImuaTaskId.Uint64()
    c.logger.Info("New task detected, scheduling for challenge resolution", "taskID", taskID)
    taskInfo, err := c.avsReader.GetTaskInfo(&bind.CallOpts{}, c.avsAddr.String(), taskID)
    if err != nil {
        c.logger.Error("Could not get task info, aborting", "taskID", taskID, "err", err)
        return
    }

    waitUntil := time.Unix(int64(taskInfo.StartingEpoch + taskInfo.TaskResponsePeriod + taskInfo.TaskStatisticalPeriod + 1), 0)
    c.logger.Info("Waiting until challenge period starts...", "taskID", taskID, "waitUntil", waitUntil)
    
	if time.Now().Before(waitUntil) {
		time.Sleep(time.Until(waitUntil))
	}

    himeraTaskDefId, ok := c.taskDefHashMap[taskInfo.Hash]
    if !ok {
        c.logger.Error("Could not find Himera task definition for hash", "hash", taskInfo.Hash.Hex())
        return
    }
    
    if err := c.ResolveChallenge(ctx, taskID, himeraTaskDefId, taskInfo.TaskInput); err != nil {
        c.logger.Error("Failed to resolve challenge", "taskID", taskID, "err", err)
    }
}

func (c *Challenger) ResolveChallenge(ctx context.Context, taskID uint64, taskDefId uint8, taskInput []byte) error {
	c.logger.Info("Resolving challenge for task", "taskID", taskID)

	operatorResponses, err := c.avsReader.GetOperatorTaskResponseList(&bind.CallOpts{}, c.avsAddr.String(), taskID)
	if err != nil { return err }
	taskInfo, err := c.avsReader.GetTaskInfo(&bind.CallOpts{}, c.avsAddr.String(), taskID)
	if err != nil { return err }
	
	totalTaskPower, ok := new(big.Int).SetString(taskInfo.TaskTotalPower, 10)
    if !ok || totalTaskPower.Cmp(big.NewInt(0)) == 0 {
        c.logger.Info("Task has no power, skipping challenge", "taskID", taskID)
        return nil
    }

	var groundTruth bool
	switch taskDefId {
	case 1: // VERIFY_EVM_AUTHORIZATION
		groundTruth, err = c.getGroundTruthForAuthTask(taskInput)
	case 2: // AUDIT_STRATEGY_EXECUTION
		groundTruth, err = c.getGroundTruthForAuditStrategyTask(taskInput)
	case 3: // AUDIT_ACCOUNT_BALANCE
		groundTruth, err = c.getGroundTruthForAuditBalanceTask(taskInput)
	default:
		return fmt.Errorf("challenge logic not implemented for task ID: %d", taskDefId)
	}
	if err != nil { return fmt.Errorf("failed to determine ground truth for task %d: %w", taskID, err) }
    c.logger.Info("Ground truth calculated", "taskID", taskID, "expectedResult", groundTruth)

	var eligibleRewardOperators []common.Address
	var eligibleSlashOperators []common.Address
	approvedPower := big.NewInt(0)
	boolType, _ := abi.NewType("bool", "", nil)
    arguments := abi.Arguments{{Type: boolType}}

	for _, resInfo := range operatorResponses {
		unpacked, err := arguments.Unpack(resInfo.TaskResponse)
        operatorResponse := !groundTruth
		if err == nil { operatorResponse = unpacked[0].(bool) }
		
		if operatorResponse == groundTruth {
			eligibleRewardOperators = append(eligibleRewardOperators, resInfo.OperatorAddress)
            approvedPower.Add(approvedPower, resInfo.Power)
		} else {
			eligibleSlashOperators = append(eligibleSlashOperators, resInfo.OperatorAddress)
		}
	}
    
    actualThresholdBig := new(big.Int).Mul(approvedPower, big.NewInt(100))
    actualThresholdBig.Div(actualThresholdBig, totalTaskPower)
    actualThreshold := uint8(actualThresholdBig.Uint64())
    isExpected := actualThreshold >= taskInfo.ThresholdPercentage

    c.logger.Info("Submitting final verdict to the chain...", "taskID", taskID, "actualThreshold", actualThreshold, "isExpected", isExpected)
	
	_, err = c.avsWriter.Challenge(ctx, taskID, actualThreshold, isExpected, eligibleRewardOperators, eligibleSlashOperators)
	if err != nil { return err }

	c.logger.Info("Challenge resolved and submitted successfully", "taskID", taskID)
	return nil
}

// --- Lógica de Negocio Específica de HIMERA para Resolver Desafíos ---

func (c *Challenger) getGroundTruthForAuthTask(taskInput []byte) (bool, error) {
	addressType, _ := abi.NewType("address", "", nil)
	bytes32Type, _ := abi.NewType("bytes32", "", nil)
	bytesType, _ := abi.NewType("bytes", "", nil)
	arguments := abi.Arguments{{Type: addressType}, {Type: bytes32Type}, {Type: bytesType}}
	unpacked, err := arguments.Unpack(taskInput)
	if err != nil { return false, err }

	userAddress := unpacked[0].(common.Address)
	messageHash := unpacked[1].([32]byte)
	signature := unpacked[2].([]byte)
	
	if len(signature) == 65 && (signature[64] == 27 || signature[64] == 28) {
		signature[64] -= 27 
	}
	pubKey, err := crypto.Ecrecover(messageHash[:], signature)
	if err != nil { return false, err }

	recoveredAddr := crypto.PubkeyToAddress(*crypto.UnmarshalPubkey(pubKey))
	return userAddress == recoveredAddr, nil
}

func (c *Challenger) getGroundTruthForAuditStrategyTask(taskInput []byte) (bool, error) {
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

    sourceRpc := os.Getenv("BASE_SEPOLIA_RPC_URL")
    bridgeContractAddress := common.HexToAddress("0x...ADDRESS_DEL_PUENTE") // Placeholder
    
    return utils.AuditEvmTransaction(sourceRpc, common.BytesToHash(sourceTxHash[:]), userEvmAA, bridgeContractAddress, expectedAmount)
}

func (c *Challenger) getGroundTruthForAuditBalanceTask(taskInput []byte) (bool, error) {
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
    
    return ethMatch && baseMatch && starknetMatch, nil
}
