// The default HTTP client's Transport does not attempt to reuse HTTP/1.0 or HTTP/1.1 TCP connections
// unless the Response Body is read to completion and is closed
// 所以，如果client希望重用该连接，1.读取完毕数据，2.调用Close方法. 我们每个函数都准从了这两条
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	doPostForm()
	doPost()
	doGet()
	doHead()
}

// The http Client and Transport guarantee that Body is always
// non-nil, even on responses without a body or responses with a zero-length body.
// It is the caller's responsibility to close Body.
// 所以，即使是HEAD操作，也需要关闭连接
func doHead() {
	resp, err := http.Head("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func doGet() {
	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	//  The client must close the response body when finished with it
	defer resp.Body.Close()
	n, err := io.Copy(ioutil.Discard, resp.Body)
	fmt.Println(n, err)
}

func doPost() {
	resp, err := http.Post("http://localhost:8888",
		"image/jpeg", strings.NewReader("helloworld"))
	if err != nil {
		panic(err)
	}
	//  The client must close the response body when finished with it
	defer resp.Body.Close()
	n, err := io.Copy(ioutil.Discard, resp.Body)
	fmt.Println(n, err)
}

func doPostForm() {
	// The Content-Type header is set to application/x-www-form-urlencoded.
	resp, err := http.PostForm("http://localhost:8888",
		url.Values{"name": {"jack"}, "age": {"19", "20"}})
	if err != nil {
		panic(err)
	}
	//  The client must close the response body when finished with it
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s%v\n", data, err)
}
