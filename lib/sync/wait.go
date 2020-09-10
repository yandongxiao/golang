// A WaitGroup waits for a collection of goroutines to finish.
// WaitGroup的主要目标：确保由协程A创建出去的协程B，以及B的子协程，都已经成功回收.
package main

import (
	"fmt"
	"sync"
)

func myfunc(wg *sync.WaitGroup, data int) {
	defer wg.Done()
	fmt.Println(data)
}

func main() {
	var wg sync.WaitGroup
	// The main goroutine calls Add to set the number of goroutines to wait for.
	// Add adds delta, which may be negative, to the WaitGroup counter.
	// If the counter becomes zero, all goroutines blocked on Wait are released.
	// If the counter goes negative, Add panics. 当counter==0时，delta必须大于0
	// NOTE: 一定要在创建协程之前调用
	// If a WaitGroup is reused to wait for several independent sets of events,
	// new Add calls must happen after all previous Wait calls have returned.
	wg.Add(100)

	// Then each of the goroutines runs and calls Done when finished.
	for i := 0; i < 100; i++ {
		// wg.Add(1)	// 也可以每创建一个协程，调用一次Add.
		go myfunc(&wg, i)
	}

	// At the same time, Wait can be used to block until all goroutines have finished.
	// If the counter becomes zero, all goroutines blocked on Wait are released.
	// NOTE: If the counter goes negative, Add panics. 所以Done操作一定要正合适
	wg.Wait()
}
