package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return resp.Header.Get("Content-Type"), nil
}

func main() {
	cType, err := contentType("https://golang.org")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cType)
}
