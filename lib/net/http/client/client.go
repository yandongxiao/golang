// A Client is higher-level than a RoundTripper (such as Transport)
//	and additionally handles HTTP details such as cookies and redirects.
// NOTICE：关于HTTP client request的修改，应该参考http.Request结构
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/cookiejar"
	"time"
)

// A Client is an HTTP client. Its zero value (DefaultClient) is a)
// usable client that uses DefaultTransport.
func zeroValueClient() {
	client := http.Client{}
	resp, err := client.Head("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	io.Copy(ioutil.Discard, resp.Body)
}

// CheckRedirect func(req *Request, via []*Request) error
// 默认策略: If CheckRedirect is nil, the Client uses its default policy, which is to stop after 10 consecutive requests.
// If CheckRedirect is not nil, the client calls it before following an HTTP redirect.
// The arguments req and via are the upcoming request and the requests made already, oldest first.
// If CheckRedirect returns an error, the Client's Get method returns both the previous Response (with its Body closed) and CheckRedirect's error (wrapped in a url.Error)
// As a special case, if CheckRedirect returns ErrUseLastResponse, then the most recent response is returned with its body unclosed, along with a nil error.
// 与http-core-module/redirect.conf搭配
func redirect(req *http.Request, via []*http.Request) error {
	fmt.Println(req.URL, via[0].URL)
	return nil
}

func doRedirect() {
	client := http.Client{
		// 用法, 参见roundTrip.go
		Transport:     http.DefaultTransport,
		CheckRedirect: redirect,
	}

	client.Get("http://localhost:8000/hello")
}

func doCookie() {
	// Set-Cookie: status=enable; expires=Tue, 05 Jul 2011 07:26:31 GMT;
	// path=/; domain=.hackr.jp;
	// The Jar is used to insert relevant cookies into every
	// outbound Request and is updated with the cookie values
	// of every inbound Response.
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}

	client := http.Client{
		Jar: jar,
	}

	resp, err := client.Get("http://localhost:8888")
	resp, err = client.Get("http://localhost:8888")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Header.Get("Set-Cookie"))
}

func doTimeout() {
	go func() {
		ls, _ := net.Listen("tcp", "localhost:8000")
		ls.Accept()
		time.Sleep(time.Second)
	}()

	client := http.Client{
		Timeout: time.Millisecond,
	}

	// Client.Timeout exceeded while awaiting headers
	resp, err := client.Get("http://localhost:8000")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(io.Copy(ioutil.Discard, resp.Body))
}

func main() {
	doTimeout()
}
