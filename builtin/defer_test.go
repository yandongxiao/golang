package buildin

import (
	"fmt"
)

func ExampleDeferOrder() {
	fmt.Println("begin")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	defer fmt.Println("defer createpanic")
	panic("try it")
	fmt.Println("end")

	// Output:
	// begin
	// defer createpanic
	// try it
}

func ExampleDeferFuncWithParam() {
	var i = 10
	defer func(v int) {
		fmt.Println(v) // 10
	}(i)
	defer func() {
		fmt.Println(i) // 20
	}()
	i = 20

	// Output:
	// 20
	// 10
}

// /////////////

func ExampleDeferOrder_2() {
	panicOrder()
	// Output:
	// 421
}

var x int

func f() int {
	x++
	return x
}

func panicOrder() {
	o := fmt.Print

	defer o(f()) // 1. x == 1, 最后执行o函数。

	defer func() {
		defer o(recover()) // 取出panic的值，等于2
		o(f())             // 4. x == 4
	}()

	defer f() // 3. x == 3

	defer recover()

	panic(f()) // 2. x == 2, NOTE：panic的形参是interface, 它会拷贝x的值。
}
