package blocks

import (
	"crypto/sha256"

	"github.com/fahyjo/blockchain/crypto"
	"github.com/fahyjo/blockchain/transactions"
)

// Block represents a block on the blockchain
type Block struct {
	Header       *Header                     // Header contains metadata about the Block
	Transactions []*transactions.Transaction // Transactions contains all the transactions to be committed with this Block
	Sig          *crypto.Signature           // Sig contains the digital signature of this Block created with the block creator's ed25519 private key
	PubKey       *crypto.PublicKey           // PubKey contains the block creator's ed25519 public key
}

// NewBlock creates a new Block struct
func NewBlock(header *Header, txs []*transactions.Transaction, sig *crypto.Signature, pubKey *crypto.PublicKey) *Block {
	return &Block{
		Header:       header,
		Transactions: txs,
		Sig:          sig,
		PubKey:       pubKey,
	}
}

// Hash hashes the given Block
// To hash a Block we hash the Header of the Block
func (block *Block) Hash() ([]byte, error) {
	header := block.Header
	b, err := EncodeHeader(header)
	if err != nil {
		return nil, err
	}
	hash := sha256.Sum256(b)
	return hash[:], nil
}

// CalculateMerkleRoot calculates the root hash of the merkle tree made of the given Block's transactions
func (block *Block) CalculateMerkleRoot() ([]byte, error) {
	// hash all transactions
	var currLevel [][]byte
	for _, tx := range block.Transactions {
		hash, err := tx.Hash()
		if err != nil {
			return nil, err
		}
		currLevel = append(currLevel, hash)
	}

	// if odd number of transactions, append the hash of the last transaction again
	if len(currLevel)%2 != 0 {
		currLevel = append(currLevel, currLevel[len(currLevel)-1])
	}

	// append and hash consecutive hashes
	// each round halves the number of hashes
	// return the only remaining hash
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
