// 调用GET方法
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

type Add struct{}

func (h Add) Add(a, b, c int) int {
	return a + b + c
}

// 调用对象object的方法
func call(object interface{}, methodName string, arguments ...interface{}) []reflect.Value {
	s := reflect.ValueOf(object)
	method := s.MethodByName(methodName) // NOTE: methodName的第一个字符必须是最大值

	input := make([]reflect.Value, 0, len(arguments))
	for _, arg := range arguments {
		input = append(input, reflect.ValueOf(arg))
	}
	fmt.Println(input)
	return method.Call(input)
}

func main() {
	handler := Handler{}
	output := call(handler, "GET", Aaa{"nihao"}, Bbb{3}, &Aaa{"pointer"})
	for _, rt := range output {
		fmt.Println(rt.Interface())
	}

	add := Add{}
	output = call(add, "Add", 1, 2, 3)
	for _, rt := range output {
		fmt.Println(rt.Interface())
	}
}
