// Design your algorithm in the first place with unbuffered channels.
// Only introduce buffering when performance is problematic.
package main

import "fmt"

func ExampleA() {
	// 默认创建的协程是不带缓存的，这导致无论是发送端还是接收端，发送或接收操作都是阻塞式.
	defer func() {
		fmt.Println(recover())
	}()
	chMsg := make(chan string)
	chMsg <- "hello" // 导致唯一的协程也被阻塞了
	// NOTE:
	// fatal error: all goroutines are asleep - deadlock!
	// NOTE: 这个错误是无法通过defer+recover捕获的，因为所有的协程都sleep了
}

func ExampleBuffered() {
	chMsg := make(chan string, 2)
	chMsg <- "hello"
	chMsg <- "world"
	close(chMsg)
	fmt.Println(<-chMsg)
	fmt.Println(<-chMsg)
	fmt.Println(<-chMsg) // chan关闭后，每次读取都会成功，且不会被阻塞，返回zero value
	// Output:
	// hello
	// world
	//
}
