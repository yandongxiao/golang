// Trailer specifies additional headers that are sent after the request body.
// 有些Header可能需要边读请求体，边计算得到。但是发送完请求体后，和发送Trailer之间是如何控制的?
// 我们应该构建自己的Reader，作为参数传递给req.Body. 这样就可以控制Trailer的值了
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	var err error
	go server()
	time.Sleep(time.Millisecond * 100)
	req, err := http.NewRequest("POST", "http://localhost:8080?name=jack&age=1", strings.NewReader("helloworld"))
	if err != nil {
		panic(err)
	}

	// For client requests Trailer must be **initialized** to a map containing the trailer keys to later send.
	// The values may be nil or their final values.
	// 虽然Trailer的值虽然还不清楚，但是必须一开始就指定好key
	req.Trailer = http.Header{
		"XXX": []string{"xx", "yy"},
		"AA":  []string{"xx", "yy"},
	}

	// The ContentLength must be 0 or -1, to send a chunked request.
	req.ContentLength = -1

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)
}

func server() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		// For server requests the Trailer map initially contains only the trailer keys, with nil values.
		fmt.Println(req.Trailer)
		// After reading from Body returns EOF, Trailer can be read again and will contain non-nil values
		io.Copy(ioutil.Discard, req.Body)
		fmt.Println(req.Trailer)
		resp.Write([]byte("server: helloworld"))
	})
	http.ListenAndServe("localhost:8080", nil)
}
