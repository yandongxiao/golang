// Package utf8 implements functions and constants to support text encoded in UTF-8.
// It includes functions to translate between runes and UTF-8 byte sequences.
// 对Unicode**字符**进行编码，变成utf-8编码；或反过来
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	r, size := utf8.DecodeRuneInString("\xe7\x95\x8c") // '界'
	fmt.Printf("%c, %d", r, size)
}
