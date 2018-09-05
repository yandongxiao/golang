package main

import (
	"expvar"
	"fmt"
	"time"
)

// NOTE: 返回值必须是interface{}
// 之所以是interface{}, 是为了适配json.Marshal接口.
func myFunc() interface{} {
	// 直接返回字符串即可
	return time.Now().String()
}

func main() {
	expvar.Publish("func", expvar.Func(myFunc))
	fmt.Println(expvar.Get("func"))
}
