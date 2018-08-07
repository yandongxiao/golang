package main

import (
	"errors"
	"fmt"
)

// errors are the last return value and have type error, a built-in interface.
// A nil value in the error position indicates that there was no error.
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message.
		// errors是一个package，里面只有New函数
		return -1, errors.New("failed to hanle 42")
	}
	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

// error的定义如下：
// type error interface {
//		Error() string
// }
// 根据golang的原则，任何实现了Error方法的类型都实现了error接口
func (err *argError) Error() string {
	return err.prob
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{400, "42 is illeagal"}
	}
	return arg + 3, nil
}

func main() {
	inputs := []int{1, 2, 42}
	for _, x := range inputs {
		if v, err := f1(x); err != nil {
			fmt.Printf("function failed to work:%v \n", err)
		} else {
			fmt.Printf("function succeed to work:%v \n", v)
		}
	}

	if _, err := f2(42); err != nil {
		ae := err.(*argError) // type assertion
		fmt.Println(ae)
	}
	// 如何放置type assertion出错时导致程序崩溃
	if _, err := f2(42); err != nil {
		if ae, ok := err.(*argError); ok {
			fmt.Println(ae)
		}
	}
}
