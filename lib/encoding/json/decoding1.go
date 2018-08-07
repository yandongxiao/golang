package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	Name string /* Why the field Name must be accessed */
	Ip   string
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	var slice ServerSlice
	str := `{"Servers":[{"Name":"shanghai", "Ip":"127.0.0.1"}]}`
	json.Unmarshal([]byte(str), &slice)

	fmt.Println(slice)
}
