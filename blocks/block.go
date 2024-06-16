package blocks

import (
	"crypto/sha256"

	"github.com/fahyjo/blockchain/crypto"
	"github.com/fahyjo/blockchain/transactions"
)

type Block struct {
	Header       *Header
	Transactions []*transactions.Transaction
	Sig          *crypto.Signature
	PubKey       *crypto.PublicKey
}

func NewBlock(header *Header, txs []*transactions.Transaction, sig *crypto.Signature, pubKey *crypto.PublicKey) *Block {
	return &Block{
		Header:       header,
		Transactions: txs,
		Sig:          sig,
		PubKey:       pubKey,
	}
}

func (block *Block) Hash() ([]byte, error) {
	header := block.Header
	b, err := EncodeHeader(header)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(b)
	return hash[:], nil
}
