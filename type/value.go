package main

import "fmt"

func main() {
	v := new(int)
	*v = 10
	fmt.Println(*v)
}
