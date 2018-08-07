package main

import "fmt"

type ST struct {
	num int
}

func main() {
	var i1, i2 interface{}
	i1 = ST{1}
	fmt.Println(i1)

	i2 = i1
	val := i2.(ST)
	val.num = 2000
	fmt.Println(i1)
	fmt.Println(i2)
	fmt.Println(val)
}
