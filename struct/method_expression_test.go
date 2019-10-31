// Methods are special functions in fact. Methods are often called
// member functions. When a type owns a method, each value of the
// type will own an **immutable** member of function type(最重要的是x.A是一个函数类型). The member name
// is the same as the method name and the type of the member is the same
// as the function declared with the form of the method declaration
// but without the receiver part.
//
// when a method is declared for a type, each value of the type will own a member
// function. Zero values are not exceptions, whether or not the zero values can be
// represented by nil.
//
// method value: have been bound to a specific receiver value.
// e.g. w.Write
// func (p []byte) (n int, err error) {	// the type of the member function
//	return w.Write(p)
// }
package main

import "fmt"

type Person struct {
	name string
}

// Each Method Corresponds To An Implicit Function
// For each method declaration, compiler will declare a corresponding implicit function for it.
// NOTE: Person.Name这种函数名称只能由编译器生成，程序员无权构造含有.的Identifier.
// 但是程序员却可以使用Person.Name(method expression).
// 将receiver parameter作为Person.Name的第一个参数，同时, 保持method body不变
// func Person.Name(p Person) string {
//		return p.name
// }
//
// For each method declared for value receiver type T,
// a corresponding method with the same name will be implictly
// declared by compiler for type *T. NOTE: 程序员不能同时定义T和*T的同名方法
// func (p *Person) Name() string {
//		return Person.Name(*p)
// }
// func (*Person).Name(p *Person) string {
//		return Person.Name(*p)
// }
//
// NOTE: In fact, compilers not only declare the two implicit functions,
// they also **rewrite** the two corresponding explicit declared methods
// to let the two methods call the two implicit functions in the method bodies.
// func (p Person) Name() string {
//		return Person.Name(p)
// }
// 最终，编译器生成了一个方法和两个函数。但是执行的内容是一致的。
//
func (p Person) Name() string {
	// 如果field是Name时，会产生错误:
	// type Person has both field and method named Name
	return p.name
}

// For each method declaration, compiler will declare
// a corresponding implicit function for it.
// func (*Person).Add(p *Person, a , b int) int {
//	 return a + b // the body is the same as the SetPages method
// }
func (p *Person) Add(a, b int) int {
	return a + b
}

func ExampleImplicitFunction() {
	var p Person = Person{name: "jack"}
	fmt.Println(Person.Name(p))
	fmt.Println((*Person).Add(nil, 10, 20))

	// Output:
	// jack
	// 30
}

func ExampleMethodExpression() {
	var foo func(p *Person, a, b int) int
	foo = (*Person).Add
	fmt.Println(foo(nil, 1, 2))
	// Output:
	// 3
}

func ExampleMethodValue() {
	var p Person
	var bar func(x, y int) int
	bar = p.Add
	fmt.Println(bar(1, 2))

	// Output:
	// 3
}

func ExampleMethodValue2() {
	// The expression p is evaluated and saved during
	// the evaluation of the method value;
	// the saved **copy** is then used as the receiver
	// in any calls, which may be executed later.
	var p Person
	p.name = "jack"
	foo := p.Name //因为Name是value recevier类型
	p.name = "alice"
	fmt.Println(foo())

	// Output:
	// jack
}
