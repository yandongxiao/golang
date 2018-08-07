package main

import "os"
import "fmt"

func main() {
	str := "/tmp/file"
	file, err := os.Open(str)
	if err != nil {
		println("can not read file:", str)
		os.Exit(-1)
	}

	data := make([]byte, 100)
	nbytes, err := file.Read(data)
	if err != nil {
		println("read failed")
		os.Exit(-1)
	}

	fmt.Println(nbytes, len(data), string(data[:nbytes]))
}
