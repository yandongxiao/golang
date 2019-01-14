//  calling methods on a nil interface value will panic at run time,'
// for there are no available declared methods can be called.
package main

import (
	"fmt"
	"io"
)

type Any interface{}
type Anything struct{}

func ExampleNil() {
	// NOTICE：判断一个ret是否等于nil，是依据具体类型是否为nil.
	any := getAny()
	if any == nil {
		fmt.Println("any is nil")
	} else {
		fmt.Println("any is not nil")
	}

	// 此时anything的类型是*Anything, 值为nil
	anything := any.(*Anything)
	if anything == nil {
		fmt.Println("anything is nil")
	} else {
		fmt.Println("anything is not nil")
	}

	// reader := io.Reader(nil).
	// 等价于 var reader io.Reader
	myfunc(nil) // NOTE: 具体类型和具体实例都是nil

	// Output:
	// any is not nil
	// anything is nil
	// reader is nil
	// not ok
}

func getAny() Any {
	return getAnything()
}

// 假设ret是一个interface instance, 它由具体类型和具体实例两部分组成。
// var ret *Anything = (*Anything)(nil) 具体类型是*Anything,
// 实例值是nil
func getAnything() *Anything {
	return (*Anything)(nil)
}

func myfunc(reader io.Reader) {
	if reader == nil {
		fmt.Println("reader is nil")
	}

	// reader == nil的情况下，切换interface type不成功也是合理的
	rc, ok := reader.(io.ReadCloser)
	if ok {
		fmt.Println(rc)
	} else {
		fmt.Println("not ok")
	}
}
