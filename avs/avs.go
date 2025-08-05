package avs

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	sdkmath "cosmossdk.io/math"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	sdkEcdsa "github.com/imua-xyz/imua-avs-sdk/crypto/ecdsa"
	sdklogging "github.com/imua-xyz/imua-avs-sdk/logging"
	"github.com/imua-xyz/imua-avs-sdk/signer"
	
	// --- INICIO DE LA CORRECCIÓN DEFINITIVA ---
	// Importamos la ruta correcta y le damos el alias que coincide con el nombre de nuestro paquete generado
	himera_avs "github.com/imua-xyz/imua-avs/contracts/bindings/avs"
	// --- FIN DE LA CORRECCIÓN DEFINITIVA ---

	"github.com/imua-xyz/imua-avs-sdk/client/txmgr"
	"github.com/imua-xyz/imua-avs/core"
	chain "github.com/imua-xyz/imua-avs/core/chainio"
	"github.com/imua-xyz/imua-avs/core/chainio/eth"
	"github.com/imua-xyz/imua-avs/types"
)

const (
	avsName       = "Himera-AVS"
	maxRetries    = 25
	retryDelay    = 6 * time.Second
	DayEpochID    = "day"
	HourEpochID   = "hour"
	MinuteEpochID = "minute"
	WeekEpochID   = "week"
)

type AvsService struct {
	logger                sdklogging.Logger
	avsWriter             chain.AvsWriter
	avsReader             chain.ChainReader
	config                types.NodeConfig
	signer                bind.SignerFn
	sender                common.Address
	ethClient             eth.EthClient
	taskResponsePeriod    uint64
	taskChallengePeriod   uint64
	thresholdPercentage   uint8
	taskStatisticalPeriod uint64
	avsEpochIdentifier    string
}

func NewAvsService(c types.NodeConfig) (*AvsService, error) {
	logger, err := sdklogging.NewZapLogger(sdklogging.Development)
	if err != nil { return nil, err }

	ethRpcClient, err := eth.NewClient(c.EthRpcUrl)
	if err != nil { return nil, err }

	chainId, err := ethRpcClient.ChainID(context.Background())
	if err != nil { return nil, err }

	ecdsaKeyPassword, _ := os.LookupEnv("AVS_ECDSA_KEY_PASSWORD")
	signer, avsSender, err := signer.SignerFromConfig(signer.Config{
		KeystorePath: c.AVSEcdsaPrivateKeyStorePath,
		Password:     ecdsaKeyPassword,
	}, chainId)
	if err != nil { return nil, err }

	logger.Info("AVS Service Signer (Owner):", "address", avsSender.String())
	
	avsAddr := common.HexToAddress(c.AVSAddress)
	code, _ := ethRpcClient.CodeAt(context.Background(), avsAddr, nil)
	
	txOpts := bind.NewKeyedTransactorWithChainID(signer.PrivateKey, chainId)

	if c.AVSAddress == "" || len(code) < 3 {
		logger.Info("AVS contract address not found or not deployed. Deploying and configuring new HimeraAvs contract...")
		
		key, err := sdkEcdsa.ReadKey(c.AVSEcdsaPrivateKeyStorePath, ecdsaKeyPassword)
		if err != nil { return nil, err }
		txOptsDeploy := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainId)
		
		// Usamos el Deployer de nuestro binding `himera_avs`
		deployedAddr, tx, _, err := himera_avs.DeployHimeraAvs(txOptsDeploy, ethRpcClient)
		if err != nil { return nil, fmt.Errorf("failed to deploy HimeraAvs contract: %w", err) }
		
		avsAddr = deployedAddr
		c.AVSAddress = deployedAddr.String()

		receipt, err := bind.WaitMined(context.Background(), ethRpcClient, tx)
		if err != nil { return nil, fmt.Errorf("failed to wait for contract deployment: %w", err) }
		if receipt.Status != 1 { return nil, fmt.Errorf("contract deployment failed") }
		
		logger.Info("HimeraAvs contract deployed", "address", avsAddr.Hex())
		
		himeraContract, err := himera_avs.NewHimeraAvs(avsAddr, ethRpcClient)
		if err != nil { return nil, err }

		tx, err = himeraContract.Initialize(txOpts, avsSender)
		if err != nil { return nil, fmt.Errorf("failed to initialize contract: %w", err) }
		bind.WaitMined(context.Background(), ethRpcClient, tx)
		logger.Info("Contract initialized.")

		tx, err = himeraContract.SetupTaskDefinitions(txOpts)
		if err != nil { return nil, fmt.Errorf("failed to setup tasks: %w", err) }
		bind.WaitMined(context.Background(), ethRpcClient, tx)
		logger.Info("Contract task definitions set up.")
	}

	txMgr := txmgr.NewSimpleTxManager(ethRpcClient, logger, signer, avsSender)
	avsWriter, err := chain.BuildChainWriter(avsAddr, ethRpcClient, logger, txMgr)
	if err != nil { return nil, err }
	avsReader, err := chain.BuildChainReader(avsAddr, ethRpcClient, logger)
	if err != nil { return nil, err }

	time.Sleep(retryDelay)
	epochId, err := avsReader.GetAVSEpochIdentifier(&bind.CallOpts{}, c.AVSAddress)
	if err != nil || epochId == "" {
		logger.Info("AVS not registered with IMUA. Registering now...")
		params := types.AVSParams{
			Sender:              avsSender,
			AvsName:             avsName,
			MinStakeAmount:      c.MinStakeAmount,
			TaskAddress:         avsAddr,
			SlashAddress:        avsAddr,
			RewardAddress:       avsAddr,
			AvsOwnerAddresses:   core.ConvertToEthAddresses(c.AvsOwnerAddresses),
			WhitelistAddresses:  core.ConvertToEthAddresses(c.WhitelistAddresses),
			AssetIDs:            c.AssetIDs,
			AvsUnbondingPeriod:  c.AvsUnbondingPeriod,
			MinSelfDelegation:   c.MinSelfDelegation,
			EpochIdentifier:     c.EpochIdentifier,
			MiniOptInOperators:  c.MiniOptInOperators,
			MinTotalStakeAmount: c.MinTotalStakeAmount,
			AvsRewardProportion: c.AvsRewardProportion,
			AvsSlashProportion:  c.AvsSlashProportion,
		}
		_, regErr := avsWriter.RegisterAVSToChain(context.Background(), params)
		if regErr != nil { return nil, fmt.Errorf("failed to register AVS with IMUA: %w", regErr) }
		
		logger.Info("AVS registered successfully. Waiting for epoch identifier to be available...")
		for i := 0; i < 5; i++ {
			time.Sleep(retryDelay)
			epochId, err = avsReader.GetAVSEpochIdentifier(&bind.CallOpts{}, c.AVSAddress)
			if err == nil && epochId != "" { break }
		}
		if epochId == "" { return nil, fmt.Errorf("could not retrieve epoch identifier after registration") }
	}
	
	return &AvsService{
		logger:                logger,
		avsWriter:             avsWriter,
		avsReader:             *avsReader,
		config:                c,
		signer:                signer,
		sender:                avsSender,
		ethClient:             ethRpcClient,
		taskResponsePeriod:    c.TaskResponsePeriod,
		taskChallengePeriod:   c.TaskChallengePeriod,
		thresholdPercentage:   c.ThresholdPercentage,
		taskStatisticalPeriod: c.TaskStatisticalPeriod,
		avsEpochIdentifier:    epochId,
	}, nil
}

func (avs *AvsService) Start(ctx context.Context, watchDir string, processedDir string, checkInterval time.Duration) error {
	avs.logger.Info("Starting HIMERA AVS Task Giver Service...", "watching", watchDir)
	ticker := time.NewTicker(checkInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			avs.logger.Info("Shutting down AVS Task Giver.")
			return nil
		case <-ticker.C:
			files, err := ioutil.ReadDir(watchDir)
			if err != nil {
				avs.logger.Error("Could not read task directory", "err", err)
				continue
			}

			if len(files) == 0 { continue }
			
			avs.logger.Info(fmt.Sprintf("Found %d new tasks to process.", len(files)))

			for _, file := range files {
				filePath := filepath.Join(watchDir, file.Name())
				avs.logger.Info("Processing task file", "file", file.Name())

				var taskDefId uint8
				var taskName string
				if strings.Contains(file.Name(), "auth") {
					taskDefId = 1; taskName = "VerifyEVMAuthorization"
				} else if strings.Contains(file.Name(), "strategy") {
					taskDefId = 2; taskName = "AuditStrategyExecution"
				} else if strings.Contains(file.Name(), "balance") {
					taskDefId = 3; taskName = "AuditAccountBalance"
				} else {
					avs.logger.Warn("Unknown task file type, skipping", "file", file.Name())
					continue
				}

				content, err := ioutil.ReadFile(filePath)
				if err != nil {
					avs.logger.Error("Could not read task file", "file", file.Name(), "err", err)
					continue
				}
				taskInputBytes, err := hexutil.Decode(strings.TrimSpace(string(content)))
				if err != nil {
					avs.logger.Error("Could not decode hex from task file", "file", file.Name(), "err", err)
					continue
				}

				if err := avs.createTaskOnChain(taskName, taskDefId, taskInputBytes); err != nil {
					avs.logger.Error("Failed to create task on-chain", "file", file.Name(), "err", err)
				} else {
					processedPath := filepath.Join(processedDir, file.Name())
					os.Rename(filePath, processedPath)
					avs.logger.Info("Task processed and file moved", "file", file.Name())
				}
			}
		}
	}
}

func (avs *AvsService) createTaskOnChain(name string, taskDefId uint8, taskInput []byte) error {
	var taskPowerTotal sdkmath.LegacyDec
	var lastErr error
    
	for attempt := 1; attempt <= maxRetries; attempt++ {
		taskPowerTotal, lastErr = avs.avsReader.GtAVSUSDValue(&bind.CallOpts{}, avs.config.AVSAddress)
		if lastErr == nil && !taskPowerTotal.IsZero() { break }
		if lastErr != nil {
			avs.logger.Error("Cannot get AVSUSDValue, retrying...", "err", lastErr, "attempt", attempt)
		} else {
			avs.logger.Info("AVS has zero stake, waiting for operators to delegate...", "attempt", attempt)
		}
		if attempt == maxRetries { return fmt.Errorf("AVS has no voting power after %d retries", maxRetries) }
		time.Sleep(retryDelay)
	}

	himeraContract, err := himera_avs.NewHimeraAvs(common.HexToAddress(avs.config.AVSAddress), avs.ethClient)
	if err != nil { return err }

	txOpts := &bind.TransactOpts{
		From:     avs.sender,
		Signer:   avs.signer,
		Context:  context.Background(),
		GasLimit: 3000000,
	}

	tx, err := himeraContract.CreateHimeraTask(txOpts, taskDefId, taskInput)
	if err != nil { return fmt.Errorf("failed to send CreateHimeraTask transaction: %w", err) }

	avs.logger.Info("Task creation transaction sent", "txHash", tx.Hash().Hex(), "taskName", name)
	return nil
}
