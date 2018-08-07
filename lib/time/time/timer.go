package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3e9)
	fmt.Println(<-timer.C) /* wait for the time eclipse */
	if timer.Stop() {
		fmt.Println("stop the timer successfully")
	} else {
		fmt.Println("the timer has expired")
	}
}
