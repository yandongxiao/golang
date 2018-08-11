package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	Name string /* Why the field Name must be accessed */
	IP   string
	Age  int
	D1   *int /* nil */
	D2   *int
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	val := 2
	slice := ServerSlice{
		[]Server{{"shanghai", "127.0.0.1", 10, nil, &val},
			{"beijing", "127.0.0.2", 20, nil, &val}}}

	data, _ := json.Marshal(slice)
	fmt.Printf("%s\n", data)
}
