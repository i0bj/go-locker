package symmetric

import (
	"crypto/rand"
	"fmt"
)

func encryptFile() {

}

func decryptFile() {

}

func keyGen() {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		fmt.Println(err)
	}

}
