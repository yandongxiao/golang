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
