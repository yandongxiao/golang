// package url 虽然定义了type Error struct，即拥有自己的Error类型
// NOTE: 它并非是interface
// 它实现的接口是net.Error
// scheme://[userinfo@]host/path[?query][#fragment]
package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

func main() {
	go func() {
		http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
			// Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/.
			//通过Path的值不能Escape原始路径值（Rawpath）主要是无法确定RawPath中的/就是/还是%2f
			fmt.Println("server: ", req.URL.Path) // /hello/你好

			// Go 1.5 introduced the RawPath field to hold the encoded form of Path.
			fmt.Println("server: ", req.URL.RawPath) // /hello%2F%E4%BD%A0%E5%A5%BD

			// 如果：url=/hello/你好, 那么String()=/hello/%E4%BD%A0%E5%A5%BD 注意hello后面是一个斜杠
			// URL's String method uses RawPath if it is a valid encoding of Path,
			// by calling the EscapedPath method 这个方法不会对+和/进行编码
			// NOTICE: 使用RawPath的前提是，它是一个有效的Path的编码.
			fmt.Println("server: ", req.URL) // /hello%2F%E4%BD%A0%E5%A5%BD
		})
		http.ListenAndServe("localhost:8080", nil)
	}()

	// 如何保证/不被编码? 使用方法url.EscapedPath
	// In general, code should call EscapedPath instead of reading u.RawPath directly.
	time.Sleep(time.Millisecond)
	http.Get("http://localhost:8080/" + url.PathEscape("hello/你好"))
	http.Get("http://localhost:8080/" + "hello/你好")
	time.Sleep(time.Millisecond)

	// The Parse function sets both Path and RawPath in the URL it returns
	// NOTICE: hello后面必须是%2F，否者这就不是一个有效的url编码. RawPath就会为空
	URL, err := url.Parse("http://localhost:8080/hello%2F%E4%BD%A0%E5%A5%BD")
	if err != nil {
		panic(err)
	}

	fmt.Println("Host:", URL.Host)
	fmt.Println("host:", URL.Hostname())
	fmt.Println("port:", URL.Port())

	fmt.Println("path:", URL.Path)
	fmt.Println("rawPath:", URL.RawPath)
	// In general there are multiple possible escaped forms of any path.
	// EscapedPath returns u.RawPath when it is a valid escaping of u.Path
	fmt.Println("escapedPath:", URL.EscapedPath())

	fmt.Println("=========")
	escape()

	fmt.Println("=========")
	server()
}

func escape() {
	// 前两种Escape的方法的区别：'+' into ' '
	escape := url.PathEscape("+/你好")
	fmt.Println("escape +你好:", escape)
	unescape, _ := url.PathUnescape(escape)
	fmt.Printf("unescape: %s: %s\n", escape, unescape)

	escape = url.QueryEscape("+/你好")
	fmt.Println("escape +你好:", escape)
	unescape, _ = url.QueryUnescape(escape)
	fmt.Printf("unescape: %s: %s\n", escape, unescape)

	// 对+和/, 一概不进行编码
	url := url.URL{
		Path: "+/你好",
	}
	fmt.Println(url.EscapedPath())
}

// server端相关的函数或方法
func server() {
	URL, _ := url.ParseRequestURI("/aaa/bbb")
	fmt.Println(URL.IsAbs()) // 测试URL是否是一个完整的URL, 从schema开始
	fmt.Println(URL)
}
