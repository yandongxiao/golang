// method value: have been bound to a specific receiver value.
// e.g. w.Write
// func (p []byte) (n int, err error) {
//	return w.Write(p)
// }
//
// method expressions: generate functions from methods of a given
// type, e.g. (*bufio.Writer).Write
// func (w *bufio.Writer, p []byte) (n int, err error) {
//		return w.Write(p)
// }
package main

import "fmt"

type Person struct {
	name string
}

// type Person has both field and method named Name
func (p Person) Name() string {
	return p.name
}

func (p *Person) Add(a, b int) int {
	return a + b
}

func ExampleMethodExpression() {
	var foo func(p *Person, a, b int) int
	foo = (*Person).Add
	fmt.Println(foo(nil, 1, 2))
	// Output:
	// 3
}

func ExampleMethodValue() {
	var p Person
	var bar func(x, y int) int
	bar = p.Add
	fmt.Println(bar(1, 2))

	// Output:
	// 3
}

func ExampleMethodValue2() {
	// The expression p is evaluated and saved during
	// the evaluation of the method value;
	// the saved **copy** is then used as the receiver
	// in any calls, which may be executed later.
	var p Person
	p.name = "jack"
	foo := p.Name
	p.name = "alice"
	fmt.Println(foo())

	// Output:
	// jack
}
