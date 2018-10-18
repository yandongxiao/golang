package main

import "fmt"

func main() {
	test1()
	test2()
}

// slice和array的区别
func test1() {
	// 类型区别
	whatIAM := func(v interface{}) {
		switch v.(type) {
		case []int:
			fmt.Println("this is a slice")
		case [3]int:
			fmt.Println("this is an array")
		}
	}
	// this is an array
	s1 := [3]int{1, 2, 3}
	whatIAM(s1)
	// this is a slice
	s2 := []int{1, 2, 3}
	whatIAM(s2)

	// NOTE: 如果data是数组类型，那么就是值传递
	var input [3]int
	func(data [3]int) {
		data[0] = 10
	}(input)
	fmt.Println(input)

	// 如果data是切片类型，那么就是引用传递
	func(data []int) {
		data[0] = 10
	}(input[:])
	fmt.Println(input)
}

func test2() {
	// 初始化的方式
	a1 := []int{1, 2, 3}
	// To create an empty slice with non-zero length, use the builtin make.
	a2 := make([]int, 3)
	fmt.Println(a1, a2)

	// slice的特性操作
	// One is the builtin append, which returns a slice containing one or more new values.
	// Note that we need to accept a return value from append as we may get a new slice value
	s := []string{"a", "b", "c"}
	s = append(s, "d")
	fmt.Println(s)
	// Slices can also be copyed
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)
}
