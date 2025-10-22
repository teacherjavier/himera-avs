package core

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/cosmos/btcutil/bech32"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"gopkg.in/yaml.v3"
	"math/big"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var (
	taskResponseType, _ = abi.NewType("tuple", "struct", []abi.ArgumentMarshaling{
		{Name: "TaskID", Type: "uint64"},
		{Name: "NumberSquared", Type: "uint64"},
	})

	// Args defines the ABI arguments for a TaskResponse.
	Args = abi.Arguments{
		{Type: taskResponseType, Name: "TaskResponse"},
	}
)

// TaskResponse represents the response to a task.
type TaskResponse struct {
	TaskID        uint64
	NumberSquared uint64
}

// BLSMessageToSign is the format string for the message signed by the BLS key.
const BLSMessageToSign = "BLS12-381 Signed Message\nChainIDWithoutRevision: %s\nAccAddressBech32: %s"

// AbiEncode encodes a TaskResponse into ABI format.
//
// Parameters:
//   - resp: The TaskResponse to encode.
//
// Returns:
//   - A byte slice representing the ABI-encoded TaskResponse.
//   - An error if the encoding fails.
func AbiEncode(resp TaskResponse) ([]byte, error) {

	abiArgs := abi.Arguments{
		{Name: "taskID", Type: mustNewType("uint64")},
		{Name: "numberSquared", Type: mustNewType("uint64")},
	}

	taskID := new(big.Int).SetUint64(resp.TaskID)
	numSquared := new(big.Int).SetUint64(resp.NumberSquared)

	return abiArgs.Pack(taskID, numSquared)
}

// AbiDecode decodes ABI-encoded data into a TaskResponse.
//
// Parameters:
//   - data: The ABI-encoded data.
//
// Returns:
//   - The decoded TaskResponse.
//   - An error if the decoding fails.
func AbiDecode(data []byte) (TaskResponse, error) {
	abiArgs := abi.Arguments{
		{Name: "taskID", Type: mustNewType("uint64")},
		{Name: "numberSquared", Type: mustNewType("uint64")},
	}

	values, err := abiArgs.UnpackValues(data)
	if err != nil {
		return TaskResponse{}, err
	}

	var (
		taskID, _    = values[0].(*big.Int)
		numSquare, _ = values[1].(*big.Int)
	)

	return TaskResponse{
		TaskID:        taskID.Uint64(),
		NumberSquared: numSquare.Uint64(),
	}, nil
}

// mustNewType creates a new ABI type and panics on error.
// This is a helper function to avoid error handling for ABI type creation.
//
// Parameters:
//   - typeStr: The string representation of the ABI type.
//
// Returns:
//   - The ABI type.
func mustNewType(typeStr string) abi.Type {
	t, err := abi.NewType(typeStr, "", nil)
	if err != nil {
		panic(err)
	}
	return t
}

// MarshalTaskResponse marshals a TaskResponse struct into JSON bytes.
//
// Parameters:
//   - h: The TaskResponse to marshal.
//
// Returns:
//   - The JSON-encoded TaskResponse.
//   - An error if the marshaling fails.
func MarshalTaskResponse(h TaskResponse) ([]byte, error) {
	return json.Marshal(h)
}

// UnmarshalTaskResponse unmarshals JSON bytes into a TaskResponse struct.
//
// Parameters:
//   - jsonData: The JSON-encoded data.
//
// Returns:
//   - The unmarshaled TaskResponse.
//   - An error if the unmarshaling fails.
func UnmarshalTaskResponse(jsonData []byte) (TaskResponse, error) {
	var taskResponse TaskResponse
	err := json.Unmarshal(jsonData, &taskResponse)
	return taskResponse, err
}

// GetTaskResponseDigestEncodeByjson returns the Keccak256 hash of a JSON-encoded TaskResponse.
// This is the digest that operators sign.
//
// Parameters:
//   - h: The TaskResponse to hash.
//
// Returns:
//   - The 32-byte Keccak256 hash of the TaskResponse.
//   - The JSON-encoded TaskResponse.
//   - An error if the marshaling fails.
func GetTaskResponseDigestEncodeByjson(h TaskResponse) ([32]byte, []byte, error) {
	jsonData, err := MarshalTaskResponse(h)
	if err != nil {
		return [32]byte{}, []byte{}, err
	}
	taskResponseDigest := crypto.Keccak256Hash(jsonData)
	return taskResponseDigest, jsonData, nil
}

// GetTaskResponseDigestEncodeByAbi returns the Keccak256 hash of an ABI-encoded TaskResponse.
//
// Parameters:
//   - h: The TaskResponse to hash.
//
// Returns:
//   - The 32-byte Keccak256 hash of the TaskResponse.
//   - The ABI-encoded TaskResponse.
//   - An error if the encoding fails.
func GetTaskResponseDigestEncodeByAbi(h TaskResponse) ([32]byte, []byte, error) {
	packed, err := Args.Pack(h)
	if err != nil {
		return [32]byte{}, []byte{}, err
	}
	taskResponseDigest := crypto.Keccak256Hash(packed)
	return taskResponseDigest, packed, nil
}

// UpdateYAMLWithComments updates a YAML file with a new value for a given key, preserving comments.
//
// Parameters:
//   - filePath: The path to the YAML file.
//   - key: The key to update.
//   - newValue: The new value for the key.
//
// Returns:
//   - An error if the file cannot be read, parsed, or written.
func UpdateYAMLWithComments(filePath, key, newValue string) error {
	if newValue == "" {
		panic("param is nil")
	}
	// Read the original YAML file content
	data, err := os.ReadFile(filepath.Clean(filePath))
	if err != nil {
		return err
	}
	// Parse YAML using yaml.v3 node parser to preserve comments
	var doc yaml.Node
	err = yaml.Unmarshal(data, &doc)
	if err != nil {
		return err
	}
	// Iterate through YAML content to find and update the specified key
	for i := 0; i < len(doc.Content[0].Content); i += 2 {
		if doc.Content[0].Content[i].Value == key {
			doc.Content[0].Content[i+1].Kind = yaml.ScalarNode
			doc.Content[0].Content[i+1].Value = newValue
			doc.Content[0].Content[i+1].Tag = "tag:yaml.org,2002:str"
			break
		}
	}

	var buf bytes.Buffer
	encoder := yaml.NewEncoder(&buf)
	encoder.SetIndent(2)
	err = encoder.Encode(&doc)
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, buf.Bytes(), 0644)
}

// GetFileInCurrentDirectory returns the full path of a file in the current directory.
//
// Parameters:
//   - filename: The name of the file.
//
// Returns:
//   - The full path to the file.
//   - An error if the file is not found or the current directory cannot be determined.
func GetFileInCurrentDirectory(filename string) (string, error) {
	// Get current working directory
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Construct full file path
	fullPath := filepath.Join(currentDir, filename)

	// Check if file exists
	_, err = os.Stat(fullPath)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("file %s not found in current directory", filename)
	}

	return fullPath, nil
}

// ConvertToEthAddresses converts a slice of hexadecimal string addresses to a slice of common.Address.
//
// Parameters:
//   - strArray: A slice of hexadecimal string addresses.
//
// Returns:
//   - A slice of common.Address.
func ConvertToEthAddresses(strArray []string) []common.Address {
	var ethAddresses []common.Address

	if len(strArray) > 0 {
		for _, str := range strArray {
			address := common.HexToAddress(str)
			ethAddresses = append(ethAddresses, address)
		}
	}

	return ethAddresses
}

// SwitchEthAddressToImAddress converts an Ethereum address to an Imua address.
//
// Parameters:
//   - ethAddress: The hexadecimal string representation of the Ethereum address.
//
// Returns:
//   - The bech32-encoded Imua address.
//   - An error if the address cannot be converted.
func SwitchEthAddressToImAddress(ethAddress string) (string, error) {
	b, err := hex.DecodeString(ethAddress[2:])
	if err != nil {
		return "", fmt.Errorf("failed to decode eth address: %w", err)
	}

	// Generate im address
	bech32Prefix := "im"
	imAddress, err := bech32.EncodeFromBase256(bech32Prefix, b)
	if err != nil {
		return "", fmt.Errorf("failed to encode bech32 address: %w", err)
	}

	return imAddress, nil
}

// ChainIDWithoutRevision returns the chainID without the revision number.
// For example, "imuachaintestnet_233-1" returns "imuachaintestnet_233".
//
// Parameters:
//   - chainID: The chain ID string.
//
// Returns:
//   - The chain ID without the revision number.
func ChainIDWithoutRevision(chainID string) string {
	if !IsRevisionFormat(chainID) {
		return chainID
	}
	splitStr := strings.Split(chainID, "-")
	return splitStr[0]
}

// IsRevisionFormat checks if a string is in the revision format (e.g., "name-1").
var IsRevisionFormat = regexp.MustCompile(`^.*[^\n-]-{1}[1-9][0-9]*$`).MatchString
