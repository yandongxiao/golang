package main

import "fmt"

var counter = 10

func main() {
	// 同时写
	go func() {
		counter++
	}()
	counter++
	fmt.Println(counter)
}
