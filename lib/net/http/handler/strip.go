package main

import "net/http"

func main() {
	http.Handle("/tmp/", http.StripPrefix("/tmp/", http.FileServer(http.Dir("/tmp"))))
	http.ListenAndServe(":8080", nil)
}
