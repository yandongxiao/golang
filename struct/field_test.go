package main

import (
	"encoding/json"
	"fmt"
)

type Child struct {
	Age int
}

type Person struct {
	Name string
	// NOTE: 如果你也想暴露child，可以将field name定义为Child
	Child Child
}

type Person2 struct {
	Name string
	Child
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

	p2 := Person2{
		Name: "jack",
		Child: Child{
			Age: 10,
		},
	}
	data, _ = json.Marshal(p2)
	fmt.Println(string(data))
	// Output:
	// {"Name":"jack","Child":{"Age":10}}
	// {"Name":"jack","Age":10}
}
