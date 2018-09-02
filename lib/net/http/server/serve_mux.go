// ServeMux is an HTTP request multiplexer.
// 原则：
//		1. Longer patterns take precedence over shorter ones
//		2. Host-specific patterns take precedence over general patterns.
//		3. ServeMux also takes care of sanitizing, 包括：url，header，repeated slashes等问题
//
// NOTE: ServeMux实现了Handler接口
//
package main

import (
	"fmt"
	"net/http"
)

func main() {
	server()
}

func server() {
	serveMux := http.NewServeMux()

	// 两种注册处理函数的方法
	// 必须在末尾加/，否则变成了完全匹配
	handlerFunc1 := func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("func world"))
	}
	serveMux.HandleFunc("/func/", handlerFunc1)
	fmt.Println("/func/", handlerFunc1)

	handlerFunc2 := func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("interface world"))
	}
	serveMux.Handle("/inter/", http.HandlerFunc(handlerFunc2))
	fmt.Println("/inter/", handlerFunc2)

	// Longer patterns take precedence over shorter ones
	serveMux.HandleFunc("/inter/hello/", handlerFunc1)
	req, _ := http.NewRequest("GET", "http://localhost:8080/inter/hello", nil)
	handler, pattern := serveMux.Handler(req)
	fmt.Println(pattern, handler) // 发起重定向

	req, _ = http.NewRequest("GET", "http://localhost:8080/inter/aa", nil)
	handler, pattern = serveMux.Handler(req)
	fmt.Println(pattern, handler)

	req, _ = http.NewRequest("GET", "http://localhost:8080/inter/hello/world", nil)
	handler, pattern = serveMux.Handler(req)
	fmt.Println(pattern, handler)

	// Host-specific patterns take precedence over general patterns.
	serveMux.HandleFunc("hello.com/", handlerFunc1)
	req, _ = http.NewRequest("GET", "http://localhost:8080/inter/hello/world", nil)
	req.Host = "hello.com"
	handler, pattern = serveMux.Handler(req)
	fmt.Println(pattern, handler)

}
