// 1. 值类型是可以比较的，包括[3]int{}, struct类型不一定可以比较
// 2. 引用类型是否可比较，跟具体的引用类型有关系.
//		slice, map, function: can only be compared to nil
//		chan: 引用同一个chan, 则相等
//		interface: 是可以比较的
package main

import (
	"fmt"
)

func test1() {}
func test2() {}

func ExampleCompareNil() {
	// slice类型不可比较
	// 注意：[3]int{}是可以比较的，[]int{}是不可以比较的。
	// invalid operation: b == c (slice can only be compared to nil))
	// var b = []int{}
	// var c = []int{}
	// println(b == c)

	// map类型不可比较
	// invalid operation: bb == cc (map can only be compared to nil))
	// var bb = map[int]int{}
	// var cc = map[int]int{}
	// println(bb == cc)

	// 函数类型不可比较
	// invalid operation: test1 == test2
	// (func can only be compared to nil))
	// println(test1 == test2)

	// Output:
	//
}

func ExampleCompareChan() {
	// 引用同一个chan, 则相等
	var b = make(chan int)
	var c = b
	fmt.Println(b == c) // true
	c = make(chan int)
	fmt.Println(b == c) // false

	// Output:
	// true
	// false
}

func ExampleComparePointer() {
	p1 := new(int)
	p2 := p1
	fmt.Println(p1 == p2)
	// Output:
	// true
}

func ExampleCompareInterface() {
	// interface的比较规则如下：
	//	1. 首先比较的两个对象必须是同一种类型
	//	2. 如果底层类型不相同，则返回false
	//	3. 如果底层类型相同，但是类型不支持比较，如map, slice, 则**panic**
	//	4. 如果支持比较，则按照底层类型的比较规则进行比较.
	// 比较的是底层类型是否相等

	type IntArray1 [3]int
	type IntArray2 []int
	type IntArray3 = [3]int
	type IntArray4 = []int
	x1 := IntArray1{1, 2, 3}
	y1 := [3]int{1, 2, 3}
	fmt.Println(interface{}(x1) == interface{}(y1))

	x2 := IntArray2{1, 2, 3}
	y2 := []int{1, 2, 3}
	fmt.Println(interface{}(x2) == interface{}(y2))

	x3 := IntArray3{1, 2, 3}
	y3 := [3]int{1, 2, 3}
	fmt.Println(interface{}(x3) == interface{}(y3))

	defer func() {
		fmt.Println(recover())
	}()

	x4 := IntArray4{1, 2, 3}
	y4 := []int{1, 2, 3}
	fmt.Println(interface{}(x4) == interface{}(y4))

	// Output:
	// false
	// false
	// true
	// runtime error: comparing uncomparable type []int
}
