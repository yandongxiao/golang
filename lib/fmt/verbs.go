package main

import "fmt"

type Person struct {
	name string
	age  int
}

func general() {
	p := Person{"jack", 10}

	// {jack 10}
	fmt.Printf("%v\n", p)
	// {name:jack age:10}
	fmt.Printf("%+v\n", p)
	// main.Person{name:"jack", age:10}
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	// 打印百分号
	fmt.Printf("%%\n")
}

func integer() {
	fmt.Printf("%t", true)
	fmt.Printf("%b\n", 10)
	fmt.Printf("%o\n", 10)
	fmt.Printf("%d\n", 10)
	fmt.Printf("%x\n", 10)
	fmt.Printf("%X\n", 10)
}

func float() {
	fmt.Printf("%e\n", 12.345)
	fmt.Printf("%E\n", 12.345)
	fmt.Printf("%f\n", 12.345)
	fmt.Printf("%F\n", 12.345) // 与%f完全一样
	// %g  %e for large exponents, %f otherwise
	fmt.Printf("%g\n", 12.345)
	fmt.Printf("%G\n", 12.345)
}

// String and slice of bytes
func slice() {
	fmt.Printf("%q\n", 97)
	fmt.Printf("%s\n", []byte{97})
	fmt.Printf("%q\n", []byte{97})
	fmt.Printf("%x\n", "hello")
	fmt.Printf("%X\n", "hello")
}

// For compound objects, the elements are printed using these rules, recursively, laid out like this:
// struct:				{field0 field1 ...}
// array, slice:		[elem0 elem1 ...]
// maps:				map[key1:value1 key2:value2]
// pointer to above:	&{}, &[], &map[]
func compound() {
	fmt.Printf("%#v\n", map[int]string{1: "hello", 2: "world"})
	fmt.Printf("%#v\n", struct {
		name string
		age  int
	}{
		name: "jack",
		age:  10,
	})
}

func main() {
	compound()
}
