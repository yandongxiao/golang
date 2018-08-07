package main

import "os"
import "fmt"

func main() {
	str, _ := os.Getwd()
	fmt.Println(str)
}
