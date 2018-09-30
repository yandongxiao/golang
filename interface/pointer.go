package main

import "fmt"

type Person struct {
	name string
}

func (p Person) get() string {
	return p.name
}

func (p *Person) set(name string) {
	p.name = name
}

func main() {
	p := Person{"jack"}
	p.get()
	p.set("lua")
	(&p).get()
	(&p).set("lua")
	fmt.Println(p)

	getter := func(getter interface {
		get() string
	}) {
		fmt.Println(getter.get())
	}
	getter(p)
	getter(&p)

	setter := func(setter interface {
		set(name string)
	}, name string) {
		setter.set(name)
	}
	setter(&p, "set")
	fmt.Println(p)

	// interface type的值有两个：具体类型和具体对象.
	// setter = p ：具体类型为Person, 具体对象是p的一个副本.
	// set操作将会作用域p的这个副本，1. 很可能与开发的愿意不符，编译出错，可以提前做出警示
	// 2. 修改这个副本对象是没有什么意义的
	setter(p, "set")
}
