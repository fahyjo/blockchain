package crypto

import "crypto/ed25519"

type Signature struct {
	sig []byte
}

func NewSignature(sig []byte) *Signature {
	return &Signature{
		sig: sig,
	}
}

func (s *Signature) Verify(pubKey *PublicKey, msg []byte) bool {
	return ed25519.Verify(pubKey.key, msg, s.sig)
}

func (s *Signature) Bytes() []byte {
	return s.sig
}
