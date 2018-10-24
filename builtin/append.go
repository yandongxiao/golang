// func append(slice []Type, elems ...Type) []Type
package main

import "fmt"

func main() {

	data := []byte{1, 2, 3}

	// 1
	data = append(data, []byte{4, 5, 6}...)
	fmt.Println(data)

	// 2
	data = append(data, 7, 8, 9)
	fmt.Println(data)

	// 3
	data = append(data, "xyz"...)
	fmt.Println(data)
}
