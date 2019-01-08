package main

import "fmt"

type Interface1 interface {
	func1()
}

type Interface2 interface {
	func2()
}

// NOTE:
type Interface3 interface {
	Interface1
	Interface2
}

type INT int

func (a INT) func1() {
	fmt.Println(a + 1)
}

func (a INT) func2() {
	fmt.Println(a + 2)
}

func ExampleAnonymousInterface() {
	v := 10
	ai := INT(v)
	var i Interface3 = ai
	i.func1()
	i.func2()

	// Output:
	// 11
	// 12
}
