package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	go server()
	time.Sleep(100 * time.Millisecond)

	resp, _ := http.Get("http://localhost:8080")
	fmt.Println(resp.Header)
}

func server() {
	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		// Header
		// Trailer
		// 两个响应头：Trailer: Name, Trailer: Age
		rw.Header().Add("Trailer", "Name")
		rw.Header().Add("Trailer", "Age")
		// To suppress implicit response headers (such as "Date"), set their value to nil
		rw.Header()["Content-Type"] = nil
		rw.Header()["Date"] = nil

		// 将HTTP Header写下去
		// If WriteHeader is not called explicitly, the first call to Write
		// will trigger an implicit WriteHeader(http.StatusOK).)
		rw.WriteHeader(http.StatusOK)

		// Changing the header map after a call to WriteHeader has no effect
		// unless the modified headers are trailers.
		rw.Header().Add("Name", "jack")
		rw.Header().Add("Age", "10")

		// Body
		// Depending on the HTTP protocol version and the client, calling Write or WriteHeader may prevent future reads on the Request.Body.
		// For HTTP/1.x requests, handlers should read any needed request body data before writing the response.
		// Once the headers have been flushed (due to either an explicit Flusher.Flush call or writing enough data to trigger a flush) the request body may be unavailable.
		rw.Write([]byte("helloworld"))

	})
	http.ListenAndServe("localhost:8080", nil)
}
