package main

import (
	"fmt"
	"time"
)

func basic() {
	// It is equivalent to NewTimer(d).C
	// If efficiency is a concern, use NewTimer instead and call Timer.Stop if the timer is no longer needed.
	ch := time.After(100 * time.Millisecond)
	fmt.Println(time.Now())
	time.Sleep(time.Second)
	fmt.Println(<-ch) // 返回值不是100ms以后的值，并非是1秒以后的值
}

func useTimer() {
	timer := time.NewTimer(100 * time.Millisecond)
	fmt.Println(time.Now())
	time.Sleep(time.Second)
	fmt.Println(<-timer.C)
}

func main() {
	basic()
	useTimer()
}
