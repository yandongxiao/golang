package main

import "os/exec"

func myroutine() {
	println("myroutine")
}

func main() {
	cmd1 := exec.Command("sleep", "1000")
	defer cmd1.Run()
	defer println("defer")
	go myroutine()
	a := 10
	b := 0
	println(a / b) /* support */
}
