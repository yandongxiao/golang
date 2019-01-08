package main

import "fmt"

func ExampleSlice() {
	var arr1 [6]int
	var slice1 = arr1[2:5] // 2, 5-2, 6-2
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
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
