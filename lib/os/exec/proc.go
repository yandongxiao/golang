// Package exec runs external commands. It wraps os.StartProcess to make it easier to remap stdin and stdout,
// connect I/O with pipes, and do other adjustments.
package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	// Command returns the Cmd struct to execute the named program
	// It sets **only** the Path and Args in the returned structure.
	// If name contains no path separators, Command uses LookPath to
	// resolve the path to a complete name if possible. Otherwise it uses name directly.
	//cmd := exec.Command("cat", "/tmp/data")
	cmd := exec.Command("cat")

	// It sets **only** the Path and Args in the returned structure.
	fmt.Println(cmd.Path)
	fmt.Println(strings.Join(cmd.Args, " "))

	// If Env is nil, Run uses the current process's environment.
	cmd.Env = os.Environ() // 所以，可忽略
	fmt.Println(cmd.Env)

	// If Dir is the empty string, Run runs the command in the calling process's current directory.
	cmd.Dir, _ = os.Getwd() // 所以，可忽略
	fmt.Println(cmd.Dir)

	// If either is nil, Run connects the corresponding file descriptor to the null device (os.DevNull).
	// If Stdin is nil, the process reads from the null device (os.DevNull).
	// 所以，默认情况下，标准输入，标准输出，标准出错都定位到了/dev/null
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	// cmd.Stdin = os.Stdin
	// 或
	stdin, _ := cmd.StdinPipe() // 之前，不能对cmd.Stdin赋值
	stdin.Write([]byte("helloworld\n"))
	stdin.Close() // 如果子进程只有等到读取EOF，才退出，而父进程未调用Close函数，Wait操作将无法返回

	// cmd.Process is the underlying process, once started.
	// 用于完成进程控制功能

	cmd.Run()
	fmt.Println(cmd.Output()) // exec: Stdout already set
}
