// What are the rules when there are two fields with the same name
// An outer name hides an inner name.
// This provides a way to override(重载) a field or method.
//
// NOTE : If the same name appears twice at the same level, it is an error
// if the name is used by the program. If it’s not used, it doesn’t matter.
// There are no rules to resolve the ambiguity; it must be fixed.
package main

import "fmt"

type Animal struct {
	Name string
}

type Dog struct {
	Animal
}

type Cat struct {
	Name string
	Animal
}

func ExampleHide() {
	dog := Dog{
		Animal: Animal{Name: "Animal"},
	}
	fmt.Println(dog)
	fmt.Println(dog.Name)
	// Output:
	// {{Animal}}
	// Animal
}

func ExampleHide2() {
	cat := Cat{
		Name:   "Cat",
		Animal: Animal{Name: "Animal"},
	}
	fmt.Println(cat)
	fmt.Println(cat.Name)
	// Output:
	// {Cat {Animal}}
	// Cat
}
