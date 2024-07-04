package crypto

import (
	"crypto/ed25519"
	"crypto/rand"
	"io"
)

type PrivateKey struct {
	key ed25519.PrivateKey
}

func NewPrivateKey(key ed25519.PrivateKey) *PrivateKey {
	return &PrivateKey{key: key}
}

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

func (p *PrivateKey) PublicKey() *PublicKey {
	return NewPublicKey(p.key.Public().(ed25519.PublicKey))
}

func (p *PrivateKey) Sign(msg []byte) *Signature {
	return &Signature{
		sig: ed25519.Sign(p.key, msg),
	}
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}
