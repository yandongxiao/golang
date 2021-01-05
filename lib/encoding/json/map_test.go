package main

import (
	"encoding/json"
	"fmt"
)

type Server2 struct {
	Name string
	Ip   string
}

func ExampleJsonToMap() {
	str := `{"name":"shanghai", "ip":"127.0.0.1"}`
	data2 := map[string]string{}
	err := json.Unmarshal([]byte(str), &data2)
	if err != nil {
		panic(err)
	}
	fmt.Println(data2)

	// Output:
	// map[ip:127.0.0.1 name:shanghai]
}
