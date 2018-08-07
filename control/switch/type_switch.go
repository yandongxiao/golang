// type switch
package main

import "fmt"

func main() {
	whatAmI := func(i interface{}) {
		// a type switch is like a regular switch statement
		// The declaration in a type switch has the same syntax as a type assertion i.(T),
		switch t := i.(type) { // but the specific type T is replaced with the keyword type.
		case bool: // but the cases in a type switch specify types (not values)
			fmt.Println("I am bool", t) // 在这个case当中，t的类型信息变为bool
		default: // 在default当中，the variable t is of the same interface type and value as i
			// fmt.Printf是如何知道t的实际类型信息为int的？
			fmt.Printf("I do not konow, %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(3)
}
