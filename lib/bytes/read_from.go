package main

// ReadFrom是io package中的一个方法，被ReaderFrom所wrap
import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	buffer := new(bytes.Buffer)
	reader := strings.NewReader("hello world")
	fmt.Println(buffer.ReadFrom(reader)) // 与io.Copy类似，EOF不会被作为error被返回
	fmt.Println(buffer.ReadFrom(reader))
}
