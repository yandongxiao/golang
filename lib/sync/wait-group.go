// A WaitGroup waits for a collection of goroutines to finish.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wait sync.WaitGroup
	// The main goroutine calls Add to set the number of goroutines to wait for.
	// wait.Add(10)

	for i := 0; i < 10; i++ {
		// 如果把wait.Add操作放在goroutine内部的一开始的位置，那么wait.Wait就可能发生在某个goroutine被创建，但是还没有开始运行
		wait.Add(1)
		go func() {
			// each of the goroutines runs and calls Done when finished.
			defer wait.Done()

			time.Sleep(time.Second)
			fmt.Println("helloworld")
		}()
	}

	// At the same time, Wait can be used to block until all goroutines have finished.
	wait.Wait()
}
