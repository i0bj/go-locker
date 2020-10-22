package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

const (
	keySize = 2048
)

func writePrivKey(priv *rsa.PrivateKey) {

	privFile := "key.priv"
	f, err := os.Create(privFile)
	if err != nil {
		log.Fatalf("Could not write to file: %v", err)
	}
	defer f.Close()
	w := gob.NewEncoder(f)
	w.Encode(priv)

}

func writePubKey(priv *rsa.PrivateKey) {
	pubFile := "key.pub"
	f, err := os.Create(pubFile)
	if err != nil {
		log.Fatalf("Could not write to file: %v", err)
	}
	defer f.Close()
	w := gob.NewEncoder(f)
	w.Encode(priv.Public())
}

// RsaGen will be used to encrypt AES keys
func main() {

	// Generates private/public key pair.
	rnd := rand.Reader
	Priv, err := rsa.GenerateKey(rnd, keySize)
	if err != nil {
		log.Println(err)
	}

	// The following functions will write
	// the private/public key to their
	// respective files.
	writePrivKey(Priv)
	writePubKey(Priv)

	fmt.Println("[+] Private/Public key pair written to file.")

}
