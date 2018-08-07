package main

import (
	"fmt"
	"reflect"
)

// 通过指针的形式修改
func test1() {
	x := 10
	// func ValueOf(i interface{}) reflect.Value
	rv := reflect.ValueOf(x)
	// x --> interface{} --> reflect. 因为x的值类型，拷贝时采用了值拷贝
	// reflect: reflect.Value.SetInt using unaddressable value
	// panic: reflect: call of reflect.Value.Elem on int Value
	// rv = rv.Elem()
	// rv.SetInt(20)

	rv = reflect.ValueOf(&x)
	rv = rv.Elem() // 必须的
	rv.SetInt(20)
	fmt.Println(rv.Interface(), x)
}

type Empty interface{}

func test2() {
	x := 10
	var v interface{} = &x // 也必须是指针形式
	rv := reflect.ValueOf(v)
	// To make it settable we need to let the Elem() function work on it which indirects through the pointer: v = v.Elem()
	// 1. 必须在SetInt之前调用；2. rv.Canset()在Elem()调用之后，会返回true, 之前为false。
	rv = rv.Elem() // 必须的
	rv.SetInt(20)
	fmt.Println(rv.Interface(), x)
}

func main() {
	test2()
}
