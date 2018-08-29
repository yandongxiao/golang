package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// A Client is an HTTP client. Its zero value (DefaultClient) is a)
// usable client that uses DefaultTransport.
func zeroValueClient() {
	client := http.Client{}
	resp, err := client.Head("http://www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
	io.Copy(ioutil.Discard, resp.Body)
}

func main() {
	zeroValueClient()
}
