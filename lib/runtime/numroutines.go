package main

import "fmt"
import "runtime"
import "time"

func myroutine() {
	fmt.Println("new routine")
}

func main() {
	fmt.Println(runtime.NumGoroutine())
	go myroutine()
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(1e9)
	/* the created routine will not be closed, event using Goexit */
	fmt.Println(runtime.NumGoroutine())
}
