package main

import "fmt"

func main() {
	var c chan struct{} // nil
	select {
	case <-c: // blocking for ever
	case c <- struct{}{}: // blocking for eveer
	default:
		fmt.Println("Go here.")
	}
}
