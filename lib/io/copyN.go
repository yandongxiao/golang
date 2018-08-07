package main

import "os"
import "fmt"
import "io"

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var src *os.File
	var dest *os.File
	var err error

	src, err = os.Open("/tmp/data")
	errCheck(err)
	defer src.Close()

	dest, err = os.Create("/tmp/data.cp")
	errCheck(err)
	defer dest.Close()

	for {
		// CopyN copies n bytes (or until an error) from src to dst. It returns the
		//	number of bytes copied and the earliest error encountered while copying.)
		// 错误返回时，size个拷贝还是有效的.
		// if 100 > file size, 返回EOF; else size==100 and err == nil
		// 假设文件大小为10， 每次copyN调用会拷贝5个字节，那么panic是发生在第三次调用.
		size, err := io.CopyN(dest, src, 5)
		errCheck(err)
		fmt.Println("size =", size)
	}
}
