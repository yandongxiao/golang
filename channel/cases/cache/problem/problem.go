// Concurrent requests are serialized by a Mutex
package problem

import "sync"

// 因为sync.Mutex, 方法必须采用pointer receiver
type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

type Func func(key string) (interface{}, error)

// result 存放函数执行的结果，如果结果有误，也会将错误保存下来
// value 的形式虽然是interface{}, 但是如果你使用引用类型（如指针），那么Get的返回值，会被同时读写，存在data race
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// 获取结果前一直 hold 住锁，会降低程序性能!
// 比如，key-x 的结果没有被缓存，key-y 的结果被缓存了。
// 如果Get(key-x)先执行，则阻塞了key-y的执行
func (memo *Memo) Get1(key string) (interface{}, error) {
	memo.mu.Lock()
	defer memo.mu.Unlock()

	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}

	return res.value, res.err
}

// 不同的 key 之间不会再被影响
// 在f(key)返回之前，同一个key的请求仍会导致f(key)被同时调用。
// 上面的问题，对于一个缓存系统来说，也是必须要解决的。尤其是当key是热点数据时。
func (memo *Memo) Get2(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()

	if !ok {
		res.value, res.err = memo.f(key)
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}
