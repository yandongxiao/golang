package main

import "fmt"
import "net/http"
import "io/ioutil"

func errcheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func method1() {
	resp, err := http.Get("http://www.baidu.com")
	errcheck(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	errcheck(err)

	fmt.Println(string(body))
}
