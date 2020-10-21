package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"log"
)

const (
	keySize = 2048
)

// RsaGen will be used to encrypt AES keys
func RsaGen() crypto.PublicKey {

	// Generates private/public key pair.
	rnd := rand.Reader
	Priv, err := rsa.GenerateKey(rnd, keySize)
	if err != nil {
		log.Println(err)
	}

	return Priv.Public()

}

func main() {
	fmt.Println(RsaGen())
}
