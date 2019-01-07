package main

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

func ExampleOrder() {
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

func ExampleDeferFuncParam() {
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

// func ExampleExit() {
// 	defer fmt.Println("hello") /* do not execute */
// 	os.Exit(1)
// 	// Output:
// 	//
// }
