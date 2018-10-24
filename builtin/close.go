// NOTE: close函数只是用来关闭chan，与文件操作符无关
package main

func main() {
	ch := make(chan struct{})
	close(ch)
}
