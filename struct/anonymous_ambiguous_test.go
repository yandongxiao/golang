package main

import "fmt"

type Empty interface {
	getAge() int
}

type Person struct {
	age int
}

type Dog struct {
	age int
}

// NOTE: struct不但可以包含匿名的struct，而且也可以包含匿名的interface
type House struct {
	Empty
	Person
}

func ExampleAmbiguous() {
	h := House{
		Person: Person{10},
	}
	fmt.Println(h)
	h.Empty = Dog{20} // 如果House.getAge实现了，该赋值是无效的
	fmt.Println(h)

	// foo(h)

	// Output:
	// {<nil> {10}}
	// {{20} {10}}
}

func foo(empty Empty) {
	println(empty.getAge())
}

func (p Person) getAge() int {
	return p.age
}

func (d Dog) getAge() int {
	return d.age
}
