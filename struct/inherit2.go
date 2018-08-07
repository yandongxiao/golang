package main

import "fmt"

type Getter interface {
	get() int
}

type Base struct {
	age int
}

func (b Base) get() int {
	return b.age
}

type Child struct {
	Base
}

func main() {
	g := Getter(new(Child)) // 可以传递一个Child, 虽然Child并没有实现get方法
	fmt.Println(g.get())
}
