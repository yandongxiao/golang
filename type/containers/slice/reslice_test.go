package main

import "fmt"

func ExampleSlice0() {
	array := [5]int{1, 2, 3, 4, 5}
	nums := array[:]
	fmt.Println(len(nums), cap(nums))

	v := nums[0:len(nums):cap(nums)]
	fmt.Println(v)

	v = nums[:]
	fmt.Println(v)
	//Output:
	// 5 5
	// [1 2 3 4 5]
	// [1 2 3 4 5]
}

func ExampleSlice1() {
	array := [6]int{0, 1, 2, 3, 4, 5}
	nums := array[:]
	v := nums[2:len(nums):cap(nums)]
	fmt.Println(v)

	v = nums[2:]
	fmt.Println(v)
	//Output:
	// [2 3 4 5]
	// [2 3 4 5]
}

func ExampleSlice2() {
	nums := []int{0, 1, 2, 3, 4, 5}
	v1 := nums[1:2] // 等价于：1, 2-1=1, cap(nums)-1=5
	fmt.Println(v1, len(v1), cap(v1))
	//Output:
	// [1] 1 5
}

func ExampleSlice3() {
	nums := []int{0, 1, 2, 3, 4, 5}
	v1 := nums[1:2:3] // 1, 2-1=1, 3-1=2
	fmt.Println(v1, len(v1), cap(v1))
	//Output:
	// [1] 1 2
}

func ExampleEmptySlice() {
	nums := []int{0, 1, 2, 3, 4, 5}
	v3 := nums[:0] // 0, 0-0=0, 6
	v3 = append(v3, 100)
	fmt.Println(v3, len(v3), cap(v3), nums) // NOTE: nums被修改了

	v4 := nums[:0:0] // make v4 as an empty slice
	v4 = append(v4, 200)
	fmt.Println(v4, len(v4), nums) // NOTE: nums未修改, cap(v4)未知
	//Output:
	// [100] 1 6 [100 1 2 3 4 5]
	// [200] 1 [100 1 2 3 4 5]
}

func ExampleOutOfRange() {
	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	nums := []int{0, 1, 2, 3, 4, 5}
	nv := nums[1:2:3]
	fmt.Println(nv[0:3]) // 注意: nv[0:2] 就是合法的
	// Output:
	// runtime error: slice bounds out of range [:3] with capacity 2
}
