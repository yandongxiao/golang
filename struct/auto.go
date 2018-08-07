// Pointer and value methods can both be called on pointer or non-pointer values
package main

import "fmt"

type Person struct {
	age int
}

func (p Person) getAge() int {
	return p.age
}

func (p *Person) setAge(age int) {
	p.age = age
}

func main() {
	p1 := Person{10}
	p2 := &Person{20}
	p1.setAge(100)
	p2.setAge(200)
	fmt.Println(p1.getAge())
	fmt.Println(p2.getAge())
}
