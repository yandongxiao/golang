package main

import "fmt"

type Empty interface {
	getAge() int
}

type APerson struct {
	age int
}

type ADog struct {
	age int
}

// NOTE: struct不但可以包含匿名的struct，而且也可以包含匿名的interface
type House struct {
	Empty
	APerson
}

func ExampleAmbiguous() {
	h := House{
		APerson: APerson{10},
	}
	fmt.Println(h)
	h.Empty = ADog{20} // 如果House.getAge实现了，该赋值是无效的
	fmt.Println(h)

	afoo(h)

	// Output:
	// {<nil> {10}}
	// {{20} {10}}
	// 20
}

func afoo(empty Empty) {
	fmt.Println(empty.getAge())
}

func (d ADog) getAge() int {
	return d.age
}
