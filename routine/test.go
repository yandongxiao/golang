// 测试：为什么需要buffer chan
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(2)
		ch <- 3
	}()

	select {
	case val := <-ch:
		fmt.Println("nihao", val)
	}
}
