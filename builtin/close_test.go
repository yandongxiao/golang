// The close built-in function closes a channel, which must be either bidirectional or send-only.
// It should be executed only by the sender, never the receiver, and has the effect of shutting
// down the channel after the last sent value is received. After the last value has been received
// from a closed channel c, any receive from c will succeed without blocking, returning the zero
// value for the channel element.
// NOTE: close函数只是用来关闭chan，与关闭文件操作无关
package main

import "fmt"

func ExampleCloseClosedChan() {
	defer func() {
		fmt.Println(recover())
	}()

	ch := make(chan struct{})
	close(ch)
	close(ch)
	//Output:
	//close of closed channel
}

func ExampleCloseNilChan() {
	defer func() {
		fmt.Println(recover())
	}()
	var ch chan int
	close(ch)
	//Output:
	//close of nil channel
}
