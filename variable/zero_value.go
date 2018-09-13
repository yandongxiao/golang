package main

import "fmt"

func main() {
	// 通过关键字var来声明变量
	// NOTICE: 在函数内声明的变量也是有初始值的, 这在其它语言中是未定义的情况
	var name string    // 初始值为空字符串
	var age int        // 初始值为0
	var salary float64 // 初始值为0
	var sex bool       // 初始值为false
	fmt.Println(name, age, salary, sex)

	// 数组和切片
	var friends [10]string // 定义了一个十个元素的数组，元素初始化为空
	fmt.Println(friends)

	// anonymous struct
	var v5 struct {
		age int
	}
	fmt.Println(v5)

	// NOTICE: 引用类型的zero value是nil. len(nil) = 0, cap(nil) = 0.
	// nil只适合赋值给引用类型, 如果赋值给值类型，则报错
	var v0 []string // NOTICE: 根据slice的数据类型，应该是(nil, 0, 0).
	var v1 *int     // 定义了一个指针，指向nil
	var v2 map[string]int
	var v3 func(a int) int
	var v4 chan int
	if v0 == nil && v1 == nil && v2 == nil && v3 == nil && v4 == nil {
		fmt.Println("引用类型的默认初始化值为nil")
	}
	fmt.Println(v0, v1, v2, v3, v4) // 虽然它们的输出值为[], map[]等，但是仍然为nil.
	v0 = make([]string, 0, 10)
	fmt.Println(v0, v0 == nil) // [] false. NOTICE： 我们不能通过fmt.Println的输出结果，判断v0是nil值还是长度为0的切片

	// 声明并定义
	// 以下三种方式均可
	var v11 float64 = 10
	var v12 = 10 // Go will infer the type of initialized variables.
	v13 := 10    // 局部变量定义的推荐形式
	fmt.Println(v11, v12, v13)
}
