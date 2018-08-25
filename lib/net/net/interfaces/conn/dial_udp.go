// 注意UDP协议并非是流式(stream-oriented)协议
package main

import (
	"fmt"
	"net"
	"time"
)

var end = make(chan bool)

func main() {

	// client
	go func() {
		conn, _ := net.Dial("udp", "localhost:8888")
		for {
			time.Sleep(time.Second)
			n, err := conn.Write([]byte("helloworld"))
			fmt.Printf("client write: %v, %v\n", n, err)
			data := make([]byte, 5)
			n, err = conn.Read(data)
			fmt.Printf("client read: %d, %s, %v\n", n, data[:n], err)
		}
	}()

	// server
	// 为什么不用net.Listen? 函数规范：The network net must be a stream-oriented network
	addr := net.UDPAddr{
		IP:   nil,  // ListenUDP listens on all available IP addresses of the local system except multicast IP addresses
		Port: 8888, // If the Port field of laddr is 0, a port number is automatically chosen.
	}
	conn, _ := net.ListenUDP("udp", &addr)
	for {
		// read first
		data := make([]byte, 5)
		n, remoteAddr, err := conn.ReadFromUDP(data)
		fmt.Printf("udpServer read: %d, %s, %#v, %#v\n", n, data[:n], remoteAddr.String(), err)

		// write
		n, err = conn.WriteToUDP([]byte("nihaoshijie"), remoteAddr)
		fmt.Printf("udpserver write, %v, %v\n", n, err)
	}
}
