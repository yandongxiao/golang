package main

import "net/http"

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		http.Error(rw, "zhibudao", 400)
	})

	http.HandleFunc("/404", func(rw http.ResponseWriter, req *http.Request) {
		http.NotFound(rw, req)
	})

	http.ListenAndServe(":8080", nil)
}
