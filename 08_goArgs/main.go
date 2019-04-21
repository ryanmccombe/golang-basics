package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	// go run main.go file.txt
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	_, _ = io.Copy(os.Stdout, f)
}

// Alternative implementation
func main2() {
	// go run main.go file.txt
	bs, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(string(bs))
}