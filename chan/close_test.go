package main

import "fmt"

func ExampleClose() {
	chMsg := make(chan string)
	chEnd := make(chan struct{}) // 如果去掉chEnd, 则不保证have received all messages会被打印

	go func() {
		for msg := range chMsg {
			fmt.Println("receive message", msg)
		}
		fmt.Println("have received all messages")
		close(chEnd)
	}()

	for i := 0; i < 3; i++ {
		chMsg <- string('a' + i)
	}
	// 如果不执行close操作，那么主协程因为receive chEnd而阻塞
	// 而另一个协程也因等待receive chMsg而阻塞，死锁
	close(chMsg)
	<-chEnd

	// Output:
	// receive message a
	// receive message b
	// receive message c
	// have received all messages
}
