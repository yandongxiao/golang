package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	// 注意是：rpc.Dial
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	// client.Call
	var reply string
	err = client.Call("HelloService.Hello", "hello", &reply)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply)
}
