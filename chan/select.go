package main

import "time"

func send() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "first"
		ch2 <- "second"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1: // 一定是先收到first
			println(msg1)
		case msg2 := <-ch2:
			println(msg2) // undefined: msg1
		}
	}
}

func recv() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		println(<-ch1)
		println(<-ch2)
	}()

	for i := 0; i < 2; i++ {
		select {
		case ch2 <- "world":
			println("send ch2")
		case ch1 <- "hello": // 一定是先发送first
			println("send ch1")
		}
	}
}

func main() {
	send()
	recv()
	time.Sleep(time.Second)

	// A common idiom used to let the main program block indefinitely while other goroutines run is to place select {} as the last statement in a main function.
	// The default clause is optional; fall through behavior, like in the normal switch, is not permitted.
	// If there are no cases, the select blocks execution forever.
	select {}
}
