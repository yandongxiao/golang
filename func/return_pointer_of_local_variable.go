package main

// golang允许你返回了一个"栈变量"的地址
// 在golang当中，你不需要关心对象地址空间是如何分配的
func ttt() *int {
	num := 10
	println(&num)
	return &num
}

func main() {
	p := ttt()
	*p = 100 * 100
	println(p, *p)
}
