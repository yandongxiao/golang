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

func (h Handler) GET(a *Aaa, ptr *Aaa) string {
	return "OK" + a.a + " ptr:" + ptr.a
}

func main() {
	handler := Handler{}
	s := reflect.ValueOf(handler)
	method := s.MethodByName("GET")

	// reflect.Type可以作为map的key
	// 由于y将x覆盖掉了，说以说明reflect.TypeOf返回的对象其实是一个对象。或者说是两个相等的对象
	x := &Aaa{"x"}
	y := &Aaa{"y"}
	m := make(map[reflect.Type]interface{})
	m[reflect.TypeOf(x)] = x
	m[reflect.TypeOf(y)] = y

	input := make([]reflect.Value, 0)
	for i := 0; i < method.Type().NumIn(); i++ {
		input = append(input, reflect.ValueOf(m[method.Type().In(i)]))
	}

	out := method.Call(input)
	for i := 0; i < method.Type().NumOut(); i++ {
		fmt.Println(out[i].Interface())
	}
}
