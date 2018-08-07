package main

import "fmt"

type T struct {
	a int
	b float32
	c string
}

func main() {
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
}

// 开发可以通过String方法，按照自己的意图打印内容
// func Printf(format string, a ...interface{}) (n int, err error)
// golang应该是利用反射来实现的
func (t *T) String() string {
	return fmt.Sprintf("%d / %f / %q", t.a, t.b, t.c)
}
