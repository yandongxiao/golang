package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	// It returns the number of bytes copied and an error if fewer bytes were read.
	// The error is EOF only if no bytes were read.
	// If an EOF happens after reading fewer than min bytes, ReadAtLeast returns ErrUnexpectedEOF.
	// On return, n >= min if and only if err == nil
	r := strings.NewReader("some io.Reader stream to be read")
	buf := make([]byte, 100)
	if n, err := io.ReadAtLeast(r, buf, 100); err != nil {
		fmt.Println(string(buf[:n]), err) // ErrUnexpectedEOF
	} else {
		fmt.Println(string(buf[:n]))
	}

	// EOF
	// The error is EOF only if no bytes were read.
	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		fmt.Println("error:", err)
	}
}
