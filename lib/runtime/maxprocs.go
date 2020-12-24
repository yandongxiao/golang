package main

import (
	"fmt"
	"runtime"
)

func ExampleB() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(3)
}
