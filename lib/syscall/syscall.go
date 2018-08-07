package main

import "syscall"

func main() {
	println(syscall.Getpid())
}
