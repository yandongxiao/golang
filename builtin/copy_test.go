// func copy(dst, src []Type) int
package main

import "fmt"

func ExampleCopy() {
	x := []byte("hello")
	copy(x, []byte("world"))
	fmt.Printf("%s\n", x)
	copy(x, "WORLD")
	fmt.Printf("%s\n", x)
	// Output:
	// world
	// WORLD
}
