package main

import (
	"fmt"
	"net/url"
)

func main() {
	// testParse()
	testParsePathChinese()
}

func testParse() {

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

func assertEqual(actual, expect interface{}) {
	if actual != expect {
		errmsg := fmt.Sprintf("expect: <%s>, actual: <%s>", expect, actual)
		panic(errmsg)
	}
}

func testParsePathChinese() {
	u, err := url.Parse("http://你好/中国/123?name=%E4%B8%AD%E5%9B%BD")
	assertEqual(err, nil)

	fmt.Printf("Host: %s\n", u.Host)
	fmt.Printf("PATH: %s\n", u.Path)
	// In general, code should call EscapedPath instead of reading u.RawPath directly.
	fmt.Printf("Raw PATH: %s\n", u.RawPath)
	fmt.Printf("RequestURI: %s\n", u.RequestURI())
	fmt.Printf("EscapePath: %s\n", u.EscapedPath())
	fmt.Printf("String: %s\n", u.String())
}
