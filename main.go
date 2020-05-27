package main

import (
	"os"
	"runtime"
)

func main() {
	var files []string
	var home string

	// Logic to determine if host is running Windows
	if runtime.GOOS == "windows" {
		home = os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
	}

}
