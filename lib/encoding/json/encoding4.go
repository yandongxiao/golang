package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	Name string /* Why the field Name must be accessed */
	Ip   string
	Next *Server
}

func (server *Server) String() (output string) {
	for {
		output += server.Name + " " + server.Ip + "\n"

		if server.Next == nil {
			break
		}
		server = server.Next
	}
	return
}

func main() {
	data1 := &Server{"shanghai", "127.0.0.1", nil}
	data1.Next = &Server{"beijing", "127.0.0.2", nil}
	encode, _ := json.Marshal(data1)
	fmt.Println(string(encode))

	decodeServer := new(Server)
	json.Unmarshal(encode, decodeServer)
	fmt.Println(decodeServer)
}
