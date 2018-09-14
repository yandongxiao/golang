package main

import "fmt"

func main() {
	fmt.Println(add1(1, 2))       // 基本语法
	fmt.Println(add2(1, 2))       // add2返回多个值
	fmt.Println(add3(1, 2, 3, 4)) // 可变参数
	x := []int{1, 2, 3, 4}
	fmt.Println(add3(x...)) // NOTICE: x...的语法要求x必须是slice类型，不可以是数组类型
}

// 基本语法
func add1(a int, b int) int {
	return a + b
}

// NOTICE: golang不支持重载
// func add1(a int64, b int64) int64 {
//   return a + b
// }

// 函数返回多个值
func add2(a int, b int) (int, error) {
	return a + b, nil
}

// 接收可变参数，nums的类型为[]int
// NOTICE: add3()是合法的，nums为空
func add3(nums ...int) (int, error) {
	// fmt.Printf("%T\n", nums)
	sum := 0
	for _, x := range nums {
		sum += x
	}
	return sum, nil
}
