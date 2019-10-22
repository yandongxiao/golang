package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 10
	pa := &a

	// 一个类型安全指针值可以被显式转换为一个非类型安全指针类型，反之亦然。
	upa := unsafe.Pointer(pa)
	// 一个uintptr值可以被显式转换为一个非类型安全指针类型，反之亦然
	p := uintptr(upa)

	fmt.Println(p)
}
