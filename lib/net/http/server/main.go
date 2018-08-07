package main

import "net/http"

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header["X-Cos-Nihao"] == nil {
		w.Write([]byte("nil"))
	} else if r.Header.Get("x-cos-nihao") == "" {
		w.Write([]byte("empty"))
	} else {
		w.Write([]byte(r.Header.Get("x-cos-nihao")))
	}
}

func main() {
	http.Handle("/", &helloHandler{})
	http.ListenAndServe(":8989", nil)
}
