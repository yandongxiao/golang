package main

import (
	"fmt"
	"io"
	"reflect"

	"./role"
)

type INT int

type Person struct {
	Name string `json:"Name"`
	INT
	role   role.Role
	reader io.Reader
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (p *Person) GetName() string {
	return p.GetName()
}

func main() {
	p := &Person{
		Name: "jack",
	}
	p.role = role.Role{0}
	p.INT = 0

	// 获取类型信息
	rt := reflect.TypeOf(p)
	fmt.Println(rt.Kind() == reflect.Ptr)
	rt = rt.Elem()
	fmt.Println(rt.Kind() == reflect.Struct)

	// 获取struct 名称
	fmt.Println(rt.Name() == "Person")

	// 遍历每一个元素的类型
	for i := 0; i < rt.NumField(); i++ {
		sf := rt.Field(i)
		// Anonymous: is an embedded field
		// sf.PkgPath 返回上层struct所在的package path
		// sf.Type.PkgPath() 返回该field类型的package path
		// sf.Type的类型是reflect.Type, 所以可以递归查找每一个类型
		fmt.Println(sf.Type, sf.Name, sf.Tag, sf.Offset, sf.Type.PkgPath(), sf.PkgPath, sf.Index, sf.Anonymous)
		if sf.Type.Kind() == reflect.Struct {
			rt := sf.Type
			for i := 0; i < rt.NumField(); i++ {
				sf := rt.Field(i)
				// If the type was predeclared (string, error) or unnamed (*T, struct{}, []int), the package path will be the empty string.
				fmt.Println("    ", sf.Type, sf.Name, sf.Tag, sf.Offset, sf.Type.PkgPath(), sf.PkgPath, sf.Index, sf.Anonymous)
			}
		}
	}

	rt = reflect.TypeOf(p)
	// rt = rt.Elem()
	for i := 0; i < rt.NumMethod(); i++ {
		sm := rt.Method(i)
		fmt.Println(sm)
	}
}
