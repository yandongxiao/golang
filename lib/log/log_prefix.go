package main

import "log"
import "fmt"

func main() {
	fmt.Println(log.Prefix())
	log.SetPrefix("prefix")
	fmt.Println(log.Prefix())
	log.Fatalln("fatal error")
}
