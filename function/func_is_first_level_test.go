package function

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// function signature: is composed of two type list, one is the input parameter type list, the other is the output result type lists.
// function type: is composed of the func keyword and a function signature literal.
// function type literal: is composed of the func keyword and a function signature literal.
//
// function prototype: func Double(n int) (result int)
// function declaration: func Double(n int) (result int){}. 匿名函数不算作函数声明，函数声明只能是package level
// When we declare a custom function, we also declared an immutable function value actually.

// 尝试回答一个问题，什么是一等公民？

// 可以：定义了一个函数类型
type MyFunc func(arg int) int

// 可以：定义一个函数类型的变量
var f MyFunc = func(arg int) int {
	return arg
}

// 可以：定义一个函数类型的常量
func myFoo(arg int) int {
	return 0
}

// 可以： 作为函数的参数
type MyBar func(f MyFunc) MyFunc

// 可以作为 struct 的成员变量
type s struct {
	myFunc MyFunc
}

func TestFuncIsFirstLevel(t *testing.T) {
	assert.True(t, true)
}
