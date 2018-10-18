package main

import "fmt"

func main() {
	// 正确1
	m := []int{1, 2, 3}
	for i := len(m) - 1; i >= 0; i-- { // NOTE: 倒序
		if m[i] <= 2 {
			m = append(m[:i], m[i+1:]...)
		}
	}
	fmt.Println(m)

	// 正确2
	m = []int{1, 2, 3}
	n := m[:0]
	for _, v := range m {
		if v > 2 {
			n = append(n, v)
		}
	}
	fmt.Println(n)

	// NOTICE: for range 的遍历次数由m的** 初始值 **决定
	m = []int{1, 2, 3}
	for i := range m {
		fmt.Printf("i=%v, v=%v\n", i, m[i])
		if i <= 2 {
			m = append(m[:i], m[i+1:]...)
		}
	}
}
