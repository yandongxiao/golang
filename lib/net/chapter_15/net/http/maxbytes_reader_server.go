package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// hello world, the web server
func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("come on")

	//	req.Body = http.MaxBytesReader(w, req.Body, 10)
	data, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(data))

	io.WriteString(w, "hello, world!hello, world!hello, world!hello, world!hello, world!\n")

}

func main() {

	// 指定当用户访问 http://www.xxx.com:mmmm/hello 的时候(注意，请不要在hello后面加上/变成hello/)
	// 调用HelloServer这个函数来处理
	http.HandleFunc("/hello", HelloServer)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
