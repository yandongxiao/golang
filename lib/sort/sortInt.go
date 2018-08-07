package main

import (
	"fmt"
	"sort"
)

func main() {
	// 排序基本数据类型
	data := []int{4, 3, 5, 1}
	sort.Ints(data)
	fmt.Println(data)

	// 一、sort.IntSlice
	// type IntSlice []int
	// IntSlice attaches the methods of Interface to []int, sorting in increasing order.
	// func (p IntSlice) Len() int
	// func (p IntSlice) Less(i, j int) bool
	// func (p IntSlice) Swap(i, j int)
	// 所以IntSlice实现了sort.Interface接口的方法
	// type Interface interface {
	//		Len is the number of elements in the collection.
	//		Len() int
	//		Less reports whether the element with
	//		index i should sort before the element with index j.
	//		Less(i, j int) bool
	//		Swap swaps the elements with indexes i and j.
	//		Swap(i, j int)
	//}
	//
	// 二、sort.Reverse
	// 根据函数声明func Reverse(data Interface) Interface)，我们知道该方法并没有执行真的排序动作，只是返回了一个新的sort.Interface.
	// 新的sort.Interface会将逆序的因子考虑进去
	//
	// 三、执行通用的排序动作
	sort.Sort(sort.Reverse(sort.IntSlice(data))) // sort.Interface
	//data = []int(rdata)
	fmt.Println(data)
}
