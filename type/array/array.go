// In Go, an array is a numbered sequence of elements of a specific length.

package main

import "fmt"

func main() {

	// Here we create an array `a` that will hold exactly
	// 5 `int`s. The type of elements and length are both
	// part of the array's type. By default an array is
	// zero-valued, which for `int`s means `0`s.
	var a [5]int
	fmt.Println("emp:", a)

	// We can set a value at an index using the
	// `array[index] = value` syntax, and get a value with
	// `array[index]`.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// The builtin `len` returns the length of an array.
	fmt.Println("len:", len(a))

	// Use this syntax to declare and initialize an array
	// in one line.
	var arrAge = [5]int{18, 20, 15, 22, 16}            // literal-1
	var arrLazy = [...]int{5, 6, 7, 8, 22}             // literal-2
	var arrKeyValue = [10]string{3: "Chris", 4: "Ron"} // literal-3
	var arr4 = [...]string{3: "Chris", 4: "Ron"}       // literal-4
	// var arrLazy = []int{5, 6, 7, 8, 22}			// 只要[]内什么都没有，返回的类型就是slice
	fmt.Printf("%T, %T, %T, %T\n", arrAge, arrLazy, arrKeyValue, arr4)

	// Array types are one-dimensional, but you can
	// compose types to build multi-dimensional data structures.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
