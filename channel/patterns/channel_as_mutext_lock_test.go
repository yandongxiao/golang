package patterns_test

import (
	"testing"

	"github.com/magiconair/properties/assert"

	"github.com/yandongxiao/go/chan/patterns"
)

func TestChannelAsMutexLock(t *testing.T) {
	// The capacity must be one.
	counter := patterns.NewCounter(0)

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			counter.Increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done

	assert.Equal(t, counter.Get(), 2000)
}
