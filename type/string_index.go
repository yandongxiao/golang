package main

import "fmt"

func main() {
	// 中文
	// 依据下标取字符串中的字符，类型为byte
	// Go语言的大多数AIP都假设字符串是byte类型构造的
	name := "李磊"
	n := len(name)
	for i := 0; i < n; i++ {
		fmt.Printf("%v, %c\n", name[i], name[i])
	}

	// ch的类型为rune, 实际是int类型
	// 关于rune的相关操作，可查阅Go标准库的unicode包
	for _, ch := range name {
		fmt.Printf("%v, %c\n", ch, ch)
	}
}
