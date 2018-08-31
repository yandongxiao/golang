package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func main() {
	go func() {
		http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
			// for log
			fmt.Println(req.RemoteAddr)

			// Usually the URL field should be used instead.
			// NOTE: It is an error to set this field in an HTTP client request.
			fmt.Println(req.RequestURI)

			fmt.Println(req.FormValue("name"))

			// ParseForm populates r.Form and r.PostForm.
			// For all requests, ParseForm parses the **raw query** from the URL and updates r.Form.
			// For POST, PUT, and PATCH requests, it also parses the request body as
			//	a form and puts the results into both r.PostForm and r.Form.
			// 所以，对于r.Form而言，请求体的Form值的优先级更高
			// For other HTTP methods, or when the Content-Type is not
			//	application/x-www-form-urlencoded, the request Body is not read, and
			//	r.PostForm is initialized to a non-nil, empty value.
			req.ParseForm()

			// Form contains the parsed form data, including both the URL field's query parameters and the POST or PUT form data.
			// This field is only available after ParseForm is called.
			// 直接从url中读取query可能存在不全的问题
			fmt.Println(req.Form)

			// PostForm contains the parsed form data from POST, PATCH, or PUT body parameters.
			// This field is only available after ParseForm is called.
			// 只包括body内的form，明确了每个key的来源
			fmt.Println(req.PostForm)

			fmt.Println(req.Host)

			resp.Write([]byte("server: helloworld"))
		})
		http.ListenAndServe("localhost:8080", nil)
	}()

	time.Sleep(time.Millisecond * 100)
	http.Post("http://localhost:8080/helloworld?name=jack", "application/x-www-form-urlencoded", strings.NewReader("age=10"))
}
