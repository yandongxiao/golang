package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func show_url(path string) {
	resp, err := http.Get("http://localhost:8080" + path)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
}

func main() {
	show_url("/show_url")
	show_url("/洞箫")
	show_url("/show_header")
}
