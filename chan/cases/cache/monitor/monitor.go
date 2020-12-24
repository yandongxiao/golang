// 使用 channel 来实现 deduplicate.go 同样的效果
// NOTE: channel不一定非得关闭，只要不阻塞协程就好。
package monitor

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

// 我们使用单独的协程，负责与远端交互。
// 它接收请求，go 请求远端，远端一旦有结果，它需要通知给请求方。
func (memo *Memo) server(f Func) {
	// NOTE: 将共享变量归属为一个协程的局部变量
	cache := make(map[string]*entry)

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

// 对外暴露的是同步接口，这个很重要 ！
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

// 执行远程调用，通知
func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

// 专门负责通知
func (e *entry) delivery(resp chan<- result) {
	<-e.ready
	resp <- e.res
}
