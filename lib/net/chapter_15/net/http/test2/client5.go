package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func show_url(path string) {
	host := "http://localhost:8080"
	form := map[string][]string{"@file": []string{"/tmp/data"}, "pwd": []string{"w", "d"}}
	resp, err := http.PostForm(host+path, form)
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
	show_url("/show_url?1=a&2=b")
}
