package main

import (
	"bytes"
	"fmt"
)

func ExampleJoin() {
	a := []byte("hello")
	b := []byte("world")
	data := [][]byte{a, b}
	sep := []byte{','}

	// 使用 %q 带来的好处
	fmt.Printf("%q\n", bytes.Join(data, sep))

	// Output:
	// "hello,world"
}
