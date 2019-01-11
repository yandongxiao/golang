// func append(slice []Type, elems ...Type) []Type
package main

import (
	"fmt"
)

func ExampleBuiltinValues() {
	// use of builtin append not in function call
	// NOTE, built-in functions can't be used as values.
	// a := append
	a([]byte(nil), "hello"...)
	//len("hello")
	// Output:
	//
}

func ExampleDiscardResults() {
	// append(([]byte)(nil), "hello"...) evaluated but not used
	// The return results of a custom function call can be all discarded together.
	// The return results of calls to built-in functions, except recover and copy,
	// can't be discarded, though they can be ignored by assigning them to some
	// blank identifiers.
	// Function calls whose results can't be discarded can't be used as deferred
	// function calls or goroutine calls.
	append([]byte(nil), "hello"...)
	//len("hello")
	// Output:
	//
}

func ExampleSlice() {
	var data []byte // NOTE: data can be nil
	// the append function doesn't require the variadic argument
	// must be a slice with the same type as the first slice argument
	// two argument slices must share the same underlying type.
	data = append(data, []byte{4, 5, 6}...)
	data = append(data, []byte(nil)...)
	data = append(data, 7, 8, 9)
	data = append(data, "xyz"...) // Sugar
	data = append(data)
	fmt.Println(data)
	//Output:
	//[4 5 6 7 8 9 120 121 122]
}

func ExampleNilMap() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()

	var m map[int]int
	m[1] = 1 // panic
	// Output:
	// assignment to entry in nil map
}
