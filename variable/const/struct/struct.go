// the underlying type of every constant is a basic type: boolean, string, or number.
// 所以可以定义time.Duration类型的常量
package main

import "math"

type Person struct {
	name string
}

const P Person = Person{"jack"}

func main() {
	// They are created at compile time, even when defined as locals in functions
	const v = 1 << 3                 // constant expression
	const v2 = math.Sin(math.Pi / 4) // 虽然等号右边的值是常量，与golang是否运行无关，但却不是常量表达式
	const v3 = add(1, 2)             // 类似，看来常量的值不能通过函数调用来获得
}

func add(a, b int) int {
	return a + b
}
