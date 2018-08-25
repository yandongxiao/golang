package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// 限制连接建立的时间
	_, err := net.DialTimeout("tcp", "localhost:8888", time.Second)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ok")
	}
}
