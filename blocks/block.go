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

func (block *Block) CalculateMerkleRoot() ([]byte, error) {
	var currLevel [][]byte
	for _, tx := range block.Transactions {
		hash, err := tx.Hash()
		if err != nil {
			return nil, err
		}
		currLevel = append(currLevel, hash)
	}

	if len(currLevel)%2 != 0 {
		currLevel = append(currLevel, currLevel[len(currLevel)-1])
	}

	for len(currLevel) > 1 {
		var nextLevel [][]byte
		for i := 0; i < len(currLevel); i += 2 {
			newHash := sha256.Sum256(append(currLevel[i], currLevel[i+1]...))
			nextLevel = append(nextLevel, newHash[:])
		}
		currLevel = nextLevel
	}
	return currLevel[0], nil
}
