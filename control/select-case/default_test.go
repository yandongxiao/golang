package main

import "fmt"

func ExampleChannelDefault() {
	var c chan struct{} // nil
	select {
	case <-c: // blocking for ever
	case c <- struct{}{}: // blocking for eveer
	default:
		fmt.Println("Go here.")
	}

	// Output:
	// Go here.
}
