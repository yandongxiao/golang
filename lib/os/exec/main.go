package main

import "os/exec"

func main() {
	cmd1 := exec.Command("sleep", "1000")
	defer cmd1.Run()
	defer println("defer")
	//	a := 10
	//	b := 0
	//	println(a / b)      /* support */

	/* not ok for signal */
	cmd2 := exec.Command("sleep", "1000")
	println("start sleep")
	cmd2.Run()
	println("helloworld")
}
