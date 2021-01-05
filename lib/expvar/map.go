package main

import (
	"expvar"
	"fmt"
)

func main3() {
	// 线程安全的
	person := expvar.NewMap("person")
	var name expvar.String
	var age expvar.Int
	person.Set("name", &name)
	person.Set("age", &age)
	name.Set("jack")
	person.Add("age", 100)
	person.Add("age", 100)
	fmt.Println(person.String())
}
