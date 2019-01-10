package main

import "fmt"

var a [5]int
var p *[7]string

// N and M are both typed constants.
const N = len(a)
const M = cap(p)

func ExampleLenCapCompile() {
	fmt.Println(N) // 5
	fmt.Println(M) // 7
	// Output:
	// 5
	// 7
}
