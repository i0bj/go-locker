package locker

import (
	"crypto/rand"
	"crypto/rsa"
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

// RsaCrypt will be used to encrypt AES keys
func RsaCrypt() {

	// Generates private/public key pair.
	PrivKey, err := rsa.GenerateKey(rand.Reader, 2048)


}
