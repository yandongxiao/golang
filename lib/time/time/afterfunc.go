package main

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("helloworld")
}

func main() {
	timer := time.AfterFunc(3e9, f)
	time.Sleep(1e9)
	if timer.Stop() {
		fmt.Println("stop the timer successfully")
	} else {
		fmt.Println("the timer has expired")
	}
}
