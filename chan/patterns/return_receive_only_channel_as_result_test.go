package patterns

import (
	"math/rand"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
)

type request struct {
	num  int
	resp chan int
}

func NewRequest(num int) *request {
	return &request{
		num:  num,
		resp: make(chan int),
	}
}

func (r *request) call() {
	// simulate a workload.
	time.Sleep(time.Millisecond * 100)
	r.resp <- 2 * r.num
	close(r.resp)
}

func Get(r *request) <-chan int {
	go r.call()
	return r.resp
}

func SumSquares(a, b int) int {
	return a + b
}

func TestReturnReceiveOnlyChannelAsResult(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	r1 := NewRequest(100)
	r2 := NewRequest(1)

	assert.Equal(t, SumSquares(<-Get(r1), <-Get(r2)), 202)
}
