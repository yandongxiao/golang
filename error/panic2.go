package main

import "fmt"

func foo() {
	defer func() {
		// r is an empty interface
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	panic("hello")
}

func main() {
	foo()
	println("world")
}
