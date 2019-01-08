package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) Foo() string {
	return "Foo"
}

func ExampleNil() {
	var p *Person
	fmt.Println(p.Foo())
	// Output:
	// Foo
}
