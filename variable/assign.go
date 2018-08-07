package main

import "fmt"

func main() {
	// 多重赋值
	a, b, c := 1, true, "helloworld"
	fmt.Println(a, b, c)

	// 匿名变量
	a, _, c = 1, false, "helloworld"
	fmt.Println(a, b, c)

	// 左右两边的值的个数必须相等
	//a = 1, 2
	//fmt.Println(a)
}
