package main

import "os"
import "fmt"

func main() {
	name, _ := os.Hostname()
	fmt.Println(name)
}
