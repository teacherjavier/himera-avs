# Hello World AVS

Welcome to the Hello World AVS! This project provides a simple, concrete example of an Actively Validated Service (AVS), demonstrating the fundamental components and their interactions. It's an excellent starting point for developers looking to understand and build their own AVSs.

## Table of Contents

- [Overview](#overview)
- [Workflow](#workflow)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
  - [Quick Start](#quick-start)
- [Private Key Management](#private-key-management)
  - [Overview](#overview-1)
  - [Prerequisites](#prerequisites-1)
  - [Steps](#steps)
  - [Key Paths](#key-paths)
  - [Key Management Best Practices](#key-management-best-practices)
  - [Verification](#verification)
  - [Important Password Configuration](#important-password-configuration)
  - [Secure Environment Variable Management](#secure-environment-variable-management)
- [Configuration](#configuration)
- [Contributing](#contributing)
- [License](#license)

## Overview

The Hello World AVS is designed to showcase the simplest functionality of an AVS. It provides a clear and concise example of how an AVS consumer can request a "Hello World" message to be generated and signed by a network of operators. This project is ideal for developers who want to learn the basics of AVS development in a hands-on manner.

## Workflow

1.  **Request:** An AVS consumer requests a "Hello World" message to be generated and signed.
2.  **Event:** The `HelloWorld` contract receives the request and emits a `NewTaskCreated` event.
3.  **Task Execution:** Registered and staked operators detect the event, generate the requested message, hash it, and sign the hash with their private keys.
4.  **Submission:** Each operator submits their signed hash back to the `HelloWorld` AVS contract.
5.  **Verification:** If the operator is registered and has the minimum required stake, the submission is accepted.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.18 or later)
- [Make](https://www.gnu.org/software/make/)

### Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/imua-xyz/imua-avs.git
    cd imua-avs
    ```
2.  Build the project:
    ```bash
    make build
    ```
    This will create the necessary binaries, including `imua-key`, `avs`, and `operator`.

### Quick Start

1.  **Import Keys:**
    ```bash
    ./imua-key import --key-type ecdsa {pri_key}
    ```
2.  **Run the AVS:**
    ```bash
    ./avs/main --config config.yaml
    ```
3.  **Run the Operator:**
    ```bash
    ./operator/main --config config.yaml
    ```

## Private Key Management

### Overview

This section provides a comprehensive guide to generating and managing private keys for the AVS system using the `imua-key` utility.

### Prerequisites

- Ensure `imua-key` is installed (use the `make imua-key` command).
- Have private keys ready for import.

### Steps

1.  **Import AVS ECDSA Private Key:**
    ```bash
    ./imua-key importKey --key-type ecdsa --private-key {avs_private_key} --output-dir tests/keys/avs.ecdsa.key.json
    ```
2.  **Import Operator ECDSA Private Key:**
    ```bash
    ./imua-key importKey --key-type ecdsa --private-key {operator_private_key} --output-dir tests/keys/operator.ecdsa.key.json
    ```
3.  **Generate or Import BLS Private Key for the operator:**
    -   **Import:**
        ```bash
        ./imua-key importKey --key-type bls --private-key {bls_private_key} --output-dir tests/keys/test.bls.key.json
        ```
    -   **Generate:**
        ```bash
        ./imua-key generate --key-type bls --num-keys 1
        ```

### Key Paths

-   **AVS Owner ECDSA Private Key:** `tests/keys/avs.ecdsa.key.json`
-   **Operator ECDSA Private Key:** `tests/keys/operator.ecdsa.key.json`
-   **BLS Private Key:** `tests/keys/test.bls.key.json`

### Key Management Best Practices

-   Keep private keys secure.
-   Never share private keys.
-   Use environment-specific key management.
-   Rotate keys periodically.

### Verification

```bash
# Verify key files exist
ls tests/keys/
```

### Important Password Configuration

If a non-empty password is used when writing private keys to JSON files, you must set the corresponding environment variables:

| Key Type           | Environment Variable          | JSON File Path                       |
| ------------------ | ----------------------------- | ------------------------------------ |
| BLS Key            | `OPERATOR_BLS_KEY_PASSWORD`     | `tests/keys/test.bls.key.json`         |
| Operator ECDSA Key | `OPERATOR_ECDSA_KEY_PASSWORD` | `tests/keys/operator.ecdsa.key.json` |
| AVS ECDSA Key      | `AVS_ECDSA_KEY_PASSWORD`      | `tests/keys/avs.ecdsa.key.json`        |

### Secure Environment Variable Management

-   **Temporary Terminal Session:**
    ```bash
    export OPERATOR_BLS_KEY_PASSWORD="strong_bls_password"
    export OPERATOR_ECDSA_KEY_PASSWORD="strong_ecdsa_password"
    export AVS_ECDSA_KEY_PASSWORD="strong_avs_password"
    ```
-   **Persistent Zsh Configuration:**
    ```bash
    # Edit ~/.zshrc configuration
    nano ~/.zshrc

    # Add secure passwords
    export OPERATOR_BLS_KEY_PASSWORD="strong_bls_password"
    export OPERATOR_ECDSA_KEY_PASSWORD="strong_ecdsa_password"
    export AVS_ECDSA_KEY_PASSWORD="strong_avs_password"

    # Reload configuration
    source ~/.zshrc
    ```

## Configuration

The `config.yaml` file contains the configuration for the AVS and operator.

### General Configuration

-   `production`: Sets the logger level (`true` for info, `false` for debug).
-   `operator_address`: The EOA address of the operator, used for sending transactions.
-   `avs_owner_address`: The EOA address of the AVS owner, used for sending transactions.
-   `avs_address`: The address of the deployed AVS contract.

### Ethereum Node Configuration

-   `eth_rpc_url`: The HTTP RPC URL of the Ethereum node.
-   `eth_ws_url`: The WebSocket URL of the Ethereum node.

### Key Configuration

-   `avs_ecdsa_private_key_store_path`: Path to the AVS ECDSA private key store.
-   `operator_ecdsa_private_key_store_path`: Path to the operator ECDSA private key store.
-   `bls_private_key_store_path`: Path to the BLS private key store.

### API Configuration

-   `node_api_ip_port_address`: The IP and port for the Node API.
-   `enable_node_api`: Enable or disable the Node API.
-   `register_operator_on_startup`: Automatically register the operator on startup.

### AVS Registration Parameters

-   `avs_name`: The name of the AVS.
-   `min_stake_amount`: The minimum stake amount for the AVS.
-   `avs_owner_addresses`: A list of AVS owner addresses with access control.
-   `asset_ids`: A list of asset IDs supported by the AVS.
-   `avs_unbonding_period`: The unbonding period for the AVS.
-   `min_self_delegation`: The minimum self-delegation for the AVS.
-   `epoch_identifier`: The epoch identifier for the AVS (`minute`, `hour`, `day`, `week`).
-   `avs_reward_address`: The AVS reward address.
-   `avs_slash_address`: The AVS slash address.
-   `task_address`: The task address.
-   `params`: A list of parameters including the minimal number of opt-in operators, minimal total stake amount, reward percentage, and slash percentage.

### Task Creation Parameters

-   `create_task_interval`: The interval in seconds for creating new tasks.
-   `task_response_period`: The period in epochs for operators to submit phase one results.
-   `task_challenge_period`: The period in epochs for anyone to raise a challenge.
-   `threshold_percentage`: The threshold percentage for a task.
-   `task_statistical_period`: The period in epochs for operators to submit phase two results.

## Contributing

Contributions are welcome! Please feel free to submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
