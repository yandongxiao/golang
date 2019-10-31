package main

import (
	"fmt"
	"unsafe"
)

// 1. these fields can be of any type, even structs themselves,
// functions or interfaces.
type SPerson struct {
	name string
	age  int
	get  func(*SPerson) int // 将函数作为一个field
	i    interface{}        // 将一个interface作为一个field
}

func ExampleInitialize() {
	// 四者等价
	p1 := &SPerson{
		name: "jack",
	}

	p2 := new(SPerson)
	p2.name = "jack"

	p3 := &SPerson{}
	p3.name = "jack"

	p4 := &SPerson{name: "jack"} // 推荐

	fmt.Println(p1, p2, p3, p4)
	// Output:
	// &{jack 0 <nil> <nil>} &{jack 0 <nil> <nil>} &{jack 0 <nil> <nil>} &{jack 0 <nil> <nil>}
}

// 3. Even when a struct contains other structs,
// structs form a continuous block in memory
// this gives a huge performance benefit.
// NOTE: The size of a struct type is the sum of the sizes of all its field types
// plus the number of some padding bytes. The padding bytes are used to align
// the memory addresses of some fields. We can learn padding and memory
// address alignments in a later article.
func ExampleContinuousBlock() {
	age := 10 // size of int == 8
	fmt.Println(unsafe.Sizeof(age))
	// 一个指针 + 一个size 各占8个字节
	name := "helloworld"
	fmt.Println(unsafe.Sizeof(name))
	// NOTE: struct的大小=每个field的大小的和.
	fmt.Println(unsafe.Sizeof(SPerson{}))
	// The size of a zero-field struct type is zero.
	fmt.Println(unsafe.Sizeof(struct{}{}))

	// Output:
	// 8
	// 16
	// 48
	// 0
}

func ExampleAccess() {
	p1 := SPerson{age: 10}
	p1.setAge(100) // 作用在p1上，p1是可访问的，所以合法
	fmt.Println(p1.getAge())

	p2 := &SPerson{age: 20}
	p2.setAge(200)
	fmt.Println(p2.getAge())

	// Output:
	// 100
	// 200
}

// when we have a struct type and define an alias type for it
// both types have the same underlying type and can be converted
// into one another,
func ExampleConversion() {
	n1 := number{1.0}
	n2 := NUMBER(n1) // 两个独立的value
	n1.f = 10.0
	fmt.Println(n1, n2)

	// cannot convert 2.3 (type float64) to type NUMBER)
	// n3 := NUMBER(2.3)
	// n4 := number(2.3)
	// fmt.Println(n3, n4)

	// OUtput:
	// {10} {1}
}

type number struct {
	f float32
}
type NUMBER number // alias type

func (p SPerson) getAge() int {
	return p.age
}

// NOTE: 如果再在pointer上重定义getAge
// 会导致method redeclared: SPerson.getAge错误
//func (p *SPerson) getAge() int {
//	return p.age
//}

func (p *SPerson) setAge(age int) {
	p.age = age
}

type Setter interface {
	setAge(age int)
}

func setAge(setter Setter) {
	setter.setAge(100)
}
