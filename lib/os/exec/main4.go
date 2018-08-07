package main

import "os/exec"

func myroutine() {
	println("myroutine")
}

func main() {
	/* wired command */
	cmd1 := exec.Command("killall", "sleep")
	defer cmd1.Run()
	defer println("defer")
	go myroutine()
	a := 10
	b := 0
	println(a / b) /* support */
}
