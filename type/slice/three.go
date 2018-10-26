// three index slice
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	nv := nums[1:2:3] // three index slice
	fmt.Println(nv)
	fmt.Println(len(nv))
	fmt.Println(cap(nv))
	fmt.Println(nv[0:2])
	// fmt.Println(nv[0:3]) // slice bounds out of range
}
