// 对于自定义的类型，可以调用len，append等函数

package main

type Stack []interface{}

// 没什么卵用
func (stack Stack) Len() int {
	return 100
}

func main() {
	var st1 Stack
	println(len(st1))
	st1 = append(st1, 100)
	println(len(st1))
}
