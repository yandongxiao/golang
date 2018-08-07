// The close built-in function closes a channel, which must be either bidirectional or send-only.
// It should be executed only by the sender, never the receiver, and has the effect of shutting down the channel after the last sent value is received.
// After the last value has been received from a closed channel c, any receive from c will succeed without blocking, returning the zero value for the channel element.
package main

func main() {
	chMsg := make(chan string)
	chEnd := make(chan bool)

	go func() {
		for {
			// NOTE: 注意接收message的同时，需要判断该channel是否已经关闭
			// It will also set more to false for a closed channel
			msg, more := <-chMsg
			if more {
				println("receive message", msg)
			} else {
				println("have received all messages")
				break
			}
		}
		chEnd <- true
	}()

	for i := 0; i < 3; i++ {
		chMsg <- string(i)
	}

	// 如果不执行close操作，那么主协程因为receive chEnd而阻塞
	// 而另一个协程也因等待receive chMsg而阻塞，死锁
	close(chMsg)
	<-chEnd
}
