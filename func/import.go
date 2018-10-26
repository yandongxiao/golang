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

	"gitlab.hz.chinac.com/cos/acoustic/utils/time"
	// 如何避免自己的代码被他人引用
	// When the go command sees an import of a package with internal in its path,
	// It verifies that the package doing the import is within the tree rooted at the *parent* of the internal directory.
	// "./b/internal/doinit2"	// use of internal package not allowed
	// package uuid // import "github.com/yandongxiao/uuid"
	// code in directory /Users/dxyan06/go/src/gitlab.hz.chinac.com/cos/acoustic/utils/uuid expects import "github.com/yandongxiao/uuid"
	// "gitlab.hz.chinac.com/cos/acoustic/utils/uuid"
)

func main() {
	fm.Println("helloworld")
	mock.Foo() // 在doinit目录下，package的名称为mock.
	fm.Println(time.Now())
	// newinit.Foo()	// doint2
	// fmt.Println(uuid.Get())
	Create("/tmp/data")
}
