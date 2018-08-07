package main

import "fmt"
import "net"
import "time"

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	errcheck(err)

	data := make([]byte, 1024)
	for {
		/* reset it when time expire, or conn.Read will alwayse fail */
		conn.SetDeadline(time.Now().Add(time.Second))
		_, err := conn.Read(data)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(data))
		}
	}
}
