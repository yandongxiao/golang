package gerror

import (
	"fmt"
	"io"
)

func ExampleNamedReturnValue() {
	func(s string) (n int, err error) {
		defer func() {
			n++
			fmt.Println(n)
		}()
		return 7, io.EOF
	}("Go")

	// Output:
	// 8
}

func ExamplePanicOrder() {
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
