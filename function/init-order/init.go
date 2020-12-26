// go run init.go ddd.go
// a common use of init functions is to verify or repair
// correctness of the program state before real execution begins.
// 在init函数中，如果出现失败，可以以panic的方式抛出.
//
// each source file can define its own init function
// 初始化顺序
//	1. all the imported packages have been initialized.
//	2. all the variable declarations in the package have
//     evaluated their initializers
//  3. init is called
//
// package 内部的多个文件可以包含init函数，一个文件内也可以定义多个init函数.
// 那么init的执行顺序？
// 1. go run ddd.go main.go, so init in add.go will be executed first
// 2. 文件内部的init的执行顺序，与init在文件中的位置有关
package main

import (
	"fmt"
	// _ "./pkg" // 这种模式只能在$GOPATH之外使用; 在Module模式下，也不赞成这样使用
)

func init() {
	fmt.Println("main init")
}

func init() {
	fmt.Println("main init2")
}

func main() {
	fmt.Println("main func")
}
