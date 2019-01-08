package main

import "fmt"

type MPerson struct {
	name string
}

func (p *MPerson) Foo() string {
	return "Foo"
}

func ExampleNilPointer() {
	var p *MPerson
	fmt.Println(p.Foo())
	// Output:
	// Foo
}
