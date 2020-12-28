package main

import "fmt"

// slice和array的区别
func ExampleSliceArray() {
	// 类型区别
	whatIAM := func(v interface{}) {
		switch v.(type) {
		case []int:
			fmt.Println("this is a slice")
		case [3]int:
			fmt.Println("this is an array")
		}
	}
	s1 := [3]int{1, 2, 3}
	whatIAM(s1)
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
	// Output:
	// this is an array
	// this is a slice
	// [0 0 0]
	// [10 0 0]
}

func ExampleInitialize() {
	// array或者slice在进行literal初始化时，可以指定{0：0, 2:20, 30}.
	// 注意30的下标值是前一个下标2+1
	ss2 := []int{0: 0, 2: 20, 30}
	fmt.Println(ss2)

	// 初始化的方式
	a1 := []int{1, 2, 3}
	a2 := make([]int, 3)
	fmt.Println(a1, a2)
	// Output:
	// [0 0 20 30]
	// [1 2 3] [0 0 0]
}
