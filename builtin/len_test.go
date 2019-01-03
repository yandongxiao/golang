// func len(v Type) int
package main

import "fmt"

func ExampleArrayLen() {
	// Array: the number of elements in v.
	fmt.Println(len([3]int{}))
	fmt.Println(len(&[3]int{}))
	var p *[3]int
	fmt.Println(len(p), p)
	//Output:
	//3
	//3
	//3 <nil>
}

func ExampleNil() {
	// Slice, or map: the number of elements in v; if v is nil, len(v) is zero
	// NOTE: 直接传递nil是不可以的
	var s []int
	var m map[int]int
	var p *[3]int
	fmt.Println(len(s), len(m), len(p), p) // 所以不能以len(p)的形式遍历所有元素
	// Output:
	// 0 0 3 <nil>
}

func ExampleChan() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(len(ch))
	close(ch)
	fmt.Println(len(ch))
	//Output:
	//1
	//1
}
