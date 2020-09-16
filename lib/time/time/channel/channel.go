package main

import (
	"fmt"
	"time"
)

func myroutine(val chan int) {
	val <- 1
}

func after() <-chan int {
	val := make(chan int)
	time.Sleep(3e9)
	go myroutine(val)
	return val
}

func main() {
	num := <-after()
	fmt.Println(num)
}
