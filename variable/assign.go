package main

import "fmt"

func main() {
	// 多重赋值
	a, b, c := 1, true, "helloworld"
	fmt.Println(a, b, c)

	// 1. 等号左边：using the usual left-to-right rule before any variables are assigned their value.
	// 2. Once everything is evaluated
	// 3. 等号右边：the right-hand-side expressions are all evaluated before any left-hand-side expressions are assigned.
	i := 0
	sa := []int{1, 2, 3}
	// 1. 等号左边：i, sa[0]; 2. 等号右边:1, 0. 3. i==1, sa[0] == 0.
	i, sa[i] = 1, i
	fmt.Println(i, sa)

	// 匿名变量
	a, _, c = 1, false, "helloworld"
	fmt.Println(a, b, c)

	// 左右两边的值的个数必须相等
	//a = 1, 2
	//fmt.Println(a)

	// 利用var定义多个变量
	var (
		b1 = 1 // 没有逗号
		b2 = 1
	)
	fmt.Println(b1, b2)
}
