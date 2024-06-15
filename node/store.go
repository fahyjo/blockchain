package node

import (
	"github.com/fahyjo/blockchain/blocks"
	"github.com/fahyjo/blockchain/transactions"
	"github.com/fahyjo/blockchain/utxos"
)

type Store struct {
	BlockStore       blocks.BlockStore
	TransactionStore transactions.TransactionStore
	UtxoStore        utxos.UTXOStore
}

func NewStore(blockStore blocks.BlockStore, transactionStore transactions.TransactionStore, utxoStore utxos.UTXOStore) *Store {
	return &Store{
		BlockStore:       blockStore,
		TransactionStore: transactionStore,
		UtxoStore:        utxoStore,
	}
}
