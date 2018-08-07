// These fields can be of any type, even structs themselves , functions or interfaces.
package main

import "fmt"

type GET func(a *Person) int

type Person struct {
	age int
	get GET         // 将一个函数作为一个field
	_   interface{} // 将一个interface作为一个field
}

func get(p *Person) int {
	return p.age
}

func main() {
	p := new(Person)
	p.age = 20
	p.get = GET(get)
	fmt.Println(p.get(p))
}
