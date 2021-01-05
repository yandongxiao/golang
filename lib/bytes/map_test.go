package main

import (
	"bytes"
	"fmt"
)

func ExampleMap() {
	output := bytes.Map(func(r rune) rune {
		return r + 1
	}, []byte("abcd"))
	fmt.Printf("%q\n", output)

	// Output:
	// "bcde"
}
