// Concurrent requests for the same key block until the first completes.
// This implementation uses a Mutex.
package memo

import (
	"sync"
	"time"
)

type Memo struct {
	f  Func
	mu sync.Mutex
	// 使用指针的特殊值nil来表示key尚未被缓存(一种特殊状态)
	cache map[string]*entry
}

type Func func(string) (interface{}, error)

type entry struct {
	res   result
	ready chan struct{}
}

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// NOTE: 先站个座位
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		// 同一个key只会产生一个函数调用
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready // 见信号取值。
	}
	return e.res.value, e.res.err
}

func main12753() {
	mem := New(do)
	go mem.Get("1")
	go mem.Get("1")
	go mem.Get("2")
	time.Sleep(time.Second * 2)
}

func do(string) (interface{}, error) {
	time.Sleep(time.Second)
	return "hello", nil
}
