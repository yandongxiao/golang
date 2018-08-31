package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

func main() {
	// NewRequest returns a Request suitable for use with Client.Do or Transport.RoundTrip.
	// If body is of type *bytes.Buffer, *bytes.Reader, or *strings.Reader, the returned request's ContentLength is set to its exact value (instead of -1))
	req, _ := http.NewRequest("POST", "http://localhost:8888", http.NoBody) // body=nil 也是可以的
	req.Body = ioutil.NopCloser(strings.NewReader("dd:wq"))

	// read and write request
	// Write writes an HTTP/1.1 request, which is the header and body, in wire format.
	buf := new(bytes.Buffer)
	req.Write(buf)
	req, _ = http.ReadRequest(bufio.NewReader(buf))
	url, _ := url.Parse("http://localhost:8888")
	req.URL = url
	req.RequestURI = "" // 必须的
	resp, err := http.DefaultClient.Do(req)
	fmt.Println(resp, err)
}

func serverHandler(resp http.ResponseWriter, req *http.Request) {
	// To create a request for use with testing a Server Handler, either use the NewRequest function
	// either use the NewRequest function in the net/http/httptest package, use ReadRequest, or manually update the Request fields.
	req = httptest.NewRequest("POST", "www.baidu.com", nil)
}
