package main

import (
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	sr := strings.NewReader("hello")

	rc := ioutil.NopCloser(sr)
	rc.Close() // 因为Close不做任何操作

	io.Copy(os.Stdout, sr)
}
