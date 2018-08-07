package main

import "os"
import "fmt"
import "io"

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	var file *os.File
	var err error

	file, err = os.Open("/tmp/data")
	errcheck(err)
	defer file.Close()

	buf := make([]byte, 100)
	size, err := file.Read(buf)
	if err != nil {
		if err == io.EOF {
			println("read to the end")
		}
	} else {
		fmt.Printf("read %d size data from file\n", size)
		buf[size-1] = 0
		fmt.Println(string(buf))
	}
}
