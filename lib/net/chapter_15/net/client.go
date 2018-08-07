package main

import "fmt"
import "net"

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
		_, err := conn.Read(data)
		errcheck(err)
		fmt.Println(string(data))
	}
}
