package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func errcheck(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	reader := strings.NewReader("hello world")

	// ReadAll reads from r until an error or EOF and returns the data it read.
	// A successful call returns err == nil, not err == EOF.
	data, err := ioutil.ReadAll(reader)
	errcheck(err)
	fmt.Printf("%s\n", data)
}
