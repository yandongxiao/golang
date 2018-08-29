// hello_world_webserver.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

// A web-address is represented by the type http.URL which has a Path field that contains the url as a string; client-requests are described by the type http.Request, which has a URL field.
// request.ParseForm(),  request.Form[“var1”]
// The Form field is in fact of type map[string][]string.
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello, %s ", req.URL.Path)
}

func main() {
	http.HandleFunc("/", HelloServer)
	// err := http.ListenAndServe("localhost:8080", nil)
	err := http.ListenAndServe("localhost:8080", http.HandlerFunc(HelloServer))
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
	// http.ListenAndServe(":8080", http.HandlerFunc(HelloServer))
}
