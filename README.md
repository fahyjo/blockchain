# Blockchain

A Go implementation of a distributed blockchain network with a Tendermint-based consensus protocol.

## Overview

This project implements a peer-to-peer blockchain network with the following key features:

- **Tendermint Consensus**: Uses a three-phase commit process (proposal, pre-vote, pre-commit) to achieve Byzantine Fault Tolerance
- **UTXO Model**: Transactions use an Unspent Transaction Output (UTXO) model similar to Bitcoin
- **ED25519 Cryptography**: Secure digital signatures for transaction validation and block creation
- **Peer Discovery**: Automatic peer discovery and network formation
- **Merkle Trees**: Efficient validation of transaction inclusion
- **Distributed Validators**: Multiple validators work together to create and validate blocks

## Architecture

The project is organized into several packages:

- `blocks`: Contains structures and logic for blockchain blocks
- `consensus`: Implements the Tendermint consensus protocol
- `crypto`: Handles cryptographic operations (keys, signatures)
- `node`: Core node implementation with networking and block processing
- `peers`: Manages peer connections and communication
- `transactions`: Defines transaction structures and validation
- `utxos`: Manages the UTXO set (unspent transaction outputs)
- `proto`: Protocol buffer definitions for network communication

## Consensus Protocol

This implementation uses a Tendermint-inspired Byzantine Fault Tolerant (BFT) consensus protocol that operates in distinct rounds, with each round progressing through three phases:

### Round Structure
- Each round has a designated **proposer** selected in a deterministic round-robin fashion from the validator set
- The validator set is predefined in the configuration and consists of nodes with special privileges
- Non-validator nodes participate in the network but don't vote in the consensus process
- A successful round results in exactly one block being committed to the blockchain

### Consensus Phases

1. **Proposal Phase**:
   - The designated proposer for the current round creates a new block containing:
     - Transactions from the mempool
     - Reference to the previous block hash
     - Merkle root of included transactions
     - Block height and timestamp
   - The proposer signs the block with their private key and broadcasts it to the network
   - Other validators verify the proposer's identity and block validity before proceeding

2. **Pre-Vote Phase**:
   - Upon receiving a valid block proposal, validators:
     - Verify the proposer is the designated proposer for the current round
     - Validate all transactions in the block
     - Check the block's Merkle root and previous block hash
     - Verify the block signature
   - If the block passes all checks, validators broadcast a Pre-Vote message containing:
     - The block ID being voted on
     - The validator's signature on the block ID
     - The validator's public key
   - The network progresses to the next phase when 2/3+ of validators have submitted valid Pre-Votes for the block

3. **Pre-Commit Phase**:
   - Once a validator sees Pre-Votes from 2/3+ of validators, it broadcasts a Pre-Commit message with:
     - The block ID being committed to
     - The validator's signature on the block ID
     - The validator's public key
   - The network finalizes the block when 2/3+ of validators have submitted valid Pre-Commits

4. **Commit and State Transition**:
   - When a node receives Pre-Commits from 2/3+ of validators, it:
     - Permanently adds the block to its blockchain
     - Updates its UTXO set by processing all transactions in the block
     - Moves to the next round with a new designated proposer

### Fault Tolerance
The system maintains correctness as long as less than 1/3 of validators are Byzantine (malicious or faulty). The threshold of 2/3 ensures:
- Network can reach consensus despite node failures
- No two conflicting blocks can be committed at the same height
- Once committed, blocks are final and cannot be reverted

## Transaction Model

This blockchain implements a UTXO (Unspent Transaction Output) model similar to Bitcoin, which offers strong security guarantees and enables simple transaction validation:

### Transaction Components

1. **Inputs**:
   - Each input references an existing UTXO by:
     - Transaction ID of the transaction that created the UTXO
     - Output index within that transaction
   - Each input includes an **UnlockingScript** containing:
     - The spender's public key
     - A digital signature created by signing the transaction hash with the spender's private key
   - Inputs must reference valid, unspent UTXOs in the UTXO set

2. **Outputs**:
   - Each output creates a new UTXO with:
     - An amount of tokens
     - A **LockingScript** that specifies conditions for spending this output
   - In the current implementation, locking scripts contain a public key hash
   - The sum of output amounts must not exceed the sum of input amounts

3. **Transaction Validation Rules**:
   - All referenced UTXOs must exist and be unspent
   - Each input must provide a valid unlocking script that satisfies its referenced UTXO's locking script
   - The transaction must be correctly signed by the owners of all input UTXOs
   - Total output amount must not exceed total input amount (no inflation)
   - No double-spending: A transaction cannot reference UTXOs already claimed by another transaction in the same block
   - Inputs and outputs must have valid formats and values (non-negative amounts)

### Transaction Processing

1. **Transaction Creation**:
   - Users create transactions by selecting UTXOs they control as inputs
   - They specify outputs to transfer value to recipients
   - They sign each input with the private key corresponding to the public key hash in the referenced UTXO's locking script

2. **Mempool Management**:
   - New transactions are validated and stored in the mempool
   - The mempool tracks which UTXOs are claimed by pending transactions
   - When blocks are committed, transactions in them are removed from the mempool
   - The mempool is regularly cleansed of invalid transactions (those with inputs that have been spent)

3. **State Changes on Block Commit**:
   - When a block is committed:
     - All UTXOs spent by transactions in the block are removed from the UTXO set
     - All new UTXOs created by transaction outputs are added to the UTXO set
   - This atomic state transition ensures consistency of the UTXO set

4. **Transaction ID**:
   - Each transaction has a unique ID generated by hashing the transaction data
   - For hashing, the unlocking scripts are temporarily set to nil to prevent transaction malleability

## Configuration

The system is configured through `config/config.json`, which defines:

- Network addresses and peer connections
- Public/private key pairs for nodes
- Storage locations for blockchain data
- Validator assignments and roles

## Running the Project

### Prerequisites

- Go 1.21 or higher
- Protocol Buffer compiler (protoc)

### Building

```
make build
```

### Running a Node

```
make run id=<node_id>
```

Where `<node_id>` corresponds to a configuration in `config/config.json`.

### Generating Protocol Buffers

```
make proto
```

## Storage

The implementation supports two storage backends for blockchain data persistence:

### In-Memory Storage
- Stores data in Go maps (key-value stores) held in memory
- Fast access but data is lost when the node shuts down
- Useful for testing, development, and running short-lived test networks
- Implemented for blocks, transactions, and UTXOs

### LevelDB Storage
- A fast key-value storage library by Google that provides an ordered mapping from string keys to string values
- Features and benefits:
  - **Persistent**: Data is stored on disk and survives node restarts
  - **Highly efficient**: Uses a Log-Structured Merge Tree (LSM-tree) architecture for optimized write performance
  - **Compression**: Uses Snappy compression to reduce storage requirements
  - **Atomic batch operations**: Multiple operations can be executed atomically
  - **Snapshots**: Point-in-time views of the entire store
  - **Range iterations**: Efficiently scan ranges of keys

#### Implementation Details
- The project implements separate LevelDB stores for:
  - `BlockStore`: Stores blocks indexed by block hash
  - `TransactionStore`: Stores transactions indexed by transaction hash
  - `UTXOStore`: Stores UTXOs indexed by UTXO ID (transaction hash + output index)
- Each node's databases are stored in configured directories under the `/db/` folder
- Data is encoded using Go's encoding/gob package before being stored in LevelDB
- Interfaces like `BlockStore`, `TransactionStore`, and `UTXOStore` abstract the storage implementation, making it easy to switch between in-memory and LevelDB backends

#### Operational Considerations
- LevelDB is optimized for sequential write operations but can handle random reads efficiently
- The blockchain design benefits from this pattern as new blocks are written sequentially
- Each node maintains its own separate LevelDB databases to ensure independence and fault isolation
- Database paths are configured per node in the `config.json` file, allowing multiple nodes to run on the same machine

## Future Work

- RESTful API for blockchain interaction
- Enhanced mempool management
- Light client implementation
- Smart contract support
- Sharding for scalability

## License

This project is a work in progress.
