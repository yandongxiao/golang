package main

import "fmt"

// In Go, each type has an underlying type. Rules:
//	1. for built-in basic types, the underlying types are themselves.
//	2. the underlying type of unsafe.Pointer is itself.
//  3. the underlying types of an unnamed type, which must be a composite type, is itself.
//  4. in a type declaration, the new declared type and the source type have the same underlying type.
func ExampleUnderlying() {
	type array [3]int

	var a [3]int
	var b array
	fmt.Println(b == array(a))

	// Output:
	// true
}
