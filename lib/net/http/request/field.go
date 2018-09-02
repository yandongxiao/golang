package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	go server()
	time.Sleep(100 * time.Millisecond)

	var err error

	req := new(http.Request)
	// GET, POST, PUT 注意都是大写
	// For client requests an empty string means GET.
	req.Method = "POST"

	// url or uri
	// URL specifies either the URI being requested (for server requests) or the URL to access (for client requests).
	// For most requests, fields other than Path and RawQuery will be empty.
	// For client requests these fields are ignored.(上下文是关于HTTP version)
	req.URL, err = url.Parse("http://localhost:8080?name=jack&age=1")
	if err != nil {
		panic(err)
	}

	// For incoming requests, the Host header is promoted to the Request.Host field and removed from the Header map.
	// For client requests, certain headers such as Content-Length and Connection are automatically written when needed and
	// values in Header may be ignored.
	req.Header = http.Header{
		"XXX": []string{"xx"},
	}

	// for client
	// a nil body means the request has no body
	// The HTTP Client's Transport is responsible for calling the Close method.
	// for server
	// For server requests the Request Body is always non-nil but will return EOF immediately when no body is present.
	// The Server will close the request body. The ServeHTTP Handler does not need to.
	// 所以，无论对于client还是server，开发者都无需关闭Body
	req.Body = ioutil.NopCloser(strings.NewReader("helloworld"))

	// If Body is present, Content-Length is <= 0 and TransferEncoding hasn't
	// been set to "identity", Write adds "Transfer-Encoding: chunked" to the
	// header. Body is closed after it is sent.
	// req.ContentLength = int64(len("helloworld"))

	// 只需要client端设置
	// setting this field prevents re-use of TCP connections between requests to the same hosts
	req.Close = true

	// For client requests Host optionally overrides the Host header to send.
	// If empty, the Request.Write method uses the value of URL.Host.
	req.Host = "yisteng"

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Println(resp.Status)
}

func server() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		// 如果client端的请求, req.ContentLength没有被设置
		// server端获取到的req.TransferEncoding=[chunked]
		fmt.Println(req.TransferEncoding)
		resp.Write([]byte("server: helloworld"))
	})
	http.ListenAndServe("localhost:8080", nil)
}
