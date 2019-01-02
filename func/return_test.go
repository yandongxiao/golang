package main

import "fmt"

func ExampleReturn() {
	add := func(a int, b int) (int, error) {
		return a + b, nil
	}
	fmt.Println(add(1, 2))
	// Output:
	// 3 <nil>
}
