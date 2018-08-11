package main

import "io"

func main() {
	_, pw := io.Pipe()
	pw.Write([]byte("hello"))
}
