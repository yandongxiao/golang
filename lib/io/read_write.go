package main

import (
	"fmt"
	"io"
	"os"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

// 关于Read和Write方法，参见README
func main() {
	src, err := os.Open("/tmp/data")
	errCheck(err)
	defer src.Close()

	dest, err := os.Create("/tmp/data.cp")
	errCheck(err)
	defer dest.Close()

	buffer := make([]byte, 1)
	for {
		if rn, rerr := src.Read(buffer); rerr != nil {
			if rn > 0 {
				if _, werr := dest.Write(buffer[:rn]); werr != nil {
					panic("读失败: " + rerr.Error() + ". 写失败：" + werr.Error())
				}
			}
			if rerr == io.EOF {
				println("读结束: " + rerr.Error())
				break
			}
		} else {
			if _, err := dest.Write(buffer[:rn]); err != nil {
				panic("写失败：" + err.Error())
			} else {
				fmt.Printf("%s", buffer[:rn])
			}
		}
	}
}
