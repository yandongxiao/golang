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
	// <nil>
	// runtime error: invalid memory address or nil pointer dereference
}

type SFoo struct{}

func (*SFoo) foo() {
	fmt.Println("*SFoo foo")
}

func ExamplePointer() {
	var s *SFoo
	p := IPerson{Foo: s}
	fmt.Println(p)
	fmt.Println(p.Foo)
	p.foo()
	// Output:
	// {<nil>}
	// <nil>
	// *SFoo foo
}
