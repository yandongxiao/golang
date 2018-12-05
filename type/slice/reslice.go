package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5}

	v1 := nums[1:2]                   // idx + len
	fmt.Println(v1, len(v1), cap(v1)) // NOTE: 1, 1, 5
	v2 := nums[1:2:cap(nums)]
	fmt.Println(v2, len(v2), cap(v2)) // NOTE: 1, 1, 5

	// three index slice
	nv := nums[1:2:3]
	fmt.Println(nv, len(nv), cap(nv)) // 1, 2-1=1, 3-1=2
	fmt.Println(nv[0:2])
	// fmt.Println(nv[0:3]) // slice bounds out of range

	v3 := nums[:0] // make v3 as an empty slice
	v3 = append(v3, 100)
	fmt.Println(v3, len(v3), cap(v3), nums) // NOTE: nums被修改了
}
