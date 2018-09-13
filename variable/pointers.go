package main

import "fmt"

func zeroptr(v *int) {
	*v = 0
}

// type *[]int does not support indexing
// 引用类型的变量本身就是指针，所以不支持*[]int
func zeroArray(arr *[3]int) {
	for i := range *arr {
		arr[i] = 0
	}
}

func main() {
	// cannot use v (type int) as type *int in argument to zeroptr
	// 不会自动进行类型转换
	// zeroptr(v)
	// println(v)
	var v int
	zeroptr(&v)
	println(v)

	arr := [3]int{1, 2, 3}
	zeroArray(&arr)
	fmt.Println(arr)
}
