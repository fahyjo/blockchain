package crypto

const (
	SeedLen = 32
)

// Keys holds a node's public and private key pair
type Keys struct {
	PrivateKey *PrivateKey // PrivateKey holds a node's private ed25519 key
	PublicKey  *PublicKey  // PublicKey holds a node's public ed25519 key
}

// NewKeys creates a new Keys struct
func NewKeys(privKey *PrivateKey, pubKey *PublicKey) *Keys {
	return &Keys{
		PrivateKey: privKey,
		PublicKey:  pubKey,
	}
}
