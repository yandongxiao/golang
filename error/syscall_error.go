package main

import "syscall"
import "fmt"

func main() {
	_, err := syscall.Open("/tmp/aaa", syscall.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(uintptr(err.(syscall.Errno)))
		fmt.Println(err.Error())
	}
}
