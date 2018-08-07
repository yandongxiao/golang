package main

import "os"
import "fmt"

func main() {
	var file *os.File
	var err error
	/* make sure file and err has the same scope */
	if file, err = os.Open("/tmp/nofile"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("read data")
	}
	defer file.Close()
}
