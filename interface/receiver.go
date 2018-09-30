package main

import "fmt"

type Adder interface {
	add(v int) int
}

type Person struct {
	age int
}

func (p Person) add(v int) int {
	p.age += v
	return p.age
}

func main() {
	p := Person{
		age: 10,
	}
	adder := Adder(p)
	fmt.Println(adder.add(1))          // 11
	fmt.Println(adder.add(1))          // 11
	fmt.Println(adder.(Person).add(1)) // 11
	fmt.Println(p.add(1))              // 11
}
