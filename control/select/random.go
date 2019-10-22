package main

import "fmt"

func main() {
	c := make(chan struct{})
	close(c)
	select {
	case c <- struct{}{}: // panic if this case is selected.
	case <-c:
	}

	fmt.Println("50% possibility to panic")
}
