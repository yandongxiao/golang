// A common idiom used to let the main program
// block indefinitely while other goroutines
// run is to place select {} as the last statement
// in a main function. The default clause is optional;
// fall through behavior, like in the normal switch, is not permitted.
// If there are no cases, the select blocks execution forever.
package main

import "fmt"

func ExampleSelectDefault() {
	ch := make(chan string) // 如果是buffered channel, 输出则变成succeed...
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

	// Output:
	// failed to send
	// failed to receive
}

func ExampleSelectOrder() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "first"
		ch2 <- "second"
	}()

Loop:
	for {
		select {
		case msg1 := <-ch1: // 一定是先收到first
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2) // undefined: msg1
			break Loop
		}
	}

	// Output:
	// first
	// second
}
