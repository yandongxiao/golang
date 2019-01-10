// 缺点：如果key的value在map中确实存在，但是值为null. 如何区分？
// Java: m.get(key)). 如果key存在，则返回对应的value；否则返回null.
// python: m.get(key) 与java保持一致；访问形式如果为m[key], 则抛出异常
package main

import "fmt"

type Person struct {
	name string
	age  int
}

func ExampleGet() {
	// NOTE: 类似golang当中，get永远不会抛出异常，如果该key不存在，
	// 返回value的一个zero-value。
	// Person 是一个值类型，它的zero-value是{"", 0}
	persons := make(map[string]Person)
	fmt.Println(persons["jack"])

	// golang中区分zero-value和key的值不存在的方法
	v, ok := persons["jack"]
	fmt.Println(v, ok)

	// range接收nil值
	var kvs map[string]Person
	for k, v := range kvs {
		fmt.Println(k, v)
	}

	// Output:
	// { 0}
	// { 0} false
}
