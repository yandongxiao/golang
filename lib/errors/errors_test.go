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
	v := MyError(str)
	return &v
}

func (err *MyError) Error() string {
	return string(*err)
}

func ExampleAs() {
	original := NewErr("first error")
	err := fmt.Errorf("third error: %w",
		fmt.Errorf("second error: %w",
			original))

	// As unwraps its first argument sequentially looking for an error that can be
	// assigned to its second argument, which must be a pointer.
	var myError *MyError
	if errors.As(err, &myError) {
		fmt.Println(myError)
	}

	// Output:
	// first error
}
