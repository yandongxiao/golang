// func len(v Type) int
package main

import "fmt"

func main() {

	// Array: the number of elements in v.
	fmt.Println("array:", len([3]int{}))

	// Pointer to array: the number of elements in *v (even if v is nil).
	var p *[3]int
	fmt.Println("pointer:", len(&[3]int{}))
	fmt.Println("pointer to nil:", len(p), p) // 3

	// Slice, or map: the number of elements in v; if v is nil, len(v) is zero
	fmt.Println("slice:", len([]int{10: 10}))     // 11
	fmt.Println("map:", len(map[int]int{10: 10})) // 1

	// String: the number of bytes in v.
	fmt.Println("string:", len("你好"))

	// Channel: the number of elements queued (unread) in the channel buffer
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println("chan:", len(ch))
	close(ch)
	fmt.Println("chan:", len(ch))

	// if v is nil, len(v) is zero.
	// NOTE: 直接传递nil是不可以的
	fmt.Println("nil:", len([]int(nil)))
}
