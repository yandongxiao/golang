// 注意TCP是流式传输，两次发送的内容，有可能会被一次性接收。反之亦然
package main

import (
	"flag"
	"fmt"
	"net"
	"syscall"
)

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		panic("usage: host port")
	}

	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	fmt.Println("hostAndPort: ", hostAndPort)
	tcpListener := initServer(hostAndPort)
	fmt.Println("listen to: ", tcpListener.Addr())

	for {
		conn, err := tcpListener.Accept()
		checkError(err)
		go connectionHandler(conn)
	}
}

func initServer(hostAndPort string) *net.TCPListener {
	var (
		tcpAddr     *net.TCPAddr
		err         error
		tcpListener *net.TCPListener
	)

	if tcpAddr, err = net.ResolveTCPAddr("tcp", hostAndPort); err != nil {
		panic(err.Error())
	}

	if tcpListener, err = net.ListenTCP("tcp", tcpAddr); err != nil {
		panic(err.Error())
	}
	return tcpListener
}

func connectionHandler(conn net.Conn) {
	defer conn.Close()

	fmt.Println(conn.RemoteAddr().String())
	sayHello(conn)
	const maxRead = 1024
	for {
		buf := make([]byte, maxRead+1)
		n, err := conn.Read(buf[:maxRead])
		switch err {
		case nil:
			fmt.Println(string(buf[:n]))
		case syscall.EAGAIN:
			continue
		default:
			break
		}
	}
}

func sayHello(conn net.Conn) {
	_, err := conn.Write([]byte("hello"))
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
