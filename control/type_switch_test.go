package main

import "fmt"

func ExampleTypeSwitch() {
	whatAmI := func(i interface{}) {
		// a type switch is like a regular switch statement
		// The declaration in a type switch has the same
		// syntax as a type assertion i.(T),
		switch t := i.(type) { // but the specific type T is replaced with the keyword type.
		case bool: // the cases in a type switch specify types (not values)
			fmt.Println("I am bool", t) // 在这个case当中，t的类型信息变为bool
		default: // 在default当中，the variable t is of the same interface type and value as i
			fmt.Printf("I do not know %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(3)

	// Output:
	// I am bool true
	// I do not know int
}

func ExampleTypeSwitch2() {
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
			// NOTE: The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println("number:", v)
		case nil:
			// NOTE: The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println(v)
		default:
			// NOTE: The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println("others:", v)
		}
		// NOTE: each variable denoted by v in the
		// last three branches is a copy of x.
	}
}
