package main

import "net"
import "fmt"
import "time"

func doClient(conn net.Conn) {
	data := make([]byte, 1024)
	for {
		conn.Read(data)
		fmt.Println("server: ", string(data))

		time.Sleep(time.Second)
		conn.Write([]byte("i am client"))
	}
}
func main() {
	/* c1 can both read and write data */
	c1, c2 := net.Pipe()
	go doClient(c1)

	data := make([]byte, 1024)
	for {
		c2.Write([]byte("I am  server"))
		c2.Read(data)
		fmt.Println("client:", string(data))
	}
}
