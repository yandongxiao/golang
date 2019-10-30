package main

import "fmt"

var x int

func f() int {
	x++
	return x
}

func ExampleTest() {
	o := fmt.Print

	defer o(f()) // 1. x == 1, 最后执行o函数。

	defer func() {
		defer o(recover()) // 取出panic的值，等于2
		o(f())             // 4. x == 4
	}()

	defer f() // 3. x == 3

	defer recover()

	panic(f()) // 2. x == 2, NOTE：panic的形参是interface, 它会拷贝x的值。

	// Output:
	// 421
}
