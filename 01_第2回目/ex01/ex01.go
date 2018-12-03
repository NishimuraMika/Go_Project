package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// Open a file.
	f, err := os.Open(`input.txt`)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Pass file descriptor returned by os.Open to ioutil.ReadAll.
	result, _ := ioutil.ReadAll(f)
	// Convert byte slice to string.
	fmt.Println(string(result))
}
