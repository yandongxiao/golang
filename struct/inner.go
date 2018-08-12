package main

import "fmt"

func main() {
	fmt.Println(struct {
		name string
		age  int
	}{
		name: "jack",
		age:  10,
	})
}
