package main

import (
	"net/http"
)

func main() {
	// 1. 先Read req.Body；2. 再操作rw；3. 函数返回后，不应再操作rw和req
	// Except for reading the body, handlers should not modify the provided Request.
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		// the panic was isolated to the active request.
		//a := 0
		//v := 10 / a
		//rw.Write([]byte(string(v)))

		// To abort a handler so the client sees an interrupted response
		// but the server doesn't log an error, panic with the value ErrAbortHandler.
		panic(http.ErrAbortHandler)
	})

	http.ListenAndServe(":8080", nil)
}
