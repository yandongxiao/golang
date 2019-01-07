// errors are the last return value and have type error,
// a built-in interface. A nil value in the error position
// indicates that there was no error.
//
// package os的PathError包含Op, Path, Err.
// 构成的error string: open /etc/passwx: no such file or directory.
// 调用os.Open的函数一般要么打印error string，要么向上返回该error.
// 所以，识别error抛出的源头是重要的.
//
// error strings should identify their origin, such as
// by having a prefix naming the operation or package that
// generated the error.
//
// naming convention: Error types **end** in "Error"
// and error variables are called (or **start** with) "err" or "Err".
package main

import (
	"errors"
	"fmt"
)

func ExampleErrors() {
	// errors.New constructs a basic error value with the given
	// error message. errors是一个package，里面只有New函数
	err := errors.New("helloworld")
	fmt.Println(err)
	// Output:
	// helloworld
}

func ExampleErrorf() {
	// fmt.Errorf可以直接返回一个error的实例
	err := fmt.Errorf("helloworld:%d", 10)
	fmt.Println(err)
	// Output:
	// helloworld:10
}

// A package can also define its own specific Error
// with additional methods, like net.Error.
type argError struct {
	arg  int
	prob string
}

func (err *argError) Error() string {
	return err.prob // NOTE: 返回的信息应该尽量全
}

func ExampleCustomize() {
	var err error
	err = &argError{400, "helloworld"}
	// 如何防止type assertion出错时导致程序崩溃
	if ae, ok := err.(*argError); ok {
		fmt.Println(ae)
	}
	// Output:
	// helloworld
}
