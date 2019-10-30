// func append(slice []Type, elems ...Type) []Type
package main

import (
	"fmt"
)

func ExampleBuiltinValues() {
	fmt.Println(string(append([]byte(nil), "hello"...)))
	// Output:
	// hello
}

func ExampleSlice() {
	var data []byte // NOTE: data can be nil
	data = append(data, []byte{4, 5, 6}...)
	data = append(data, []byte(nil)...)
	data = append(data, 7, 8, 9)
	data = append(data, "xyz"...) // Sugar
	data = append(data)
	fmt.Println(data)
	//Output:
	//[4 5 6 7 8 9 120 121 122]
}

func ExampleShare() {
	data := make([]int, 0, 10)
	data2 := data[5:10:10]
	for i := 0; i < 10; i++ {
		data = append(data, i)
	}
	fmt.Println(data2)
	//Output:
	//[5 6 7 8 9]
}
