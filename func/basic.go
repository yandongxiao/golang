package main

import "fmt"

func main() {
	fmt.Println(add1(1, 2))
	fmt.Println(add2(1, 2))
	fmt.Println(add3(1, 2, 3, 4))
	// 传递一个slice
	x := []int{1, 2, 3, 4}
	fmt.Println(add3(x...))

	seq1 := intSeq()
	seq2 := intSeq()
	fmt.Println(seq1()) // 1
	fmt.Println(seq2()) // 1
	fmt.Println(seq1()) // 2
}

// 基本语法
func add1(a int, b int) int {
	return a + b
}

// golang不支持重载
//func add1(a int64, b int64) int64 {
//	return a + b
//}

// 函数返回多个值
func add2(a int, b int) (int, error) {
	return a + b, nil
}

// 接收多个形参
func add3(nums ...int) (int, error) {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	return sum, nil
}

// clouse
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
