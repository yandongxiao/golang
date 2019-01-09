package main

import "fmt"

func ExampleTypeDefinition1() {
	// 1. a new defined type and its respective source type in
	// type definitions are two distinct types.
	// 3. types can be defined within function bodies.
	type INT int
	var a INT = 10
	fmt.Println(a)

	// Output:
	// 10
}

func ExampleTypeDefinition2() {
	type (
		INT   int
		INT32 int32
	)
	var a INT = 10
	var b INT32 = 10
	fmt.Println(a, b)

	// Output:
	// 10 10
}

func ExampleTypeDefinition3() {
	// two types defined in two type definitions are always two distinct types.
	type (
		INT   int32
		INT32 int32
	)
	// invalid operation: INT(10) == INT32(10) (mismatched types INT and INT32)
	// fmt.Println(INT(10) == INT32(10))
}
