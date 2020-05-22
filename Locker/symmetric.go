package symmetric

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// Generates ranndom key used to encrypt each file
func keyGen() *[]byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println(err)
	}
	return &key
}

// Encrypt func encrypts files using our random key and file.
// the result will be an AES cipher block
func encrypt(data []byte, key []byte) (*[]byte, error) {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// Galois Counter Mode being applied to AES cipher block.
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}
	// Number used once
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return &cipherText, nil

}

func encryptFile(filepath string)
