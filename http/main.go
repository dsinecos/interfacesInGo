package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	// Making a HTTP request using Golang

	res, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, res.Body)
}
