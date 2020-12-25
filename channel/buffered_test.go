// 注意：Design your algorithm in the first place with unbuffered channels.
// 注意：Only introduce buffering when performance is problematic.
package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestUnbufferedChannel(t *testing.T) {
	defer func() {
		if v := recover(); v != nil {
			assert.Equal(t, 1, 1)
		}
		assert.Equal(t, 1, 2)
	}()

	// 默认创建的协程是不带缓存的，这导致无论是发送端还是接收端，发送或接收操作都是阻塞式.
	chMsg := make(chan string)
	chMsg <- "hello" // 导致唯一的协程也被阻塞了
	assert.Equal(t, 1, 2)
}

func TestBufferedChannel(t *testing.T) {
	chMsg := make(chan string, 2)
	chMsg <- "hello"
	chMsg <- "world"
	close(chMsg)

	assert.Equal(t, <-chMsg, "hello")
	assert.Equal(t, <-chMsg, "world")
	assert.Equal(t, <-chMsg, "")
}
