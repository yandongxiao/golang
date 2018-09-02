// func ListenAndServe(addr string, handler Handler) error
// 1. Handler是一个接口，只包含一个方法。ServeHTTP(ResponseWriter, *Request)
// 2. 所以，该Handler相当于一个Router，由它来转发给相应的处理函数。
// 3. http.ServeMux就是一个Router

package main

import (
	"fmt"
	"net"
	"net/http"
)

func listenAndServe() {

	// 注册在DefaultServeMux内
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		// Accepted connections are configured to enable TCP keep-alives. 注意与http的Keep-Alive的区别
	})

	// Handler is typically nil, in which case the DefaultServeMux is used.
	// http.ListenAndServe的实现就两行代码：
	//	server := &Server{Addr: addr, Handler: handler}
	//	return server.ListenAndServe()
	http.ListenAndServe(":8080", nil)
}

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		rw.Header().Add("Connection", "close")
		rw.Write([]byte("hello"))
		// Accepted connections are configured to enable TCP keep-alives. 注意与http的Keep-Alive的区别
	})

	// The zero value for Server is a valid configuration.
	server := http.Server{}

	// 如果add不存在，则监听:80端口.
	// panic: listen tcp :80: bind: permission denied
	server.Addr = ":8080"

	// http.ServeMux如果匹配失败，不会去check什么当前文件系统是否存在同名文件
	// 直接返回404 not found
	server.Handler = http.DefaultServeMux

	// ReadHeaderTimeout is the amount of time allowed to read request headers.
	// 本质是是TCP连接conn.SetReadDeadline
	// The connection's read deadline is reset after reading the headers.
	// 重置：如果没有设置server.ReadTimeout，则调用c.r.setInfiniteReadLimit.
	//       如果设置了server.ReadTimeout，则设置读超时
	server.ReadHeaderTimeout = 0
	// 全局设置
	server.ReadTimeout = 0
	// 这是从读取完Header后，开始计时，这是与server.ReadTimeout的区别
	server.WriteTimeout = 0
	// IdleTimeout is the maximum amount of time to wait for the next
	// request when keep-alives are enabled.
	server.IdleTimeout = 0

	// new --> active --> close
	// new --> active --> idle
	server.ConnState = cb

	// ListenAndServe always returns a non-nil error. 因为平时不返回
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func cb(conn net.Conn, st http.ConnState) {
	fmt.Println(conn, st)
}
