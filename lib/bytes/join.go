package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := []byte("hello")
	b := []byte("world")
	data := [][]byte{a, b}
	sep := []byte{','}
	fmt.Println(string(bytes.Join(data, sep)))
}
