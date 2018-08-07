package main

import "net"
import "fmt"
import "time"

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

func handleconnection(conn net.Conn) {
	for {
		_, err := conn.Write([]byte("i am a server"))
		if err != nil {
			fmt.Println("failed to write data to client", err)
			conn.Close()
			break
		}
		time.Sleep(time.Second)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	errcheck(err)

	for {
		conn, err := ln.Accept()
		errcheck(err)
		fmt.Println(conn.LocalAddr()) /* ipv6 first */
		fmt.Println(conn.RemoteAddr())
		go handleconnection(conn)
	}
}
