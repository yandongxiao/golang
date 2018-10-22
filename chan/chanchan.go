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
