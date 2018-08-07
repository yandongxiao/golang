package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	name string
	age  int
}

func main() {
	age := 10
	fmt.Println(unsafe.Sizeof(age))

	// 一个指针 + 一个size 各占8个字节
	name := "helloworld"
	fmt.Println(unsafe.Sizeof(name))

	fmt.Println(unsafe.Sizeof(Person{}))
}
