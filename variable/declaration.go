package main

import "fmt"

func main() {
	// 通过关键字var来声明变量
	// NOTICE: 在函数内声明的变量是有初始值的。其它语言是未定义的情况
	var name string    // 初始值为空字符串
	var age int        // 初始值为0
	var salary float64 // 初始值为0
	var sex bool       // 初始值为false
	fmt.Println(name, age, salary, sex)

	// 数组和切片
	var friends [10]string // 定义了一个十个元素的数组，元素初始化为空
	fmt.Println(friends)

	// 函数内定义并初始化结构体
	var v5 struct {
		age int
	}
	fmt.Println(v5)

	var v0 []string // 根据slice的数据类型，应该是(nil, 0, 0)
	var v1 *int     // 定义了一个指针，指向nil(NULL)
	var v2 map[string]int
	var v3 func(a int) int
	var v4 chan int
	if v0 == nil && v1 == nil && v2 == nil && v3 == nil && v4 == nil {
		fmt.Println("引用类型的默认初始化值为nil")
	}
	fmt.Println(v0, v1, v2, v3, v4) // 虽然它们的输出值为[], map[]等，但是仍然为nil

	// 定义一组变量
	var (
		name2   string // 不允许再使用name变量，否则golang会报告重定义错误
		age2    int    // 注意:没有逗号
		salary2 float64
		sex2    bool
	) // 处于单独的一行
	fmt.Println(name2, age2, salary2, sex2)

	// 声明并定义
	// 以下三种方式均可
	var v11 float64 = 10
	var v12 = 10 // Go will infer the type of initialized variables.
	v13 := 10    // 局部变量定义的推荐形式
	fmt.Println(v11, v12, v13)
}
