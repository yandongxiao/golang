package main

import "fmt"

func main() {
	a := 100 /* automatic type inference */
	b := "nihao"
	c := true

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
