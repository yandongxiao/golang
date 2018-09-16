// What are the rules when there are two fields with the same name
// An outer name hides an inner name. This provides a way to override a field or method.
//
// If the same name appears twice at the same level, it is an error if the name is used by the program.
// If it’s not used, it doesn’t matter.
// There are no rules to resolve the ambiguity; it must be fixed.
package main

import "fmt"

type Empty interface {
	getAge() int
}

type Person struct {
	age int
}

type Dog struct {
	age int
}

// 解释：1. 为什么House没有实现getAge接口，编译时却没有出错, foo(h).
//		 2. 为什么为Person类增加了getAge接口后，编译时却报错了
//
// 假如House只有Empty.
// struct之间只存在一种关系，那就是组合关系。所以匿名接口Empty的本质含义是struct House拥有一个Empty接口的成员变量
// 所以支持语法: h.Empty = nil. 同时也支持h.Empty = Dog{20}语法. 所以foo(h)在编译器没有报错.
//
// 增加Person以后:
//	h.getAge就会出现歧义.  Person.getAge 和 Empty.getAge 两个都实现该方法。
//	golang开始谋求通过House.getAge来解决，如果发现该方法未实现，则编译时报错
//  所以, 上面的歧义通过实现House自己的getAge即可解决.
//
type House struct {
	Empty // 接口也可以匿名
	Person
}

func main() {
	h := House{
		Person: Person{10},
	}
	fmt.Println(h)
	h.Empty = Dog{20} // 如果House.getAge实现了，该赋值是无效的

	// ambiguous selector h.age
	// h.age = 10
	println(h.Person.age)
	foo(h)
}

func foo(empty Empty) {
	println(empty.getAge())
}

func (p Person) getAge() int {
	return p.age
}

func (d Dog) getAge() int {
	return d.age
}

func (h House) getAge() int {
	return h.Person.age
}
