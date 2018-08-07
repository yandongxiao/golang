package main

type Person struct {
	age int
}

// 重新创建一种类型，并在之上创建方法
// 同时，可以使用P(p1)的方式进行类型转换
type P Person

func (p P) get() int {
	return p.age
}

func test1() {
	p1 := Person{
		age: 10,
	}

	p2 := P(p1)
	println(p2.get())
}

/*
// The receiver type must be of the form T or *T where T is a type name.
// The type denoted by T is called the receiver base type; it must not be a pointer or interface type and it must be declared in the same package as the method.
type PP *Person

func (p PP) get() int {
	return p.age
}

func test2() {
	p1 := &Person{
		age: 10,
	}

	p2 := PP(p1)
	println(p2.get())
}
*/
func main() {
	test2()
}
