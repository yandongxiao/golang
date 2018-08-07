// 协程没有成功捕获deadlock的问题
package main

import (
	"fmt"
)

func test() {
	ch := make(chan int)
	defer func() {
		fmt.Println("----")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ch <- 1
}

func main() {
	go test()
	select {}
}
