package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func showURL(w http.ResponseWriter, req *http.Request) {
	//	req.ParseForm()
	req.ParseMultipartForm(1024 * 1024) /* it will call parseform */

	fmt.Println("showURL")
	fmt.Println("path: ", req.URL.Path)
	fmt.Println("raw path: ", req.URL.RawPath) /* raw path is NULL */
	fmt.Println("url:", req.URL)               /* equal to PATH */
	fmt.Println("host:", req.URL.Host)         /* why host is empty */
	w.Write([]byte("helloworld"))

	fmt.Println("URI: ", req.RequestURI)        /* /show_url */
	fmt.Println("remoteaddr: ", req.RemoteAddr) /* 127.0.0.1:57119 */

	/* this is header info */
	fmt.Println("proto: ", req.Proto)
	fmt.Println("ProtoMajor: ", req.ProtoMajor)
	fmt.Println("ProtoMinor: ", req.ProtoMinor)

	fmt.Println("host:", req.Host)
	fmt.Println("method:", req.Method)
	fmt.Println("header:", req.Header) /* some few useless message */

	fmt.Println("Form: ", req.Form)
	fmt.Println("ContentLength: ", req.ContentLength)
	fmt.Println("PostForm: ", req.PostForm)
	fmt.Println("ParseMultipartForm: ", req.MultipartForm)
	fmt.Println("PostFormValue: ", req.PostFormValue("name"))
	data, _ := ioutil.ReadAll(req.Body)
	fmt.Println("body: ", string(data))
	fmt.Println("useragent: ", req.UserAgent())
	name, pass, ok := req.BasicAuth()
	if ok {
		fmt.Println("username:", name)
		fmt.Println("password:", pass)
	}
	fmt.Println("Referer:", req.Referer())
}

/* not important, offen empty array */
func show_transfer_encoding(w http.ResponseWriter, req *http.Request) {
	fmt.Println("show_transfer_encoding")
	fmt.Println(req.TransferEncoding)
	w.Write([]byte("helloworld"))
}

/* few client server or proxy support it */
func show_header(w http.ResponseWriter, req *http.Request) {
	fmt.Println("show_header")
	fmt.Println(req.Trailer)

}

func main() {
	http.HandleFunc("/show_url", showURL) /* you can use chinese */
	http.ListenAndServe("localhost:8080", nil)
}
