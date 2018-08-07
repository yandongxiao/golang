package main

import "fmt"

type Server struct {
	name string
	ip   string
}

type ServerSlice struct {
	servers []Server
}

func main() {
	data1 := &Server{"nihao", "127"}
	fmt.Println(data1)

	data2 := []Server{{"nihao", "127"}, {"nihao", "128"}}
	fmt.Println(data2)

	data3 := ServerSlice{[]Server{{"nihao", "127"}, {"nihao", "128"}}}
	fmt.Println(data3)
}
