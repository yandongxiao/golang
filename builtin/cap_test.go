// 你可以认为map的capacity是无限大
package main

import "fmt"

//	if v is nil, cap(v) is zero.
func ExampleNil() {
	var v []int
	fmt.Println(cap(v))
	//Output:
	// 0
}

//	Array: the number of elements in v (same as len(v)).
func ExampleArray() {
	a := [3]int{1, 2, 3}
	fmt.Println(cap(a))
	//Output:
	// 3
}

//	Pointer to array: the number of elements in *v (same as len(v)).
func ExamplePointer() {
	p := &[3]int{1, 2, 3}
	fmt.Println(cap(p))
	//Output:
	// 3
}

//	Slice: the maximum length the slice can reach when resliced;
func ExampleCapSlice() {
	a := [3]int{1, 2, 3}
	s := a[:2]
	fmt.Println(cap(s))
	//Output:
	// 3
}

//	Channel: the channel buffer capacity, in units of elements;
func ExampleChan() {
	ch := make(chan int, 1)
	fmt.Println(cap(ch))
	//Output:
	//1
}
