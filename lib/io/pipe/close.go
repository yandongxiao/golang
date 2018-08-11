// Close will complete once pending I/O is done.
package main

import (
	"fmt"
	"io"
	"time"
)

func main() {
	pr, pw := io.Pipe()

	go func() {
		fmt.Println(pw.Write([]byte("hello")))
		pr.Close()
		fmt.Println(pw.Write([]byte("world")))
	}()

	// 这个Sleep看似只与主协程有关系，但是因为golang pipe是同步读写
	// 所以新协程的Write操作也会被阻塞
	time.Sleep(time.Millisecond * 100)

	fmt.Println(pr.Read(make([]byte, 100)))
	fmt.Println(pr.Read(make([]byte, 100)))
}
