package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// int既然是predeclared identifiers，那么我们可以定义重名的identifier，隐藏predeclared identifiers
	int := 10
	fmt.Println(int)

	p := Person{
		Name: "jack",
		Child: Child{
			Age: 10,
		},
	}

	data, _ := json.Marshal(p)
	fmt.Println(string(data))
}

type Person struct {
	Name string
	// NOTICE: 如果你也想暴露child，可以将field name定义为Child
	Child Child
}

type Child struct {
	Age int
}
