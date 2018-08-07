package main

import "fmt"

func main() {
	ch := make(chan string)

	select {
	case ch <- "hello":
		fmt.Println("succeed to send")
	default:
		fmt.Println("failed to send")
	}

	select {
	case msg := <-ch:
		fmt.Println("succeed to receive", msg)
	default:
		fmt.Println("failed to receive")
	}
}
