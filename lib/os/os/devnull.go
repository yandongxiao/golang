package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.OpenFile(os.DevNull, os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	fmt.Println(io.Copy(f, strings.NewReader("helloworld")))
}
