package test

import "testing"

func TestAdd_WithCleanup(t *testing.T) {
	defer func() {
		t.Log("defer-1") // 4
	}()

	// 它们是在defer之后执行
	t.Cleanup(func() {
		t.Log("clean up") // 6
	})

	defer func() {
		t.Log("defer-2") // 3
	}()

	t.Log("before handle") // 1
	handle(t)
	t.Log("after handle") // 2
}

func handle(t *testing.T) {
	// t.Cleanup与defer的区别: 不是在handle函数返回之前被调用, 而是在TestAdd_WithCleanup结束前被调用
	t.Cleanup(func() {
		t.Log("handle: clean up") // 5
	})
}
