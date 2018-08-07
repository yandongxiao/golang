package main

import (
	"fmt"
	"testing"
)

/* NOTE: the manner to call benchmark */
func main() {
	fmt.Println(" sync", testing.Benchmark(BenchmarkChannelSync).String())
	fmt.Println("buffered", testing.Benchmark(BenchmarkChannelBuffered).String())
}

func BenchmarkChannelSync(b *testing.B) {
	ch := make(chan int)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for _ = range ch {
	}
}

/* channel buffer : 1280 = 128 */
func BenchmarkChannelBuffered(b *testing.B) {
	ch := make(chan int, 1280)
	go func() {
		for i := 0; i < b.N; i++ {
			ch <- i
		}
		close(ch)
	}()
	for _ = range ch {
	}
}
