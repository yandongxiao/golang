// the underlying type of every constant is a basic type: boolean, string, or number.
// 所以可以定义time.Duration类型的常量
package main

type Person struct {
	name string
}

// 不能定义 struct 常量
// const initializer Person literal is not a constant
// const P Person = Person{"jack"}

func main() {
	// They are created at compile time, even when defined as locals in functions
	const v = 1 << 3 // constant expression

	// 虽然等号右边的值是常量，与golang是否运行无关，但却不是常量表达式
	// const v2 = math.Sin(math.Pi / 4)

	// 常量的值不能通过函数调用来获得
	// const v3 = add(1, 2)
}

func add(a, b int) int {
	return a + b
}
