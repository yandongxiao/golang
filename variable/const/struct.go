// The constant data can only be of type boolean, number or string
package main

import "fmt"

type Person struct {
	name string
}

const P Person = Person{"jack"}

func main() {
	fmt.Println(P)
}
