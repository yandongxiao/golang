package main

import "fmt"

func ExampleTypeDefinition1() {
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
