package main

import "fmt"

// a function signature: is composed of two type list, one is the input parameter type list, the other is the output result type lists.
// a function type: is composed of the func keyword and a function signature literal.
// a function type literal: is composed of the func keyword and a function signature literal.
// a function prototype: func Double(n int) (result int).
// a function declaration: func Double(n int) (result int){}. 匿名函数不算作函数声明，函数声明只能是package level
// When we declare a custom function, we also declared an immutable function value acutally.
type MyFunc func(arg int) int

type I interface {
	f(arg int) int
}

func main() {
	// a function literal: a function type literal + {}
	var f MyFunc = func(arg int) int {
		return arg
	}

	fmt.Println(f(210))
}
