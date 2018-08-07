// the values must be given in field-order. 按照顺序初始化
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	p := Person{
		"jack", // 与field定义的顺序保持一致
		10,
	}
	fmt.Println(p)
}
