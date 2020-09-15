package locker

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
)

var (
	err       error
	PrivKey   *rsa.PrivateKey
	PubKey    *rsa.PublicKey
	Msg       []byte
	PlainTxt  []byte
	CipherTxt []byte
	Signature []byte
	Label     []byte
)

func RsaCrypt() {

	// Generates privaate/public key pair.
	PrivKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatal(err)
	}
	// Holds the public key out of the newly created public/private key pair.
	PubKey := &PrivKey.PublicKey

}
