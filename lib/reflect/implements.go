package main

import (
	"fmt"
	"reflect"
)

type Person struct{}

func (p Person) set() {

}

type Setter interface {
	set()
}

func main() {
	p := Person{}
	s := reflect.TypeOf((*Setter)(nil)).Elem()
	fmt.Println(reflect.TypeOf(p).Implements(s))
	fmt.Println(reflect.TypeOf(&p).Implements(s))
}
