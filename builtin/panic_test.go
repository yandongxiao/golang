// panic工作原理:
// The panic built-in function stops normal execution of the
// current goroutine. When a function F calls panic, normal
// execution of F stops immediately. Any functions whose
// execution was deferred by F are run in the usual way, and
// then F returns to its caller. To the caller G, the invocation
// of F then behaves like a call to panic, terminating G's execution
// and running any deferred functions.
// NOTICE: 由上而知，defer语句的内容一定会被执行。如果没有调用recover函数，
// 错误将会沿着调用堆栈向上抛
//
// This continues until all functions in the executing goroutine
// have stopped, in reverse order. At that point, the program is
// terminated and the error condition is reported, including the
// value of the argument to panic.
// NOTE: 如果一个协程的最顶层都panic了，那么它将不只是影响该协程，
// 整个程序都会退出, 并且无法在其它协程中捕获
//
// This termination sequence is called panicking and can be controlled
// by the built-in function recover.
//
// package 如何使用 panic:
// The convention(惯例) in the Go libraries is that even when a package
// uses panic internally, a recover is done so that its external API
// still presents explicit error return values.
// 1. always recover from panic in your package: no explicit panic()
//    should be allowed to cross a package boundary
// 2. return errors as error values to the callers of your package.
// 3. 发生panic意味着that something **impossible** has happened.
//    在 init 中，初始化 package 失败时，可抛出异常
package buildin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPanic(t *testing.T) {
	// NOTE: If the deferred function has any return values, they are discarded
	// when the function completes. 所以，传递给 defer 的函数，可以有返回参数
	//
	// If a deferred function value evaluates to nil, execution panics
	// when the function is invoked, not when the "defer" statement is executed.
	// 重点有两个：1. 调用 nil 函数后会导致 panic 发生； 2. 发生panic的时机是在"nil函数"被执行时
	defer func() {
		assert.True(t, true)
		// 捕获panic的一般做法：
		// The recover built-in function allows a program to manage
		// behavior of a panicking goroutine. Executing a call to
		// recover inside a **deferred** function stops the panicking
		// sequence by restoring normal execution and retrieves the
		// error value passed to the call of panic.
		//
		// recover 返回Nil的情况：
		// If recover is called outside the deferred function it will
		// not stop a panicking sequence. In this case, or when the
		// goroutine is not panicking, or if the argument supplied
		// to panic was nil, recover returns nil.
		//
		// Thus the return value from recover reports whether the goroutine
		// is panicking. see builtin/recover_test.go
		if r := recover(); r != nil {
			assert.True(t, true)
		}
	}()

	panic("hello")
	assert.True(t, false)
}

func TestReturnValue(t *testing.T) {
	assert.True(t, foo() == 0)
}

func foo() int {
	defer func() {
		fmt.Println(recover())
	}()
	panic("hello")
	// 命名返回值和非命名返回值，对于Go语言来说是一致的
	// 即，程序在一开始就为函数声明了返回值的变量
	// 所以，return语句的本质看来更像是一个赋值操作
	return 200
}
