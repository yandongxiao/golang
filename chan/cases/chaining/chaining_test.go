/**
// The -8 suffix relates to the value of GOMAXPROCS that was used to run this test.
// 每次for循环运行的时间, 分别是62102506, 129597788, 54183

goos: darwin
goarch: amd64
pkg: github.com/yandongxiao/go/chan/cases/chaining
BenchmarkMethod1-8   	      20	  62102506 ns/op
BenchmarkMethod2-8   	       9	 129597788 ns/op
BenchmarkMethod3-8   	   22206	     54183 ns/op
PASS
ok  	github.com/yandongxiao/go/chan/cases/chaining	6.723s
*/
package chaining

import (
	"testing"
)

var ngoroutine = 100000

// 一个协程搭配一个chan，上一个协程工作完毕以后，触发下一个协程开始工作
func BenchmarkMethod1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		channels := make([]chan int, ngoroutine)
		for i := 0; i < ngoroutine; i++ {
			channels[i] = make(chan int)
			go func(i int) {
				if i == 0 {
					channels[i] <- 1
				} else {
					channels[i] <- 1 + <-channels[i-1]
				}
			}(i)
		}

		if ngoroutine != <-channels[ngoroutine-1] {
			b.Fatal("the result is wrong")
		}
	}
}

func BenchmarkMethod2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var init, in, out chan int
		// 初始化out
		init = make(chan int)
		out = init
		for i := 0; i < ngoroutine; i++ {
			// 每次遍历创建一个新的channel，它作为编号为 n 的 goroutine 的输出
			in, out = out, make(chan int)
			go func(in <-chan int, out chan<- int) {
				out <- 1 + <-in
			}(in, out)
		}

		// NOTE: 所有协程都处于阻塞状态，等待执行. 这是与 TestMethod1 的最重要的区别
		init <- 0
		if ngoroutine != <-out {
			b.Fatal("the result is wrong")
		}
	}
}

// 每个并发协程的工作量如果很少的话
// 如果创建大量的协程, 协程间的通信，会成为性能的瓶颈
func BenchmarkMethod3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		end := make(chan int)
		go func() {
			sum := 0
			for i := 0; i < ngoroutine; i++ {
				sum += 1
			}
			end <- sum
		}()
		if ngoroutine != <-end {
			b.Fatal("the result is wrong")
		}
	}
}
