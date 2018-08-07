// The panic built-in function stops normal execution of the current goroutine.
// When a function F calls panic, normal execution of F stops immediately. Any functions whose execution was deferred by F are run in the usual way, and then F returns to its caller.
// To the caller G, the invocation of F then behaves like a call to panic, terminating G's execution and running any deferred functions.
// This continues until all functions in the executing goroutine have stopped, in reverse order.
// At that point, the program is terminated(如果一个协程的最顶层都panic了，那么它将不只是影响该协程，整个程序都会退出, 并且无法捕获)
//		and the error condition is reported, including the value of the argument to panic.
// This termination sequence is called panicking and can be controlled by the built-in function recover.
package main

import (
	"fmt"
	"time"
)

func main() {

	// 无法捕获其它协程的panic

	go func() {

		defer func() {
			fmt.Println("--", recover())
		}()

		// creates a run-time error
		// has a RuntimeError()-method, to distinguish it from a normal error
		panic("go routine crashed")
	}()

	time.Sleep(time.Second)
}
