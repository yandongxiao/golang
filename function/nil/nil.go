// 注意：It is fatal error to call a nil function to start a new goroutine.
//      The fatal error is not recoverable and will make the whole program crash.
// For other situations, calls to nil function values will produce recoverable panics, including deferred function calls.
package main

import (
	"fmt"
)

func ExampleNilGoroutine() {
	var f func()
	defer func() {
		fmt.Println(recover()) // 无法恢复
	}()
	go f()
}

func main() {
	var f func()
	defer func() {
		fmt.Println(recover())
		fmt.Println("hello")
	}()
	f()
}
