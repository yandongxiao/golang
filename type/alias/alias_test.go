package main

import "fmt"

// 注意与 type INT int 的区别
type INT = int

func ExampleAlias() {
	var v1 int = 100
	var v2 int = 100
	// NOTE: INT 和 int是完全一样的类型
	fmt.Println(v1 + v2)

	// Output:
	// 200
}
