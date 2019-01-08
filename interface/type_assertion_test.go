// A type assertion provides access to an interface value's
// underlying concrete value.
package main

import (
	"fmt"
	"io"
	"os"
)

func ExampleTypeAssertion() {
	v := 10
	var i interface{} = v
	i = 20 // NOTE: 该操作并没有修改原始变量v的值
	fmt.Println(i)
	fmt.Println(v)

	// If i does not hold a T, the statement will trigger a panic.
	// If i holds a T, then t will be the underlying value and
	// ok will be true.
	if num, ok := i.(int); ok {
		fmt.Printf("num=%d\n", num)
	}

	// If not, ok will be false and t will be the zero value of
	// type T, and no panic occurs.
	if num, ok := i.(int32); !ok {
		fmt.Printf("num=%d\n", num)
	}

	// Output:
	// 20
	// 10
	// num=20
	// num=0
}

func ExampleInterface() {
	// 如果接口类型之间存在包含与被包含的关系，可以不借助type assertion，
	// 直接进行赋值. type assertion: interfaceA 和 Interface B之间
	// 是否可以转换，只与底层的类型变量是否同时实现了他们的接口有关.
	f, _ := os.Open("/tmp/dd")
	var empty io.Writer = f
	s, ok := empty.(io.Reader)
	fmt.Printf("%T, %v\n", s, ok) // 返回具体的类型

	// Output:
	// *os.File, true
}
