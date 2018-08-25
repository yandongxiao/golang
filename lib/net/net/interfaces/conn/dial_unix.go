package unix

import (
	"fmt"
	"net"
)

func main() {

	// client
	/*
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
	*/
	// server
	// 借助linux命令nc，可以模拟客户端.nc -p 31337 -w 5 host.example.com 42
	Listener, _ := net.Listen("unix", "/tmp/aaabc.socket")
	defer Listener.Close() ///tmp/aaabc.socket
	conn, _ := Listener.Accept()
	for {
		data := make([]byte, 100)
		conn.Write([]byte("helloworld\n"))
		_, err := conn.Read(data)
		if err != nil {
			break
		}
		fmt.Printf("%s\n", data)
	}
	conn.Close()
}
