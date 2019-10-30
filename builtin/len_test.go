// func len(v Type) int
package main

import "fmt"

func ExampleArrayLen() {
	fmt.Println(len([3]int{}))
	fmt.Println(len(&[3]int{}))
	var p *[3]int
	fmt.Println(len(p), p)
	//Output:
	//3
	//3
	//3 <nil>
}

func ExampleLenNil() {
	var s []int
	var m map[int]int
	var p *[3]int
	// NOTE: 所以不能以len(p)的形式遍历所有元素
	fmt.Println(len(s), len(m), len(p), p)
	// Output:
	// 0 0 3 <nil>
}

func ExampleLenChan() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println(len(ch))
	close(ch)
	fmt.Println(len(ch))
	<-ch
	fmt.Println(len(ch))
	//Output:
	//1
	//1
	//0
}
