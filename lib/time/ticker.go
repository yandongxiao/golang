package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for now := range ticker.C {
			// 如果处理时间超过了100毫秒，那么期间产生的ticker event会丢失
			time.Sleep(500 * time.Millisecond)
			fmt.Println(now)
		}
	}()

	time.Sleep(time.Second)
	ticker.Stop()
}
