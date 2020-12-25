package main

import "fmt"

func ExampleTypeSwitch() {
	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		// Here, v is declared once, but it denotes
		// different varialbes in different branches.
		switch v := x.(type) {
		case []int: // a type literal
			// The type of v is "[]int" in this branch.
			fmt.Println("int slice:", v)
		case string: // one type name
			// The type of v is "string" in this branch.
			fmt.Println("string:", v)
		case int, float64, int32: // multiple type names
			// 注意: The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println("number:", v)
		case nil:
			// 注意: The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println(v)
		default:
			// 注意: The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println("others:", v)
		}
		// 注意: each variable denoted by v in the
		// last three branches is a copy of x.
	}

	// Output:
	// number: 456
	// string: abc
	// others: true
	// number: 0.33
	// number: 789
	// int slice: [1 2 3]
	// others: map[]
	// <nil>
}
