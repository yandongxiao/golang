package main

import "fmt"

func ExampleBufferedChannel() {
	ch := make(chan int, 1)
	select {
	case <-ch:
		fmt.Println("receive from chan")
	case ch <- 10:
		fmt.Println("send to chan")
	}

	// Output:
	// send to chan
}
