package main

import (
	"fmt"
	"reflect"
)

// tag尽量要成pair出现. `json:"level"`
// omitempty表示在打印时若该项为empty则不打印，应将其放在双引号内. json:"max_instance,omitempty"
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	personType := reflect.TypeOf(Person{})
	for i := 0; i < personType.NumField(); i++ {
		fmt.Println(personType.Field(i).Tag)
	}
}
