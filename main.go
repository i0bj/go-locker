package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// Enumerate all directories in ROOT
	files, _ := ioutil.ReadDir("/")

	for _, f := range files {
		fmt.Println(f.Name())
	}

}
