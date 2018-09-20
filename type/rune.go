// A []rune conversion applied to a UTF-8-encoded string
package main

import "fmt"

func main() {

	fmt.Println(string(1234567))

	s := "\u0061\u4f60\u597d" // 直接给定UNICODE字符或者二进制形式
	rs := []rune(s)
	for _, r := range rs {
		fmt.Printf("%c", r)
	}
	fmt.Println()
	s = string(rs)
	fmt.Println(s)
}
