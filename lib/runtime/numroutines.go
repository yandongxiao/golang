package main

import "fmt"
import "runtime"
import "time"

func myroutine() {
	fmt.Println("new routine")
}

func main() {
	// get the number of logical CPUs available for the current program.
	// Each logical CPU can only execute one goroutine at any given time.
	// 8: 表示最多并行执行8个协程
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	go myroutine()
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(1e9)
	/* the created routine will not be closed, event using Goexit */
	fmt.Println(runtime.NumGoroutine())
}
