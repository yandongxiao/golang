package main

import (
	"log"
	"net"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "index.data")
}

func main() {

	http.HandleFunc("/hello", HelloServer)

	l, e := net.Listen("tcp", ":9090")
	if e != nil {
		log.Fatal("Listen: ", e)
	}

	err := http.Serve(l, nil)
	if err != nil {
		log.Fatal("Serve: ", err)
	}
}
