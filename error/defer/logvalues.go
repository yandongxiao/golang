// defer_logvalues.go
package main

import (
	"io"
	"log"
)

// 说明defer的执行是在return或panic之后
// 使用命名返回值的一个好处是在defer函数执行时，可以获取返回值
func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

func main() {
	func1("Go")
}

// Output: 2011/10/04 10:46:11 func1("Go") = 7, EOF
