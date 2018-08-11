package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// CopyN的描述
// If dst implements the ReaderFrom interface, the copy is implemented using it.

// Copy的描述
// If src implements the WriterTo interface, the copy is implemented by
// calling src.WriteTo(dst). Otherwise, if dst implements the ReaderFrom
//interface, the copy is implemented by calling dst.ReadFrom(src).

// CopyN为什么不能用src.WriteTo(dst) ？ CopyN的工作原理如下：
// 1. Reader --> LimitedReader
// 2. Copy(dst, LimitReader(src, n))
// LimitedReader并没有实现WriteTo接口！

func main() {
	// On return, written == n if and only if err == nil.
	// if 100 > file size, 返回EOF; else size==100 and err == nil
	reader := strings.NewReader("helloworld")
	// func CopyN(dst Writer, src Reader, n int64) (written int64, err error)
	fmt.Println(io.CopyN(os.Stdout, reader, 101))
}
