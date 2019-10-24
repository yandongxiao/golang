// After the last value has been received from a closed channel c,
// any receive from c will succeed without blocking, returning the zero
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

func ExampleReceiveFromClosedChan() {
	ch := make(chan int, 1)
	ch <- 10
	close(ch)
	fmt.Println(<-ch)

	val, has := <-ch
	fmt.Println(val, has)
	//Output:
	// 10
	// 0 false
}
