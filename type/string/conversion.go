// A []rune conversion applied to a UTF-8-encoded string
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 数字(字符) --> 字符串
	fmt.Println(string(0x4f60))
	c := '你'
	fmt.Printf("%T %v\n", c, string(c)) // NOTE: 即使c:='a', c的类型也是rune

	// 字符串 --> 数字
	s := "你"
	fmt.Printf("0x%x\n", []rune(s)[0])

	// 字符串 --> 数组
	s = "\u0061\u4f60\u597d" // 直接给定**UNICODE**字符
	rs := []rune(s)
	fmt.Println(rs)

	// 数组 --> 字符串
	s = string(rs)
	fmt.Println(s)

	/* cannot convert name (type string) to type []int */
	name := "你好"
	println(len([]int32(name)))
	println(len([]rune(name)))
	println(utf8.RuneCountInString(name))
	println(len([]byte(name)))
	println(len([]uint8(name)))
	// cannot convert name (type string) to type []int8
	// println(len([]int8(name)))
}
