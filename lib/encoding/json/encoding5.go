package main

import (
	"encoding/json"
	"fmt"
)

/* This is not OK */
type Server struct {
	Name string /* Why the field Name must be accessed */
	IP   string
	Next *Server
	Prev *Server
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
	data1 := &Server{"shanghai", "127.0.0.1", nil, nil}
	data2 := &Server{"beijing", "127.0.0.2", nil, nil}
	data1.Next = data2
	data2.Prev = data1

	encode, _ := json.Marshal(data1)
	fmt.Printf("%s\n", encode)

	decodeServer := new(Server)
	json.Unmarshal(encode, decodeServer)
	fmt.Println(decodeServer)
}
