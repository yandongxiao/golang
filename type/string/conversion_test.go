// A []rune conversion applied to a UTF-8-encoded string
//
// When a string is converted to a rune slice, the bytes stored in
// the string will be viewed as successive UTF-8 encoding byte sequence
// representations of many Unicode code points.
//
// When a string is converted to a byte slice, the resultant byte slice is
// just a deep copy of the underlying byte sequence of the string.
//
// case: a string is converted to a byte slice
// NOTE: A memory allocation is needed to store the deep copy in each of such
// conversions. The reason why a deep copy is essential is slice elements
// are mutable but the bytes stored in strings are immutable, so a byte
// slice and a string can't share byte elements.
//
// Conversions between byte slices and rune slices are not supported directly in Go
// unicode/utf8 或 use the Runes function in the bytes standard package
// 使用string作为中间结果的方法，需要两次深度拷贝，不见得是好方法
//
// 编译器优化，防止深度拷贝的几种情况：
// 	1. for i, b := range []byte(str)
//	2. m[string(key)] = "value"
//	3. if string(x) != string(y) {
//	4. s = (" " + string(x) + string(y))[1:]
//	   at least one of concatenated string values is a non-blank string constant.
package main

import (
	"fmt"
	"unicode/utf8"
)

func ExampleNumber2String() {
	fmt.Println(string(0x4f60))
	c := '你' // NOTE: 即使c:='a', c的类型也是rune
	fmt.Printf("%T %v\n", c, string(c))

	// Output:
	// 你
	// int32 你
}

func ExampleString2Number() {
	s := "你好"
	fmt.Printf("%v\n", []rune(s)[0])

	// Output:
	// 20320
}

func ExampleString2Slice() {
	s := "\u0061\u4f60\u597d" // 直接给定**UNICODE**字符
	rs := []rune(s)
	fmt.Println(rs, len(rs))
	// 数组 --> 字符串
	s = string(rs)
	fmt.Println(s)
	// Output:
	// [97 20320 22909] 3
	// a你好
}

func ExampleConvert() {
	/* cannot convert name (type string) to type []int */
	name := "你好"
	fmt.Println(len([]int32(name)))
	fmt.Println(len([]rune(name)))
	fmt.Println(utf8.RuneCountInString(name))
	fmt.Println(len([]byte(name)))
	fmt.Println(len(name))
	fmt.Println(len([]uint8(name)))
	// cannot convert name (type string) to type []int8
	// fmt.Println(len([]int8(name)))
	// Output:
	// 2
	// 2
	// 2
	// 6
	// 6
	// 6
}
