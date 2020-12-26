package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type addFunc func(a, b int) int

// 为什么要使用 var add 的方式，而不直接使用 func add 呢 ？
// func add 声明了一个函数常量
var add = func(a, b int) int {
	if a == 0 {
		panic("a==0")
	}
	return a + b
}

func init() {
	add = foo(add)
}

// 高阶函数foo的作用: 为函数 add 增加 defer + recover 的保护
// 这里隐藏了高阶函数的一个应用，面向切面编程(AOP)：
//	1. foo可以为所有函数提供 defer + recover 的保护
//  2. 在init中添加add := foo(add)语句，
//     原始的add再也不能被访问到，必须访问封装后的函数
func foo(add addFunc) addFunc {
	return func(a, b int) (c int) {
		defer func() {
			if err := recover(); err != nil {
				c = -1
			}
		}()
		return add(a, b)
	}
}

func TestAOP(t *testing.T) {
	assert.True(t, add(1, 2) == 3)
	assert.True(t, add(0, 2) == -1)
}
