// NOTICE: the examples in this package assume a Unix system. They may not run on Windows

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// 根据Linux系统的
	cmd := exec.Command("ls", "/tmp")

	// Path is the path of the command to run.
	// This is the only field that must be set to a non-zero value
	// If Path is relative, it is evaluated relative to Dir.
	fmt.Println(cmd.Path)

	// Args holds command line arguments, including the command as Args[0].
	// If the Args field is empty or nil, Run uses {Path}.
	// In typical use, both Path and Args are set by calling Command
	fmt.Println(cmd.Args)

	// Dir specifies the working directory of the command.
	// If Dir is the empty string, Run runs the command in the calling process's current directory.
	cmd.Dir = "/tmp"
	fmt.Println(cmd.Dir)
	data, _ := cmd.Output()
	fmt.Println(string(data))
	fmt.Println(cmd.Dir)
	data, _ = cmd.Output()
	fmt.Println(string(data))
}
