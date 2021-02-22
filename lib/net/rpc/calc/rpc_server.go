// Package main
package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"

	"github.com/yandongxiao/go/lib/net/rpc/calc"
)

func main() {
	v := new(calc.Args)
	v.N = 10
	v.M = 10

	// rpc.Register函数调用会将对象类型中所有满足RPC规则的对象方法注册为RPC函数
	// A server registers an object, making it visible as a **service** with the name of
	// the type of the object(对象的类型的名称，比如Args).
	// After registration, exported methods of the object will be accessible remotely.
	// A server may register multiple objects (services) of different types
	// but it is an error to register multiple objects of the same type.
	rpc.Register(v)
	rpc.HandleHTTP()

	// NOTE: RPC和HTTP服务是可以在一个应用中共存的
	// 通过前缀/_goRPC_/来区分RPC和HTTP两种服务
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	http.Serve(listener, nil)
}
