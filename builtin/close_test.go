// After the last value has been received from a closed channel c,
// any receive from c will succeed without blocking, returning the zero
// value for the channel element.
// NOTE: close函数只是用来关闭chan，与关闭文件操作无关
package buildin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCloseClosedChannel(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()

	ch := make(chan struct{})
	close(ch)
	close(ch) // panic here
	assert.Equal(t, 1, 2)
}

func TestCloseNilChan(t *testing.T) {
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()
	var ch chan int
	close(ch) // panic here
	assert.Equal(t, 1, 2)
}

func TestReceiveFromClosedChan(t *testing.T) {
	ch := make(chan int, 1)
	ch <- 10
	close(ch)

	val, ok := <-ch
	assert.Equal(t, val, 10)
	assert.Equal(t, ok, true)

	val, ok = <-ch
	assert.Equal(t, val, 0)
	assert.Equal(t, ok, false)
}
