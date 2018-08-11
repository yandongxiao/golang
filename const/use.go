package main

import "fmt"

type Sex int

const (
	MAN   = Sex(0)
	WOMAN = Sex(1)
)

func (sex Sex) String() string {
	if sex == MAN {
		return "man"
	}
	return "woman"
}

type Person struct {
	name string
	sex  Sex
}

func main() {
	p := Person{
		name: "jack",
		sex:  MAN,
	}

	fmt.Println(p)
	fmt.Printf("{name:%v, sex=%v}", p.name, p.sex)
}
