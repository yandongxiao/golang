package main

import "net/http"

func main() {
	http.Handle("/", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
