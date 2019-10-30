package main

import "fmt"

func ExampleOrder() {
	defer func() {
		fmt.Println(recover())
	}()
	m := []int{1, 2, 3}
	for i := range m { // NOTE: for range 的遍历次数由m的** 初始值 **决定
		if i <= 2 {
			m = append(m[:i], m[i+1:]...)
		}
	}
	// Output:
	// runtime error: slice bounds out of range
}

func ExampleReverseOrder() {
	// 正确1
	m := []int{1, 2, 3}
	for i := len(m) - 1; i >= 0; i-- { // NOTE: 倒序
		if m[i] <= 2 {
			m = append(m[:i], m[i+1:]...)
		}
	}
	fmt.Println(m)

	// Output:
	// [3]
}

func ExampleAdd() {
	// 正确2
	m := []int{1, 2, 3}
	n := m[:0] // len(n)=0, cap(n)=3, 与m共享同一份数据(没有浪费内存)
	for _, v := range m {
		if v > 2 {
			n = append(n, v) // 通过构造一个新的slice来达到删除元素的目的
		}
	}
	fmt.Println(n)

	// Output:
	// [3]
}

func ExampleMap() {
	// This is safe
	// 另外两种情况:
	// If map entries that have not yet been reached are removed during iteration, the corresponding iteration values will not be produced.
	// If map entries are created during iteration, that entry may be produced during the iteration or may be skipped.
	mm := map[string]int{"jack": 10, "bob": 20}
	for key := range mm {
		if key == "jack" {
			delete(mm, key)
		}
	}
	fmt.Println(mm)

	// Output:
	// map[bob:20]
}
