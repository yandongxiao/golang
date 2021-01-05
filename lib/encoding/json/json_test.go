package main_test

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func ExampleIgnoreField() {
	type Persone struct {
		Name string
		Age  int `json:"-"`
	}

	p := Persone{
		Name: "jack",
		Age:  10,
	}

	// encode
	var buffer bytes.Buffer
	err := json.NewEncoder(&buffer).Encode(p)
	if err != nil {
		panic(err)
	}
	fmt.Print(buffer.String()) // 不能用fmt.Println

	var dp Persone
	err = json.NewDecoder(&buffer).Decode(&dp)
	if err != nil {
		panic(err)
	}
	fmt.Println(dp)

	// Output:
	// {"Name":"jack"}
	// {jack 0}
}
