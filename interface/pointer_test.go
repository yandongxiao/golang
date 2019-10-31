package main

import "fmt"

type Person struct {
	name string
}

func (p Person) get() string {
	return p.name
}

func (p *Person) set(name string) {
	p.name = name
}

func ExampleGetSet() {
	p := Person{"jack"}
	fmt.Println(p.get())
	p.set("lua")
	(&p).set("lua")
	fmt.Println((&p).get())

	// Output:
	// jack
	// lua
}

func ExampleGetter() {
	p := Person{"jack"}
	getter := func(getter interface {
		get() string
	}) {
		fmt.Println(getter.get())
	}
	getter(p)
	getter(&p)
	// Output:
	// jack
	// jack
}

func ExampleSetter() {
	p := Person{"jack"}
	setter := func(setter interface {
		set(name string)
	}, name string) {
		setter.set(name)
	}
	setter(&p, "set") //setter(p, "set") 会导致编译出错
	fmt.Println(p)

	// Output:
	// {set}
}
