package crypto

import (
	"crypto/ed25519"
	"crypto/sha256"
)

type PublicKey struct {
	key ed25519.PublicKey
}

func NewPublicKey(key ed25519.PublicKey) *PublicKey {
	return &PublicKey{
		key: key,
	}
}

func (p *PublicKey) Hash() []byte {
	hash := sha256.Sum256(p.key)
	return hash[:]
}
