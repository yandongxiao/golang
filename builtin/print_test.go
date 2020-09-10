package main

import "fmt"

// 内置print和fmt.Print函数不是同步函数(synced)。需要使用log.Print函数
func ExamplePrint() {
	fmt.Println("hello world")
	//Output:
	// hello world
}
