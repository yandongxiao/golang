package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Example_timerExpire() {
	// 一旦创建 timer 以后，timer 就开始工作，并在一秒后触发事件，并向timer.C中写时间
	timer := time.NewTimer(time.Millisecond * 10)
	<-timer.C
	fmt.Println("timer expired")

	// 注意：这个分支的代码没有被执行到
	// Stop prevents the Timer from firing.
	// It returns true if the call stops the timer, false if the timer has already
	// expired(属于这种情况) or been stopped.
	if stopped := timer.Stop(); stopped {
		fmt.Println("timer is stopped")
	}

	// Output:
	// timer expired
}

func Test_timerStop(t *testing.T) {
	// 一旦创建 timer 以后，timer 就开始工作，并在一秒后触发事件，并向timer.C中写时间
	timer := time.NewTimer(time.Millisecond)

	// 一旦执行 close the channel，会发生广播操作
	// 一旦执行 timer.C = nil, 那么 read 操作将会被永久性阻塞
	// Stop does not close the channel, to prevent a read from the channel succeeding(接下来) incorrectly.
	// To ensure the channel is empty after a call to Stop, check the return value and drain(排空) the channel.
	time.Sleep(time.Millisecond * 10) // 让 timer expired
	if !timer.Stop() {
		fmt.Println(<-timer.C) // 说明 timer.C != nil
		// fmt.Println(<-timer.C) // 再次调用会被阻塞, 说明 timer.C 没有被 close
		assert.True(t, true)
		return
	}
	assert.True(t, false)
}
