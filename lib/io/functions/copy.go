package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// test1 基本用法
func test1() {
	r := strings.NewReader("some io.Reader stream to be read\n")
	// A successful Copy returns err == nil, not err == EOF.
	if n, err := io.Copy(os.Stdout, r); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("read: ", n)
	}
}

// test2 拷贝大文件
// Copy 操作的基本实现，读上来，缓存数据(默认32KB)，写数据。 一次读对应一次写
// func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)
//	自己设定缓存大小.	CoptBuffer(src, dest, nil) == Copy(dst, src)
func test2() {
	var src *os.File
	var dest *os.File
	var err error

	src, err = os.Open("/tmp/data")
	errCheck(err)
	defer src.Close()

	dest, err = os.Create("/tmp/data.cp")
	errCheck(err)
	defer dest.Close()

	// 查看返回值我们可知，一次Copy可以拷贝很大的文件
	// func Copy(dst Writer, src Reader) (written int64, err error)
	_, err = io.Copy(dest, src)
	errCheck(err)
}
