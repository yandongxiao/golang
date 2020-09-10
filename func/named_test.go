package main

import "fmt"

func ExampleBlankIdentifier() {
	// multiple functions can be declared with names as the blank
	// identifier _, in which cases, the declared functions can never be called.
}

func ExampleShadow() {
	v := func() (val int) {
		// no new variables on left side of :=
		// 说明，returned named value 并没有处在一个更大的scope之中
		// val := 10

		if 1 == 1 {
			val := 10
			return val
		}

		if 1 == 1 {
			// the Go 1 compilers disallow return statements
			// without arguments if any of the named return
			// values is shadowed at the point of the return statement.
			// return
			// val := 10
			// return
		}

		return
	}()
	fmt.Println(v)
	// Output:
	// 10
}
