// 不同类型的变量之间是不能进行计算，即便是int32和int64之间
// 常量可以是无类型的, 变量一定是有类型的
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
	nums := IntSlice{1, 2, 3, 4}
	fmt.Println(nums[0])
	fmt.Println(nums.sum())

	nums2 := [4]int{1, 2, 3, 4}
	// nums2.sum undefined (type []int has no field or method sum))
	// fmt.Println(nums2.sum())

	// 从数据层面说：IntSlice 和 []int 是同一种类型, 因为可以相互赋值
	// 从操作层面说：IntSlice 和 []int 不是同一种类型, IntSlice多了一些操作方法。
	// 这其实满足了给任意对象构造方法的述求！！
	nums = nums2
	fmt.Println(nums.sum())
}
