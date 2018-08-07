package main

import "fmt"

type INT int

func mustCase() {
	num1 := 10
	num2 := 1.2
	var num3 INT = 30
	println(num1 + int(num2))
	println(num1 + int(num3)) /* must cast */
}

func noNeedToCast() {
	type INTS []int
	var a = INTS{10}
	var b = []int{20}
	a = b // 这里又不是必须的
	b = a
	fmt.Println(a, b)
}

func main() {
	mustCase()
	noNeedToCast()
}
