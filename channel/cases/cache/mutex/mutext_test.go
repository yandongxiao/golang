package mutex_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/yandongxiao/go/chan/cases/cache/mutex"
)

var sema = make(chan struct{}, 20)

/**
BenchmarkGet2
BenchmarkGet2-8   	 3055131	       367 ns/op
*/
func BenchmarkGet2(b *testing.B) {
	memo := mutex.New(foo)
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			select {
			case sema <- struct{}{}:
			}
			<-sema
			memo.Get(fmt.Sprintf("%v", rand.Intn(1000)))
		}(i)
	}
	wg.Wait()
}

func foo(key string) (interface{}, error) {
	time.Sleep(10 * time.Millisecond)
	return nil, nil
}
