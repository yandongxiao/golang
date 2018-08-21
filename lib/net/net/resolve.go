package main

import (
	"fmt"
	"net"
)

func main() {
	// 间接域名解析
	addr, _ := net.ResolveTCPAddr("tcp", "baidu.com:80")
	fmt.Println(addr)

	// 直接域名解析
	addrs, _ := net.LookupHost("baidu.com")
	fmt.Println(addrs)
}
