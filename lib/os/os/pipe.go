// 参见 io.Pipe，比较两种pipe的差异
package main

import "os"
import "fmt"
import "time"

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

func input(in *os.File) {
	for {
		buf := make([]byte, 1)
		fmt.Println(in.Read(buf))
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	in, out, err := os.Pipe()
	errcheck(err)
	go input(in)
	out.Write([]byte("nihao"))
	// 如果os.Close()函数不存在，input协程在读取完管道内的数据后（异步读取），阻塞
	// 如果os.Close()函数存在，read返回0, io.EOF （来自io.Pipe.README.md的介绍）
	out.Close()
	time.Sleep(time.Second)
}
