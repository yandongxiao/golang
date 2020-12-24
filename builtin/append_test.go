// func append(slice []Type, elems ...Type) []Type
package buildin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendToNil(t *testing.T) {
	var b []byte
	b = append(b, "hello"...) // sugar
	assert.Equal(t, b, []byte("hello"))
}

func TestAppendFromNil(t *testing.T) {
	var data []byte
	data = append(data, []byte{4, 5, 6}...)

	// append 语法处理了零值的情况
	data = append(data, []byte(nil)...)
	data = append(data)

	data = append(data, 7, 8, 9)
	data = append(data, "xyz"...)
	data = append(data)
	assert.Equal(t, data, []byte{4, 5, 6, 7, 8, 9, 'x', 'y', 'z'})
}

func TestAppendShare(t *testing.T) {
	data := make([]int, 0, 10)
	assert.Equal(t, len(data), 0)
	assert.Equal(t, cap(data), 10)

	data2 := data[5:10:10] // len=10-5, cap=10-5
	for i := 0; i < 10; i++ {
		data = append(data, i)
	}
	assert.Equal(t, data2, []int{5, 6, 7, 8, 9})
	fmt.Println(data2)
}

// 采用 Example 更合理
// NOTE: for-range 这种删除元素的方式不可取
func ExampleRemoveElementByAppend() {
	defer func() {
		fmt.Println(recover())
	}()

	m := []int{1, 2, 3}

	// NOTE: for range 的遍历次数由 m 的** 初始值 **决定
	for i := range m {
		if i <= 2 {
			fmt.Println(m)
			m = append(m[:i], m[i+1:]...)
			fmt.Println(m)
		}
	}

	// Output:
	// [1 2 3]
	// [2 3]
	// [2 3]
	// [2]
	// [2]
	// runtime error: slice bounds out of range [3:1]
}

// 原地删除 - 1
// NOTE: 倒序遍历时，删除当前元素是安全的
func ExampleRemoveElementByReverseOrder() {
	// 正确1
	m := []int{1, 2, 3}
	for i := len(m) - 1; i >= 0; i-- {
		if i <= 2 {
			m = append(m[:i], m[i+1:]...)
			fmt.Println(m)
		}
	}

	// Output:
	// [1 2]
	// [1]
	// []
}

// 原地删除 - 2
func TestRemoveElementBySharedMemory(t *testing.T) {
	m := []int{1, 2, 3}
	n := m[:0] // len(n)=0, cap(n)=3, 与m共享同一份数据(没有浪费内存)
	for _, v := range m {
		if v > 2 { // 相当于一个过滤器
			n = append(n, v)
		}
	}
	assert.Equal(t, n, []int{3})
}
