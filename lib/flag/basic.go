// ./basic --name='alice'
package main

import (
	"flag"
)

var (
	name string
	age  int
	sex  bool
)

func init() {
	flag.IntVar(&age, "age", 1, "your age")
	flag.StringVar(&name, "name", "jack", "your name")
	flag.BoolVar(&sex, "sex", false, "your sex")
}

func main() {
	flag.Parse()
	println("name=", name)
	println("age=", age)
	println("sex=", sex)
}
