// A []rune conversion applied to a UTF-8-encoded string
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
