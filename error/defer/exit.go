package main

import "os"

//import "os"

func main() {
	defer println("hello") /* do not execute */
	os.Exit(1)
}
