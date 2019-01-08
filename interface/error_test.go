package main

import "fmt"

type bar struct {
}

func (b *bar) Error() string {
	return "bar"
}

// error is an interface, and *bar implement it
// rule: if err == nil; means the underlying type is nil,
// and the underlying value is nil.
// err = b; means the underlying type is not nil.
func foo() error {
	var b *bar
	return b // b == nil
}

func ExampleError() {
	err := foo()
	if err != nil {
		b := err.(*bar)
		fmt.Println(b == nil)
		fmt.Println(b)
	}

	// Output:
	// true
	// bar
}
