package main

import "fmt"

func main() {
	// Regardless of the verb, if an operand is an interface value, the
	// internal concrete value is used, not the interface itself. Thus:
	var i interface{} = 23
	fmt.Printf("%d\n", i) // 注意用的是%d
}
