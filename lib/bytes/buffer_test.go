package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func ExampleBuffer() {
	// 也可以使用 strings.Builder
	var buffer bytes.Buffer
	f := strings.NewReader("helloworld")
	if _, err := io.Copy(&buffer, f); err != nil {
		panic(err)
	}
	fmt.Println(buffer.String())

	// Output:
	// helloworld
}
