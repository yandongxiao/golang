package main

import (
	"fmt"
	"log"
	"net/rpc"

	"./calc"
)

const serverAddress = "localhost"

func main() {
	// The convenience function Dial (DialHTTP) performs both steps for a raw network connection. (创建TCP连接，new http.Client)
	// The resulting Client object has two methods, Call(同步) and Go(异步), that specify the service and method to call, a pointer containing the
	// arguments, and a pointer to receive the result parameters. 请求参数不一定是指针形式吧
	// NOTE：rpc.DialHTTP
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

	// call Add, 使用注册对象的值
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
