// simple_webserver.go
package main

import (
	"io"
	"net/http"
)

const form = `<html><body><form action="#" method="post" name="bar">
		      <input type="text" name="in"/>
			  <input type="submit" value="Submit"/>
			  </form></html></body>`

/* handle a simple get request */
func SimpleServer(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
}

// 在一个独立的协程当中工作，但这个函数绝对不是协程的开始
// response.Body.Close如果没有被调用，连接就不会被复用；而且对端如果关闭了连接，连接状态将会变为CLOSE_WAIT状态
// 如Handler在处理请求的同时，需要作为客户端，请求另外的服务B。则，，
// 1. http.Client不能在Handler里面创建，这样每次就会建立一个新的连接到服务B；而连接默认是长链接，即便是调用resp.Resp.Close也只是将连接放到了连接池当中. 解决办法就是共享一个http.Client变量
// 2. handler中希望使用短连接, 则需要设置http.Client.Transpor.DisableKeepAlives == true，那么resp.Body.Close被调用时，将直接关闭该连接
//
func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		/* display the form to the user */
		io.WriteString(w, form)
	case "POST":
		request.ParseForm()
		//io.WriteString(w, request.Form["in"][0])
		io.WriteString(w, request.FormValue("in")) //更加值的调用的方法
	}
}

func main() {
	// 此种方法是将路由回调函数，注册到了默认路由容器DefaultServeMux中
	http.HandleFunc("/test1", SimpleServer) //跟RPC类似的注册方法
	http.HandleFunc("/test2", FormServer)

	// nil表示希望使用默认路由容器DefaultServeMux
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}
