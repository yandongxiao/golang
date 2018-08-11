// 证明写读操作是一对多的关系
package main

import (
	"fmt"
	"io"
)

func main() {
	r, w := io.Pipe()

	go func() {
		fmt.Fprint(w, "hello")
		fmt.Fprint(w, "world")
		w.Close()
	}()

	// 一对多
	fmt.Println(r.Read(make([]byte, 3)))
	fmt.Println(r.Read(make([]byte, 2)))

	// 一对一
	fmt.Println(r.Read(make([]byte, 100)))
}
