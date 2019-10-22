// NOTE: fatal error: all goroutines are asleep - deadlock! 该错误信息+堆栈信息都打印在了stderr上
// NOTE: 该错误是无法通过defer+recover恢复过来的，因为所有的协程都处于asleep状态
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
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	go test()
	select {}
}
