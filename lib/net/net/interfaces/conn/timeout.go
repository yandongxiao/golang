package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

var end = make(chan bool)

func main() {

	go func() {
		// client
		conn, _ := net.Dial("tcp", "localhost:8888")
		data := make([]byte, 1)
		for {
			// An idle timeout can be implemented by repeatedly extending
			// the deadline after successful Read or Write calls.
			conn.SetDeadline(time.Now().Add(time.Second))
			n, err := conn.Read(data)
			if n > 0 {
				fmt.Printf("%v", string(data))
			} else {
				fmt.Println()
			}

			if err == io.EOF {
				break
			} else if err != nil {
				// panic: read tcp 127.0.0.1:63149->127.0.0.1:8888: i/o timeout
				panic(err)
			}
		}

		end <- true
	}()

	Listener, _ := net.Listen("tcp", "localhost:8888")
	conn, _ := Listener.Accept()
	time.Sleep(time.Second * 2)
	conn.Write([]byte("helloworld"))
	conn.Close()
	<-end
}
