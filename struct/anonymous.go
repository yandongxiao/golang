// What are the rules when there are two fields with the same name
// An outer name hides an inner name. This provides a way to override a field or method.
// If the same name appears twice at the same level, it is an error if the name is used by the program. (If it’s not used, it doesn’t matter.) There are no rules to resolve the ambiguity; it must be fixed.

package main

import "fmt"

type Person struct {
	age int
}

type Dog struct {
	age int
}

type House struct {
	Person
	Dog
}

func main() {
	h := House{Person{10}, Dog{1}}
	fmt.Println(h)

	// ambiguous selector h.age
	// h.age = 10
}
