package symmetric

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type AESCIPHER struct {
	bs  []byte
	key []byte
}

// Function to create 32 byte hash.
func newHash(key []byte) []byte {
	key := []byte("dfkjasfajsfhdjhadiawe110")
	hasher := sha256.New()
	hasher.Write(key)
	fmt.Println(hex.EncodeToString(hasher.Sum(key)))
}
