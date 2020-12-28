package main

import "fmt"

func ExampleIdentical() {
	// Two unnamed struct types are identical only if they have the
	// same sequence of field declarations. Two field declarations are
	// identical only if their respective names, their respective types
	// and their respective tags are all identical.
	// Please note, two non-exported struct field names from different packages are always viewed as two different names.
	v1 := struct {
		a int
	}{}

	v2 := struct {
		a int
	}{}

	fmt.Println(v1 == v2) // 注意：可以比较
	// Output:
	// true
}

func ExampleAssignments() {
	// Two struct values can be assigned to each other only if their
	// types are identical or the types of the two struct values have
	// the identical underlying type (considering field tags)
	// and at least one of the two types is an non-defined type.
	type SPerson struct {
		name string
	}
	var v1 SPerson

	var v2 = struct {
		name string
	}{}

	v2 = v1
	fmt.Println(v1 == v2) // 可以比较
	// Output:
	// true
}

// struct, interface 是否可比较，要看具体的field
// invalid operation: v1 == v2 (struct containing []string cannot be compared)
// func ExampleComparisons() {
//	type SPerson struct {
//		name []string
//	}
//
//	var v1, v2 SPerson
//	fmt.Println(v1 == v2)
// }
