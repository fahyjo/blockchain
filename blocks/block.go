package blocks

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"

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

func EncodeBlock(b *Block) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(b)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func DecodeBlock(b []byte) (*Block, error) {
	var block Block
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err := dec.Decode(&block)
	if err != nil {
		return nil, err
	}
	return &block, nil
}
