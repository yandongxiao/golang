package main

import (
	"io"
	"log"
	"os"
	"strings"
)

// MultiReader returns a Reader that's the logical concatenation of the provided input readers.
// They're read sequentially.
// Once all inputs have returned EOF, Read will return EOF.
// If any of the readers return a non-nil, non-EOF error, Read will return that error.
func main() {
	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")
	r := io.MultiReader(r1, r2, r3)

	if _, err := io.Copy(os.Stdout, r); err != nil {
		log.Fatal(err)
	}
}
