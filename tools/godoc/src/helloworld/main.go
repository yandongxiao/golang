// godoc -http :8080
// http://localhost:8080/src/helloworld/main.go 注意：需要设置GOPATH
// Godoc的使用
//	1. 在package, const, type, func等关键字上面并且紧邻关键字的注释才会被展示
//	2. 有效的关键字注释不应该超过3行
//	3. Package的注释如果超过3行, 应该放在当前包目录下一个单独的文件中, 如:doc.go
//	4. 如果当前包目录下包含多个Package注释的go文件(包括doc.go), 那么按照文件名的字母数序优先显示

/*
如果包含很多行注释，采用块注释，参见net.go
	1. Comments do not need extra formatting such as banners of stars.
	2. The comments are uninterpreted plain text. 所以，虽然可以抽取的文档将会在WEB页面展示，但是嵌入<h1>等标签却是不起作用的
	3. 对于注释内容是不加修改的，所以保证注释的拼写正确，标点符号正确
	4. Every exported (capitalized) name in a program should have a doc comment.
	5. 注释符//后面要加空格, 例如: // xxx
	6. package以Package name为注释的开头

abc

属于abc的段落. http://www.baidu.com

def

属于def的段落
*/
package helloworld

import "fmt"

// INT is an alias of int
type INT int // type, const, func以名称为注释的开头

// 不推荐在此处进行注释
var Val = 10   // Var is an global varibale
const NUM = 10 // NUM is an const

// Foo will print "helloworld" string.
//	1. type, const, func以名称为注释的开头
//
// BUG(abc): 以BUG(abc)开头的注释, 将被识别为已知bug, 显示在bugs区域
func Foo() { // 此处的注释不会被展示
	fmt.Println("helloworld")
}

// A ReadError reports an error encountered while reading input.
//
// Deprecated: Use foo instead. (被废弃的函数注释)
func Bar() {
	fmt.Println("helloworld")
}
