package main

import (
	"fmt"
	"unsafe"
)

type GET func(a *Person) int

// These fields can be of any type, even structs themselves , functions or interfaces.
type Person struct {
	name string
	age  int
	get  GET         // 将一个函数作为一个field
	i    interface{} // 将一个interface作为一个field
}

func main() {
	initialize()
	getSize()
	getAccess()
}

func initialize() {
	// 四者等价
	p1 := &Person{
		name: "jack",
	}

	// struct 是 value type，所以使用new来创建新的对象
	p2 := new(Person)
	p2.name = "jack"

	p3 := &Person{}
	p3.name = "jack"

	p4 := &Person{name: "jack"} // 推荐

	fmt.Printf("%T %v\n", p1, p1)
	fmt.Printf("%T %v\n", p2, p2)
	fmt.Printf("%T %v\n", p3, p3)
	fmt.Printf("%T %v\n", p4, p4)
}

// Even when a struct contains other structs, structs form a continuous block in memory
// this gives a huge performance benefit.
func getSize() {
	age := 10 // size of int == 8
	fmt.Println(unsafe.Sizeof(age))
	// 一个指针 + 一个size 各占8个字节
	name := "helloworld"
	fmt.Println(unsafe.Sizeof(name))
	// NOTE: struct的大小=每个field的大小的和.
	fmt.Println(unsafe.Sizeof(Person{}))
}

// The rule about pointers vs. values for receivers is that value methods can be invoked on pointers and values,
// but pointer methods can only be invoked on pointers. 依据上面的规则，p1.setAge(100)岂不是应该调用失败？
// 先要搞清楚什么是addressable变量：其实就是开发者可以访问到的变量，不论是指针还是值
// pointer receiver 方法意味着会对对象进行修改，如果被修改的变量是addressable的，那么就是合法调用, 不论是指针还是值。
func getAccess() {
	p1 := Person{age: 10}
	p1.setAge(100) // 作用在p1上，p1是可访问的，所以合法
	fmt.Println(p1.getAge())

	p2 := &Person{age: 20}
	p2.setAge(200)
	fmt.Println(p2.getAge())

	// cannot use p1 (type Person) as type Setter in argument to setAge
	// 解释：setter = p1, 即赋值给interface变量时，会创建一个p1的一个副本，修改操作将会应用在p1的副本上。
	//       p1副本是不可访问的。注意，通过type assertion的方式返回的是p1副本的副本.
	// setAge(p1)
}

func (p Person) getAge() int {
	return p.age
}

// 如果再在pointer上重定义getAge，会导致method redeclared: Person.getAge错误
//func (p *Person) getAge() int {
//	return p.age
//}

func (p *Person) setAge(age int) {
	p.age = age
}

type Setter interface {
	setAge(age int)
}

func setAge(setter Setter) {
	setter.setAge(100)
}
