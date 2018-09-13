package main

import (
	"fmt"
	"math"
)

func main() {
	// 常量可以是任意精度
	const n = 100000000000000000000
	const d = 3e30 / n
	// constant 100000000000000000000 overflows int64
	// fmt.Println(int64(n))
	// 如果将常量直接赋值为interface{}, 那么常量的类型为int
	// constant 100000000000000000000 overflows int
	// fmt.Println(n)

	// 两个常量之间的运算，可以不考虑类型问题. 例如虽然n值很大，但是还是能计算出d的值
	fmt.Println(d)

	// **常量是没有类型的**, Sin方法需要的是float64类型
	fmt.Println(math.Sin(d))

	// 整数变量和浮点数常量进行计算
	var num1 = 10 // 如果没有指定类型信息，则num1的数据类型为int
	// println(num1 + 10.2) // constant 10.2 truncated to integer. 还是会报错
	println(num1 + 10.0) // ok

	// 也可以强制为常量声明一个类型
	const NUM2 int = 10
	var num2 = 10.0
	// invalid operation: num2 + NUM2 (mismatched types float64 and int))
	println(num2 + NUM2)
}
