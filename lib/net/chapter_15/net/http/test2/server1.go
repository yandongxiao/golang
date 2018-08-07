package main

import (
	"net/http"
)

func sayHelloWorld(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("helloworld"))
}

func main() {
	http.HandleFunc("/test", sayHelloWorld)
	http.ListenAndServe("localhost:8080", nil)
}
