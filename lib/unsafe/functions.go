package main

import (
	"fmt"
	"unsafe"
)

type S struct {
	v1 int8
	v2 int16
	v3 int32
	v4 int64
}

func main() {
	s := S{}
	fmt.Println(unsafe.Sizeof(s))      // 16
	fmt.Println(unsafe.Offsetof(s.v4)) // 8
	fmt.Println(unsafe.Alignof(s.v4))  // 8
}
