// go test -race race.go    // to test the package
// go run -race race.go  // to run the source file
// go build -race race.go; ./race	// to build the command
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
