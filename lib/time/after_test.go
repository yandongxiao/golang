package main

import (
	"fmt"
	"time"
)

// 注意：不能在for{} 中，使用 select + time.After
// 每次调用 time.After() 都相当于创建了一个 time.Timer 对象，同时启动一个协程，等待一段时间
// 该时间段内，time.Timer 对象是不会被GC的。如果 select{} 内部有不阻塞的channel操作，
// 那么会导致创建大量无用的 time.Timer 对象. 导致内存泄露
func ExampleTimeAfter() {
	select {
	case <-time.After(100 * time.Millisecond):
		fmt.Println("time.After")
	}
	// Output:
	// time.After
}

// It is equivalent to NewTimer(d).C
// If efficiency is a concern, use NewTimer instead and call Timer.Stop if the timer is no longer needed.
func ExampleNewTimer() {
	timer := time.NewTimer(100 * time.Millisecond)
	select {
	case <-timer.C:
		fmt.Println("time.NewTimer")
	}

	// Output:
	// time.NewTimer
}
