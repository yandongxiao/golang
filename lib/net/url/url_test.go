package main

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func assertEqual(actual, expect interface{}) {
	if actual != expect {
		errmsg := fmt.Sprintf("expect: <%s>, actual: <%s>", expect, actual)
		panic(errmsg)
	}
}

// positive case

//
// 目的：
//  熟悉Parse函数
//
// 场景：
//  调用Parse函数，并熟悉url.URL的各种方法，见TestParse
//  Path中含有中文字符, 见TestParsePathChinese
//  URL中含有编码过的字符，见TestParsePathEscaped
//
func TestParse(t *testing.T) {

	// Parse
	u, err := url.Parse("http://cos.chinac.com/aa/bb/cc?name=%E4%B8%AD%E5%9B%BD&age=10&name=dav")
	assertEqual(err, nil)

	// Scheme
	fmt.Println(u.Scheme)

	// Query
	for key, vals := range u.Query() {
		fmt.Printf("%v:", key)
		for _, val := range vals {
			fmt.Printf("  %v", val)
		}
		fmt.Println()
	}

	// Request URI 与 Request Path之间的区别
	fmt.Printf("request uri: %s\n", u.RequestURI())
	fmt.Printf("request path: %s\n", u.EscapedPath())

	// Host
	fmt.Printf("host: %v\n", u.Host)
}

// EscapePath的范围只限于Path部分
func TestParsePathChinese(t *testing.T) {
	u, err := url.Parse("http://你好/中国/123?name=%E4%B8%AD%E5%9B%BD")
	assertEqual(err, nil)

	fmt.Printf("Host: %s\n", u.Host)
	fmt.Printf("PATH: %s\n", u.Path)
	fmt.Printf("Raw PATH: %s\n", u.RawPath)
	fmt.Printf("RequestURI: %s\n", u.RequestURI())
	fmt.Printf("EscapePath: %s\n", u.EscapedPath())
	fmt.Printf("String: %s\n", u.String())
}

//u.Path 的内容已经是解码过的
func TestParsePathEscaped(t *testing.T) {
	u, err := url.Parse("http://你好/%E4%B8%AD%E5%9B%BD/123?name=小明")
	assertEqual(err, nil)

	fmt.Printf("Path: %s\n", u.Path)
	fmt.Printf("Raw Path: %s\n", u.RawPath)
}

func TestPathQueryEscaped(t *testing.T) {
	u, err := url.Parse("http://你好/%E4%B8%AD%E5%9B%BD/123?name=%E4%B8%AD%E5%9B%BDYY&age=10")
	assertEqual(err, nil)

	// Query
	for key, vals := range u.Query() {
		fmt.Printf("%v:", key)
		for _, val := range vals {
			fmt.Printf("  %v", val)
		}
		fmt.Println()
	}

	// Request URI 与 Request Path之间的区别
	querys, err := url.ParseQuery("name=%E4%B8%AD%E5%9B%BD&age=10")
	for key, vals := range querys {
		fmt.Printf("%s:", key)
		for _, val := range vals {
			fmt.Printf(" %s", val)
		}
		fmt.Println()
	}
}

//url package 的最主要作用是可以作为request.URL的值
//值的注意的是：网络传输过程中传输的是/,而不是%2f
func TestUsage(t *testing.T) {
	var err error
	request := http.Request{}
	request.URL, err = url.Parse("http://baidu.com/path%2fto%2f文件?name=火龙")
	assertEqual(err, nil)
	request.Method = http.MethodGet
	clinet := http.Client{}
	resp, err := clinet.Do(&request)
	//此种方法也行，也会主动进行escapse操作
	//resp, err := http.Get("http://baidu.com/path%2fto%2f文件?name=火龙")

	assertEqual(err, nil)
	fmt.Println(resp.StatusCode)
}

func TestOpaque(t *testing.T) {
	var err error
	request := http.Request{}
	u := url.URL{}
	u.Opaque = "/path%2fto%2f文件" //完全没有被编码
	u.Scheme = "http"
	u.Host = "baidu.com"
	query := u.Query()
	query.Set("name", "火龙")

	request.URL = &u
	assertEqual(err, nil)
	request.Method = http.MethodGet
	clinet := http.Client{}
	resp, err := clinet.Do(&request)
	//此种方法也行，也会主动进行escapse操作
	//resp, err := http.Get("http://baidu.com/path%2fto%2f文件?name=火龙")

	assertEqual(err, nil)
	fmt.Println(resp.StatusCode)
}
