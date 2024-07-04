package main

import (
	"encoding/hex"
	"fmt"

	"github.com/fahyjo/blockchain/crypto"
)

func main() {
	for i := 0; i < 15; i++ {
		privKey, err := crypto.NewPrivateKey()
		if err != nil {
			panic(err)
		}
		pubKey := privKey.PublicKey()

		fmt.Printf("Public key: %s\n", hex.EncodeToString(pubKey.Bytes()))
		fmt.Printf("Private key: %s\n", hex.EncodeToString(privKey.Bytes()))
	}
}
