package main

import (
	"flag"
	"fmt"
	"time"
)

var ngoroutine = flag.Int("n", 100000, "how many goroutines")

// 一遍创建channel的同时，一遍计算结果
func method1() {
	chs := make([]chan int, *ngoroutine+1)
	for i := 0; i < *ngoroutine; i++ {
		chs[i] = make(chan int)
		go func(i int) {
			if i == 0 {
				chs[i] <- 1
			} else {
				chs[i] <- (1 + <-chs[i-1])
			}
		}(i)
	}

	fmt.Println(<-chs[*ngoroutine-1])
}

func method2() {
	var init, in, out chan int
	init = make(chan int)
	out = init
	for i := 0; i < *ngoroutine; i++ {
		in, out = out, make(chan int)
		go func(in <-chan int, out chan<- int) {
			out <- 1 + <-in
		}(in, out)
	}

	// 所有协程都处于阻塞状态，等待执行
	init <- 0
	fmt.Println(<-out)
}

// 事实证明：每个并发协程的工作量如果很少的话
// 如果创建大量的协程, 协程间的通信，会成为性能的瓶颈
func method3() {
	end := make(chan struct{})
	go func() {
		sum := 0
		for i := 0; i < *ngoroutine; i++ {
			sum += 1
		}
		fmt.Println(sum)
		end <- struct{}{}
	}()
	<-end
}

func main() {
	flag.Parse()

	wrapper := func(f func()) {

		begin := time.Now()
		defer func() {
			end := time.Now()
			fmt.Println(end.Sub(begin))
		}()

		f()
	}

	wrapper(method1)
	wrapper(method2)
	wrapper(method3)
}
