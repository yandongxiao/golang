// Concurrent requests for the same key block until the first completes.
// This implementation uses a Mutex.
package mutex

import (
	"sync"
)

// 要解决该问题，可以在数据模型上定义出这么一种状态：该key确实不存在，但是已经有worker在处理了。
// 紧接着就有另外一个问题：如果worker处理完毕以后，它如何通知其它协程呢？ 所以，还要有消息传递机制在里面。
type Memo struct {
	f  Func
	mu sync.Mutex
	// 于是，我们可以用 *result 来表示那种状态。 即，使用指针的特殊值nil来表示key尚未被缓存
	// 但是，因为要实现通知机制，所以我们用 entry 对 result 包了一层。
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
		// 这条语句在执行时，肯定没有其它协程在读取该值。关键在 e.ready
		e.res.value, e.res.err = memo.f(key)
		close(e.ready)
	} else {
		memo.mu.Unlock()
		<-e.ready // 见信号取值。注意，不是去判断e.res的值是否为nil.
	}
	return e.res.value, e.res.err
}
