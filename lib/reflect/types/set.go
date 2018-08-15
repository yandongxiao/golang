package main

import (
	"fmt"
	"reflect"
)

type Empty interface{}

func set(v Empty, field string, newVal Empty) {
	rv := reflect.ValueOf(v)
	if !rv.CanSet() {
		rv = rv.Elem()
	}
	rv.FieldByName(field).Set(reflect.ValueOf(newVal))
}

type Person struct {
	Name string
	Age  int
}

func main() {
	var p Person
	set(&p, "Name", "jack")
	set(&p, "Age", 10)
	fmt.Println(p)
}
