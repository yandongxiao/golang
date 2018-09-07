package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"

	"./calc"
)

func main() {
	// A server registers an object, making it visible as a service
	// with the name of the type of the object.
	// After registration, exported methods of the object will be accessible remotely.
	// A server may register multiple objects (services) of different types
	// but it is an error to register multiple objects of the same type.
	v := new(calc.Args)
	v.N = 10
	v.M = 10

	rpc.Register(v)
	rpc.HandleHTTP()

	// 以下是标准的HTTP服务
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)
}
