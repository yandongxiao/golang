package main

import "fmt"

// NOTE: struct不但可以包含匿名的struct，而且也可以包含匿名的interface
// 在 struct 中添加 interface 作为成员的技术，叫做 call dispatch
//
// 如果 House 实现了 getAge 方法，h.getAge的调用肯定会调用该实现
type House struct {
	AgeInterface
	AnonymousPerson
}

type AgeInterface interface {
	getAge() int
}

// 解释：为什么Person不能实现getAge接口?
type AnonymousPerson struct {
	age int
}

type AnonymousDog struct {
	age int
}

func (d AnonymousDog) getAge() int {
	return d.age
}

func ExampleAmbiguous() {
	h := House{
		AnonymousPerson: AnonymousPerson{10},
	}
	fmt.Println(h)

	h.AgeInterface = AnonymousDog{20}
	fmt.Println(h)

	foo := func(i AgeInterface) {
		fmt.Println(i.getAge())
	}
	foo(h)

	// Output:
	// {<nil> {10}}
	// {{20} {10}}
	// 20
}
