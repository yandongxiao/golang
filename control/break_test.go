// A "break" statement terminates execution of the innermost "for", "switch", or "select" statement within the same function.
// If there is a label, it must be that of an enclosing "for", "switch", or "select" statement
package main

import "fmt"

func ExampleBasic() {
	var i, j int
OuterLoop:
	for i = 0; i < 10; i++ {
		for j = 0; j < 10; j++ {
			if i == 5 && j == 5 {
				break OuterLoop
			}
		}
	}

	fmt.Println(i, j)
	// Output:
	// 5 5
}
