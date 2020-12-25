package patterns

type counter struct {
	n     int
	mutex chan struct{}
}

func NewCounter(n int) counter {
	return counter{
		n:     n,
		mutex: make(chan struct{}, 1), // 1 buffered channel 是关键
	}
}

func (c *counter) Increase() {
	c.mutex <- struct{}{}
	defer func() {
		<-c.mutex
	}()
	c.n++
}

// 自己实现的mutex好处是，可以定义值类型的方法
func (c counter) Get() int {
	c.mutex <- struct{}{}
	defer func() {
		<-c.mutex
	}()
	return c.n
}
