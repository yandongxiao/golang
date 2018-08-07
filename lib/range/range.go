package main

import "fmt"

func main() {
	// slice
	// range on arrays and slices provides both the index and value for each entry.
	for idx, x := range []int{1, 2, 3} {
		fmt.Printf("%d %d\n", idx, x)
	}

	// map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%v %v\n", k, v)
	}
	// 也可以只是遍历key
	for k := range kvs {
		fmt.Printf("%v\n", k)
	}

	// range on strings iterates over Unicode code points.
	// The first value is the starting byte index of the rune and the second the rune itself.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
