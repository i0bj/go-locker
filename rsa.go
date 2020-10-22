package main

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

// Bit size of key used to generate key pair.
const (
	keySize = 2048
)

// Writes private key to file which will be placed in current
// directory.
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

// Writes public key to file which will be placed in current
// directory.
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
