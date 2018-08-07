// Go supports user-defined or custom types in the form of alias types or structs.
package main

import "fmt"

type INT int

type Person struct {
	name string
}

func main() {
	a := INT(3)
	p := Person{
		name: "hello",
	}

	fmt.Println(a, p)
}
