package main

// 测试函数参数i的值
// If the defer has arguments they are evaluated at the line of the defer-statement.
func foo() {
	var i = 10
	defer func(v int) {
		println(v) // 10
	}(i)
	i = 20
}

func bar() {
	var i = 10
	defer func() {
		println(i) // 20
	}()
	i = 20
}

func main() {
	foo()
	bar()
}
