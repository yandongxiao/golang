package test_test

import (
	"fmt"

	"github.com/yandongxiao/golang-learning/test"
)

// The entire test file is presented as the example when it contains a single
// example function, at least one other function, type, variable, or constant
// declaration, and no test or benchmark functions.
// NOTE: godoc -http=:8080, 应该是暴露出整个项目，无论当前工作目录是哪里
// http://localhost:8080/pkg/github.com/yandongxiao/golang-learning/test/
func ExampleAdd() {
	fmt.Println(test.Add(1, 2))
	fmt.Println(test.Add(10, 20))
	//Output:
	//3
	//30
}
