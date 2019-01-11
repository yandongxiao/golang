// func copy(dst, src []Type) int
package main

import "fmt"

func ExampleCopyFast() {
	a := []int{1, 2, 3}
	b := append(a[:0:0], a...)
	fmt.Println(b)
	// Output
	// [1 2 3]
}

func ExampleCopyUnderlying() {
	// the types of the two slices are not required to be identical,
	// but their element types must be identical. In other words,
	// the two argument slices must share the same underlying type.
	type bytes []byte
	x := make(bytes, 10)
	copy(x, []byte("hello"))
	copy(x, "world")           // Sugar
	fmt.Println(string(x[:5])) // :5是必须的
	// Output:
	// world
}

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
