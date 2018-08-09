package main

import (
	"flag"
	"fmt"
	"time"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

func foo(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	flag.Parse()
	ch := make(chan int)
	left, right := ch, ch
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go foo(left, right)
	}

	begin := time.Now()
	right <- 0
	result := <-ch
	end := time.Now()
	fmt.Println(end.Sub(begin))
	fmt.Println(result)
}
