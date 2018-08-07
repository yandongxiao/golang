package main

import "os"
import "fmt"
import "io"

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// func Copy(dst Writer, src Reader) (written int64, err error)
// 查看返回值我们可知，一次Copy可以拷贝很大的文件
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
		size, err := io.Copy(dest, src)
		errCheck(err)
		if size == 0 {
			println("读取完毕")
			break
		} else {
			fmt.Println("size =", size)
		}
	}
}
