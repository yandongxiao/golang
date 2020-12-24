// 这是默写的结果
package monitor2

type Fun func(key string) (interface{}, error)

type Cache struct {
	Requests chan Request
}

type Request struct {
	key  string
	resp chan Result
}

type Result struct {
	value interface{}
	err   error
}

type Entry struct {
	end chan struct{}
	res Result
}

func New(f Fun) *Cache {
	c := &Cache{Requests: make(chan Request)}
	go c.Server(f)
	return c
}

func (c *Cache) Server(f Fun) {
	entries := make(map[string]*Entry)
	for req := range c.Requests {
		entry, ok := entries[req.key]
		if !ok {
			entry = &Entry{end: make(chan struct{})}
			entries[req.key] = entry
			go entry.call(f, req.key)
		}
		go entry.notify(req.resp)
	}
}

func (e *Entry) call(f Fun, key string) {
	e.res.value, e.res.err = f(key)
	close(e.end)
}

func (e *Entry) notify(resp chan Result) {
	<-e.end
	resp <- e.res
}

func (c *Cache) Get2(key string) (interface{}, error) {
	request := Request{
		key:  key,
		resp: make(chan Result),
	}

	c.Requests <- request
	res := <-request.resp
	return res.value, res.err
}
