package main

import (
	"io"
	"os"
)

func main() {
	io.WriteString(os.Stdout, "hw")
}
