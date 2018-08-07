package main

import "fmt"

type ST struct {
	name string
	age  int
}

func (st *ST) Name() {
	fmt.Println(st.name)
}

func main() {
	person := &ST{"niha", 10}
	person.Name()
}
