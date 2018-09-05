// All published variables.
// var (
//	  mutex   sync.RWMutex
//	  vars    = make(map[string]Var)
//	  varKeys []string // sorted
// )
// var helloRequests = expvar.NewInt("hello-requests") 也是注册在vars和varKeys中
package main

import (
	"bytes"
	"expvar"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type Counter struct {
	// NOTE: 这时候，就需要自己来确保线程安全了
	n int
}

// This makes Counter satisfy the expvar.Var interface.
// String returns a valid JSON value for the variable. 注意不是json格式
// 如果是string的话，需要是"string", 参考json.Marshal("nihao")
func (ctr *Counter) String() string { return fmt.Sprintf("%d", ctr.n) }

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET": // increment n
		ctr.n++
	case "POST": // set n to posted value
		buf := new(bytes.Buffer)
		io.Copy(buf, req.Body)
		body := buf.String()
		if n, err := strconv.Atoi(body); err != nil {
			fmt.Fprintf(w, "bad POST: %v\nbody: [%v]\n", err, body)
		} else {
			ctr.n = n
			fmt.Fprint(w, "counter reset\n")
		}
	}
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func main() {
	ctr := new(Counter)
	expvar.Publish("counter", ctr)
	http.Handle("/counter", ctr)
	http.ListenAndServe(":8080", nil)
}
