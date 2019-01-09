package main

import "fmt"

func ExampleBasicType() {
	// 通过关键字var来声明变量
	// NOTE: 在函数内声明的变量也是有初始值的, 这在其它语言中是未定义的情况
	var name string    // 初始值为空字符串
	var age int        // 初始值为0
	var salary float64 // 初始值为0
	var sex bool       // 初始值为false
	fmt.Println(name, age, salary, sex)
	// Output:
	//  0 0 false
}

func ExampleSlice() {
	// 数组和切片
	var friends [10]string // 定义了一个十个元素的数组，元素初始化为空
	fmt.Println(friends)
	// Output:
	// [         ]
}

func ExampleStruct() {
	// anonymous struct
	var v5 struct {
		age int
	}
	fmt.Println(v5)
	// Output:
	// {0}
}

func ExampleReferenceType() {
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
	// Output:
	// 引用类型的默认初始化值为nil
}
