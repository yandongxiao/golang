package main

import (
	"errors"
	"fmt"
)

func ExampleWrap() {
	// wrap error
	// *fmt.errorString
	err := errors.New("first error")

	// *fmt.wrapError
	err = fmt.Errorf("second error: %w", err)
	err = fmt.Errorf("third error: %w", err)
	fmt.Println(err)

	// Output:
	// third error: second error: first error
}

func ExampleUnwrap() {
	err := fmt.Errorf("third error: %w",
		fmt.Errorf("second error: %w",
			errors.New("first error")))
	for err != nil {
		fmt.Println(err)
		err = errors.Unwrap(err)
	}

	// Output:
	// third error: second error: first error
	// second error: first error
	// first error
}

func ExampleIs() {
	original := errors.New("first error")

	// 不能传递, errors.New("first error")。因为两个不同的指针的比较，肯定不等。
	// 或者，实现 Is 方法
	err := fmt.Errorf("third error: %w",
		fmt.Errorf("second error: %w",
			original))

	if errors.Is(err, original) {
		fmt.Println(original)
	}

	// Output:
	// first error
}

type MyError string

func NewErr(str string) error {
	return MyError(str)
}

func (err MyError) Error() string {
	return string(err)
}

func ExampleAs() {
	// 可以传递, NewErr("first error")。因为比较的是字符串。
	// 或者，实现 Is 方法
	original := NewErr("first error")
	err := fmt.Errorf("third error: %w",
		fmt.Errorf("second error: %w",
			NewErr("first error")))

	if errors.Is(err, original) {
		fmt.Println(original)
	}

	// As unwraps its first argument sequentially looking for an error that can be
	// assigned to its second argument, which must be a pointer.
	var myError MyError
	if errors.As(err, &myError) {
		fmt.Println(myError)
	}

	// Output:
	// first error
	// first error
}
