package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "done"
	}()

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case t := <-time.After(time.Second):
		fmt.Println(t)
	}
}
