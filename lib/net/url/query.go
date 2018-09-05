package main

import (
	"fmt"
	"net/url"
)

func main() {
	// type Values map[string][]string
	// 从server的角度来看，第一步应该是执行QueryUnescape解码，接下来再是解析
	escape := url.QueryEscape("name=jack&name=bob&age=10&from=北京")
	fmt.Println("escape +你好:", escape)
	unescape, _ := url.QueryUnescape(escape)
	fmt.Printf("unescape: %s: %s\n", escape, unescape)

	values, err := url.ParseQuery(unescape)
	if err != nil {
		panic(err)
	}
	for key, vals := range values {
		fmt.Println(key, vals)
	}
}
