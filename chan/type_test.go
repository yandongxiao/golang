// NOTE: all types can be used in a channel.(包括channel)
package main

import (
	"fmt"
	"time"
)

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

func ExampleInterface() {
	c := make(chan interface{})
	go func() {
		val := <-c
		fmt.Println(val.(int))
	}()
	c <- 3
	time.Sleep(time.Second)
	// Output:
	// 3
}
