package main

import (
	"fmt"
	"os"
)

func main() {
	attr := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
	}
	var args []string = []string{"/bin/sleep", "100"}

	process, err := os.StartProcess(args[0], args, &attr)
	if err != nil {
		panic(err)
	}

	fmt.Println(process.Pid)
}
