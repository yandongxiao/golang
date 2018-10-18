// 1. 值类型是可以比较的，包括[3]int{}
// 2. 引用类型是否可比较，跟具体的引用类型有关系.
//		slice, map, function: can only be compared to nil
//		chan: 引用同一个chan, 则相等
//		interface: 是可以比较的
package main

func test1() {}
func test2() {}

func main() {
	// slice类型不可比较
	// 注意：[3]int{}是可以比较的，[]int{}是不可以比较的。
	// invalid operation: b == c (slice can only be compared to nil))
	// var b = []int{}
	// var c = []int{}
	// println(b == c)

	// map类型不可比较
	// invalid operation: bb == cc (map can only be compared to nil))
	// var bb = map[int]int{}
	// var cc = map[int]int{}
	// println(bb == cc)

	// 函数类型不可比较
	// invalid operation: test1 == test2 (func can only be compared to nil))
	// println(test1 == test2)

	// 引用同一个chan, 则相等
	var b = make(chan int)
	var c = b
	println(b == c) // true
	c = make(chan int)
	println(b == c) // false

	// interface的比较规则如下：
	//	1. 首先比较的两个对象必须是同一种类型，即实现了相同的接口
	//	2. 如果底层类型不相同，则返回false
	//	3. 如果底层类型相同，但是类型不支持比较，如map, slice, 则**panic**
	//	4. 如果支持比较，则按照底层类型的比较规则进行比较.
	// 更多信息参见type/type/rune.go
	var d interface{} = [3]int{1, 2, 3}
	var e interface{} = [3]int{1, 2, 3}
	println(d == e) // true

	type IntArray [3]int
	x := IntArray{1, 2, 3}
	y := [3]int{1, 2, 3}
	var f interface{} = x
	var g interface{} = y
	println(x == y) //true
	println(f == g) //false, 满足了第二条规则
}
