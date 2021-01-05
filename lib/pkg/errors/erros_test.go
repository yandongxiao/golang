// 新生成一个错误, 带堆栈信息
// func New(message string) error

// 只附加新的信息
// func WithMessage(err error, message string) error

// 只附加调用堆栈信息
// func WithStack(err error) error

//同时附加堆栈和信息
// func Wrap(err error, message string) error

package main

import (
	"fmt"
	"testing"

	stderrors "errors"

	"github.com/pkg/errors"
)

func ExampleWrap() {
	e := errors.Wrap(stderrors.New("err: a"), "read failed")
	fmt.Println(e)

	// Output:
	// read failed: err: a
}

func ExampleNew() {
	e := errors.New("err: a")
	fmt.Println(e)

	// Output:
	// err: a
}

func TestStack(t *testing.T) {
	e := errors.New("err: a")
	t.Logf("%+v", e)
}