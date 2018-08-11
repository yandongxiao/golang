package main

import "fmt"
import "io/ioutil"

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// 与ReadAll相比，入参有区别
	entries, err := ioutil.ReadFile("/tmp/data")
	errcheck(err)
	fmt.Printf("%s\n", entries)
}
