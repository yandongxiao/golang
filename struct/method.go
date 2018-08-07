// a Go method is a function that acts on variable of a certain type, called the receiver.
// ype分为两类，一种是value type(如type INT int), 另一种是pointer Type(如 type PINT *int).
// it cannot be a pointer type, but it can be a pointer to any of the allowed types. 所以，PINT类型是不允许被作为recevier的
// variable: 一个recevier可以是(a INT) 或 (p *INT) 两种变形，其中，(a INT)定义的函数也可以被pointer recevier使用.
//
// a method is a special kind of function. 注意，方法的本质是函数
// The receiver type can be (almost) anything, not only a struct type: any type can have methods, even a function type or alias types for int, bool, string or array.

package main

import "fmt"

type Person struct {
	name string
	age  int
}

// golang不允许有方法名重载，但是不同的类型可以有同名的方法
// 类型定义和方法定义无需在同一个文件内, 只需要在同一个package内，即可
func (p Person) getName1() string {
	return "object"
}

func (p *Person) getName2() string {
	return "pointer"
}

func (p Person) setName1(name string) {
	p.name = name
}

func (p *Person) setName2(name string) {
	p.name = name
}

func main() {
	p := Person{"jack", 10}
	// Go automatically handles conversion between values and pointers for method calls.
	println(p.getName1())
	println(p.getName2())
	println((&p).getName1())
	println((&p).getName2())

	// You may want to use a pointer receiver type to avoid copying on method calls
	// or to allow the method to mutate the receiving struct.
	p.setName1("value")
	fmt.Println(p)
	(&p).setName1("value") // 不起作用。method本质上是函数，receiver也是通过参数的形式传递的!!
	fmt.Println(p)
	p.setName2("pointer")
	fmt.Println(p)
}
