package test_test

import (
	"fmt"

	"github.com/yandongxiao/golang-learning/test"
)

// To achieve this we can use a "whole file example."
// A whole file example is a file that ends in _test.go
// and contains exactly one example function, no test or
// benchmark functions, and at least one other package-level
// declaration. When displaying such examples godoc will show
// the entire file.
func Example() {
	fmt.Println(test.Add(1, 2))
	fmt.Println(test.Add(10, 20))
	//Output:
	//3
	//30
}
