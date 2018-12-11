package main

import "fmt"

func ExampleCloseClosedChan() {
	defer func() {
		fmt.Println(recover())
	}()

	// NOTE: close函数只是用来关闭chan，与文件操作符无关
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
