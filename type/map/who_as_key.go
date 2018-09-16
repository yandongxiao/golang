package main

import "fmt"

func arrayAsKey() {
	a := [10]int{1, 2, 3}
	b := [10]int{1, 2, 3}
	vals := make(map[[10]int]int)
	vals[b] = 1
	println(a == b)
	fmt.Println(vals)
}

// (slice can only be compared to nil)
/*
func sliceNotAsKey() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	println(a == b)
}
*/

// (map can only be compared to nil)
/*
func mapNotAsKey() {
	a := map[int]int{}
	b := map[int]int{}
	println(a == b)
}
*/

// (map can only be compared to nil)
func chanAsKey() {
	a := make(chan int)
	b := make(chan int)
	println(a == b) // 注意是false
}

// 注意：
// 如果只想以name作为Person的key的话是比较困难的
type Person struct {
	name string
	age  int
}

func structAsKey() {
	a := Person{name: "jack", age: 10}
	b := Person{name: "jack", age: 20}
	println(a == b)
	vals := map[Person]int{a: 1}
	fmt.Println(vals[b])
}

func main() {
	chanAsKey()
	structAsKey()
}
