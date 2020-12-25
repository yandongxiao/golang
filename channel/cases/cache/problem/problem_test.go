package problem_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/yandongxiao/go/chan/cases/cache/problem"
)

var sema = make(chan struct{}, 20)

/**
pkg: github.com/yandongxiao/go/chan/cases/cache/problem
BenchmarkGet1
BenchmarkGet1-8   	     100	  10409561 ns/op
*/
func BenchmarkGet1(b *testing.B) {
	memo := problem.New(foo)
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			select {
			case sema <- struct{}{}:
			}
			<-sema
			memo.Get1(fmt.Sprintf("%v", rand.Intn(1000)))
		}(i)
	}
	wg.Wait()
}

/**
pkg: github.com/yandongxiao/go/chan/cases/cache/problem
BenchmarkGet2
BenchmarkGet2-8   	 2717475	       421 ns/op
*/
func BenchmarkGet2(b *testing.B) {
	memo := problem.New(foo)
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			select {
			case sema <- struct{}{}:
			}
			<-sema
			memo.Get2(fmt.Sprintf("%v", rand.Intn(1000)))
		}(i)
	}
	wg.Wait()
}

func foo(key string) (interface{}, error) {
	time.Sleep(10 * time.Millisecond)
	return nil, nil
}
