// Getwd 返回当前进程的工作目录
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Getwd())

	os.Chdir("/tmp")

	fmt.Println(os.Getwd())
}
