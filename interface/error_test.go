package main

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type bar struct{}

func (b *bar) Error() string {
	return "bar"
}

// 这个函数的返回值是error interface类型, 返回值不为nil !
// error is an interface, and *bar implement it
// rule: if err == nil; means the underlying type is nil, and the underlying value is nil.
// err = b; means the underlying type is not nil.
func foo() error {
	var b *bar
	return b // b == nil
}

func TestNilError(t *testing.T) {
	err := foo()
	if err != nil {
		b := err.(*bar)
		assert.True(t, b == nil)
	} else {
		assert.True(t, false)
	}

	// Output:
	// true
	// bar
}

func TestNilInterfaceTypeSwitch(t *testing.T) {
	// reader := io.Reader(nil).
	// 等价于 var reader io.Reader
	myfunc(t, nil) // NOTE: 具体类型和具体实例都是nil
	// Output:
	// reader is nil
	// not ok
}

func myfunc(t *testing.T, reader io.Reader) {
	if reader == nil {
		assert.True(t, true)
	}

	// reader == nil的情况下，切换interface type不成功也是合理的
	_, ok := reader.(io.ReadCloser)
	if ok {
		assert.True(t, false)
	} else {
		assert.True(t, true)
	}
}
