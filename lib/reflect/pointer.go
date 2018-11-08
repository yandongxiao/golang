package main

import (
	"fmt"
	"reflect"
)

func main() {
	p := &struct {
		Name string
	}{
		Name: "jack",
	}

	fmt.Println(reflect.ValueOf(p).Kind())
	fmt.Println(reflect.ValueOf(p).CanSet()) // NOTE: false
}
