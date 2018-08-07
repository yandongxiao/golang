// The convention in the Go libraries is that even when a package uses panic internally, a recover is done so that its external API still presents explicit error return values.
// 1. always recover from panic in your package: no explicit panic() should be allowed to cross a package boundary
// 2. return errors as error values to the callers of your package.

package main

import "fmt"

func main() {
	defer func() {
		// The recover built-in function allows a program to manage behavior of a panicking goroutine.
		// Executing a call to recover inside a **deferred** function
		// stops the panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic.
		//
		// 返回值为nil的情况
		// If recover is called outside the deferred function it will not stop a panicking sequence. In this case, or when the goroutine is not panicking, or if the argument supplied to panic was nil, recover returns nil
		// Thus the return value from recover reports whether the goroutine is panicking.
		if r := recover(); r != nil { // r is an empty interface
			fmt.Println(r)
		}
	}()

	panic("hello")
	// 一旦发生错误，这个发生panic的函数的后面的语句不再会被执行
	println("world") // unreachable code
}
