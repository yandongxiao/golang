package main

import "fmt"

func main() {
	var n int16 = 16
	var m int32 = 32

	// golang不支持隐式类型转换
	// compiler error: cannot use n (type int16) as type int32 in assignment
	//m = n
	m = int32(n)

	fmt.Printf("32 bit int is:  %d\n", m)
	fmt.Printf("16 bit int is:  %d\n", n)
}
