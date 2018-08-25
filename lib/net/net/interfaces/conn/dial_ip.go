// 需要使用root权限运行
// port number:
// In the Internet Protocol version 4 (IPv4) [RFC791] there is a field
// called "Protocol" to identify the next level protocol.  This is an 8 bit field.
// https://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml
// 所以，你必须使用一个没有被占用的protocol number
package main

import (
	"bytes"
	"fmt"
	"net"
	"time"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// client
	go func() {
		// 为什么不是net.Dial? 它通过conn.Read读取到的信息包含了IP报文头部，不方便.
		// 为什么使用两个IP地址：
		//		1. 函数DialIP和ListenIP的参数，可以是ip4:50000, 也可以是ip4
		//		2. 所以如果是同一个IP地址，ip server在发送响应packet时，消息自然又会被server收到, 造成死循环
		laddr, _ := net.ResolveIPAddr("ip", "127.0.0.1")
		raddr, _ := net.ResolveIPAddr("ip", "10.245.208.17")
		conn, err := net.DialIP("ip4:200", laddr, raddr)
		//conn, err := net.Dial("ip4", "10.245.208.17")
		checkError(err)
		fmt.Println(conn.LocalAddr().Network(), conn.LocalAddr().String())
		data := make([]byte, 100)
		for {
			time.Sleep(time.Second)
			conn.Write([]byte("hello"))
			n, remoteAddr, err := conn.ReadFromIP(data)
			fmt.Printf("client read: %d, %s, %#v, %v\n", n, data[:n], remoteAddr.String(), err)
		}
	}()

	// server
	addr, err := net.ResolveIPAddr("ip", "10.245.208.17")
	checkError(err)
	conn, err := net.ListenIP("ip4:200", addr)
	checkError(err)
	for {
		data := make([]byte, 100)
		n, remoteAddr, err := conn.ReadFrom(data)
		fmt.Printf("server read: %d, %s, %#v, %v\n", n, data[:n], remoteAddr.String(), err)
		if bytes.Equal(data[:n], []byte("hello")) {
			conn.WriteToIP([]byte("world"), remoteAddr.(*net.IPAddr))
		} else {
			conn.WriteToIP([]byte("hello"), remoteAddr.(*net.IPAddr))
		}
	}
}
