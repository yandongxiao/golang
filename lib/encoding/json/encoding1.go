package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	Name string /* Why the field Name must be accessed */
	IP   string
}

type ServerSlice struct {
	Servers []Server
}

func main() {
	slice := ServerSlice{[]Server{{"shanghai", "127.0.0.1"}, {"beijing", "127.0.0.2"}}}

	data, _ := json.Marshal(slice)
	fmt.Printf("%s\n", data)
}
