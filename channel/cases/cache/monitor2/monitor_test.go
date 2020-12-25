package monitor2_test

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/yandongxiao/go/channel/cases/cache/monitor2"
)

var sema = make(chan struct{}, 20)

/**
pkg: github.com/yandongxiao/go/chan/cases/cache/monitor
BenchmarkGet2
BenchmarkGet2-8   	  532741	      2060 ns/op
这种方式并不快
*/
func BenchmarkGet2(b *testing.B) {
	memo := monitor2.New(foo)
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
