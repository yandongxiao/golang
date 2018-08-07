package main

import (
	"fmt"
	"reflect"
)

type Aaa struct {
	a string
}

type Bbb struct {
	b int
}

type Handler struct{}

func (h Handler) GET(a Aaa, b Bbb, ptr *Aaa) string {
	return "OK" + a.a + " ptr:" + ptr.a
}

func main() {
	handler := Handler{}
	s := reflect.ValueOf(handler)
	method := s.MethodByName("GET")

	x := Aaa{"x"}
	y := &Aaa{"y"}
	z := Bbb{2}
	m := make(map[reflect.Type]interface{})
	m[reflect.TypeOf(x)] = x
	m[reflect.TypeOf(y)] = y
	m[reflect.TypeOf(z)] = z

	input := make([]reflect.Value, 0)
	for i := 0; i < method.Type().NumIn(); i++ {
		input = append(input, reflect.ValueOf(m[method.Type().In(i)]))
	}

	out := method.Call(input)
	for i := 0; i < method.Type().NumOut(); i++ {
		fmt.Println(out[i].Interface())
	}
}
