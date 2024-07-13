package crypto

import "crypto/ed25519"

// Signature holds a digital signature
type Signature struct {
	sig []byte // sig is the digital signature created with a node's ed25519 private key
}

// NewSignature creates a new Signature struct
func NewSignature(sig []byte) *Signature {
	return &Signature{
		sig: sig,
	}
}

// Verify verifies that the ed25519 private key corresponding to the given ed25519 public key was used to sign the given message to produce the given signature
func (s *Signature) Verify(pubKey *PublicKey, msg []byte) bool {
	return ed25519.Verify(pubKey.key, msg, s.sig)
}

// Bytes returns the digital signature
func (s *Signature) Bytes() []byte {
	return s.sig
}
