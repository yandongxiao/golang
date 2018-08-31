package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	go server()
	time.Sleep(100 * time.Millisecond)
	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)

	// Context returns the request's context. To change the context, use WithContext.
	// The returned context is always non-nil; it defaults to the background context.
	// For outgoing client requests, the context controls cancelation.
	// For incoming server requests, the context is canceled when the
	//   client's connection closes, the request is canceled (with HTTP/2), or when the ServeHTTP method returns.
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func server() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		time.Sleep(time.Second)
		resp.Write([]byte("server: helloworld"))
	})
	http.ListenAndServe("localhost:8080", nil)
}
