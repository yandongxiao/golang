// 打印当前协程的堆栈信息
package main

import (
	"runtime/debug"
	"time"
)

func main() {
	go a()
	m1()
}
func m1() {
	m2()
}
func m2() {
	m3()
}
func m3() {
	time.Sleep(time.Second * 5)
}
func a() {
	debug.PrintStack()
	time.Sleep(time.Second * 5)
}
