package main

import (
	"bufio"
	"os"
)

func main() {

	// ReadFrom reads data from r until EOF or error. The return value n is the number of bytes read.
	// Any error except io.EOF encountered during the read is also returned.
	//
	// The Copy function uses ReaderFrom if available
	reader, _ := os.Open("/tmp/data")
	writer, _ := os.Create("/tmp/data.cp")
	bufWriter := bufio.NewWriter(writer)
	bufWriter.ReadFrom(reader)
	bufWriter.Flush()
}
