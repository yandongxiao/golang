package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// 判断io.Copy是否读取完毕的方法是判断n是否等于0
// Because Copy is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.
// Copy可以拷贝的大小是int64!!
func main() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	if n, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("read: ", n)
	}

	if n, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("read: ", n)
	}
}
