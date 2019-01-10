package main

import "fmt"

func ExampleNil() {
	// if v[k] is used as a destination value in an assignment and v is a nil map, a panic will occur at run time.
	// if v[k] is used to retrieve the element value corresponding key k in map v, then no panics will occur, even if v is a nil map.
	var v map[int]int
	fmt.Println(v[0])
	// Output:
	// 0
}

func ExampleAddressability() {
	m := map[int]bool{1: true}
	_ = m
	// Elements of map values are always unaddressable.
	// _ = &map[int]bool{1: true}[1]
	// _ = &m[1]
}

func ExampleModifyElement1() {
	// Unlike most other unaddressable values, which direct parts can not be modified,
	// the direct part of a map element values can be modified, but can only be modified
	// (overwritten) as a whole.
	type T struct{ age int }
	mt := map[string]T{}
	mt["John"] = T{age: 29} // modify it as a whole
	ma := map[int][5]int{}
	ma[1] = [5]int{1: 789} // modify it as a whole
	// ma[1][1] = 123      // error: cannot assign to a[1][1]
	// mt["John"].age = 30 // cannot assign to struct field. mt["John"].age in map.
	_ = mt
	_ = ma
	// Output:
	//
}

func ExampleModifyElement2() {
	ma := map[int][]int{}
	ma[1] = []int{1: 789} // modify it as a whole
	ma[1][1] = 100
	fmt.Println(ma[1][1])
	// Output:
	// 100
}

func ExampleShare() {
	// If a map is assigned to another map, then the two maps will share
	// all (underlying) elements. Appending elements into (or deleting elements
	// from) one map will refect on the other map.
	// Like map assignments, if a slice is assigned to another slice, they will
	// share all (underlying) elements. Their respective lengths and capacities
	// equal to each other. However, if the length/capacity of one slice changes
	// later, the change will not reflect on the other slice.
	type INTMAP map[int]int
	v1 := make(INTMAP)
	v1[1] = 1
	v2 := v1
	v1[2] = 2
	fmt.Println(len(v2))
	// Output:
	// 2
}

func ExampleMap() {
	// 定义并初始化
	m1 := map[string][]string{
		"name": []string{"nihao"},
	}
	fmt.Println(m1)

	// add
	// To create an empty map, use the builtin make
	persons := make(map[string][]string)
	persons["k1"] = []string{"1"}
	persons["k2"] = nil
	fmt.Println(persons)

	// get 这个方法的返回值很有考究
	if persons["k2"] == nil {
		fmt.Println("如何区分：k2在map当中，值为nil；和k2不在map当中")
	}

	// 区分上述情况的方法
	// The optional second return value when getting a value from a map indicates if the key was present in the map
	// This can be used to disambiguate between missing keys and keys with zero values like 0 or "".
	v, exist := persons["k2"]
	fmt.Println(v, exist)
	v, exist = persons["k3"]
	fmt.Println(v, exist)

	// delete
	delete(persons, "k1")
	delete(persons, "k111")
	fmt.Println(persons)

	// Output:
	// map[name:[nihao]]
	// map[k1:[1] k2:[]]
	// 如何区分：k2在map当中，值为nil；和k2不在map当中
	// [] true
	// [] false
	// map[k2:[]]
}
