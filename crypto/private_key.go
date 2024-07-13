package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"
)

// PrivateKey holds a node's ed25519 private key
type PrivateKey struct {
	key ed25519.PrivateKey // key a node's ed25519 private key
}

// NewPrivateKey creates a new PrivateKey struct
func NewPrivateKey(key ed25519.PrivateKey) *PrivateKey {
	return &PrivateKey{key: key}
}

// GenerateNewPrivateKey generates a new private key from a random seed, returns a new PrivateKey struct
func GenerateNewPrivateKey() (*PrivateKey, error) {
	seed := make([]byte, SeedLen)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		return nil, err
	}
	return &PrivateKey{
		key: ed25519.NewKeyFromSeed(seed),
	}, nil
}

// PublicKey creates a PublicKey struct containing the ed25519 public key corresponding to a node's ed25519 private key
func (p *PrivateKey) PublicKey() *PublicKey {
	return NewPublicKey(p.key.Public().(ed25519.PublicKey))
}

// Sign signs the given message with a node's ed25519 private key
func (p *PrivateKey) Sign(msg []byte) *Signature {
	return &Signature{
		sig: ed25519.Sign(p.key, msg),
	}
}

// Bytes returns a node's ed25519 private key
func (p *PrivateKey) Bytes() []byte {
	return p.key
}
