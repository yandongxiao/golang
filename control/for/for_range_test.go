package main

import "fmt"

func ExampleChan_2() {
	// 支持 map, slice, array, pointer
	// 在for执行开始之前，已经明确了遍历的次数。中间删除slice元素的方式不可取
	for i, v := range []int{1, 2, 3} {
		fmt.Printf("%d %d, ", i, v)
	}
	// Output:
	// 0 1, 1 2, 2 3,
}

func ExampleForRange2() {
	// for range 简化形式
	for range []int{1, 2, 3} {
		fmt.Print("--")
	}
	// Output:
	// ------
}

func ExampleNil() {
	// NOTE: range 的参数值可以是nil! 但不支持 for range nil
	var strs []string
	for i := range strs {
		println(strs[i])
	}
	// Output:
	//
}

func ExampleChan() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 1
	close(ch)
	for v := range ch {
		fmt.Printf("%d", v)
	}
	// Output:
	// 121
}

func ExampleArray() {
	// can also with dereferencing *a to get back to the array(不会对数组进行复制)
	// NOTE: 即使p==nil, len(p) 仍然等于3，所以，不要使用i := 0; i < len(p); i++的形式，遍历指向数组的指针
	array := [3]int{7, 8, 9}
	for i := range array {
		array[i] = 0
	}
	fmt.Printf("%v\n", array[0])
	// Output:
	// 0
}

// NOTE: 支持指向数组的指针
func ExamplePointer() {
	// can also with dereferencing *a to get back to the array(不会对数组进行复制)
	// NOTE: 即使p==nil, len(p) 仍然等于3，所以，不要使用i := 0; i < len(p); i++的形式，遍历指向数组的指针
	p := &[3]int{7, 8, 9}
	for i := range p {
		p[i] = 0
	}
	fmt.Printf("%v\n", p[0])
	// Output:
	// 0
}
