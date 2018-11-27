package test_test

import (
	"fmt"

	"github.com/yandongxiao/golang-learning/test"
)

// The package also runs and verifies example code. Example functions
// may include a concluding line comment that begins with "Output:"
// and is compared with the standard output of the function when
// the tests are run.
// Example functions without output comments are compiled but not executed.
// naming convention: Example, ExampleF, ExampleT, ExampleT_M, ExampleF_suffix
// The entire test file is presented as the example when it contains a single
// example function, at least one other function, type, variable, or constant
// declaration, and no test or benchmark functions.
// godoc -http=:8080
// http://localhost:8080/pkg/github.com/yandongxiao/golang-learning/test/
func ExampleAdd() {
	fmt.Println(test.Add(1, 2))
	fmt.Println(test.Add(10, 20))
	//Output:
	//3
	//30
}
