package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	// LimitReader returns a Reader that reads from r but stops with EOF after n bytes.
	// golang的另外一个设计哲学：基于interface编程，以下两行是接口层的转换. 所以接口的方法需要有明确的含义
	r := strings.NewReader("some io.Reader stream to be read\n")
	lr := io.LimitReader(r, 4) // 只会读取前四个字节，并返回

	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}
