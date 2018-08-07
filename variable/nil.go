package main

type Person struct {
	name string
	age  int
}

func main() {
	// 引用类型还包括interface和function
	var s []int
	var m map[int]int
	var c chan int
	if s == nil && m == nil && c == nil {
		println("引用类型初始化值为nil")
	}

	// Java或python当中的map结构：在get之前，需要确定确实存在这个key. 否则抛出异常
	// 所以，java和python允许key的值为NULL/null等类型
	// golang当中，get永远不会抛出异常，如果该key不存在，返回value的一个zero-value。
	// Person 是一个值类型，它的zero-value是{"", 0}
	persons := make(map[string]Person)
	if persons["jack"].name == "" {
		println("因为struct是值类型，persons返回了一个zero-value")
	}

	// golang中区分zero-value和key的值不存在的方法
	if _, ok := persons["jack"]; !ok {
		println("map中不存在该条记录")
	}

	// nil只适合赋值给引用类型
	// cannot use nil as type Person in assignment
	// persons["jack"] = nil
}
