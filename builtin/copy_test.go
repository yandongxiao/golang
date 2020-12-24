package buildin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeepCopySlice(t *testing.T) {
	a := []int{1, 2, 3}
	b := append(a[:0:0], a...) // append申请新的内存块的时机
	b[0] = 100
	assert.Equal(t, a, []int{1, 2, 3})
	assert.Equal(t, b, []int{100, 2, 3})
}

func TestCopyByteSlice(t *testing.T) {
	x := []byte("hello")
	y := make([]byte, 1)
	copy(y, x)
	assert.Equal(t, y, []byte{'h'})
}

func TestCopyFromNil(t *testing.T) {
	x := []byte("hello")
	copy(x, []byte(nil))
	assert.Equal(t, x, []byte("hello"))
}

func TestCopyToNil(t *testing.T) {
	x := []byte("hello")
	var y []byte
	copy(y, x)
	assert.Nil(t, y)
}

func TestCopyOverlap(t *testing.T) {
	x := []byte("hello")
	copy(x, x[2:])
	assert.Equal(t, x, []byte("llolo"))

	// 解释为什么第五个字符是l, 而不是h
	// 按照指针的理念:
	// llo, hello
	// hlo, hehlo
	// heo, heheo
	// heh, heheh
	x = []byte("hello")
	copy(x[2:], x)
	assert.Equal(t, x, []byte("hehel"))

	x = []byte("hello")
	copy(x[1:], x)
	assert.Equal(t, x, []byte("hhell"))
}

// NOTE: As a special case, it also will copy
// bytes from a string to a slice of bytes
func TestCopyStringToBytes(t *testing.T) {
	x := []byte("hello")
	copy(x, "WORLD")
	assert.Equal(t, x, []byte("WORLD"))
}
