package main

import "fmt"

func ExampleGetAddrOfStringElmt() {
	// 不能⽤用序号获取字节元素指针，&s[i] 非法
	// 编译错误
	// s := "helloworld"
	// fmt.Println(&s[0])
}

func ExampleStringAsFile() {
	s2 := `文件内容
hello
world`
	fmt.Println(s2)
	// Output:
	// 文件内容
	// hello
	// world
}

func ExampleSubString() {
	v := "helloworld"
	fmt.Println(v[2:4])
	// Output:
	// ll
}
