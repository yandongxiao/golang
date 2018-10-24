// func copy(dst, src []Type) int
package main

import "fmt"

func main() {
	x := []byte("hello")

	// 1
	copy(x, []byte("world"))
	fmt.Printf("%s\n", x)

	// 2
	copy(x, "WORLD")
	fmt.Printf("%s\n", x)
}
