// func append(slice []Type, elems ...Type) []Type
package main

import (
	"fmt"
)

func ExampleSlice() {
	var data []byte // NOTE: data can be nil
	// the append function doesn't require the variadic argument
	// must be a slice with the same type as the first slice argument
	// two argument slices must share the same underlying type.
	data = append(data, []byte{4, 5, 6}...)
	data = append(data, []byte(nil)...)
	data = append(data, 7, 8, 9)
	data = append(data, "xyz"...)
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
