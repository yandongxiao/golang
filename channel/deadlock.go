// NOTE: fatal error: all goroutines are asleep - deadlock! 该错误信息+堆栈信息都打印在了stderr上
// NOTE: 该错误是无法通过defer+recover恢复过来的，因为所有的协程都处于asleep状态
package channel

import (
	"fmt"
)

func test() {
	ch := make(chan int)
	defer func() {
		fmt.Println("----")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ch <- 1
}

// 没有Output注释，所以这个Example也不会执行，安全
func ExampleC() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	go test()
	select {} // block forever
}
