package main

import "fmt"

func ExampleExistedElement() {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	delete(m, 1)
	// NOTE: the order is unpredictable even if the same loop is run multiple times with the same map
	for k, v := range m {
		fmt.Println(k, v)
	}
	// Unordered output:
	// 2 2
	// 3 3
}

func ExampleNotExistedElement() {
	m := map[int]int{1: 1, 2: 2}
	delete(m, 10)
	for k, v := range m {
		fmt.Println(k, v)
	}
	// Unordered output:
	// 2 2
	// 1 1
}

func ExampleNilMap() {
	var m map[int]int
	delete(m, 10)
	fmt.Println(m)
	// Output:
	// map[]
}
