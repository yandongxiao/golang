package main

type Adder interface {
	add(v int)
	get() int
}

type Person struct {
	age int
}

func (p Person) add(v int) {
	p.age += v
}

func (p Person) get() int {
	return p.age
}

func (p *Person) set(n int) {
	p.age = n
}

func main() {
	// Adder的数据类型是receiver + method pointer, receiver内存储的数据类型，数据大小都是不确定的
	// adder := Adder(p) 指定了receiver的数据类型，类型为Person，所以此时receiver存储了一个全新的与p无关的类型对象
	// p.age = 100 证明了上面的猜测
	var p Person
	adder := Adder(p)
	p.age = 100 //
	// 连续调用两次Person的add方法，并没有导致p的成员变量age的值有所改变
	// 通过接口调用的，可以类比通过person.add调用, Person类型通过值传递
	adder.add(1)
	println(adder.get()) // not 1

	p.set(200) // 自动引用和自动解引用都是针对类型对象调用方法的情况下
	println(p.age)
}
