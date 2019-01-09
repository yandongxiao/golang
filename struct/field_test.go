package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	// NOTE: 如果你也想暴露child，可以将field name定义为Child
	Child Child
}

type Child struct {
	Age int
}

func ExampleFieldName() {
	p := Person{
		Name: "jack",
		Child: Child{
			Age: 10,
		},
	}

	data, _ := json.Marshal(p)
	fmt.Println(string(data))
	// Output:
	// {"Name":"jack","Child":{"Age":10}}
}
