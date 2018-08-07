// 1. Pointer methods can be called with pointers.
// 2. Value methods can be called with values.
// 3. Value-receiver methods can be called with pointer values because they can be dereferenced first.
// 4. Pointer-receiver methods cannot be called with values, however, because the value stored inside an interface(receiver类型是Person对象，pointer方法调用作用在该对象上是没有意义的) has no address.
//
package main

import "fmt"

func foo(num int) {
	num++
}

func test1() {
	v := new(int)
	*v = 20
	// cannot use v (type *int) as type int in argument to foo
	// NOTE：不能直接传递v
	foo(*v)
	fmt.Println(*v)
}

type Person struct {
	age int
}

func (p Person) get() int {
	return p.age
}

func (p *Person) set(n int) {
	p.age = n
}

func test2() {
	p := new(Person)
	// NOTE: 以下两行是自动解引用
	p.age = 20
	fmt.Println(p.get())

}

type Getter interface {
	get() int
}

func test3() {
	p := new(Person)
	p.age = 20
	getter := Getter(p) // NOTE: 自动解引用
	println(getter.get())
}

type Gser interface {
	Getter
	set(int)
}

func test4() {
	p := new(Person)
	// 虽然Person get的recevier类型是value, set的recevier类型是pointer
	// 但是get方法会被认为value和receiver类型都实现了.
	sger := Gser(p)
	sger.set(20)
	println(sger.get())
	println(p.get())
}

func main() {
	test4()
}
