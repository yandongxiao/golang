package main

import (
	"fmt"
	"log"
	"net/rpc"

	"./calc"
)

const serverAddress = "localhost"

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}

	// call Multiply
	args := &calc.Args{7, 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d\n", args.N, args.M, reply)

	// call Add
	err = client.Call("Args.Add", 10, &reply)
	if err != nil {
		log.Fatal("add error", err)
	}
	fmt.Printf("add(remote object, 10) = %d", reply)

	// async
	divCall := client.Go("Args.Add", 20, &reply, nil)
	<-divCall.Done
	fmt.Printf("add(remote object, 20) = %d", reply)
}
