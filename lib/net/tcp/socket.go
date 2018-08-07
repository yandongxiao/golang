package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	var (
		host   = "www.apache.org"
		port   = "80"
		remote = host + ":" + port
		msg    = "GET / \n"
		data   = make([]uint8, 40)
	)

	conn, err := net.Dial("tcp", remote)
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	n, err := io.WriteString(conn, msg)
	if err != nil {
		panic(err.Error())
	} else if n != len(msg) {
		panic("write not completely")
	}

	for {
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(string(data[:n]))
	}
}
