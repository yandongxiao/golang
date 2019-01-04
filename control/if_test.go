package main

import "fmt"

func ExampleIf() {
	// A statement can precede conditionals;
	// any variables declared in this statement are available in all branches
	// NOTE: if 后面跟的语句形成了一个新的scope.
	num := -10
	if num := 1; num < 0 {
		fmt.Printf("%d is negtive\n", num)
	} else if num == 0 {
		fmt.Printf("%d is zero\n", num)
	} else {
		fmt.Printf("%d is positive\n", num)
	}
	fmt.Println(num)

	// Output:
	// 1 is positive
	// -10
}
