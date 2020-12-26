// A type assertion provides access to an interface value's
// underlying concrete value.
package main

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeAssertion(t *testing.T) {
	v := 10
	var i interface{} = v
	i = 20 // NOTE: 该操作并没有修改原始变量v的值
	assert.True(t, v == 10)
	assert.True(t, i == 20)

	// If i does not hold a T, the statement will trigger a panic.
	// If i holds a T, then t will be the underlying value and
	// ok will be true.
	if num, ok := i.(int); ok {
		assert.True(t, ok)
		assert.True(t, num == 20)
	} else {
		assert.True(t, false)
	}

	// If not, ok will be false and t will be the zero value of
	// type T, and no panic occurs.
	if _, ok := i.(int32); !ok {
		assert.False(t, ok)
	} else {
		assert.True(t, false)
	}
}

func TestTypeAssertion_2(t *testing.T) {
	f, err := os.Open("/tmp/dd")
	assert.Nil(t, err)

	var writeCloser io.WriteCloser
	writeCloser = f

	// 如果接口类型之间存在包含与被包含的关系，可以不借助 type assertion，
	// 直接进行赋值. 注意，反向不可以，需要使用type assertion
	var writer io.Writer
	writer = writeCloser

	// 注意：type assertion: interfaceA 和 InterfaceB 之间
	// 是否可以转换，只与底层的类型变量是否同时实现了他们的接口有关。
	s, ok := writer.(io.Reader)

	assert.True(t, true)
	fmt.Printf("%T, %v\n", s, ok) // 返回具体的类型
}
