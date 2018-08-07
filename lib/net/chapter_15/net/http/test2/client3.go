package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func show_url(path string) {
	content := "i am a client"
	host := "http://localhost:8080"
	resp, err := http.Post(host+path,
		"image/jpeg",
		strings.NewReader(content))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("read data ", string(data))
}

func main() {
	show_url("/")
}
