package main

import "log"
import "fmt"

func main() {
	fmt.Println(log.Flags())
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Fatalln("file not found")
}
