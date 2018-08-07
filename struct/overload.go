//	./receiver.go:14: method redeclared: Person.getAge
//			method(Person) func() int
//		    method(*Person) func() int
//
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) getAge() int {
	return p.age
}

func (p *Person) getAge() int {
	return p.age
}

func main() {
	p := new(Person)
	fmt.Println(p.getAge())
}
