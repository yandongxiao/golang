// 常量可以是无类型的, 变量一定是有类型的
// 对于value type类型来说，type INT int和int是不同的类型，需要强制类型转换；
// 对于reference type(还有数组类型)，例如slice、map，只要它们底层的数据结构是一致的，
// 无需进行强制类型转换。
package main

import "fmt"

// type IntSlice [4]int  	// 数组类型也无需进行强制类型转换
type IntSlice []int

func (nums IntSlice) sum() int {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	return sum
}

func ExampleValueCast() {
	// golang不支持隐式类型转换
	// compiler error: cannot use n (type int16) as type
	// int32 in assignment
	//m = n
	var n int16 = 16
	var m int32 = 32
	m = int32(n)
	fmt.Println(m, n)

	// Output:
	// 16 16
}

func ExampleStructCast() {
	type S struct{ name string }
	s := S{name: "jack"}
	type NS S
	var ns NS
	// cannot use s (type S) as type NS in assignment
	ns = NS(s)
	fmt.Println(ns.name)
	// Output:
	// jack
}

func ExampleSliceCast() {
	// nums2.sum undefined (type []int has no field or method sum)
	// fmt.Println(nums2.sum())
	// 从数据层面说：IntSlice 和 []int 的底层数据类型是完全一致的.
	// 从操作层面说：IntSlice 和 []int 不是同一种类型, IntSlice多了一些操作方法。
	// 这其实满足了给任意对象构造方法的诉求！！
	var nums IntSlice
	nums2 := []int{1, 2, 3, 4}
	nums = nums2 // NOTE: 可以直接赋值
	nums2[0] = 11
	fmt.Println(nums.sum())
	// Output:
	// 20
}

func ExampleMapCast() {
	m := map[int]int{1: 1}
	type Map map[int]int
	var nm Map
	nm = m
	fmt.Println(nm[1])
	// Output:
	// 1
}

func ExampleChanCast() {
	type Chan chan int
	ch := make(chan int, 1)
	var nch Chan
	nch = ch
	ch <- 1
	fmt.Println(<-nch)
	// Output:
	// 1
}
