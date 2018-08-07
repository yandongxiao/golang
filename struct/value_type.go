// They are also value types and so are constructed with the new function.
package main

import "fmt"

type Person struct {
	age int
}

func set(p Person, age int) {
	p.age = age
}

func main() {
	p := Person{age: 10}
	set(p, 20)
	fmt.Println(p)
}
