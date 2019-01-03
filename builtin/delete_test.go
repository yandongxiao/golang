package main

import "fmt"

func ExampleExistedElement() {
	m := map[int]int{1: 1, 2: 2, 3: 3}
	delete(m, 1)
	// NOTE: the order is unpredictable even if the same loop is run multiple times with the same map
	fmt.Println(len(m)) // 如果打印整个map，则无法保证元素的顺序
	// Output:
	// 2
}

func ExampleNotExistedElement() {
	m := map[int]int{1: 1, 2: 2}
	delete(m, 10)
	fmt.Println(len(m))
	// Output:
	// 2
}

func ExampleNilMap() {
	var m map[int]int
	delete(m, 10)
	fmt.Println(m)
	// Output:
	// map[]
}
