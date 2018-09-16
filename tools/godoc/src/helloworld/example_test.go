// 包名是当前包名 + _test, 如: strings_test
package helloworld_test

import (
	"fmt"
)

// 此注释将会被展示在页面上
// 此函数将被展示在OverView区域
func Example() {
	fmt.Println("Hello OverView")
	// Output:
	// Hello OverView
}

func ExampleFoo_basic() {
	fmt.Println("helloworld")
	// Output:
	// helloworld
}

func ExampleFoo_advance() {
}
