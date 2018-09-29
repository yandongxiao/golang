package main

import "fmt"

func main() {
	ch := make(chan string)

	select {
	case ch <- "hello": // NOTICE：执行的是default表达式，该表达式根本未执行
		fmt.Println("succeed to send")
	default:
		fmt.Println("failed to send")
	}

	select {
	case msg := <-ch: // NOTICE：不会被执行
		fmt.Println("succeed to receive", msg)
	default:
		fmt.Println("failed to receive")
	}
}
