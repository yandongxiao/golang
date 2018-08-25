// The method for resolving domain names varies by operating system.
// On Unix systems
// pure Go resolver: sends DNS requests directly to the servers listed in /etc/resolv.conf
// cgo-based resolver: calls C library routines such as getaddrinfo and getnameinfo.
// 区别： By default the pure Go resolver is used, because a blocked DNS request
//		  consumes only a goroutine, while a blocked C call consumes an operating system thread.
//
// export GODEBUG=netdns=go+1    # force pure Go resolver, 并输出信息
// export GODEBUG=netdns=cgo   # force cgo resolver

package main

import (
	"fmt"
	"net"
)

func main() {
	// 间接域名解析
	// NOTE: 如果/etc/hosts配置了域名地址，则尝试与baidu.com:80建立连接，连接建立过程中，程序被block住
	// To bypass the host resolver, use a custom Resolver.
	conn, _ := net.Dial("tcp", "baidu.com:80")
	fmt.Printf("net.Dial: %+v\n", conn.RemoteAddr())
	// NOTE：还存在逆向的函数: func JoinHostPort(host, port string) string)
	fmt.Println(net.SplitHostPort(conn.RemoteAddr().String()))

	// 直接域名解析
	addr, _ := net.ResolveTCPAddr("tcp", "baidu.com:80")
	fmt.Println("net.ResolveTCPAddr", addr)

	// LookupHost looks up the given host using the local resolver.
	addrs, _ := net.LookupHost("baidu.com")
	fmt.Println("net.LookupHost", addrs)

	// LookupAddr performs a reverse lookup for the given address
	// When using the host C library resolver, at most one result will be returned. 由此可见用的是pure go resolver
	addrs, _ = net.LookupAddr("127.0.0.1")
	fmt.Println("net.LookupAddr", addrs)

	// type IP []byte
	ips, _ := net.LookupIP("baidu.com")
	fmt.Println("net.LookupIP", ips)
}
