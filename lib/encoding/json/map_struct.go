package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	Name string
	Ip   string
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	var slice Server
	str := `{"name":"shanghai", "ip":"127.0.0.1"}`
	json.Unmarshal([]byte(str), &slice)
	fmt.Println(slice)

	data2 := map[string]string{}
	err := json.Unmarshal([]byte(str), &data2)
	if err != nil {
		panic(err)
	}
	fmt.Println(data2)
}
