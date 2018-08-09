package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	buffer := new(bytes.Buffer)

	var f *os.File
	var err error
	if f, err = os.Open("/tmp/data"); err != nil {
		panic(err)
	}

	if n, err := io.Copy(buffer, f); err != nil {
		panic(err)
	} else {
		println("read: ", n)
	}
	println(string(buffer.Bytes()))
}
