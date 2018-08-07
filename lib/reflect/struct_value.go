package main

import (
	"fmt"
	"io"
	"reflect"

	"./role"
)

type INT int

type Person struct {
	Name   string `json:"Name"`
	age    int
	role   role.Role
	reader io.Reader
}

func (p *Person) GetName() string {
	return p.Name
}

func main() {
	p := &Person{
		Name: "jack",
		age:  10,
	}

	// 获取类型信息
	rv := reflect.ValueOf(p)

	// panic: reflect: call of reflect.Value.NumField on ptr Value
	rv = rv.Elem()

	for i := 0; i < rv.NumField(); i++ {
		new := rv.Field(i) // NOTE: 返回的是新的reflect.Value类型，而rt.Field()返回的是一个StructField类型
		// It panics if the Value was obtained by accessing unexported struct fields.
		//fmt.Println(new.Interface())
		switch new.Kind() {
		case reflect.String:
			fmt.Println("string", new.String())
		case reflect.Int:
			fmt.Println("int", new.Int()) // 获取unexported的值的方法
			// panic: reflect: reflect.Value.SetInt using value obtained using unexported field
			// new.SetInt(1000)
		case reflect.Struct:
			fmt.Println("struct")
		case reflect.Interface:
			fmt.Println("interface")
		}
	}

	rv = reflect.ValueOf(p)
	method := rv.Method(0)

	// 如何遍历一个方法
	fmt.Println(method.Type().NumIn())
	fmt.Println(method.Call([]reflect.Value{}))
}
