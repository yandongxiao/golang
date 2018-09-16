// go fmt main.go
// 编码规范最大的问题在于: how to approach this Utopia without a long prescriptive style guide.	--> gofmt
//
//	We use tabs for indentation and gofmt emits them by default. Use spaces only if you must.
//	Go has no line length limit.
//	Go needs fewer parentheses than C and Java.
package main

import "fmt"

func main() {
	x := 1
	y := 1
	z := x<<1 + y<<2 // the operator precedence hierarchy is shorter and clearer. 空格很好的标记了运算符的优先级
	fmt.Println(z)

	// control structures (if, for, switch) do not have parentheses in their syntax.
	if x != 5 && y != 2 { // If a line feels too long, wrap it and indent with an extra tab.
		panic("fail to init")
	}
}
