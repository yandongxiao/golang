// 使用手册: go doc net/http/pprof | less
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// pprof.StartCPUProfile(w) CPU Profiling的时机是发送请求时
	// 在请求参数中带着seconds参数，表示CPU Profiling的执行时间
	// 返回CPU Profiling信息给客户端
	// go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
	// go tool pprof http://localhost:6060/debug/pprof/heap
	log.Println(http.ListenAndServe("localhost:6060", nil))
}
