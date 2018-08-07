package main

import (
	"fmt"
)

func foo() (err error) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 捕获到异常，并打印错误
		}
	}()
	panic("have unexpected error") // 这里传递的是一个字符串
}

func main() {
	foo()
}
