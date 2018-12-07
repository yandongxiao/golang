// 常量可以是无类型的, 变量一定是有类型的
// 对于value type类型来说，type INT int 和int是两种不同的类型，所以需要强制类型转换；
// NOTE: 但是对于reference type，例如slice、map，只要它们底层的数据结构是一致的，无需进行强制类型转换。
package main

import "fmt"

type IntSlice [4]int

func (nums IntSlice) sum() int {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	return sum
}

func main() {
	// golang不支持隐式类型转换
	// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n
	var n int16 = 16
	var m int32 = 32
	m = int32(n)
	fmt.Printf("32 bit int is:  %d\n", m)
	fmt.Printf("16 bit int is:  %d\n", n)

	// nums2.sum undefined (type []int has no field or method sum))
	// fmt.Println(nums2.sum())
	// 从数据层面说：IntSlice 和 []int 的底层数据类型是完全一致的.
	// 从操作层面说：IntSlice 和 []int 不是同一种类型, IntSlice多了一些操作方法。
	// 这其实满足了给任意对象构造方法的述求！！
	nums := IntSlice{1, 2, 3, 4}
	fmt.Println(nums[0])
	fmt.Println(nums.sum())

	nums2 := [4]int{1, 2, 3, 4}
	nums = nums2 // NOTICE: 可以直接赋值
	nums2 = nums // NOTICE: 可以直接赋值
	fmt.Println(nums.sum())

	a := rune(1000)
	var b int32 = a
	println(b)
}
