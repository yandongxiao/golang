package buildin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecover(t *testing.T) {
	v := recover()
	assert.Nil(t, v)

	defer func() {
		if v := recover(); v != nil {
			assert.Equal(t, v, "hello")
		}
	}()
	panic("hello")
	assert.Equal(t, 1, 2)
}

func TestRecover_Duplicate(t *testing.T) {
	defer func() {
		if v := recover(); v != nil {
			assert.Equal(t, 1, 2)
		}
	}()

	defer func() {
		if v := recover(); v != nil {
			assert.Equal(t, v, "hello")
		} else {
			assert.Equal(t, 1, 2)
		}

	}()

	panic("hello")
}

// If recover is called outside the deferred function it will
// not stop a panicking sequence. In this case, or when the
// goroutine is not panicking, or if the argument supplied
// to panic was nil, recover returns nil.
func TestRecoverNil(t *testing.T) {
	defer func() {
		if v := recover(); v != nil {
			assert.Equal(t, 1, 2)
		}
	}()

	panic(nil)
	assert.Equal(t, 1, 2)
}
