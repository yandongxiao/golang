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

func TestRecoverNil(t *testing.T) {
	defer func() {
		if v := recover(); v != nil {
			assert.Equal(t, 1, 2)
		}
	}()

	panic(nil)
	assert.Equal(t, 1, 2)
}
