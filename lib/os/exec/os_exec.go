package main

import "os/exec"

func main() {
	cmd := exec.Command("ls", "/tmp")
	data, _ := cmd.Output()
	println(string(data))
}
