// 使用channel来实现deduplicate.go同样的效果
// NOTE: channel不一定非得关闭，只要不阻塞协程就好。
package memo

type Func func(string) (interface{}, error)

type Memo struct {
	requests chan request
}

type request struct {
	key  string
	resp chan result
}

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry) // NOTE: 将共享变量归属为一个协程的局部变量
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.delivery(req.resp) // 不能把server给阻塞住
	}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	req := request{key: key, resp: make(chan result)}
	memo.requests <- req
	res := <-req.resp
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

type entry struct {
	res   result
	ready chan struct{}
}

// 没有返回值，返回值是信号
func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

// 输入和删除都是channel
func (e *entry) delivery(resp chan<- result) {
	<-e.ready
	resp <- e.res
}
