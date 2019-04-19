package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Sends stuff to something that implements the writer interface
	// from something that implements the reader interface
	// _, _ = io.Copy(os.Stdout, resp.Body)

	customWriter := logWriter {}
	_, _ = io.Copy(customWriter, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}