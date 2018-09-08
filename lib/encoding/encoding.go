// Package encoding defines **interfaces** shared by other packages that convert data to and from byte-level and textual representations.
// NOTE: package encoding定义了编码和解码的接口规范. 共计四个接口，每个接口只有一个方法
// Standard types that implement these interfaces include time.Time and net.IP.
package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	// binary encode
	binary, err := now.MarshalBinary()
	if err != nil {
		panic(err)
	}

	// binary decode
	t := new(time.Time)
	err = t.UnmarshalBinary(binary)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)

	// text encoding
	// NOTE: 对对象进行UTF-8编码
	text, _ := now.MarshalText()
	fmt.Println(string(text))

	// text decoding
	t.UnmarshalText(text)
	fmt.Println(t)
}
