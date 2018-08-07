package main

// 导入相对路径下的一个package。
// NOTE: doint 既是目录b下面的一个子目录, 也是这个目录下的pacakge的名称
import _ "./b/doinit"

// 实验：如果目录doinit2下的package名称是newinit，即目录名称和package名称不一致
// b/doinit2/newinit: no such file or directory
// import _ "./b/doinit2/newinit"

func main() {
	println("go")
}
