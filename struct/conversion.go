// When we have a struct type and define an alias type for it, both types have the same underlying type and can be converted into one another,
// but also note the compile-error cases which denote impossible assignments or conversions:

package main

import "fmt"

type number struct {
	f float32
}
type NUMBER number // alias type

func main() {
	n1 := number{1.0}
	n2 := NUMBER(n1)
	fmt.Println(n1, n2) // ok

	// cannot convert 2.3 (type float64) to type NUMBER)
	// n3 := NUMBER(2.3)
	// n4 := number(2.3)
	// fmt.Println(n3, n4)
}
