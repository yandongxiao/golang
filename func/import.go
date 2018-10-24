package main

import (
	fm "fmt" // alias import

	. "os" // NOTE: 不推荐, 存在兼容性问题

	_ "fmt" // 运行package fmt的init方法

	// 导入相对路径下的一个package。
	// 一般要求: doint 既是一个目录, 也是这个目录下的pacakge的名称
	// NOTE：以上规则非强制
	// 标准做法 mock "./b/doinit"
	"./b/doinit"
)

func main() {
	fm.Println("helloworld")
	mock.Foo() // 在doinit目录下，package的名称为mock.
	Create("/tmp/data")
}
