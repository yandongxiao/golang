// tcp or udp
// For TCP and UDP networks, addresses have the form host:port.
// Dial("tcp", "192.0.2.1:80")	基本形式
// Dial("tcp", "golang.org:http")	间接进行域名解析
// Dial("tcp", "[fe80::1%lo0]:80")	同时支持ipv4和ipv6
// Dial("tcp", ":80") 以IPV4为例：0.0.0.0 as a target address, In practice connecting to 0.0.0.0 is equivalent to connecting to localhost.
//								  When binding, they’ll receive packets addressed to any IPv4 address on the system. 即绑定了该host的所有网络地址
//
// ip
// For IP networks, the network must be "ip", "ip4" or "ip6" followed by a colon and a protocol number or name
// the addr must be a literal IP address. 不支持域名解析
// Dial("ip4:1", "192.0.2.1")
// Dial("ip6:ipv6-icmp", "2001:db8::1")
//
// unix
// For Unix networks, the address must be a file system path.
//
// NOTE: If the host is resolved to multiple addresses, Dial will try each address in order until one succeeds.
package main

import (
	"fmt"
	"io"
	"net"
)

var end = make(chan bool)

func main() {

	// client
	go func() {
		// network: tcp, tcp4, tcp6, udp, udp4, udp6, ip, ip4, ip6, unix, unixgram, unixpacket
		conn, _ := net.Dial("tcp", "localhost:8888")
		data := make([]byte, 1)
		for {
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

	// server
	Listener, _ := net.Listen("tcp", "localhost:8888")
	defer Listener.Close() // Already Accepted connections are not closed.
	conn, _ := Listener.Accept()
	conn.Write([]byte("helloworld"))
	conn.Close()
	<-end
}
