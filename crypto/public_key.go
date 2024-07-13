package crypto

import (
	"crypto/ed25519"
	"crypto/sha256"
)

// PublicKey holds a node's ed25519 public key
type PublicKey struct {
	key ed25519.PublicKey // a node's ed25519 public key
}

// NewPublicKey creates a new PublicKey struct
func NewPublicKey(key ed25519.PublicKey) *PublicKey {
	return &PublicKey{
		key: key,
	}
}

// Hash hashes a node's ed25519 public key
func (p *PublicKey) Hash() []byte {
	hash := sha256.Sum256(p.key)
	return hash[:]
}

// Bytes returns a node's ed25519 public key
func (p *PublicKey) Bytes() []byte {
	return p.key
}
