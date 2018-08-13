// fmt函数调用是不会抛出异常的
// func Println(a ...interface{}) (n int, err error)
package main

import "fmt"

func wrongVerb() {
	fmt.Printf("%d\n", "dsa") // 内容还是打印出来了，%!d(string=dsa)
	fmt.Println("helloworld") // format error并没有导致程序崩溃
}

type INT int

func (i INT) String() string {
	panic("BAD")
}

func panicString() {
	fmt.Println(fmt.Println(INT(10))) // 15 <nil> 最外层的输出
}

func main() {
	panicString()
	fmt.Println("helloworld")
}
