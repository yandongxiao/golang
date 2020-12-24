package buildin

import "fmt"

// 内置 print 和 fmt.Print 函数不是同步函数(synced)。需要使用 log.Print 函数
func ExamplePrint() {
	fmt.Println("hello world")
	// Output:
	// hello world
}
