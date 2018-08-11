package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

// MultiWriter creates a writer that duplicates its writes to all the provided writers, similar to the Unix tee(1) command
func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")

	var buf1, buf2 bytes.Buffer
	w := io.MultiWriter(&buf1, &buf2)

	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}

	fmt.Print(buf1.String())
	fmt.Print(buf2.String())
}
