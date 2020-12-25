package main

import "fmt"

func ExampleTrySendAndReceive() {
	c := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c <- v:
		default: // go here if c is full.
		}
	}
	tryReceive := func() string {
		select {
		case v := <-c:
			return v
		default:
			return "-" // go here if c is empty.
		}
	}
	trySend("Hello!")
	trySend("Hi!")
	trySend("Bye!")           // fail to send, but will not blocked.
	fmt.Println(tryReceive()) // Hello!
	fmt.Println(tryReceive()) // Hi!
	fmt.Println(tryReceive()) // -

	// Output:
	// Hello!
	// Hi!
	// -
}
