package main

import fm "fmt" // alias import

import . "os" // 不推荐

import _ "fmt" // 运行package fmt的init方法

// 导入相对路径下的一个package。
// NOTE: doint 既是目录b下面的一个子目录, 也是这个目录下的pacakge的名称
import _ "./b/doinit"

func main() {
	fm.Println("helloworld")
	Create("/tmp/data")
}
