// all types can be used in a channel. A channel is in fact a typed message queue
// so channels are first class objects: they can be stored in variables, passed
// as arguments to functions, returned from functions and sent themselves over channels.
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan interface{})
	go func() {
		val := <-c
		fmt.Println(val.(int))
	}()
	c <- 3
	time.Sleep(time.Second)
}
