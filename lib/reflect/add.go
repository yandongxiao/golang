package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Empty interface{}

// add 的参数是一个interface{}, 但是interface{}和数字无法直接相加，语法上就过不去
func add(v Empty) {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Interface ||
		rv.Kind() == reflect.Ptr {
		if !rv.CanSet() {
			rv = rv.Elem()
		}
	} else {
		panic(errors.New("the value can not be set"))
	}

	fmt.Printf("%T\n", rv.Interface())
	if rv.Kind() == reflect.Int ||
		rv.Kind() == reflect.Int8 ||
		rv.Kind() == reflect.Int16 ||
		rv.Kind() == reflect.Int32 ||
		rv.Kind() == reflect.Int64 ||
		rv.Kind() == reflect.Float32 ||
		rv.Kind() == reflect.Float64 {

		v := rv.Int()
		v++
		rv.SetInt(v)
	}
}

// 反射函数不是用来检查类型的，这也是reflect pacakge的很多函数都可能会panic。
// 参数检查是调用者的事情
func add2(v Empty) {
	rv := reflect.ValueOf(v)
	if !rv.CanSet() {
		rv = rv.Elem()
	}
	fmt.Printf("%T\n", rv.Interface())

	switch rv.Kind() {
	case reflect.Int:
		i := rv.Int()
		i++
		rv.SetInt(i)
	case reflect.Float64:
		f := rv.Float()
		f++
		rv.SetFloat(f)
	case reflect.Interface:
		v = reflect.ValueOf(rv.Interface()).Int() + 1
		rv.Set(reflect.ValueOf(v))
	}
}

type INT int

func main() {
	// 自定义
	var num INT = 1
	add(&num) // 传递指针
	fmt.Println(num)

	// 任何对象都存在它对应的指针类型，例如slice, map, interface
	// &errors.New("helloworld")这种创建指向interface的指针，在语法上是错误的。
	v := Empty(num)
	add2(&v)
	fmt.Println(v)
	fmt.Printf("%T %v", v, v) // interface v 指向了一个新的元素，因为之前的元素是没有办法改变的.
}
