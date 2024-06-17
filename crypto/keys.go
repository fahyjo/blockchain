package crypto

const (
	SeedLen = 32
)

type Keys struct {
	PrivateKey *PrivateKey
	PublicKey  *PublicKey
}

func NewKeys(privKey *PrivateKey, pubKey *PublicKey) *Keys {
	return &Keys{
		PrivateKey: privKey,
		PublicKey:  pubKey,
	}
}
