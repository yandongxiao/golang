package main

import (
	"fmt"
	"net/http"
	"strings"
)

func main() {
	json := `{"content":"hello,world"}`
	b := strings.NewReader(json)

	http.Post("http://localhost:8888/hello", "image/jpeg", b)
	fmt.Println("post ok")
}
