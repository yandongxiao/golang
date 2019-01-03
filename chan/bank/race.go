package main

import "fmt"

var counter = 10

func main() {
	go func() {
		counter++
	}()
	counter++
	fmt.Println(counter)
}
