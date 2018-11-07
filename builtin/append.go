// func append(slice []Type, elems ...Type) []Type
package main

import "fmt"

func slice() {

	var data []byte

	// 0
	data = append(data, []byte{4, 5, 6}...)
	fmt.Println(data)

	// 0.1
	data = append(data, []byte(nil)...)
	fmt.Println(data)

	// 2
	data = append(data, 7, 8, 9)
	fmt.Println(data)

	// 3
	data = append(data, "xyz"...)
	fmt.Println(data)
}

func main() {
	slice()

	// for map
	// NOTE: panic: assignment to entry in nil map
	var m map[int]int
	m[1] = 1
}
