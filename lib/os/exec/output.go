package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("ls", "/tmp")
	cmd.Run()
	// Output runs the command and returns its standard output.
	// 可见，Output方法对子进程的标准输出是做了修改的
	fmt.Println(cmd.Output()) // exec: already started

	// Stderr holds a subset of the standard error output from the
	// Cmd.Output method if standard error was not otherwise being collected.

	cmd = exec.Command("ls", "/tmpd")
	if _, err := cmd.Output(); err != nil {
		fmt.Println(err)
		fmt.Printf("%s\n", err.(*exec.ExitError).Stderr)
	}

	cmd = exec.Command("sh", "-c", "echo stdout; echo 1>&2 stderr")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}
