package main

import (
	"fmt"
	"time"
)

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
