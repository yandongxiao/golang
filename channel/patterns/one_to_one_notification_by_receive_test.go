package patterns

import (
	"crypto/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneToOneNotificationByReceive(t *testing.T) {
	// 初始化随机数
	values := make([]byte, 1024*1024)
	if _, err := rand.Read(values); err != nil {
		t.Fatal(err)
	}

	// The capacity of the signal channel can also be one.
	// If this is true, then a value must be sent to the channel before
	// creating the following goroutine.
	done := make(chan struct{})

	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		<-done // by receive
	}()
	// Blocked here, wait for a notification.
	done <- struct{}{}

	for i := range values {
		if i == 0 {
			continue
		}
		assert.True(t, values[i] >= values[i-1])
	}
}

func TestOneToOneNotificationByReceive_2(t *testing.T) {
	// 初始化随机数
	values := make([]byte, 1024*1024)
	if _, err := rand.Read(values); err != nil {
		t.Fatal(err)
	}

	// buffered channel 带来的好处是解耦。即worker做完事情后可以理解退出。
	done := make(chan struct{}, 1)
	done <- struct{}{}

	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		<-done
	}()

	// Blocked here, wait for a notification.
	done <- struct{}{}

	for i := range values {
		if i == 0 {
			continue
		}
		assert.True(t, values[i] >= values[i-1])
	}
}
