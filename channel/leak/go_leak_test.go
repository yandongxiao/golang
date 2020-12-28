package main

import (
	"fmt"
	"net"
	"time"
)

func broadcastMsg(msg string, addrs []string) error {
	errc := make(chan error)
	quit := make(chan struct{})

	defer close(quit)

	for _, addr := range addrs {
		go func(addr string) {
			select {
			// 虽然开始执行sendMsg中的内容，但是不代表走了这个分支。
			case errc <- sendMsg(msg, addr):
				fmt.Println("done")
			case <-quit:
				fmt.Println("quit")
			}
		}(addr)
	}

	// 正常逻辑是要检查 errc 中是否有错误，这里直接退出，触发defer的运行
	// for _ = range addrs {
	// 	if err := <-errc; err != nil {
	// 		return err
	// 	}
	// }
	return nil
}

func sendMsg(msg, addr string) error {
	fmt.Println("进入发送消息的函数")
	time.Sleep(time.Second)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = fmt.Fprint(conn, msg)
	return err
}

func ExampleGoroutineLeak() {
	addr := []string{"localhost:8080", "http://google.com"}

	err := broadcastMsg("hi", addr)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 等待其它协程的退出！
	time.Sleep(time.Second * 2)

	// Output:
	// 进入发送消息的函数
	// 进入发送消息的函数
	// quit
	// quit
}
