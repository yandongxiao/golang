package main

import (
	"io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {

	http.Redirect(w, req, "world", http.StatusFound)

}

func WorldServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "world server")

}

func main() {

	http.HandleFunc("/hello", HelloServer)
	http.HandleFunc("/world", WorldServer)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
