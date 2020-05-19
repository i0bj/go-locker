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

/* encrypt func encrypts files using our random key and file.
the result will be an AES cipher block */
func encrypt(data []byte, key []byte) *[]byte {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	cipherText := gcm.Seal(nonce, nonce, data)
	return &cipherText

}

func encryptFile(filepath string)
