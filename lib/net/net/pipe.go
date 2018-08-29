package main

import (
	"fmt"
	"net"
)

func main() {
	// conn1, conn2 均可读写
	conn1, conn2 := net.Pipe()
	go func() {
		data := make([]byte, 1)
		for {
			fmt.Println(conn1.Read(data))
			fmt.Printf("%s", data)
		}
	}()

	// 为什么不用end := make(chan bool) 来表示结束？
	// Pipe creates a synchronous, in-memory, full duplex network connection;
	// Reads on one end are matched with writes on the other,
	// copying data directly between the two; there is no internal buffering.
	conn2.Write([]byte("helloworld"))
}
