// rune is an alias for int32 and is equivalent to
// int32 in all ways. It is used, by convention, to
// distinguish character values from integer values.
// type rune = int32
// type byte = uint8
package main

import "fmt"

func test1() {}
func test2() {}

func main() {
	// NOTE: 不能将rune等价为type rune int32
	m := int32(1)
	n := rune(1)
	fmt.Println(m == n) //true

	// type INT32 int32
	// q := INT32(1)
	// fmt.Println(m == q) // 编译类错误

	// Output:
	// true
}
