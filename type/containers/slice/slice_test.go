// for _, buf := range buffers. 如果buffers中元素很多，则**不建议**使用这种方法
package main

import "fmt"
import "unsafe"

func ExampleSliceNil() {
	var a []int
	b := []int(nil)
	c := []int{}

	fmt.Println(a == nil && b == nil && unsafe.Sizeof(c) == 24)
	//Output:
	// true
}

func ExampleSliceAddressability() {
	// Elements of any slice value are always addressable, whether
	// or not that slice value is addressable.
	ps0 := &[]string{"Go", "C"}[0]
	fmt.Println(*ps0) // Go

	// Elements of addressable array values are also addressable.
	// Elements of unaddressable array values are also unaddressable.
	// The reason is each array value only consists of one direct part.
	// _ = &[3]int{2, 3, 5}[0]

	// Output:
	// Go

}

func ExampleSlice() {
	// 对切片的要求，可见low不一定要比len(baseContainer)小
	// 0 <= low <= high <= cap(baseContainer)        // two-index form
	// 0 <= low <= high <= max <= cap(baseContainer) // three-index form
	var arr1 [6]int
	var slice1 = arr1[2:5] // 2, 5-2, 6-2
	var i int32            // 作为下标，不一定非得是int类型
	for i = 0; i < 6; i++ {
		arr1[i] = int(i)
	}
	fmt.Println("array state")
	fmt.Println(len(arr1))
	fmt.Println(arr1)

	fmt.Println("slice state")
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(slice1)

	// grow the slice:
	fmt.Println("grow slice")
	slice1 = slice1[0:4]
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	fmt.Println(slice1)
	// grow the slice beyond capacity:
	// slice1 = slice1[0:7 ] // panic: runtime error: slice bounds out of range

	fmt.Println("grow slice2")
	slice1 = append(slice1, 1) // 执行append之后，arr1和slice1再无关系
	slice1[0] = 100
	fmt.Println(slice1)
	fmt.Println(arr1)

	// Output:
	// array state
	// 6
	// [0 1 2 3 4 5]
	// slice state
	// 3
	// 4
	// [2 3 4]
	// grow slice
	// 4
	// 4
	// [2 3 4 5]
	// grow slice2
	// [100 3 4 5 1]
	// [0 1 2 3 4 5]
}
