// In Go, we can (explicitly) declare a method for type T and *T,
// where T must satisfiy 4 conditions:
//	1. T must be a defined type;
//	2. T must be defined in the same package as the method declaration;
//	3. T must not be a pointer type;
//	4. T must not be an interface type.
// Type T and *T are called the receiver type.
// Type T is called the receiver base types of all methods.
// The receiver of type *T are called pointer receiver, non-pointer receivers are called value receivers.
//
// we can also declare methods for alias types of the T and *T types specified above.
// The effect is the same as declaring methods for the T and *T types themselves.
//
// we can never (explicitly) declare methods for built-in basic types, interface types, non-defined types.
// 两种特殊情况除外：the pointer types *T，embeds other types which have methods.
//
// A method prototype can be viewed as a function prototype without the func keyword.
// a function prototype: func Double(n int) (result int)
// method declaration: the func keyword + a receiver parameter declartion
//					   + a method prototype + a method (function) body.
// Each type has a method set: all the method prototypes of the methods declared,
//							   either explicitly or implicitly, for the type.
// method set 关系：subset，superset, identical
// non-exported method names, which start with lower-case letters, from different packages
// will be always viewed as two different method names, even if the two method
// names are the same in literal.
//
// Receiver Arguments Are Passed By Copy
// the modifications on the dirct part of a receiver argument in a method call will not be
// reflected to the outside of the method.
//
package main

import "fmt"

type MPerson struct {
	name string
}

func (p *MPerson) Foo() string {
	return "Foo"
}

// Method names can be the blank identifier _.
// A type can have multiple methods with the blank identifier as names.
// But such methods can never be called.
func (p *MPerson) _() string {
	return ""
}

func ExampleNilPointer() {
	var p *MPerson
	fmt.Println(p.Foo())
	// Output:
	// Foo
}
