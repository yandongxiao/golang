package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(100 * time.Millisecond)
	i := 0
	for range ticker.C {
		// 如果处理时间超过了100毫秒，那么期间产生的ticker event会丢失
		// time.Sleep(500 * time.Millisecond)
		i++
		if i > 3 {
			return
		}
		assert.True(t, true)
	}
	assert.True(t, false)
}
