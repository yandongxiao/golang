// 如何避免自己的代码被他人引用
// When the go command sees an import of a package with internal
// in its path, It verifies that the package doing the import
// is within the tree rooted at the *parent* of the internal
// directory.
// 参见newinit.Foo()
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
	// "./b/internal/doinit2"
)

func ExampleImport() {
	// NOTE: mock init, 在init中的打印信息是无法追踪的
	fm.Println("helloworld")
	mock.Foo() // 在doinit目录下，package的名称为mock.
	// newinit.Foo() // doint2
	Create("/tmp/data")

	// Output:
	// helloworld
	// mock Foo
}
