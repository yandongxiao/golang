package main

import "fmt"

type addf func(a, b int) int

var add = func(a, b int) int {
	if a == 0 {
		panic("a==0")
	}
	return a + b
}

func init() {
	add = foo(add)
}

// 高阶函数foo的作用: 为参数add增加defer+recver的保护
// 这里隐藏了高阶函数的一个应用，面向切面编程(AOP)：
//	1. foo可以为所有函数提供defer+recover的保护
//  2. 在init中添加add := foo(add)语句，
//     原始的add再也不能被访问到，必须访问封装后的函数
func foo(add addf) addf {
	return func(a, b int) (c int) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("---", err)
				c = -1
			}
		}()
		return add(a, b)
	}
}

func main() {
	fmt.Println(add(0, 2))
}
