// func copy(dst, src []Type) int
package main

import "fmt"

func ExampleCopyBytes() {
	x := []byte("hello")
	copy(x, []byte("world"))
	fmt.Printf("%s\n", x)
	// Output:
	// world
}

func ExampleCopyFromNil() {
	x := []byte("hello")
	copy(x, []byte(nil))
	fmt.Printf("%s\n", x)
	// Output:
	// hello
}

func ExampleCopyTONil() {
	var x []byte
	copy(x, []byte("hello"))
	fmt.Println(x == nil)
	// Output:
	// true
}

func ExampleCopyOverlap() {
	x := []byte("hello")
	copy(x, x[2:])
	fmt.Printf("%s\n", x)
	// Output:
	// llolo
}

// NOTE: As a special case, it also will copy
// bytes from a string to a slice of bytes
func ExampleCopyString() {
	x := []byte("hello")
	copy(x, "WORLD")
	fmt.Printf("%s\n", x)
	// Output:
	// WORLD
}
