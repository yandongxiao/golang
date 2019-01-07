// 1. 函数是一等变量（first class），可以作为参数或者返回值.
// 2. 返回函数的函数称为高阶函数

package main

import "fmt"

func ExampleVariable() {
	foo := adder()
	bar := adder()
	fmt.Println(foo(1))
	fmt.Println(bar(1))
	fmt.Println(foo(20))
	fmt.Println(bar(20))
	fmt.Println(foo(300))
	// Output:
	// 1
	// 1
	// 21
	// 21
	// 321
}

func adder() func(int) int {
	// 闭包变量不会随着高阶函数Adder的返回而消失
	// 相反，它可以被返回的匿名函数访问
	num := 0
	return func(elm int) int {
		num += elm
		return num
	}
}
