// NOTE: nc是一个非常简单的可以模拟TCP/UDP的客户端或者服务端的命令
package main

import (
	"fmt"
	"net"
)

func main() {
	// server
	// 借助nc模拟客户端. nc -U /tmp/aaabc.socket
	listener, err := net.Listen("unix", "/tmp/aaabc.socket") // 如果文件已存在则返回失败
	if err != nil {
		fmt.Println(err)
	}
	// 退出时删除该文件tmp/aaabc.socket, 否则
	// listen unix /tmp/aaabc.socket: bind: address already in use
	// A SIGHUP, SIGINT(CTRL+C), or SIGTERM signal causes the program to exit. 所以直接CTRL+C时，该文件不会被删除
	defer listener.Close()

	conn, _ := listener.Accept()
	for {
		data := make([]byte, 100)
		conn.Write([]byte("helloworld\n"))
		_, err := conn.Read(data)
		if err != nil {
			break
		}
		fmt.Printf("%s", data)
	}
	conn.Close()
}
