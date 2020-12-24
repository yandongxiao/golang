package buildin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// make 分配的内存是不确定的
func TestMakeMap(t *testing.T) {
	// creates a new empty map with enough space to hold a small
	// number of entries without reallocating memory again.
	// The small number is compiler dependent.
	m := make(map[int]int)
	assert.Equal(t, len(m), 0)

	// map 不支持 cap 操作
	// assert.True(t, cap(m) > 0)
	// creates a new empty map which is allocated with enough space
	// to hold at least n entries without reallocating memory again.
	m = make(map[int]int, 10)
	assert.Equal(t, len(m), 0)
}

func TestMakeSlice(t *testing.T) {
	// the capacity of the new created slice is the same as its length.
	m := make([]int, 10)
	assert.Equal(t, len(m), 10)

	m = make([]int, 0, 10)
	assert.Equal(t, len(m), 0)
}
