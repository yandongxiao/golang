package main

import "fmt"

// NOTE: golang不支持重载
func ExampleBasic() {
	add := func(a int, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2))
	// Output:
	// 3
}

func ExampleVariadic() {
	add := func(nums ...int) int { // nums的类型为[]int
		sum := 0
		for _, x := range nums {
			sum += x
		}
		return sum
	}

	fmt.Println(add(1, 2, 3, 4))
	fmt.Println(add()) // add()是合法的，nums == nil
	x := []int{1, 2, 3, 4}
	fmt.Println(add(x...)) // NOTE: x...的语法要求x必须是slice类型，不可以是数组类型
	// Output:
	// 10
	// 0
	// 10
}
