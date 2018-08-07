package main

import (
	"flag"
	"fmt"
	"time"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

func f(left, right chan int) { left <- 1 + <-right }

func main() {
	flag.Parse()
	leftmost := make(chan int)
	var left, right chan int = nil, leftmost
	for i := 0; i < *ngoroutine; i++ {
		left, right = right, make(chan int)
		go f(left, right)
	}

	start := time.Now()
	right <- 0      // bang!
	x := <-leftmost // wait for completion
	end := time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println(x)  // 100000, ongeveer 1,5 s
}
