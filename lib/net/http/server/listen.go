package main

import (
	"io"
	"log"
	"net"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hellWorld server")
}

func main() {
	/* lwayse the first step */
	http.HandleFunc("/hello", HelloServer)

	// 首先，创建用tcp协议监听8888端口
	l, e := net.Listen("tcp", ":9090")
	if e != nil {
		log.Fatal("Listen: ", e)
	}

	// 然后在监听的这个端口上启用http服务进行http服务
	err := http.Serve(l, nil)
	if err != nil {
		log.Fatal("Serve: ", err)
	}
}
