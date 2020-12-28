package main

import "fmt"

func ExampleTypeAliasDeclaration() {
	// NOTE: INT 和 int是完全一样的类型
	type INT = int // Go1.9新增提醒， 注意与type INT int的区别
	var v1 int = 100
	var v2 INT = 100
	fmt.Println(v1 + v2)
	// Output:
	// 200
}

func ExampleRune() {
	// rune is an alias for int32 and is equivalent to
	// int32 in all ways. It is used, by convention, to
	// distinguish character values from integer values.
	// type rune = int32
	// type byte = uint8
	// NOTE: 不能将rune等价为type rune int32
	m := int32(1)
	n := rune(1)
	fmt.Println(m == n) // true

	// Output:
	// true
}
