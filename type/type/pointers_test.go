// the unsafe standard package can be used to break the restrictions made for pointers
// in Go. The unsafe.Pointer type is like the void* in C. However, generally, the unsafe
// ways are not recommended to be used in general Go programming.
//	1. Go Pointer Values Don't Support Arithmetic Operations
//  2. A Pointer Value Can't Be Converted To An Arbitrary Pointer Type
//  3. A Pointer Value Can't Be Compared With Values Of An Arbitrary Pointer Type
//  4. A Pointer Value Can't Be Assigned To Pointer Values Of Other Pointer Types
package main

import "fmt"

func ExampleValueAddress() {
	// 1. the start address of the memory segment occupied
	// by the direct part of the value.
	// 2. Type T is called the base type of pointer type *T.
	num := 10
	println(&num)
	// Output:
	//
}

func ExampleAddressable() {
	// The built-in new function can be used to allocate memory for a value of **any** type
	var a *int = new(int) // create an anonymous variable
	*a = 10
	// we can use the expression &num to take the address of num(a variable)
	// 不能获取地址的情况：
	// bytes in strings, map elements, dynamic values of interface values (exposed by type assertions)
	// constant values, literal values, package level functions, methods (used as function values)
	// intermediate values
	//	  function calls, explicit value conversions
	//    all sorts of operations, excluding pointer dereference operations, but including:
	//		channel receive operations
	//		sub-string operations
	//		sub-slice operations
	//		addition, subtraction, multiplication, and division, etc.
	//    &T{}是一个语法糖：tmp := T{}; (&tmp)
	// 可以获取地址的情况：
	//	variables
	//	fields of addressable structs
	//	elements of addressable arrays
	//  elements of any slices (whether the slices are addressable or not)
	//  pointer dereference operations
	var num int = 10
	var b *int = &num
	fmt.Println(*a == *b)
	// Output:
	// true
}

func ExampleDereference() {
	defer func() {
		fmt.Println(recover())
	}()
	var a *int
	_ = *a
	// Output:
	// runtime error: invalid memory address or nil pointer dereference
}

func ExampleNamedPointerType() {
	// We can declare named pointer types, but generally, it’s not recommended.
	// 1. Golang不允许在Pointer Type上定义方法
	// 2. Unnamed pointer types have better readabilities than named ones.
	// 3. the base type of *int(or PINT) is int
	// 4. Two non-defined pointer types with the same base type are the same type.
	type PINT *int
	var num int

	var a *int = &num
	var b *int = &num
	// The conditions to assign a pointer value to another pointer value are
	// the same as the conditions to compare a pointer value to another pointer
	// value, which are listed below. --> 如2， 可以隐式转换
	// Type T1 and T2 are both non-defined pointer types and the underlying types
	// of their base types are identical (ignoring struct tags). --> 显式转换
	var c PINT = a
	fmt.Println(a == b)
	// Two Go pointer values can only be compared if either of the following
	// three conditions are satisfied.
	// 1. The types of the two Go pointers are identical.
	// 2. One pointer value can be implicitly converted to the pointer type of
	//    the other. In other words, the underlying types of the two types must
	//    be identical and either of the two types of the two Go pointers must
	//    be an undefined type.
	// 3. compared to nil
	fmt.Println(a == c)
	// Output:
	// true
	// true
}

func ExampleNamedPointerType2() {
	type PINT1 *int
	type PINT2 *int

	var num int
	var p1 PINT1 = &num
	var p2 PINT2 = PINT2(p1)
	fmt.Println(p1 == PINT1(p2))
	// Output:
	// true
}

func ExampleNamedPointerType3() {
	type MyInt int64
	type Ta *int64
	type Tb *MyInt

	var a Ta
	var b Tb
	//fmt.Println(a == Ta(b))
	fmt.Println(a == Ta((*int64)((*MyInt)(b))))
	// Output:
	// true
}

func newInt() *int {
	a := 3
	return &a
}

func ExamplePointer() {
	// Return Pointers Of Local Variables Is Safe In Go
	fmt.Println(*newInt())
	// Output:
	// 3
}
