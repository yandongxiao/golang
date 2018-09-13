// scope等级：
//	global level or package level
//	local level
//	structure level
//  下层可以隐藏上层变量
package main

import "fmt"

func test1() {
	m := 10
	// 在同一级block内，下面的操作只会定义变量n
	m, n := 20, 30
	fmt.Println(m, n)
}

func test2() {
	m := 10
	// 在下面的block内，重新定义了m和n，所以上面的m变量被隐藏了
	// structure level
	if m, n := 20, 30; n == 30 {
		fmt.Println(m, n)
	}
	fmt.Println(m) // m=10
}

func main() {
	test1()
	test2()
}
