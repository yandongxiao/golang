// func copy(dst, src []Type) int
package main

import "fmt"

func ExampleDeepCopy() {
	a := []int{1, 2, 3}
	// append申请新的内存块的时机
	b := append(a[:0:0], a...)
	a[0] = 100
	fmt.Println(b)
	// Output
	// [1 2 3]
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
