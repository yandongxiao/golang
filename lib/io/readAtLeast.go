package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// If an EOF happens after reading fewer than min bytes, ReadAtLeast returns ErrUnexpectedEOF.
	// It returns the number of bytes copied and an error if fewer bytes were read.
	// On return, n >= min if and only if err == nil
	r := strings.NewReader("some io.Reader stream to be read")
	buf := make([]byte, 100)
	if n, err := io.ReadAtLeast(r, buf, 100); err != nil {
		fmt.Println(string(buf[:n]), err)
	} else {
		fmt.Println(string(buf[:n]))
	}

	// EOF
	// The error is EOF only if no bytes were read.
	if _, err := io.ReadAtLeast(r, buf, 4); err == io.EOF {
		fmt.Println("error:", err)
	}

	// buffer smaller than minimal read size.
	// If min is greater than the length of buf, ReadAtLeast returns ErrShortBuffer.
	shortBuf := make([]byte, 3)
	r = strings.NewReader("some io.Reader stream to be read")
	if _, err := io.ReadAtLeast(r, shortBuf, 4); err == io.ErrShortBuffer {
		fmt.Println("error:", err)
	}

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)
	if _, err := io.ReadAtLeast(r, longBuf, 64); err != nil {
		fmt.Println("error:", err)
	}
}
