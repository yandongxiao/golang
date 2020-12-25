// race的影响：
// NOTE: memory usage may increase by 5-10x and execution time by 2-20x.

// The race detector writes its report to a file named log_path.pid
// The special names stdout and stderr cause reports to be written
// to standard output and standard error, respectively.
// export GORACE="log_path=/tmp/gr.err"
// export GORACE="log_path=stdout"
// go run -race race.go

// go test -race race.go    // to test the package
// go run -race race.go  // to run the source file
// go build -race race.go; ./race	// to build the command
package main

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	mu sync.Mutex
	a  = 10
)

// NOTE: DATA RACE
func alice() {
	go func() {
		a = 20
	}()
	fmt.Println(a)
}

// NOTE: DATA RACE on loop counter
func bob() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i) // Not the 'i' you are looking for.
			wg.Done()
		}()
	}
	wg.Wait()
}

// 这个也是常见错误，主协程和子协程尽量不用共用变量，比如 error
// ParallelWrite writes data to file1 and file2, returns the errors.
func cindy(data []byte) chan error {
	res := make(chan error, 2)
	f1, err := os.Create("/tmp/xxfile1")
	if err != nil {
		res <- err
	} else {
		go func() {
			// This err is shared with the main goroutine,
			// so the write races with the write below.
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}

	f2, err := os.Create("/tmo/xxfile2")
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}

	return res
}

/**
BenchmarkAtomicValue
BenchmarkAtomicValue-8   	  873064	      1353 ns/op
*/
// NOTE: 通过atomic.Value实现GET、SET的原子化
func BenchmarkAtomicValue(b *testing.B) {
	f, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}

	for i := 0; i < b.N; i++ {
		var val atomic.Value
		val.Store(a)
		go func() {
			b := 20
			val.Store(b)
		}()
		fmt.Fprint(f, val.Load())
	}
}

/**
BenchmarkMutex
BenchmarkMutex-8   	  741564	      1543 ns/op
*/
func BenchmarkMutex(b *testing.B) {
	f, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}
	for i := 0; i < b.N; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			a = 20
		}()

		mu.Lock()
		fmt.Fprint(f, a)
		mu.Unlock()
	}
}
