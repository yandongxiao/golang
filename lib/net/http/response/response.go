package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	go server()
	time.Sleep(100 * time.Millisecond)

	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		panic(err)
	}

	// 协议版本、状态码
	fmt.Println(resp.Proto)
	fmt.Println(resp.ProtoMajor)
	fmt.Println(resp.ProtoMinor)
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)

	// 响应头部
	// Get返回Value数组([]string)中的第一个元素
	// Values duplicated by other fields in this struct (e.g., ContentLength) are omitted from Header.
	// Keys in the map are canonicalized.
	// 以Content-Type为例，单词之间以横杠连接，第一个字符大写。
	fmt.Println(resp.Header.Get("Xxx"))
	// -1 indicates that the length is unknown.
	fmt.Println(resp.ContentLength)
	// Transfer-Encoding的值：chunked,identity(不知道Body大小，但是仍是一次性传输所有数据)
	// 区别：http://www.cnblogs.com/jcli/archive/2012/10/19/2730440.html
	fmt.Println(resp.TransferEncoding)
	// HTTP/1.1 版本的默认连接都是持久连接。Server返回Connection:close时，resp.Close=true
	// NOTE: The value is **advice** for clients: neither ReadResponse nor Response.Write ever closes a connection.
	fmt.Println(resp.Close)

	// 响应体
	// The http Client and Transport guarantee that Body is always non-nil,
	// even on responses without a body or responses with a zero-length body.
	// NOTE：如果客户端也使用net/http，那么只有当服务端
	//       1.读取Body所有数据；2. 关闭连接，客户端才会reuse该连接
	io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()

	// Request is the request that was sent to obtain this Response.
	// Request's Body is nil
	// This is only populated for Client requests.
	fmt.Println(resp.Request)
}

func server() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		// NOTE: 发送的HTTP响应报文中，是Xxx
		resp.Header().Add("XXX", "xxx")
		resp.Header().Add("XXX", "yyy")
		resp.Header().Add("Connection", "close")

		// 如何以chunked的方式发送
		// The trick appears to be that you simply need to call Flusher.Flush() after each chunk is written.
		flush := resp.(http.Flusher)
		io.Copy(resp, strings.NewReader("hello"))
		flush.Flush()
		io.Copy(resp, strings.NewReader("world\n"))
	})
	http.ListenAndServe("localhost:8080", nil)
}
