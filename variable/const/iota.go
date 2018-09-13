package main

import "fmt"

func main() {
	// itoa
	// golang没有显示的enum关键字，使用const+iota进行模拟
	// 1. 每次出现const关键字时，iota被重置为0
	const c0 = iota // c0 == 0
	const c1 = iota // c1 == 0
	// 2. 在下一个const出现之前，每出现一个iota，其代表的数字就会自动增加1
	const (
		c2 = iota // 0
		c3 = iota // 1
	)
	fmt.Println(c0, c1, c2, c3)

	const (
		c00 = iota
		c01
		c02
		c03
	)
	fmt.Println(c00, c01, c02, c03)
}
