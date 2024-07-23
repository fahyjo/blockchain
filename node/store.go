package node

import (
	"github.com/fahyjo/blockchain/blocks"
	"github.com/fahyjo/blockchain/transactions"
	"github.com/fahyjo/blockchain/utxos"
)

// Store holds a node's BlockStore, TransactionStore, and UtxoStore
type Store struct {
	BlockStore       blocks.BlockStore             // BlockStore holds every block committed to the blockchain
	TransactionStore transactions.TransactionStore // TransactionStore holds every transaction committed to the blockchain
	UtxoStore        utxos.UTXOStore               // UtxoStore holds all utxos
}

// NewStore creates a new Store struct
func NewStore(blockStore blocks.BlockStore, transactionStore transactions.TransactionStore, utxoStore utxos.UTXOStore) *Store {
	return &Store{
		BlockStore:       blockStore,
		TransactionStore: transactionStore,
		UtxoStore:        utxoStore,
	}
}
