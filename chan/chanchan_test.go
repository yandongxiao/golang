// NOTE: all types can be used in a channel.
// A channel is in fact a typed message queue so channels are first class objects
package main

import "fmt"

func ExampleChan() {
	ch1 := make(chan int)
	ch2 := make(chan chan int)
	go func() {
		ch := <-ch2
		ch <- 1
		close(ch)
	}()

	ch2 <- ch1
	for v := range ch1 {
		fmt.Println(v)
	}

	// Output:
	// 1
}
