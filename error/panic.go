// panic工作原理:
// The panic built-in function stops normal execution of the current goroutine.
// When a function F calls panic, normal execution of F stops immediately.
// Any functions whose execution was deferred by F are run in the usual way, and then F returns to its caller.
// To the caller G, the invocation of F then behaves like a call to panic, terminating G's execution and running any deferred functions.
// NOTICE: 由上而知，defer语句的内容一定会被执行。如果没有调用recover函数，错误将会沿着调用堆栈向上抛
// This continues until all functions in the executing goroutine have stopped, in reverse order.
// At that point, the program is terminated and the error condition is reported, including the value of the argument to panic.
// NOTICE: 如果一个协程的最顶层都panic了，那么它将不只是影响该协程，整个程序都会退出, 并且无法在其它协程中捕获)
// This termination sequence is called panicking and can be controlled by the built-in function recover.
//
// package如何使用panic:
// The convention in the Go libraries is that even when a package uses panic internally,
// a recover is done so that its external API still presents explicit error return values.
// 1. always recover from panic in your package: no explicit panic() should be allowed to cross a package boundary
// 2. return errors as error values to the callers of your package.
//
// NOTICE:  发生panic意味着that something **impossible** has happened. 在init中，初始化package失败时，可抛出异常
package main

import "fmt"

func main() {
	defer func() {
		// NOTICE: 即使是在recover调用之前
		fmt.Println("helloworld")

		// 捕获panic的一般做法：
		// The recover built-in function allows a program to manage behavior of a panicking goroutine.
		// Executing a call to recover inside a **deferred** function
		// stops the panicking sequence by restoring normal execution and retrieves the error value passed to the call of panic.
		//
		// 返回值为nil的情况
		// If recover is called outside the deferred function it will not stop a panicking sequence.
		// In this case, or when the goroutine is not panicking, or if the argument supplied to panic was nil, recover returns nil
		// Thus the return value from recover reports whether the goroutine is panicking.
		//
		// A call to recover stops the unwinding and returns the argument passed to panic.
		//  run unaffected by the panicking state. 如果fmt.Println发生panic，则它会被fmt.Println的defer语句捕获, 执行recover, 消灭fmt内部的panic
		if r := recover(); r != nil { // r is an empty interface.
			fmt.Println(r)
		}
	}()

	panic("hello")
	// NOTICE: 一旦发生错误，这个发生panic的函数的后面的语句不再会被执行
	println("world") // unreachable code
}
