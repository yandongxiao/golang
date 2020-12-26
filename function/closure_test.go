package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClosureVariable(t *testing.T) {
	foo := adder()
	bar := adder()
	assert.True(t, foo(1) == 1)
	assert.True(t, bar(1) == 1)
	assert.True(t, foo(20) == 21)
	assert.True(t, bar(20) == 21)
	assert.True(t, foo(300) == 321)
}

// 函数是一等变量（first class），可以作为参数或者返回值.
// 返回函数的函数称为高阶函数
func adder() func(int) int {
	// 闭包变量不会随着高阶函数Adder的返回而消失
	// 相反，它可以被返回的匿名函数访问
	num := 0
	return func(elm int) int {
		num += elm
		return num
	}
}
