package main

import "fmt"

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
