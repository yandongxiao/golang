package main

type Interface1 interface {
	func1()
}

type Interface2 interface {
	func2()
}

type Interface3 interface {
	Interface1
	Interface2
}

type INT int

func (a INT) func1() {
	println(a + 1)
}

func (a INT) func2() {
	println(a + 2)
}

func main() {
	v := 10
	ai := INT(v)
	var i Interface3 = ai
	i.func1()
	i.func2()
}
