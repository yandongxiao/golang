package main

import "fmt"

type Base struct {
	age int
}

type Child struct {
	Base
}

func ExampleInherit2() {
	child := Child{
		Base: Base{
			age: 20,
		},
	}

	// 解释为什么输出Base的值20？
	// 1. 直接调用foo(c)会出错，因为存在类型转换的问题
	// 2. golang中的继承的本质是composite
	// 所以child.Base是将它内部的Base变量赋值给foo的参数，
	// 所以肯定只会调用Base的方法
	fmt.Println(foo(child.Base))

	// 解释为什么输出20
	// child.getMore调用的本质是方法名寻找,
	// 调用对应field的方法的过程(而调用方法的本质是调用函数)
	child.getMore()

	// Output:
	// 20
	// 20
}

func foo(b Base) int {
	return b.get()
}

func (b Base) get() int {
	return b.age
}

func (c Child) get() int {
	return 200
}

func (b Base) getMore() {
	fmt.Println(b.get())
}
