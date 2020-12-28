package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Base struct {
	age int
}

func (b Base) get() int {
	return b.age
}

func (b Base) getMore() int {
	return b.get()
}

type Child struct {
	Base
}

func (c Child) get() int {
	return 200
}

func TestInherit(t *testing.T) {
	child := Child{Base: Base{age: 20}}
	foo := func(b Base) int {
		return b.age
	}

	// 解释为什么输出Base的值20？
	// 1. 直接调用foo(c)会出错，因为存在类型转换的问题。也简介说明了，Base和Child本质上是两种类型，不存在继承关系。
	// 2. child 将它内部的 Base 变量赋值 给foo 的参数
	// 3. 所以肯定只会调用Base的方法
	v := foo(child.Base)
	assert.True(t, v == 20)

	// 解释为什么输出20, child.getMore() == child.Base.getMore()
	v = child.getMore()
	assert.True(t, v == 20)

	// Output:
	// 20
	// 20
}
