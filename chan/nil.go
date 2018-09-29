// 关闭一个空值管道会引发panic
// 向/从一个空值的管道发送/接收数据，会导致一直等待下去
package main

import "time"

func main() {
	var ch chan int
	// close a nil chan
	// panic: close of nil channel
	// close(ch)

	// send on nil chan
	// fatal error: all goroutines are asleep - deadlock
	// ch <- 1

	// receive on nil chan
	// fatal error: all goroutines are asleep - deadlock!
	go func() { // 这个routine没什么卵用
		time.Sleep(time.Second)
		ch = make(chan int)
		ch <- 10
	}()
	<-ch
}
