package types

// NodeConfig represents the configuration for an AVS operator node.
type NodeConfig struct {
	// Production specifies the logger level (true for info, false for debug).
	Production bool `yaml:"production"`
	// AVSOwnerAddress is the EOA address of the AVS owner.
	AVSOwnerAddress string `yaml:"avs_owner_address"`
	// OperatorAddress is the EOA address of the operator.
	OperatorAddress string `yaml:"operator_address"`
	// AVSAddress is the address of the AVS contract.
	AVSAddress string `yaml:"avs_address"`
	// EthRpcUrl is the HTTP RPC URL of the Ethereum node.
	EthRpcUrl string `yaml:"eth_rpc_url"`
	// EthWsUrl is the WebSocket URL of the Ethereum node.
	EthWsUrl string `yaml:"eth_ws_url"`
	// BlsPrivateKeyStorePath is the path to the BLS private key store.
	BlsPrivateKeyStorePath string `yaml:"bls_private_key_store_path"`
	// OperatorEcdsaPrivateKeyStorePath is the path to the operator's ECDSA private key store.
	OperatorEcdsaPrivateKeyStorePath string `yaml:"operator_ecdsa_private_key_store_path"`
	// AVSEcdsaPrivateKeyStorePath is the path to the AVS's ECDSA private key store.
	AVSEcdsaPrivateKeyStorePath string `yaml:"avs_ecdsa_private_key_store_path"`
	// RegisterOperatorOnStartup specifies whether to automatically register the operator on startup.
	RegisterOperatorOnStartup bool `yaml:"register_operator_on_startup"`
	// NodeApiIpPortAddress is the IP and port for the Node API.
	NodeApiIpPortAddress string `yaml:"node_api_ip_port_address"`
	// EnableNodeApi specifies whether to enable the Node API.
	EnableNodeApi bool `yaml:"enable_node_api"`

	// AvsName is the name of the AVS.
	AvsName string `yaml:"avs_name"`
	// MinStakeAmount is the minimum stake amount for the AVS.
	MinStakeAmount uint64 `yaml:"min_stake_amount"`
	// AvsOwnerAddresses is a list of AVS owner addresses with access control.
	AvsOwnerAddresses []string `yaml:"avs_owner_addresses"`
	// WhitelistAddresses is a list of whitelisted addresses.
	WhitelistAddresses []string `yaml:"whitelist_addresses"`
	// AssetIDs is a list of asset IDs supported by the AVS.
	AssetIDs []string `yaml:"asset_ids"`
	// AvsUnbondingPeriod is the unbonding period for the AVS.
	AvsUnbondingPeriod uint64 `yaml:"avs_unbonding_period"`
	// MinSelfDelegation is the minimum self-delegation for the AVS.
	MinSelfDelegation uint64 `yaml:"min_self_delegation"`
	// EpochIdentifier is the epoch identifier for the AVS (e.g., "minute", "hour", "day", "week").
	EpochIdentifier string `yaml:"epoch_identifier"`
	// TaskAddress is the address of the task contract.
	TaskAddress string `yaml:"task_address"`
	// AVSRewardAddress is the AVS reward address.
	AVSRewardAddress string `yaml:"avs_reward_address"`
	// AVSSlashAddress is the AVS slash address.
	AVSSlashAddress string `yaml:"avs_slash_address"`

	// CreateTaskInterval is the interval in seconds for creating new tasks.
	CreateTaskInterval int64 `yaml:"create_task_interval"`
	// TaskResponsePeriod is the period in epochs for operators to submit phase one results.
	TaskResponsePeriod uint64 `yaml:"task_response_period"`
	// TaskChallengePeriod is the period in epochs for anyone to raise a challenge.
	TaskChallengePeriod uint64 `yaml:"task_challenge_period"`
	// ThresholdPercentage is the threshold percentage for a task.
	ThresholdPercentage uint8 `yaml:"threshold_percentage"`
	// TaskStatisticalPeriod is the period in epochs for operators to submit phase two results.
	TaskStatisticalPeriod uint64 `yaml:"task_statistical_period"`
	// MiniOptInOperators is the minimum number of opt-in operators required.
	MiniOptInOperators uint64 `yaml:"mini_opt_in_operators"`
	// MinTotalStakeAmount is the minimum total amount of stake by all operators.
	MinTotalStakeAmount uint64 `yaml:"min_total_stake_amount"`
	// AvsRewardProportion is the proportion of the reward for the AVS.
	AvsRewardProportion uint64 `yaml:"avs_reward_proportion"`
	// AvsSlashProportion is the proportion of the slash for the AVS.
	AvsSlashProportion uint64 `yaml:"avs_slash_proportion"`

	// DepositAmount is the amount to deposit.
	DepositAmount int64 `yaml:"deposit_amount"`
	// DelegateAmount is the amount to delegate.
	DelegateAmount int64 `yaml:"delegate_amount"`
	// Staker is the address of the staker.
	Staker string `yaml:"staker"`
}
