package main

import "fmt"

func main() {
	// go将会推断出变量a的类型
	var a = "helloworld"
	fmt.Println(a)

	// 声明变量时指定类型, 同时定义多个变量
	// 多重赋值
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 同时定义多个变量的第二种方法
	var (
		b1 = 1
		b2 = 1
	)
	fmt.Println(b1, b2)

	// go将会推断出变量d的类型
	var d = true
	fmt.Println(d)

	// 如果声明变量时没有指定初始化值，变量被赋值为zero-value
	var e int
	fmt.Println(e)
}
