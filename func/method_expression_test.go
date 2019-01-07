package main

import "fmt"

type Person struct{}

func (p *Person) Add(a, b int) int {
	return a + b
}

func ExampleMethodExpression() {
	var foo func(p *Person, a, b int) int
	foo = (*Person).Add
	fmt.Println(foo(nil, 1, 2))

	var p Person
	var bar func(x, y int) int
	bar = p.Add
	fmt.Println(bar(1, 2))

	// Output:
	// 3
	// 3
}
