package main

import "os"

func main() {
	attr := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   "/tmp",
		Env:   []string{"PATH=/usr/bin", "DD=XX"},
	}
	var args []string = []string{"/tmp/show_env.sh"}
	os.StartProcess(args[0], args, &attr)
}
