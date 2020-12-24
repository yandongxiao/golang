package main

import "syscall"

func ExampleA() {
	println(syscall.Getpid())
}
