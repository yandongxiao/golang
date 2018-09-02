// RoundTripper is an interface representing the ability to execute **a single HTTP transaction**
// RoundTrip(*Request) (*Response, error)
//
// type Transport struct 实现了接口RoundTripper， 而DefaultTransport是类型Transport的一个实例
//
// DefaultTransport is the default implementation of Transport and is used by DefaultClient
// It establishes network connections as needed and caches them for reuse by subsequent calls
//
// http.Get、http.Head, http.Post等函数, 内部实现都是调用DefaultClient的方法
//
// 综上：package http的关键数据结构之间的关系如下：
// 1. http.Client只是负责构造http.Request;
// 2. RoundTripper的实现负责管理所有的连接, 接收请求，返回响应http.Response
package main

import (
	"errors"
	"net/http"
)

func main() {
	customRoundTrip()
}

func defaultRoundTrip() {

	req, err := http.NewRequest("GET", "http://www.baidu.com", nil)
	if err != nil {
		panic(err)
	}

	// NOTE: 不建议直接使用, 使用DefaultClient
	// RoundTrip should not attempt to interpret the response. In
	// particular, RoundTrip must return err == nil if it obtained
	// a response, regardless of the response's HTTP status code.
	// Similarly, RoundTrip should not attempt to handle higher-level protocol details such as redirects, authentication, or cookies.
	// RoundTrip should not modify the request, except for consuming and closing the Request's Body. 所以我们不需要关闭http.Request的Body
	// 但是RoundTrip关闭Request.Body的实际可能是在函数返回之后
	resp, err := http.DefaultTransport.RoundTrip(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

func customRoundTrip() {
	client := http.Client{
		Transport: new(AuthTransport),
	}
	resp, err := client.Head("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}

type AuthTransport struct {
}

func (auth *AuthTransport) RoundTrip(req *http.Request) (resp *http.Response, err error) {
	name := req.Header.Get("name")
	pass := req.Header.Get("pass")
	if name != "root" || pass != "pass" {
		err = errors.New("please input your name and password")
		return
	}
	return http.DefaultTransport.RoundTrip(req)
}
