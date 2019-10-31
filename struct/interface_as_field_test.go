package main

import "fmt"

type Foo interface {
	foo()
}

type IPerson struct {
	Foo
}

func ExampleNil() {
	defer func() {
		fmt.Println(recover())
	}()
	p := IPerson{}
	fmt.Println(p)
	fmt.Println(p.Foo)
	p.foo()
	// Output:
	// {<nil>}
	// <nil>		// 这个<>没有什么特别的含义
	// runtime error: invalid memory address or nil pointer dereference
}

type SFoo struct{}

func (*SFoo) foo() {
	fmt.Println("*SFoo foo")
}

func ExamplePointer() {
	var s *SFoo // s==nil
	p := IPerson{Foo: s}
	fmt.Println(p)
	fmt.Println(p.Foo)
	p.foo()
	// Output:
	// {<nil>}
	// <nil>	// 这个<>没有什么特别的含义
	// *SFoo foo
}
