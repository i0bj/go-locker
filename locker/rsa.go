package locker

import (
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

func RsaCrypt() {

}
