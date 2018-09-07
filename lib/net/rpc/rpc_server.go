package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"./calc"
)

func main() {
	// A server registers an object(可以使用被注册的对象), making it visible as a **service**
	// with the name of the type of the object. (注册服务，名称为类型名称)
	// After registration, exported methods of the object will be accessible remotely.
	// A server may register multiple objects (services) of different types
	// but it is an error to register multiple objects of the same type.
	v := new(calc.Args)
	v.N = 10
	v.M = 10

	rpc.Register(v)
	rpc.HandleHTTP()

	// NOTE: RPC和HTTP服务是可以在一个应用中共存的
	// 通过前缀/_goRPC_/来区分RPC和HTTP两种服务（毕竟它们的协议是不同的, 但是协议的开头部分又是兼容的）
	// More typically it will create a network listener and call Accept or,
	// for an HTTP listener, HandleHTTP and http.Serve.
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	go http.Serve(listener, nil)
}
