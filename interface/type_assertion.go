// A type assertion provides access to an interface value's underlying concrete value.
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// concrete value --> interface
	v := 10
	var i interface{} = v
	i = 20 // NOTICE: 该操作并没有修改原始变量v的值

	// If i does not hold a T, the statement will trigger a panic.
	// If i holds a T, then t will be the underlying value and ok will be true.
	if num, ok := i.(int); ok {
		println("num=", num)
	}

	// If not, ok will be false and t will be the zero value of type T, and no panic occurs.
	if _, ok := i.(int32); !ok {
		println("this is not a int32 type")
	}

	// interface --> concrete value
	fmt.Println(i)
	fmt.Println(v)

	// 如果接口类型之间存在包含与被包含的关系，可以不借助type assertion，直接进行赋值
	// type assertion: interfaceA 和 Interface B之间是否可以转换，只与底层的类型变量是否同时实现了他们的接口有关
	f, _ := os.Open("/tmp/dd")
	var empty io.Writer = f
	s, ok := empty.(io.Reader)
	fmt.Printf("%T, %v, %v\n", s, s, ok)
}
