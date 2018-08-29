package main

import (
	"fmt"
	"net"
)

func main() {
	var listener net.Listener
	var err error

	if listener, err = net.Listen("tcp", "localhost:50000"); err != nil {
		panic("Listen: " + err.Error())
	}
	defer listener.Close()

	addr := listener.Addr()
	fmt.Println("Listen on ", addr)

	for {
		var conn net.Conn
		if conn, err = listener.Accept(); err != nil {
			panic("Accept: " + err.Error())
		}

		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 25)
		if n, err := conn.Read(buf); err != nil {
			panic(err)
		} else {
			fmt.Printf("%v", string(buf[:n]))
		}
	}
}
