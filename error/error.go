package main

import (
	"fmt"
)

// errors are the last return value and have type error, a built-in interface.
// A nil value in the error position indicates that there was no error.
//
// package os的PathError包含Op, Path, Err. 构成的error string: open /etc/passwx: no such file or directory.
// 所以，error string的信息是很丰富的，调用os.Open的函数一般要么打印error string，要么向上返回该error. 所以，识别error抛出的源头是重要的.
// error strings should identify their origin, such as by having a prefix naming the operation or package that generated the error.
func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New constructs a basic error value with the given error message.
		// errors是一个package，里面只有New函数
		// return -1, errors.New("package: failed to hanle 42")

		// fmt.Errorf可以直接返回一个error的实例
		return -1, fmt.Errorf("package: failed to hanle %d", arg)
	}
	return arg + 3, nil
}

// 1. A package can also define its own specific Error with additional methods, like net.Error.
// 2. naming convention: Error types **end** in "Error" and error variables are called (or **start** with) "err" or "Err".
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
	return err.prob // NOTICE: 返回的信息应该尽量全
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
			fmt.Printf("function succeed to work: %v \n", v)
		}
	}

	if _, err := f2(42); err != nil {
		ae := err.(*argError) // type assertion
		fmt.Println(ae)
	}
	// 如何防止type assertion出错时导致程序崩溃
	if _, err := f2(42); err != nil {
		if ae, ok := err.(*argError); ok {
			fmt.Println(ae)
		}
	}
}
