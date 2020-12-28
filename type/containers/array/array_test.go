// In Go, an array is a numbered sequence of elements of a specific length.
// Array elements can neither be appended nor deleted, though elements of
// addressable arrays can be modified.

package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroArray(t *testing.T) {
	a := [0]int{}
	assert.Equal(t, len(a), 0)
}

func TestArrayModification(t *testing.T) {
	a := [1]struct {
		name string
	}{{"jack"}}

	// 注意：1. The ranged container is a copy of aContainer. Please note, only the direct part of aContainer is copied.
	// 注意：The copied container direct part is anonymous, so there are no ways to modify it. 你引用不到这个array
	// All key-element pairs will be assigned to the same iteration variable pair.
	for _, s := range a {
		s.name = "bob"
	}
	assert.Equal(t, a[0].name, "jack")

	// 取地址的好处是省去了数组的拷贝
	// 但是s还是与数组a的元素无关
	for _, s := range &a {
		s.name = "bob"
	}
	assert.Equal(t, a[0].name, "jack")
}

func TestArrayModification2(t *testing.T) {
	a := [1][]int{
		{1, 2, 3},
	}

	// s 和 a[0] 共享了底层数据，但是append操作以后，它们可能指向了不同的内存块
	for _, s := range a {
		s = append(s, 4, 5, 6)
	}
	assert.Equal(t, 3, len(a[0]))
	assert.Equal(t, 3, cap(a[0]))
}

func ExamplePointerIter() {
	var p *[2]int // nil

	for i := range p { // okay
		fmt.Println(i)
	}

	defer func() {
		fmt.Println(recover())
	}()
	for i, n := range p { // panic, NOTE, 非指针类型的情况下，不会发生panic
		fmt.Println(i, n)
	}
	// Output:
	// 0
	// 1
	// runtime error: invalid memory address or nil pointer dereference
}

func ExampleInitilize() {
	// Use this syntax to declare and initialize an array in one line.
	var arrAge = [5]int{18, 20, 15, 22, 16}            // literal-1
	var arrLazy = [...]int{5, 6, 7, 8, 22}             // literal-2
	var arrKeyValue = [10]string{3: "Chris", 4: "Ron"} // literal-3
	var arr4 = [...]string{3: "Chris", 4: "Ron"}       // literal-4
	// var arrLazy = []int{5, 6, 7, 8, 22}			// 只要[]内什么都没有，返回的类型就是slice
	fmt.Printf("%T, %T, %T, %T\n", arrAge, arrLazy, arrKeyValue, arr4)

	// Output:
	// [5]int, [5]int, [10]string, [5]string
}
