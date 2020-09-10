/*
for InitSimpleStatement; Condition; PostSimpleStatement {
	// do something
}
*/
package main

import "fmt"

func ExampleFormat1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}
	// Output:
	// 0123456789
}

// InitSimpleStatement; Condition; PostSimpleStatement中的任意一个都可以忽略，如果只保留Condition，它两边的分号可省略
func ExampleFormat2() {
	i := 1
	for i < 10 {
		fmt.Printf("%d", i)
		i++
	}
	// Output:
	// 123456789
}

func ExampleFormat3() {
	for { // 等价于 for true {
		fmt.Println("helloworld")
	}
	// Output
	// helloworld
}
