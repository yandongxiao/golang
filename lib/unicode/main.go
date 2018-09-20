// Package unicode provides data and functions to test some properties of Unicode code points.
// 对Unicode字符进行检查
package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(unicode.IsPrint('a'))
	fmt.Println(unicode.IsPrint('\t'))
}
