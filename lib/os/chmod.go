package main

import "fmt"
import "os"

func main() {
	file, err := os.Open("/tmp/data")
	if err != nil {
		panic(err)
	}
	file.Chmod(os.ModePerm)
	fmt.Println("set ok")
}
