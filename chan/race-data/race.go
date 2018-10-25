// go test -race mypkg    // to test the package
// go run -race mysrc.go  // to run the source file
// go build -race mycmd   // to build the command
// go install -race mypkg // to install the package
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var (
	mu sync.Mutex
	a  = 10
)

// NOTE: WARNING: DATA RACE
func alice() {
	go func() {
		a = 20
	}()
	fmt.Println(a)
}

func bob() {
	var val atomic.Value
	val.Store(a)
	go func() {
		b := 20
		val.Store(b)
	}()

	time.Sleep(time.Millisecond * 100)
	fmt.Println(val.Load())
}

func jack() {
	go func() {
		mu.Lock()
		defer mu.Unlock()
		a = 20
	}()

	mu.Lock()
	b := a
	mu.Unlock()
	fmt.Println(b)
}

func main() {
	alice()
	// bob()
	// jack()
}
