package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println(exec.LookPath("ls"))
}
