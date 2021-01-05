// Package expvar provides a standardized interface to public variables,
// such as operation counters in servers. It exposes these variables via
// HTTP at /debug/vars in JSON format. 统计类信息使用expvar
//
// /debug/vars的输出有两个内置的记录：cmdline, memstats

package main

import (
	// The package is sometimes only imported for the side effect of
	// registering its HTTP handler and the above variables（指cmdline和memstats）
	"expvar"
	"fmt"
	"io"
	"net/http"
)

// hello world, the web server
var helloRequests = expvar.NewInt("hello-requests")

// Operations to set or modify these public variables are atomic.
func HelloServer(w http.ResponseWriter, req *http.Request) {
	helloRequests.Add(1)
	io.WriteString(w, "hello, world!\n")
}

func do(kv expvar.KeyValue) {
	key := kv.Key
	val := kv.Value
	fmt.Println(key, val)
}

func main1() {
	expvar.Do(do)
	//http.HandleFunc("/go/hello", HelloServer)
	//http.ListenAndServe(":8080", nil)
}
