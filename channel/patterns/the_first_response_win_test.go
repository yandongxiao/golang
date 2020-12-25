package patterns

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TheFirstResponseWin struct{}

func (thrw TheFirstResponseWin) call(resp chan<- int) {
	rb := rand.Intn(3) + 1
	// Simulate a workload, and sleep 1s/2s/3s.
	time.Sleep(time.Duration(rb) * 100 * time.Millisecond)
	resp <- rb
}

// 应用场景: 希望尽快拿到响应结果，那么多个协程同时请求一个服务。
func TestTheFirstResponseWin(t *testing.T) {
	thrw := TheFirstResponseWin{}
	rand.Seed(time.Now().UnixNano())

	// NOTE: c must be a buffered channel.
	var wg sync.WaitGroup
	resp := make(chan int, 5)
	for i := 0; i < cap(resp); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			thrw.call(resp)
		}()
	}
	wg.Wait()
	close(resp)

	// Only the first response will be used.
	max := -1
	for v := range resp {
		assert.True(t, max <= v)
		max = v
	}
}
