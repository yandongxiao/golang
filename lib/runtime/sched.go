package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	go func() {
		fmt.Println("routine start")
		time.Sleep(time.Second)
		fmt.Println("routine end")
	}()

	// Gosched yields the processor, allowing other goroutines to run.
	// It does not suspend the current goroutine, so execution resumes automatically.
	runtime.Gosched()
	fmt.Println("main end")
}
