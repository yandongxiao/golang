package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := 10
	// reflect package中两个重要的数据类型，reflect.Type和reflect.Value
	rv := reflect.ValueOf(v)
	rt := reflect.TypeOf(v)
	// reflect.Value 到 reflect.Type 的转换
	fmt.Println(rv.Type() == rt)
	// 类型检查
	fmt.Println(rv.Kind() == reflect.Int)
	fmt.Println(rt.Kind() == reflect.Int)
	// 转换为正常值
	new := rv.Int() // 返回的是int64，而非int
	fmt.Println(new, rv.Interface().(int))
}
