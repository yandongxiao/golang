// The close built-in function closes a channel, which must be either bidirectional or send-only.
// It should be executed only by the sender, never the receiver, and has the effect of shutting down the channel after the last sent value is received.
// After the last value has been received from a closed channel c, any receive from c will succeed without blocking, returning the zero value for the channel element.
package main

import "time"

func method1() {
	chMsg := make(chan string)
	chEnd := make(chan struct{}) // 如果去掉chEnd, 则不保证have received all messages会被打印

	go func() {
		for {
			// NOTE: 注意接收message的同时，需要判断该channel是否已经关闭
			// It will also set ok to false for a closed channel
			if msg, ok := <-chMsg; ok {
				println("receive message", msg)
			} else {
				println("have received all messages")
				break
			}
		}
		chEnd <- struct{}{}
	}()

	for i := 0; i < 3; i++ {
		chMsg <- string(i)
	}

	// 如果不执行close操作，那么主协程因为receive chEnd而阻塞
	// 而另一个协程也因等待receive chMsg而阻塞，死锁
	close(chMsg)
	<-chEnd
}

// 改进方法
func method2() {
	chMsg := make(chan string)
	chEnd := make(chan struct{})

	go func() {
		for msg := range chMsg {
			println("receive message", msg)
		}
		println("have received all messages")
		close(chEnd)
	}()

	for i := 0; i < 3; i++ {
		chMsg <- string(i)
	}

	close(chMsg)
	<-chEnd
}

func closeClosedChan() {
	ch := make(chan int)
	go func() {
		close(ch)
		close(ch) // panic: close of closed channel
		ch <- 100 // panic: send on closed channel
	}()
	<-ch
	time.Sleep(time.Second)
}

func main() {
	method2()
	closeClosedChan()
}
