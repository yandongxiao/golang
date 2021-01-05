package time

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("helloworld")
}

func ExampleAfterFunc() {
	// AfterFunc waits for the duration to elapse and then calls f
	// in its own goroutine. It returns a Timer that can
	// be used to cancel the call using its Stop method.
	timer := time.AfterFunc(10*time.Millisecond, f)
	time.Sleep(time.Second)

	// For a timer created with AfterFunc(d, f), if t.Stop returns false, then the timer
	// has already expired and the function f has been started in its own goroutine;
	// Stop does not wait for f to complete before returning.
	if timer.Stop() {
		fmt.Println("stop the timer successfully")
	} else {
		fmt.Println("the timer has expired")
	}

	// Output:
	// helloworld
	// the timer has expired
}

func ExampleAfterFunc_2() {
	timer := time.AfterFunc(10*time.Millisecond, f)
	if timer.Stop() {
		fmt.Println("stop the timer successfully")
	} else {
		fmt.Println("the timer has expired")
	}

	// Output:
	// stop the timer successfully
}
