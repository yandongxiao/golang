package main

import "fmt"

func main() {
	// 默认创建的协程是不带缓存的，这导致无论是发送端还是接收端，发送或接收的方式都是阻塞式的.
	chMsg := make(chan string)
	chEnd := make(chan bool)

	go func() {
		fmt.Println(<-chMsg)
		fmt.Println(<-chMsg)
		fmt.Println(<-chMsg)
		fmt.Println(<-chMsg)
		chEnd <- true
	}()

	chMsg <- "hello"
	chMsg <- "world"

	// 发送close消息
	// The close built-in function closes a channel, which must be either bidirectional or send-only. 说明接收端是没有权利close一个chan的
	// 证明这句话：any receive from c will succeed without blocking, returning the zero value for the channel element
	close(chMsg)

	// chEnd是必须的，否则主协程退出以后，整个进程就终止了。新创建的协程没有执行完毕就会退出
	fmt.Println(<-chEnd)
}
