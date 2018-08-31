package main

import (
	"net/http"
	"net/url"
)

func main() {
	req := new(http.Request)
	// GET, POST, PUT 注意都是大写
	// For client requests an empty string means GET.
	req.Method = "GET"

	// url or uri
	// URL specifies either the URI being requested (for server requests) or the URL to access (for client requests).
	// For most requests, fields other than Path and RawQuery will be empty.
	// For client requests these fields are ignored.(上下文是关于HTTP version)
	req.URL = &url.URL{
		Path:     "www.baidu.com",
		RawQuery: "name=jack&age=1",
	}

	// server:
	// For incoming requests, the Host header is promoted to the Request.Host field and removed from the Header map.
	// client:
	// For client requests, certain headers such as Content-Length and Connection are automatically written when needed and
	// values in Header may be ignored.
	req.Header = http.Header{
		"XXX": []string{"xx"},
	}

	// a nil body means the request has no body
	// The HTTP Client's Transport is responsible for calling the Close method.
	req.Body = nil
}
