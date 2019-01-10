package main

import "fmt"

func ExampleMakeMap1() {
	// creates a new empty map with enough space to hold a small
	// number of entries without reallocating memory again.
	// The small number is compiler dependent.
	m := make(map[int]int)
	fmt.Println(len(m))
	// Output:
	// 0
}

func ExampleMakeMap2() {
	// creates a new empty map which is allocated with enough space
	// to hold at least n entries without reallocating memory again.
	m := make(map[int]int, 10)
	fmt.Println(len(m))
	// Output:
	// 0
}

func ExampleMakeSlice1() {
	// the capacity of the new created slice is the same as its length.
	m := make([]int, 10)
	fmt.Println(len(m))
	// Output:
	// 10
}

func ExampleMakeSlice2() {
	m := make([]int, 0, 10)
	fmt.Println(len(m))
	// Output:
	// 0
}
