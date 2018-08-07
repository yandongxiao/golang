package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	writer, _ := os.Create("/tmp/data")
	defer writer.Close()

	// Writer implements buffering for an io.Writer object.
	// If an error occurs writing to a Writer, no more data will be accepted and all subsequent writes will return the error.
	bufWriter := bufio.NewWriter(writer)
	fmt.Println(bufWriter.Available()) // Available returns how many bytes are unused in the buffer.
	fmt.Println(bufWriter.Buffered())  // Buffered returns the number of bytes that have been written into the current buffer.
	bufWriter.Write([]byte("helloworld"))
	fmt.Println(bufWriter.Available())
	fmt.Println(bufWriter.Buffered())
	// After all data has been written, the client should call the Flush method to guarantee all data has been forwarded to the underlying io.Writer.
	// 并不保证写入到文件
	bufWriter.Flush()
}
