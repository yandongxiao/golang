// all types can be used in a channel. A channel is in fact a typed message queue
// so channels are first class objects: they can be stored in variables, passed
// as arguments to functions, returned from functions and sent themselves over channels.
package main

import "fmt"

func main() {
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
}
