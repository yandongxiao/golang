// The cap built-in function returns the capacity of v, according to its type:
//	Array: the number of elements in v (same as len(v)).
//	Pointer to array: the number of elements in *v (same as len(v)).
//	Slice: the maximum length the slice can reach when resliced;
//	if v is nil, cap(v) is zero.
//	Channel: the channel buffer capacity, in units of elements;
//	if v is nil, cap(v) is zero.
// For some arguments, such as a simple array expression, the result can be a
// constant. See the Go language specification's "Length and capacity" section for
// details. 一般来讲，函数作用于常量，返回值不是常量
// map不能作为cap函数的参数，你可以认为map的capacity是无限大
package buildin

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceCap(t *testing.T) {
	var v []int
	assert.Equal(t, len(v), 0)
	assert.Equal(t, cap(v), 0)
}

func TestArrayCap(t *testing.T) {
	a := [3]int{1, 2, 3}
	assert.Equal(t, cap(a), 3)
}

func TestArrayPointerCap(t *testing.T) {
	p := &[3]int{1, 2, 3}
	fmt.Println(cap(p))

	// invalid argument q (type *[]int) for cap
	// q := &[]int{1, 2, 3}
	// fmt.Println(cap(q))
}

func TestReslice(t *testing.T) {
	a := [3]int{1, 2, 3}
	s := a[:2]
	assert.Equal(t, cap(s), 3)

	s = a[:]
	assert.Equal(t, cap(s), 3)

	s = a[0:0]
	assert.Equal(t, cap(s), 3)

	s = a[0:0:0]
	assert.Equal(t, cap(s), 0)
}

func TestChanCap(t *testing.T) {
	ch0 := make(chan int, 0)
	assert.Equal(t, cap(ch0), 0)

	ch1 := make(chan int, 1)
	assert.Equal(t, cap(ch1), 1)
}
