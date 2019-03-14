package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	// Making a HTTP request using Golang

	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	lw := logWriter{}

	// io.Copy(os.Stdout, res.Body)
	io.Copy(lw, res.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))

	return len(bs), nil
}
