/**
5.145447453s
10000000
122.3107ms
10000000
60.465192ms
10000000
3.623827ms
10000000
*/
// 假设，一开始我们使用channel同步方案
// 如果换成sync.Mutex方案，效率提升了40倍
// 接下来，换成atomic方案，效率是刚才的两倍
// 最后，换成没有同步的方案，效率是刚才的20倍
// 结论：每个goroutine执行很少的代码时，不划算
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func addOne(val int, resp chan<- int) {
	resp <- val + 1
}

func ExampleB() {
	now := time.Now()
	sum := 0
	resp := make(chan int)
	for i := 0; i < 10000000; i++ {
		go addOne(sum, resp)
		sum = <-resp
	}
	fmt.Println(time.Now().Sub(now))
	fmt.Println(sum)

	now = time.Now()
	sum = 0
	var mutex sync.Mutex
	for i := 0; i < 10000000; i++ {
		mutex.Lock()
		sum += 1
		mutex.Unlock()
	}
	fmt.Println(time.Now().Sub(now))
	fmt.Println(sum)

	sum32 := int32(0)
	now = time.Now()
	for i := 0; i < 10000000; i++ {
		atomic.AddInt32(&sum32, 1)
	}
	fmt.Println(time.Now().Sub(now))
	fmt.Println(sum32)

	sum = 0
	now = time.Now()
	for i := 0; i < 10000000; i++ {
		sum += 1
	}
	fmt.Println(time.Now().Sub(now))
	fmt.Println(sum)
}
