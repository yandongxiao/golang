// What are the rules when there are two fields with the same name
// An outer name hides an inner name. This provides a way to override a field or method.
//
// If the same name appears twice at the same level, it is an error if the name is used by the program.
// If it’s not used, it doesn’t matter.
// There are no rules to resolve the ambiguity; it must be fixed.
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

type House struct {
	//	Empty // 接口也可以匿名
	Person
	Dog
}

func main() {
	h := House{
		Person: Person{10},
	}
	fmt.Println(h)

	// ambiguous selector h.age
	// h.age = 10
	println(h.Person.age)
	foo(h)
}

func foo(empty Empty) {
	println(empty.getAge())
}
