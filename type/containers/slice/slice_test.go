// for _, buf := range buffers. 如果buffers中元素很多，则**不建议**使用这种方法
package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNilSlice(t *testing.T) {
	var a []int
	b := []int(nil)
	c := []int{}

	assert.True(t, a == nil)
	assert.True(t, b == nil)
	assert.True(t, c != nil)
}

// 一个微小的差别
func TestNilSliceIndex(t *testing.T) {
	var a []int
	if a[:0] != nil {
		t.Fatal("a[:0] != nil")
	}

	b := []int{1}
	if b[:0] == nil {
		t.Fatal("b[:0] == nil")
	}
}

func TestSliceAddressability(t *testing.T) {
	// Elements of any slice value are always addressable, whether
	// or not that slice value is addressable.
	ps0 := &[]string{"Go", "C"}[0]
	assert.True(t, ps0 != nil)

	// 数组
	// Elements of addressable array values are also addressable.
	// Elements of unaddressable array values are also unaddressable.
	// The reason is each array value only consists of one direct part.
	// _ = &[3]int{2, 3, 5}[0]
}

func TestSlice(t *testing.T) {
	// 对切片的要求，可见low不一定要比len(baseContainer)小
	// 0 <= low <= high <= cap(baseContainer)        // two-index form
	// 0 <= low <= high <= max <= cap(baseContainer) // three-index form
	var arr1 [6]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	assert.Equal(t, 6, len(arr1))

	var slice1 = arr1[2:5] // 2, 5-2, 6-2
	assert.Equal(t, 3, len(slice1))
	assert.Equal(t, 4, cap(slice1))

	// grow the slice:
	slice1 = slice1[0:4]
	assert.Equal(t, 4, len(slice1))
	assert.Equal(t, 4, cap(slice1))
	// grow the slice beyond capacity:
	// slice1 = slice1[0:7 ] // panic: runtime error: slice bounds out of range

	slice1 = append(slice1, 1) // 执行append之后，arr1和slice1再无关系
	slice1[0] = 100
	assert.NotEqual(t, slice1[0], arr1[0])
	fmt.Println(slice1)
}
