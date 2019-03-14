package main

import (
	"fmt"
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

	// fmt.Println(res.Status)
	printHeader(res.Header)

	bs := make([]byte, 99999)

	res.Body.Read(bs)
	fmt.Println(string(bs))

}

func printHeader(m http.Header) {
	for key, headerValues := range m {
		fmt.Println(key)

		for _, value := range headerValues {
			fmt.Println("\t" + value)
		}
	}
}
