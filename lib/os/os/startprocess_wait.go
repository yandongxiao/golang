package main

import (
	"fmt"
	"os"
)

func errcheck(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	attr := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	args := []string{"/bin/sleep", "1"}

	process, err := os.StartProcess(args[0], args, &attr)
	errcheck(err)

	stat, err := process.Wait()
	errcheck(err)

	fmt.Println(stat.Exited())
	fmt.Println(stat.Pid())
	fmt.Println(stat.String())
	fmt.Println(stat.Success())
	fmt.Println(stat.SysUsage())
	fmt.Println(stat.SystemTime())
	fmt.Println(stat.UserTime())
}
