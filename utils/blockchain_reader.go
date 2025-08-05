// utils/blockchain_reader.go
package utils

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/NethermindEth/juno/core/felt"
	"github.com/NethermindEth/starknet.go/rpc"
	starknetutils "github.com/NethermindEth/starknet.go/utils"
)

const erc20BalanceOfABI = `[{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"stateMutability":"view","type":"function"}]`

// Direcciones de contratos en Sepolia Testnet
const starknetStrkAddress = "0x04718f5a0fc34cc1af16a1cdee98ffb20c31f5cd61d6ab07201858f4287c938d"

// GetEvmBalance consulta el balance de un token ERC20 en cualquier cadena EVM.
func GetEvmBalance(rpcUrl string, tokenContractAddress common.Address, userAddress common.Address) (*big.Int, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil { return nil, fmt.Errorf("failed to connect to EVM client at %s: %w", rpcUrl, err) }
	defer client.Close()

	parsedABI, err := abi.JSON(strings.NewReader(erc20BalanceOfABI))
	if err != nil { return nil, fmt.Errorf("failed to parse ABI: %w", err) }

	calldata, err := parsedABI.Pack("balanceOf", userAddress)
	if err != nil { return nil, fmt.Errorf("failed to pack calldata: %w", err) }

	msg := ethereum.CallMsg{To: &tokenContractAddress, Data: calldata}
	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil { return nil, fmt.Errorf("failed to call token contract: %w", err) }

	balance := new(big.Int).SetBytes(result)
	return balance, nil
}

// GetStarknetStrkBalance consulta el balance de STRK en Starknet.
func GetStarknetStrkBalance(rpcUrl string, userAddress string) (*big.Int, error) {
	provider, err := rpc.NewProvider(rpcUrl)
	if err != nil { return nil, fmt.Errorf("failed to connect to Starknet client: %w", err) }

	contractAddrFelt, err := starknetutils.HexToFelt(starknetStrkAddress)
	if err != nil { return nil, err }
	userAddrFelt, err := starknetutils.HexToFelt(userAddress)
	if err != nil { return nil, err }
	
	callData := []*felt.Felt{ userAddrFelt }

	tx := rpc.FunctionCall{
		ContractAddress:    contractAddrFelt,
		EntryPointSelector: starknetutils.GetSelectorFromNameFelt("balanceOf"),
		Calldata:           callData,
	}

	result, err := provider.Call(context.Background(), tx, rpc.BlockID{Tag: "latest"})
	if err != nil { return nil, fmt.Errorf("failed to call Starknet STRK contract: %w", err) }
    
    if len(result) < 2 { return big.NewInt(0), nil }

	low := result[0]
	high := result[1]
	balance := starknetutils.FeltToBigInt(high)
	balance.Lsh(balance, 128)
	balance.Add(balance, starknetutils.FeltToBigInt(low))

	return balance, nil
}

// AuditEvmTransaction verifica una transferencia ERC20 en los logs de una transacción.
func AuditEvmTransaction(rpcUrl string, txHash common.Hash, expectedFrom, expectedTo common.Address, expectedAmount *big.Int) (bool, error) {
    client, err := ethclient.Dial(rpcUrl)
	if err != nil { return false, err }
	defer client.Close()

    receipt, err := client.TransactionReceipt(context.Background(), txHash)
    if err != nil { return false, err }

    transferEventTopic := crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))

    for _, vLog := range receipt.Logs {
        if len(vLog.Topics) == 3 && vLog.Topics[0] == transferEventTopic {
            from := common.BytesToAddress(vLog.Topics[1].Bytes())
            to := common.BytesToAddress(vLog.Topics[2].Bytes())
            amount := new(big.Int).SetBytes(vLog.Data)

            if from == expectedFrom && to == expectedTo && amount.Cmp(expectedAmount) >= 0 {
                return true, nil
            }
        }
    }
    return false, fmt.Errorf("expected transfer event not found in transaction logs")
}
