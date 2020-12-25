package patterns

import (
	"crypto/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOneToOneNotification(t *testing.T) {
	// 初始化随机数
	values := make([]byte, 1024*1024)
	if _, err := rand.Read(values); err != nil {
		t.Fatal(err)
	}

	// 一对一的通知：使用unbuffered channel, 进行数据传递
	done := make(chan struct{})
	// The sorting goroutine. channel作为输入和输出参数
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{}
	}()

	<-done

	for i := range values {
		if i == 0 {
			continue
		}
		assert.True(t, values[i] >= values[i-1])
	}
}
