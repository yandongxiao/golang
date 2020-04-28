// code that uses goroutines can be called from init routines and global initialization expressions
package main

import "fmt"

var IntChan = func() <-chan int {
	ch := make(chan int)
	// 在这里创建了协程!
	go func() {
		ch <- 1
	}()
	return ch
}()

func init() {
	ch := make(chan int)
	// 在这里创建了协程
	go func() {
		ch <- 1 + <-IntChan
	}()
	fmt.Println(<-ch)
}

func main() {}
