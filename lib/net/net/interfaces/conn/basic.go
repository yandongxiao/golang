// package net中最重要的接口之一
// Conn is a generic stream-oriented network connection.
// Multiple goroutines may invoke methods on a Conn simultaneously.
package main

import (
	"fmt"
	"io"
	"net"
)

var end = make(chan bool)

func main() {

	go func() {
		// client
		conn, _ := net.Dial("tcp", "localhost:8888")
		data := make([]byte, 1)
		for {
			// 如果对端没有写任何数据，但连接未关闭，则读阻塞
			// 对于阻塞的进一步解释：
			//		主协程因等待chan end而阻塞，client 协程因读数据而阻塞，但是系统并没有因此而发生panic，为什么？
			//		由于网络通信存在延迟，conn.Read造成的阻塞有可能在未来的某个时间自动unblocked;
			//		而且对端如果关闭连接，读操作也不再会被阻塞，而是出错返回
			n, err := conn.Read(data)
			if n > 0 {
				fmt.Printf("%v", string(data))
			} else {
				fmt.Println()
			}

			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}
		}

		end <- true
	}()

	Listener, _ := net.Listen("tcp", "localhost:8888")
	conn, _ := Listener.Accept()
	conn.Write([]byte("helloworld"))
	conn.Close()
	<-end
}
